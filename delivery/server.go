package delivery

import (
	"Kelompok-2/dompet-online/config"
	"Kelompok-2/dompet-online/delivery/controller"
	"Kelompok-2/dompet-online/delivery/middleware"
	"Kelompok-2/dompet-online/docs"
	"Kelompok-2/dompet-online/manager"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) Run() {
	s.initMiddleware()
	s.initControllers()
	s.swagDocs()
	err := s.engine.Run()
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddleware() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
}

func (s *Server) swagDocs() {
	docs.SwaggerInfo.Title = "dompet-online"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.engine.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Server) initControllers() {
	// Inisialisasi Controller
	controller.NewUserController(s.ucManager.UserUseCase(), s.ucManager.WalletUseCase(), s.engine).AuthRoute()
	controller.NewUserController(s.ucManager.UserUseCase(), s.ucManager.WalletUseCase(), s.engine).UsersRoute()
	controller.NewWalletController(s.ucManager.WalletUseCase(), s.engine).Route()
	controller.NewTransactionController(s.ucManager.TransactionUseCase(), s.engine).Route()
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	infraManager, err := manager.NewInfraManagerConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}

	// Instance Repo
	rm := manager.NewRepoManager(infraManager)

	// Instance UC
	ucm := manager.NewUseCaseManager(rm)

	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	log := logrus.New()

	// Controller
	engine := gin.Default()
	return &Server{
		ucManager: ucm,
		engine:    engine,
		host:      host,
		log:       log,
	}
}
