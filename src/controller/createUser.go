package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaovictorcruz/golang-api/src/configuration/rest_err"
	"github.com/joaovictorcruz/golang-api/src/controller/dtos/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Algum campo está errado, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}