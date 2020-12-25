package main

import (
	"app/db"
	"app/server"
)

func main() {
	db.Init()
	server.Init()
}
