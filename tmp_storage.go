package userstorage

import (
	"fmt"
	"io"
	"crypto/rand"
	"time"
	"errors"
)

// newUUID generates a random UUID according to RFC 4122
func newUUIDString() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}


type TmpStorage struct {
	users []User
}

func (s *TmpStorage) String() string {
	res := "[ "
	for _, elem := range s.users {
		res += elem.String() + ";"
	}
	return res + " ]"
}

func NewTmpStorage() (*TmpStorage) {
	return &TmpStorage{ make([]User, 0, 10) }
}

func (s *TmpStorage) AddUser(login string) error { // if already exists!! --> sqlite
	userUUID, err := newUUIDString()
	if err != nil {
		return err
	}
	userTime := time.Now()
	s.users = append(s.users, User{userUUID, login, userTime})
	logger.Printf("User added; Users: %s\n", s.users)
	return nil
}

func (s *TmpStorage) FindUser(login string) (User, error) {
	for _, user := range s.users {
		if user.login == login {
			logger.Printf("User %s found\n", login)
			return user, nil
		}
	}
	logger.Printf("User %s not found\n", login)
	return User{}, errors.New("User does not exist")
}

func (s *TmpStorage) RenameUser(oldLogin, newLogin string) error {
	for i, user := range s.users {
		if user.login == oldLogin {
			logger.Printf("User %s renamed to %s\n", oldLogin, newLogin)
			s.users[i].login = newLogin
			return nil
		}
	}
	logger.Printf("User %s not found\n", oldLogin)
	return errors.New("user not found")
}