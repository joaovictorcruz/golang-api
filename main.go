package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaovictorcruz/golang-api/src/controller/routes"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	router.Run(":8080")
	
	if err = router.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}