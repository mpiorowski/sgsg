package main

import "fmt"

// TODO - email template as proto enum
func getTemplate(template string, html []string) (subject string, body string) {
	if template == "WELCOME" {
		subject = "Welcome to our site!"
		body = fmt.Sprintf(`
            <html>
                <body>
                    <h1>Welcome to our site!</h1>
                    <p>Thanks for signing up!</p>
                </body>
            </html>
        `)
	} else if template == "CONTACT" {
		subject = "Contact"
		body = fmt.Sprintf(
			`
            <div>You received a contact from %v.</div>
            <div>%v</div>
            <div>Thanks, Your Team</div>
            `, html[0], html[1])
		return subject, body
	} else if template == "MESSAGE" {
		subject = "Message"
		body = fmt.Sprintf(
			`
            <div>Your test message.</div>
            <div>%v</div>
            `, html[0])
		return subject, body
	}
	return "", ""
}
