package userstorage

import (
	"fmt"
	"net/http"
	"time"
)


type User struct {
	uuid  string
	login string
	regDate time.Time
}

func (u User) String() string {
	return fmt.Sprintf("{'uuid':%s, 'login':%s, 'regDate':%s }", u.uuid, u.login, u.regDate.Format("2006-01-02"))
}

type CommonStorage interface {
	AddUser(login string) error
	FindUser(login string) (User, error)
	RenameUser(oldLogin, newLogin string) error
}

type Storage struct {
	// storage TmpStorage
	storage CommonStorage
}

func (s *Storage) Add(r *http.Request, args *AddIn, reply *AddOut) error {
	err := s.storage.AddUser(args.Login)
	if err != nil {
		reply.Result = "error"
		return err
	}
	reply.Result = "ok"
	return nil
}

func (s *Storage) Get(r *http.Request, args *GetIn, reply *GetOut) error {
	user, err := s.storage.FindUser(args.Login)
	if err != nil {
		return err
	}
	reply.Login   = user.login
	reply.UUID    = user.uuid
	reply.RegDate = user.regDate.Format("2006-01-02")
	return nil
}

func (s *Storage) Rename(r *http.Request, args *RenameIn, reply *RenameOut) error {
	err := s.storage.RenameUser(args.OldLogin, args.NewLogin)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}