package kindle

import (
	"net/mail"
	"net/smtp"
	"github.com/seppo0010/rss-to-kindle/utils"

	"github.com/scorredoira/email"
)

//Send ...
func Send(server string, port string, fromEmail string, password string, toEmail string, filePath string) {
	m := email.NewMessage("New file", "")
	m.From = mail.Address{Name: "From", Address: fromEmail}
	m.To = []string{toEmail}

	err := m.Attach(filePath)
	utils.ExitIfErr(err)

	auth := smtp.PlainAuth("", fromEmail, password, server)
	err = email.Send(server+":"+port, auth, m)
	utils.ExitIfErr(err)
}
