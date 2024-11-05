package main

import (
	"log"

	"github.com/PhuPhuoc/curanest_exe_be/api"
	"github.com/PhuPhuoc/curanest_exe_be/db/mysql"
)

//	@title		Curanest API for EXE
//	@version	1.0

func main() {
	db, appport := mysql.ConnectDB()
	server_port := ":" + appport
	sv := api.InitServer(server_port, db)
	if err := sv.RunApp(); err != nil {
		log.Fatal("server (main): ", err)
	}
}
