package main

import (
	"Kelompok-2/dompet-online/delivery"
	_ "github.com/lib/pq"
)

// @title           dompet-online
// @version         1.0

// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @schemes http
func main() {
	delivery.NewServer().Run()
}
