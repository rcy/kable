// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package api

import (
	"database/sql"
	"time"

	"oj/element/gradient"
)

type Attempt struct {
	ID        int64
	CreatedAt time.Time
	QuizID    int64
	UserID    int64
}

type Bot struct {
	ID          int64
	CreatedAt   time.Time
	OwnerID     int64
	Name        string
	Description string
	AssistantID string
	Published   bool
}

type Code struct {
	ID        int64
	CreatedAt time.Time
	Code      interface{}
	Nonce     interface{}
	Email     string
}

type Delivery struct {
	ID          int64
	CreatedAt   time.Time
	MessageID   int64
	RoomID      int64
	RecipientID int64
	SenderID    int64
	SentAt      sql.NullTime
}

type Friend struct {
	ID        int64
	CreatedAt time.Time
	AID       int64
	BID       int64
	BRole     string
}

type Gradient struct {
	ID        int64
	CreatedAt string
	UserID    int64
	Gradient  gradient.Gradient
}

type Image struct {
	ID        int64
	CreatedAt time.Time
	Url       interface{}
	UserID    interface{}
}

type KidsCode struct {
	ID        int64
	CreatedAt time.Time
	Code      interface{}
	Nonce     interface{}
	UserID    int64
}

type KidsParent struct {
	ID        int64
	CreatedAt time.Time
	KidID     int64
	ParentID  int64
}

type Message struct {
	ID        int64
	CreatedAt time.Time
	SenderID  int64
	RoomID    string
	Body      string
}

type MigrationVersion struct {
	Version sql.NullInt64
}

type Postcard struct {
	ID        int64
	CreatedAt time.Time
	Sender    int64
	Recipient int64
	Subject   string
	Body      string
	State     string
}

type Question struct {
	ID        int64
	CreatedAt string
	QuizID    int64
	Text      string
	Answer    string
}

type Quiz struct {
	ID          int64
	CreatedAt   time.Time
	Name        interface{}
	Description interface{}
	Published   bool
}

type Response struct {
	ID         int64
	CreatedAt  time.Time
	QuizID     interface{}
	UserID     interface{}
	AttemptID  interface{}
	QuestionID interface{}
	Text       interface{}
}

type Room struct {
	ID        int64
	CreatedAt time.Time
	Key       string
}

type RoomUser struct {
	ID        int64
	CreatedAt time.Time
	RoomID    int64
	UserID    int64
}

type Session struct {
	ID     int64
	UserID int64
	Key    interface{}
}

type Thread struct {
	ID          int64
	CreatedAt   time.Time
	ThreadID    string
	AssistantID string
	UserID      int64
}

type User struct {
	ID           int64
	CreatedAt    time.Time
	Username     string
	Email        sql.NullString
	AvatarURL    string
	IsParent     bool
	Bio          string
	BecomeUserID sql.NullInt64
	Admin        bool
}
