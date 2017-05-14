package models

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

type Message struct {
	ID             int
	RequestID      string
	Chat           *Chat
	Text           string
	From           *User
	NewChatMember  *User
	LeftChatMember *User
	IsBotLeft      bool
	IsBotAdded     bool
}

func toString(model interface{}) string {
	b, err := json.Marshal(model)
	if err != nil {
		return fmt.Sprintf("cannot represent as json: %s", err)
	}
	return string(b)
}

func (m Message) String() string {
	return toString(&m)
}

func (m *Message) ToCommand() (string, string) {
	text := strings.TrimSpace(m.Text)
	parts := strings.SplitN(text, " ", 2)
	cmd := strings.TrimSpace(parts[0])
	args := ""
	if len(parts) > 1 {
		args = strings.TrimSpace(parts[1])
	}
	return cmd, args
}

type Chat struct {
	ID        int
	IsPrivate bool
	Title     string
}

func (c Chat) String() string {
	return toString(&c)
}

type User struct {
	ID       int
	PMID     int
	Username string
	Name     string
}

func (u User) String() string {
	return toString(&u)
}

type Notification struct {
	ID        string
	RequestID string
	Text      string
	ReadyAt   time.Time
	User      *User
	MessageID int
	ChatID    int
}

func NewNotification(user *User, msgID, chatID, readyDelay int, text, requestID string) *Notification {
	return &Notification{
		ID:        uuid.NewV4().String(),
		RequestID: requestID,
		Text:      text,
		ReadyAt:   time.Now().Add(time.Second * time.Duration(readyDelay)),
		User:      user,
		MessageID: msgID,
		ChatID:    chatID,
	}
}

func (n Notification) String() string {
	return toString(&n)
}
