package main

import (
	"work/db"
	"work/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}