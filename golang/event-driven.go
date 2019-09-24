// events/user_created.go
package events

import "time"

var UserCreated userCreated

// UserCreatedPayload is the data for when a user is created
type UserCreatedPayload struct {
    Email string
    Time  time.Time
}

type userCreated struct {
    handlers []interface{ Handle(UserCreatedPayload) }
}

// Register adds an event handler for this event
func (u *userCreated) Register(handler interface{ Handle(UserCreatedPayload) }) {
    u.handlers = append(u.handlers, handler)
}

// Trigger sends out an event with the payload
func (u userCreated) Trigger(payload UserCreatedPayload) {
    for _, handler := range u.handlers {
        go handler.Handle(payload)
    }
}

////////////////////////////////////////////////////////////
// create another event for when a user is deleted:

// events/user_deleted.go
package events

import "time"

var UserDeleted userDeleted

// UserDeletedPayload is the data for when a user is Deleted
type UserDeletedPayload struct {
    Email string
    Time  time.Time
}

type userDeleted struct {
    handlers []interface{ Handle(UserDeletedPayload) }
}

// Register adds an event handler for this event
func (u *userDeleted) Register(handler interface{ Handle(UserDeletedPayload) }) {
    u.handlers = append(u.handlers, handler)
}

// Trigger sends out an event with the payload
func (u userDeleted) Trigger(payload UserDeletedPayload) {
    for _, handler := range u.handlers {
        go handler.Handle(payload)
    }
}

////////////////////////////////////////////////////////////
// Listening for Events

// create_notifier.go
package main

import (
    "time"

    "github.com/stephenafamo/demo/events"
)

func init() {
    createNotifier := userCreatedNotifier{
        adminEmail: "the.boss@example.com",
        slackHook: "https://hooks.slack.com/services/...",
    }

    events.UserCreated.Register(createNotifier)
}

type userCreatedNotifier struct{
    adminEmail string
    slackHook string
}

func (u userCreatedNotifier) notifyAdmin(email string, time time.Time) {
    // send a message to the admin that a user was created
}

func (u userCreatedNotifier) sendToSlack(email string, time time.Time) {
    // send to a slack webhook that a user was created
}

func (u userCreatedNotifier) Handle(payload events.UserCreatedPayload) {
    // Do something with our payload
    u.notifyAdmin(payload.Email, payload.Time)
    u.sendToSlack(payload.Email, payload.Time)
}

////////////////////////////////////////////////////////////
// add another listener that does the same when a user is deleted.

// delete_notifier.go
package main

import (
    "time"

    "github.com/stephenafamo/demo/events"
)

func init() {
    createNotifier := userCreatedNotifier{
        adminEmail: "the.boss@example.com",
        slackHook: "https://hooks.slack.com/services/...",
    }

    events.UserCreated.Register(createNotifier)
}

type userCreatedNotifier struct{
    adminEmail string
    slackHook string
}

func (u userCreatedNotifier) notifyAdmin(email string, time time.Time) {
    // send a message to the admin that a user was created
}

func (u userCreatedNotifier) sendToSlack(email string, time time.Time) {
    // send to a slack webhook that a user was created
}

func (u userCreatedNotifier) Handle(payload events.UserCreatedPayload) {
    // Do something with our payload
    u.notifyAdmin(payload.Email, payload.Time)
    u.sendToSlack(payload.Email, payload.Time)
}

////////////////////////////////////////////////////////////
// Triggering Events

// Now that we have our listeners set up, we can then trigger these events from our
// auth package (or anywhere else).

// auth.go
package auth

import (
    "time"

    "github.com/stephenafamo/demo/events"
    // Other imported packages
)

func CreateUser() {
    // ...
    events.UserCreated.Trigger(events.UserCreatedPayload{
        Email: "new.user@example.com",
        Time: time.Now(),
    })
    // ...
}

func DeleteUser() {
    // ...
    events.UserDeleted.Trigger(events.UserDeletedPayload{
        Email: "deleted.user@example.com",
        Time: time.Now(),
    })
    // ...
}
