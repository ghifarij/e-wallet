package main

import (
	"Kelompok-2/dompet-online/delivery"
	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
