package main

import (
	"fmt"
	"net/http"
	"os"

	"am.com/pakages/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type EmailBody struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Html    string `json:"html"`
}

type ResponseType struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func home(context *gin.Context) {

	context.JSON(200, gin.H{
		"status": "true",
		"author": "Atanu Debnath",
		"about":  "Email server",
	})
}

func send(context *gin.Context) {
	var emailBody EmailBody
	if err := context.BindJSON(&emailBody); err != nil {
		return
	}

	var successfulres ResponseType
	successfulres.Status = "true"
	successfulres.Message = "email send successful"

	var unsuccessfulres ResponseType
	unsuccessfulres.Status = "false"
	unsuccessfulres.Message = "email not send successful"

	isEmailSend := util.SendEmail(emailBody.To, emailBody.Subject, emailBody.Html)

	if isEmailSend {
		context.IndentedJSON(http.StatusAccepted, successfulres)
	} else {
		context.IndentedJSON(http.StatusAccepted, unsuccessfulres)
	}

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	//router := gin.Default()
	router.GET("/", home)
	router.POST("/send", send)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("error: %s", err)
	}

}
