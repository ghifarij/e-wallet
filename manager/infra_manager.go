package manager

import (
	"Kelompok-2/dompet-online/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.DbConfig.Host,
		i.cfg.DbConfig.Port,
		i.cfg.DbConfig.User,
		i.cfg.DbConfig.Password,
		i.cfg.DbConfig.Name,
	)
	db, err := sql.Open(i.cfg.DbConfig.Driver, dsn)
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

// constructor
func NewInfraManagerConnection(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
