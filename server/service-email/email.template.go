package main

import "fmt"

func getTemplate(template string, html []string) (subject string, body string) {
	if template == "CONTACT" {
		subject = "Contact"
		body = fmt.Sprintf(
			`
            <div>You received a contact from %v.</div>
            <div>%v</div>
            <div>Thanks, Your Team</div>
            `, html[0], html[1])
		return subject, body
	}
	return "", ""
}
