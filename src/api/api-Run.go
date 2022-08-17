package main

import (
	apiLogin "TestApi/src/api/login"
	apiUser "TestApi/src/api/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/account_register", apiLogin.AccountRegister)
	r.POST("/sayHello", apiUser.SayHello)

	r.Run("127.0.0.1:8080")
}
