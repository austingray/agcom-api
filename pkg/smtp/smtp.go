package smtp

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type payload struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Htmlbody string `json:"htmlbody"`
}

// SendEmail sends an SMTP email using OhMySMTP to the selected recipient with the selected template
func SendEmail(recipient, templateName string) {
	data := payload{
		From:     "no-reply@austingray.com",
		To:       recipient,
		Subject:  "Welcome to austingray.com",
		Htmlbody: registrationHTML(),
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://app.ohmysmtp.com/api/v1/send", body)
	if err != nil {
		// handle err
	}
	smtpAPIKey := os.Getenv("SMTP_API_KEY")
	log.Println(smtpAPIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ohmysmtp-Server-Token", smtpAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func registrationHTML() string {
	return `
	<h1>austingray.com</h1>
	<p>An account has just been created using this email address. Welcome!!!</p>
	`
}
