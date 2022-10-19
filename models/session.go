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
	row := ss.DB.QueryRow(`
		UPDATE sessions 
		SET token_hash = $2
		WHERE user_id = $1
		RETURNING id;
    	`, userID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err == sql.ErrNoRows {
		row = ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES ($1, $2)
		RETURNING id;
    `, session.UserID, session.TokenHash)
		err = row.Scan(&session.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &session, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionsService.User
	// 1. Hash the session tokens
	tokenHash := ss.hash(token)
	// 2. query for the session with that hash
	var user User
	row := ss.DB.QueryRow(`
		SELECT user_id FROM sessions 
		WHERE token_hash = '$1';`, tokenHash)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	//3. Using the UserID from the session, we need to query for that user
	row = ss.DB.QueryRow(`
		SELECT email, password_hash from users 
		WHERE id = $1`, user.ID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	//4. Return the User
	return &user, nil
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
