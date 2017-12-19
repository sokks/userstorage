package userstorage

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"time"

	//"github.com/mattn/go-sqlite3"
)

type SqliteStorage struct{
	database *sql.DB
}

func NewSqliteStorage(sqliteFilepath string) *SqliteStorage {
	database, err := sql.Open("sqlite3", sqliteFilepath)
	if err != nil {
		logger.Fatalln("Can't connect SQLite database")
	}
	return &SqliteStorage{ database }
}

func (s *SqliteStorage) AddUser(login string) (error) {
	if login == "" {
		return errors.New("Incorrect login: empty")
	}
	if len(login) > 50 {
		return errors.New("Incorrect login: too long")
	}
	userUUID, err := newUUID()
	if err != nil {
		return err
	}
	userTime := time.Now()
	rows, _ := s.database.Query("SELECT login FROM 'users' WHERE login = ?", login)
	defer rows.Close()
	if rows.Next() {
		return errors.New("User already exists")
    }
	statement, _ := s.database.Prepare("INSERT INTO 'users' (uuid, login, registration_date) VALUES (?, ?, ?)")
    _, err = statement.Exec(userUUID, login, userTime)
	if err != nil {
		return err
	}

	logger.Printf("User %s added\n", login)
	return nil
}

func (s *SqliteStorage) FindUser(login string) (User, error) {
	rows, _ := s.database.Query("SELECT uuid, login, registration_date FROM 'users' WHERE login = ?", login)
	defer rows.Close()
	var user User
	if rows.Next() {
		uuid := make([]byte, 100)
		var login string
		var date time.Time
		rows.Scan(&uuid, &login, &date)
		user = User{string(uuid), login, date}
	} else {
		logger.Printf("User %s not found\n", login)
		return User{}, errors.New("User does not exist")
	}
	
	if rows.Next() { // not supposed to be generated
		logger.Printf("Multiple users %s found\n", login)		
		return User{}, errors.New("More than one user found")
	}

	return user, nil
}

func (s *SqliteStorage) RenameUser(oldLogin, newLogin string) error {
	if newLogin == "" {
		return errors.New("Incorrect login: empty")
	}
	if len(newLogin) > 50 {
		return errors.New("Incorrect login: too long")
	}
	
	rows, _ := s.database.Query("SELECT login FROM 'users' WHERE login = ?", newLogin)
	defer rows.Close()
	if rows.Next() {
		return errors.New("User already exists")
    }

	statement, _ := s.database.Prepare("UPDATE 'users' SET login = ? WHERE login = ?")
    res, err := statement.Exec(newLogin, oldLogin)
	if err != nil {
		return err
	}

	n, _ := res.RowsAffected()
	if n < 1 {
		logger.Printf("User %s not found\n", oldLogin)
		return errors.New("user not found")
	}

	return nil
}

func newUUID() ([]byte, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return make([]byte, 0), err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return []byte(fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])), nil
}