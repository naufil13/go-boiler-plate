package controllers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type ToEmail struct {
	Name  string
	Email string
}

func SendMailSimple(c *gin.Context) {

	subject := "Mail Testng Server - 4"
	templatePath := "./templates/email.html"

	emails := []ToEmail{
		{Name: "Naufil khan", Email: "naufil.dys@gmail.com"},
		{Name: "Ali khan", Email: "fahim.naseem@lathran.com"},
		{Name: "Salman khan", Email: "salman.moosa@lathran.com"},
		{Name: "Aamir khan 2", Email: "naufil.khan@lathran.com"},
	}
	//to := []string{"naufil.khan@lathran.com"}

	//err := SendMail(subject, templatePath, to)

	err := sendGoMail(subject, templatePath, emails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email sending failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email has sent successfully",
	})

}

func sendGoMail(subject string, templatePath string, emails []ToEmail) error {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "Naufil khan"})
	if err != nil {
		log.Println("Invalid template file")
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM_ADDRESS"))
	for _, t := range emails {
		m.SetHeader("To", t.Email)
	}

	m.SetAddressHeader("Cc", "developer.naufil@gmail.com", "Naufil")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())
	m.Attach("./assets/file-upload/Top Level.pdf")

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Println("Invalid to convert MAIL_PORT")
		return err
	}

	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), port, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}
	return err
}

func SendMail(subject string, templatePath string, to []string) error {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "Naufil khan"})

	if err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"",
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_HOST"),
	)

	headers := "\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := "Subject: " + subject + " \n " + headers + "\n\n" + body.String()

	errMail := smtp.SendMail(
		os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
		auth,
		os.Getenv("MAIL_FROM_ADDRESS"),
		to,
		[]byte(msg),
	)

	return errMail
}
