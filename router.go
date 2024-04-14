package goappointments

import (
	"errors"
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println(errors.New("please set PORT environment variable"))
		port = "8080"
		fmt.Println("Defaulting to port ", port)

	}

	err := InitRouter().Run(":" + port)
	if err != nil {
		return
	}

}
func InitRouter() *gin.Engine {

	r := gin.New()

	return r
}
