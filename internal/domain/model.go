package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	LoginPageURL = "https://nanext.alcnanext.jp/anetn/student/stlogin/index/nit-ariake/"
	TopPageURL   = "https://nanext.alcnanext.jp/anetn/Student/StTop"
)

var (
	ErrInvalidID       = errors.New("ID は s + 5 桁の形式でなければならない")
	ErrInvalidPassword = errors.New("パスワードが空である")

	reOnlyDigits = regexp.MustCompile(`^\d{5}$`)
	reStudentID  = regexp.MustCompile(`^s\d{5}$`)
)

type Credentials struct {
	ID       string
	Password string
}

func NewCredentials(rawID string, password string) (Credentials, error) {
	normalizedID := NormalizeID(rawID)
	if !reStudentID.MatchString(normalizedID) {
		return Credentials{}, fmt.Errorf("%w: %s", ErrInvalidID, rawID)
	}
	if strings.TrimSpace(password) == "" {
		return Credentials{}, ErrInvalidPassword
	}

	return Credentials{
		ID:       normalizedID,
		Password: password,
	}, nil
}

func NormalizeID(rawID string) string {
	id := strings.TrimSpace(rawID)
	if reOnlyDigits.MatchString(id) {
		return "s" + id
	}
	return id
}
