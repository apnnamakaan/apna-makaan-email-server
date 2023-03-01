package main

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"nvc.com/events/services"
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

func Home(context *gin.Context) {

	context.JSON(200, gin.H{
		"status": "true",
		"author": "Atanu Debnath",
		"about":  "Mail server",
	})
}

func SendMail(context *gin.Context) {
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

	isEmailSend := emailervice.SendMail(emailBody.To, emailBody.Subject, emailBody.Html)

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

	router := gin.Default()
	router.GET("/", Home)
	router.POST("/send", SendMail)
	router.Run("localhost:8083")

}
