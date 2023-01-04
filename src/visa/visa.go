package visa

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-rod/rod"
)

func Login(page *rod.Page) error {
	// Env variables
	email := os.Getenv("VISA_EMAIL")
	password := os.Getenv("VISA_PASSWORD")

	if email == "" || password == "" {
		return errors.New("email or password not provided")
	}

	// Elements
	emailInput := page.MustElement("#user_email")
	passwordInput := page.MustElement("#user_password")
	checkbox := page.MustElement("div.radio-checkbox-group > label.icheck-label")
	submitButton := page.MustElement(`[type="submit"]`)

	emailInput.MustInput(email)
	passwordInput.MustInput(password)
	if !checkbox.MustProperty("checked").Bool() {
		checkbox.MustClick()
	}
	submitButton.MustClick()

	return nil
}

func Navigate(page *rod.Page) {
	continueButton := page.MustElement("div > ul > li > a.button.primary.small")
	continueButton.MustClick()

	paymentAccordion := page.MustElement("section#forms > ul.accordion > li:first-child")
	paymentAccordion.MustClick()

	paymentButton := page.MustElement("p > a.button.small.primary.small-only-expanded")
	paymentButton.MustClick()
}

func IsAvailable(page *rod.Page) (error, bool) {
	err := rod.Try(func() {
		page.Timeout(time.Second * 3).MustElement("div.noPaymentAcceptedMessage > h3")
	})

	if errors.Is(err, context.DeadlineExceeded) {
		return nil, true
	} else if err != nil {
		return err, false
	}

	return nil, false
}
