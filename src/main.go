package main

import (
	"log"

	"github.com/bvsaur/visa-checker/src/config"
	"github.com/bvsaur/visa-checker/src/mail"
	"github.com/bvsaur/visa-checker/src/visa"
	"github.com/go-rod/rod"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	browser := rod.New().MustConnect().NoDefaultDevice()
	page := browser.MustPage("https://ais.usvisa-info.com/es-pe/niv/users/sign_in").MustWindowFullscreen()

	err = visa.Login(page)
	if err != nil {
		log.Fatal(err)
	}

	visa.Navigate(page)

	err, isAvailable := visa.IsAvailable(page)
	if err != nil {
		log.Fatal(err)
	}

	id, err := mail.Send(isAvailable)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Mailgun ID: %s", id)
}
