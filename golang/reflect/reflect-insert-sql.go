package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Message struct {
	ID         uint64    `db:"id"`
	Channel    string    `db:"channel"`
	UserName   string    `db:"user_name"`
	UserID     string    `db:"user_id"`
	UserAvatar string    `db:"user_avatar"`
	Message    string    `db:"message"`
	RawMessage string    `db:"message_raw"`
	MessageID  string    `db:"message_id"`
	Stamp      time.Time `db:"stamp"`
}

func insert(table string, data interface{}) string {
	message_value := reflect.ValueOf(data)
	if message_value.Kind() == reflect.Ptr {
		message_value = message_value.Elem()
	}

	message_fields := make([]string, message_value.NumField())

	for i := 0; i < len(message_fields); i++ {
		fieldType := message_value.Type().Field(i)
		message_fields[i] = fieldType.Tag.Get("db")
	}
	sql := "insert into " + table + " set"
	for _, tagFull := range message_fields {
		if tagFull != "" && tagFull != "-" {
			tag := strings.Split(tagFull, ",")
			sql = sql + " " + tag[0] + "=:" + tag[0] + ","
		}
	}
	return sql[:len(sql)-1]
}

func main() {
	message := &Message{
		ID:       1,
		Channel:  "#common",
		UserName: "titpetric",
		Stamp:    time.Now(),
	}

	fmt.Println(insert("messages", message))
}

// https://play.golang.org/p/KcuTIWa3S1F
// insert into messages set id=:id, channel=:channel, user_name=:user_name, 
//    user_id=:user_id, user_avatar=:user_avatar, message=:message, 
//    message_raw=:message_raw, message_id=:message_id, stamp=:stamp
