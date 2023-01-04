package mail

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func Send(isAvailable bool) (string, error) {

	mailgunDomain := os.Getenv("MAILGUN_DOMAIN")
	mailgunApiKey := os.Getenv("MAILGUN_API_KEY")

	sender := "Visa Checker <brunoveranoc@gmail.com>"
	subject := getSubject(isAvailable)
	body := getBody(isAvailable)
	recipient := "brunoveranoc@gmail.com"

	mg := mailgun.NewMailgun(mailgunDomain, mailgunApiKey)
	m := mg.NewMessage(sender, subject, "", recipient)
	m.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

func getSubject(isAvailable bool) string {
	if isAvailable {
		return "[Visa Checker] ¡YA PUEDES SACAR TU CITA!"
	}
	return "[Visa Checker] Aún no hay citas"
}

func getBody(isAvailable bool) string {
	var message string
	if isAvailable {
		message = "Ya puedes sacar tu cita de la visa!!!!!!"
	} else {
		message = "Aún no hay citas para la visa."
	}

	var body string = `
	<html>
	<body>
		<h1>%s</h1>
	</body>
	</html>
	`

	return fmt.Sprintf(body, message)
}
