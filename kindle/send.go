package kindle

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/fatih/color"
	"github.com/scorredoira/email"
)

//Send ...
func Send(server string, port string, fromEmail string, password string, toEmail string, filePath string) {
	color.New(color.FgCyan).Println("Sending to your kindle email...")
	m := email.NewMessage("New file", "")
	m.From = mail.Address{Name: "From", Address: fromEmail}
	m.To = []string{toEmail}

	if err := m.Attach(filePath); err != nil {
		log.Fatal(err)
	}

	auth := smtp.PlainAuth("", fromEmail, password, server)
	if err := email.Send(server+":"+port, auth, m); err != nil {
		log.Fatal(err)
	}
	color.New(color.FgCyan).Println("File sent successfully.")
}
