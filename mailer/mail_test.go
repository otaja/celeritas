package mailer

import "testing"

func TestMailer_SendSMTPMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Jan",
		To:          "you@here.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	err := Mailer.SendSMTPMessage(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestMailer_SendUsingAChannel(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Jan",
		To:          "you@here.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	Mailer.Jobs <- msg
	res := <-Mailer.Results

	if res.Error != nil {
		t.Error("failed to send over channel")
	}

	msg.To = "not_an_email_adress"
	Mailer.Jobs <- msg
	res = <-Mailer.Results
	if res.Error == nil {
		t.Error("no error recived but there should one in email adress")
	}

}

func TestMailer_SendUsingApi(t *testing.T) {
	msg := Message{
		To:          "you@here.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	Mailer.API = "unknown"
	Mailer.APIKey = "ABCSGG"
	Mailer.APIUrl = "fake.com"

	err := Mailer.SendUsingApi(msg, "unknown")
	if err == nil {
		t.Error(err)
	}

	Mailer.API = ""
	Mailer.APIKey = ""
	Mailer.APIUrl = ""
}

func TestMailer_BuildHtmlMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Jan",
		To:          "you@here.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	_, err := Mailer.buildHTMLMessage(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestMailer_BuildPlainlMessage(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Jan",
		To:          "you@here.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	_, err := Mailer.buildPlainTextMessage(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestMailer_Send(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Jan",
		To:          "you@here.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}

	err := Mailer.Send(msg)
	if err != nil {
		t.Error(err)
	}

	Mailer.API = "unknown"
	Mailer.APIKey = "ABCSGG"
	Mailer.APIUrl = "fake.com"

	err = Mailer.Send(msg)
	if err == nil {
		t.Error(err)
	}

	Mailer.API = ""
	Mailer.APIKey = ""
	Mailer.APIUrl = ""

}

func TestMailer_ChooseApi(t *testing.T) {
	msg := Message{
		From:        "me@here.com",
		FromName:    "Jan",
		To:          "you@here.com",
		Subject:     "test",
		Template:    "test",
		Attachments: []string{"./testdata/mail/test.html.tmpl"},
	}
	Mailer.API = "unknown"
	err := Mailer.ChooseAPI(msg)
	if err == nil {
		t.Error(err)
	}
}
