package mailer

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(recipientEmail string, msg string) error {

	// Load the .env file into environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return err
	}
	SMTP_PORT := os.Getenv("SMTP_PORT")
	SMTP_HOST := os.Getenv("SMTP_HOST")
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")

	// Convert the string to an integer using strconv.Atoi()
	SMTP_PORT_INT, err := strconv.Atoi(SMTP_PORT)
	if err != nil {
		fmt.Println("Error converting port to integer:", err)
		return err
	}

	m := gomail.NewMessage()

	// Set the sender and recipient
	m.SetHeader("From", SMTP_USERNAME)
	m.SetHeader("To", recipientEmail)

	m.SetHeader("Subject", "Bible!")
	// m.SetBody("text/html", "Hello <b>chris</b> and <i>Mi</i>!")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(SMTP_HOST, SMTP_PORT_INT, SMTP_USERNAME, SMTP_PASSWORD)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}
