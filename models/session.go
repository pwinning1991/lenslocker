package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session.
	// When looking up a session this will be empty
	// as we only store the hash of a session token
	// in our database
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: Create the session token
	// TODO: Implement sessionService.Create
	return nil, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionsService.User
	return nil, nil
}