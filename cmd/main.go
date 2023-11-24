package main

import (
	"lib/internal/routers"
	"lib/pkg"
	"log"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	mainRouter := routers.RegisterRouter(db)
	server := pkg.InitServer(mainRouter)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
