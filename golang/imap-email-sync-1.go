package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type EmailAccount struct {
	Server   string
	Username string
	Password string
}

var accounts = []EmailAccount{}

type EmailMessage struct {
	MsgNo       uint32
	Date        string
	From        string
	FromAddress string
	Subject     string
	Body        string
	MessageId   string
}

func main() {
	db := openDatabase()
	defer db.Close()

	for _, account := range accounts {
		emails, err := FetchMails(account, 10)
		if err == nil {
			SaveMails(db, emails)
		}
	}
}

func SaveMails(db *sql.DB, emails []*EmailMessage) error {
	stmt, err := db.Prepare("INSERT INTO emails SET `msgno`=?,`date`=?,`from`=?,`from_address`=?,`subject`=?,`body`=?,`message_id`=?")
	if err != nil {
		return err
	}

	for _, email := range emails {
		if ExistMail(db, email.MessageId) {
			continue
		}

		_, err := stmt.Exec(email.MsgNo, email.Date, email.From, email.FromAddress, email.Subject, email.Body, email.MessageId)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func ExistMail(db *sql.DB, msgid string) bool {
	var exists bool

	row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM emails WHERE message_id='" + msgid + "')")
	if err := row.Scan(&exists); err != nil {
		log.Println(err)
	}

	return exists
}

func FetchMails(account EmailAccount, num uint32) ([]*EmailMessage, error) {
	// Connect to server
	c, err := client.DialTLS(account.Server, nil)
	if err != nil {
		return nil, err
	}

	// Don't forget to logout
	defer c.Logout()

	// Login
	err = c.Login(account.Username, account.Password)
	if err != nil {
		return nil, err
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		return nil, err
	}

	// Get the last NUM messages
	from := uint32(1)
	to := mbox.Messages
	done := make(chan error, 1)
	if mbox.Messages > num {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - num - 1
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, num)
	go func() {
		var section imap.BodySectionName
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages)
	}()

	emails := make([]*EmailMessage, 0)

	for msg := range messages {
		if msg.Envelope == nil {
			continue
		}

		// Print email subject
		log.Println(msg.Envelope.Subject)

		email, err := DecodeEmail(msg)
		if err != nil {
			log.Println(err)
			continue
		}

		emails = append(emails, email)
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return emails, nil
}

func DecodeEmail(msg *imap.Message) (*EmailMessage, error) {
	email := new(EmailMessage)
	email.MsgNo = msg.SeqNum
	email.Date = msg.Envelope.Date.Format("2006-01-02T15:04:05")
	email.From = msg.Envelope.From[0].PersonalName
	email.FromAddress = msg.Envelope.From[0].MailboxName + "@" + msg.Envelope.From[0].HostName
	email.Subject = msg.Envelope.Subject
	email.MessageId = msg.Envelope.MessageId[1 : len(msg.Envelope.MessageId)-1]

	// Get email body
	sectionName, err := imap.ParseBodySectionName(imap.FetchItem("BODY[]"))
	if err != nil {
		return nil, err
	}
	r := msg.GetBody(sectionName)
	if r == nil {
		return nil, err
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		return nil, err
	}

	// header := mr.Header

	// Process each message's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			body, _ := ioutil.ReadAll(p.Body)
			email.Body = string(body)
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			log.Printf("# Attachment: %v\n", filename)
		}
	}

	return email, nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		//panic(e)
	}
}

func openDatabase() *sql.DB {
	host := "localhost"
	port := "3306"
	username := "root"
	//password := ""
	name := "test"

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, host, port, name)

	db, err := sql.Open("mysql", dsn)
	check(err)

	err = db.Ping()
	check(err)

	return db
}
