package mailer

import "testing"

func TestMail_SendSMTPMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	err := mailer.SendSMTPMessage(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestMail_SendUsingChan(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Joe",
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	mailer.Jobs <- msg
	res := <-mailer.Results
	if res.Error != nil {
		t.Error("failed to send over channel")
	}

	msg.To = "not_an_email_address"
	mailer.Jobs <- msg
	res = <-mailer.Results
	if res.Error == nil {
		t.Error("expected an error but got none")
	}
}

func TestMail_SendUsingAPI(t *testing.T) {
	msg := Message{
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	mailer.API = "unknown"
	mailer.APIKey = "greatkey"
	mailer.APIUrl = "https://www.fake.com"

	err := mailer.SendUsingAPI(msg, "unknown")
	if err == nil {
		t.Error("expected an error, got none")
	}
	mailer.API = ""
	mailer.APIKey = ""
	mailer.APIUrl = ""
}

func TestMail_buildHTMLMessage(t *testing.T) {
	msg := Message{
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	_, err := mailer.buildHTMLMessage(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestMail_buildPlainMessage(t *testing.T) {
	msg := Message{
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	_, err := mailer.buildPlainTextMessage(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestMail_Send(t *testing.T) {
	msg := Message{
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	err := mailer.Send(msg)
	if err != nil {
		t.Error(err)
	}

	mailer.API = "unknown"
	mailer.APIKey = "greatkey"
	mailer.APIUrl = "https://www.fake.com"

	err = mailer.Send(msg)
	if err == nil {
		t.Error("expected an error but got none")
	}
	mailer.API = ""
	mailer.APIKey = ""
	mailer.APIUrl = ""
}

func TestMail_ChooseAPI(t *testing.T) {
	msg := Message{
		To:          "you@there.com",
		Subject:     "Test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}
	mailer.API = "unknown"
	err := mailer.ChooseAPI(msg)
	if err == nil {
		t.Error("exptected an error but got none")
	}
}
