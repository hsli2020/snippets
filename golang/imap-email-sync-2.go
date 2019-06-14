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
		conn, err := ConnectServer(account.Server, account.Username, account.Password)
		if err != nil {
			continue
		}
		mails, err := FetchMails(conn, 10)
		if err == nil {
			SaveMails(db, conn, mails)
		}
		conn.Logout()
	}
}

func ConnectServer(server, username, password string) (*client.Client, error) {
	c, err := client.DialTLS(server, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Login(username, password); err != nil {
		if err2 := c.Logout(); err2 != nil {
			return nil, fmt.Errorf("Error while logging in to %v: %v\n(followup error: %v)", server, err, err2)
		}
		return nil, fmt.Errorf("Error while logging in to %v: %v", server, err)
	}

	log.Printf("Connected to %v as user %v.", server, username)

	return c, nil
}

func FetchMails(conn *client.Client, num uint32) (map[uint32]string, error) {
	// Select INBOX
	mbox, err := conn.Select("INBOX", false)
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
		done <- conn.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, imap.FetchUid}, messages)
	}()

	emails := make(map[uint32]string)

	for msg := range messages {
		if msg.Envelope == nil {
			continue
		}
		emails[msg.SeqNum] = msg.Envelope.MessageId[1 : len(msg.Envelope.MessageId)-1]
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return emails, nil
}

func FetchMessage(conn *client.Client, seqNo uint32) (*EmailMessage, error) {
	messages := make(chan *imap.Message)
	done := make(chan error, 1)

	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNo)

	go func() {
		var section imap.BodySectionName
		done <- conn.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages)
	}()

	var email *EmailMessage
	var err error

	for msg := range messages {
		if msg.Envelope == nil {
			continue
		}

		// Print email subject
		log.Printf("[%d] %s\n", msg.SeqNum, msg.Envelope.Subject)

		email, err = DecodeEmail(msg)
		if err != nil {
			log.Printf("? %s\n", err)
			continue
		}
	}

	if err = <-done; err != nil {
		return nil, err
	}

	return email, nil
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

func SaveMails(db *sql.DB, conn *client.Client, emails map[uint32]string) error {
	stmt, err := db.Prepare("INSERT INTO emails SET `msgno`=?,`date`=?,`from`=?,`from_address`=?,`subject`=?,`body`=?,`message_id`=?")
	if err != nil {
		return err
	}

	for seqNo, messageId := range emails {
		if ExistMail(db, messageId) {
			continue
		}

		email, err := FetchMessage(conn, seqNo)
		if err != nil {
			continue
		}

		_, err = stmt.Exec(email.MsgNo, email.Date, email.From, email.FromAddress, email.Subject, email.Body, email.MessageId)
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
