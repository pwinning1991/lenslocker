package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/pwinning1991/lenslocker/rand"
)

const (
	// The minimum number of bytes to be used for each session token
	MinBytesPerToken = 32
)

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
	// Determines the how many bytes to yse when generating each session token
	// If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used instead
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	// TODO: Store the session in our DB
	row := ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES ($1, $2)
		RETURNING id;
    `, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &session, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionsService.User
	return nil, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

// TODO: look at implementing the tokenmanger and removing the logic from the
// create function above
//type TokenManager struct{}
//
//func (tm TokenManager) New() (token, tokenHash string, err error) {
//	bytesPerToken := SessionService{}.BytesPerToken
//	if bytesPerToken < MinBytesPerToken {
//		bytesPerToken = MinBytesPerToken
//	}
//	token, err = rand.String(bytesPerToken)
//	if err != nil {
//		return "", "", fmt.Errorf("tokenManager: %w", err)
//	}
//	tokenHash = SessionService.hash(token)
//
//	return token, tokenHash, nil
//}
