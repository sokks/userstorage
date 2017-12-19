package userstorage

import "net/http"

type StorageService interface {
	Add(r *http.Request, args *AddIn, reply *AddOut) error
	Get(r *http.Request, args *GetIn, reply *GetOut) error
	Rename(r *http.Request, args *RenameIn, reply *RenameOut) error
}

type AddIn struct {
	Login string
}

type AddOut struct {
	Result string
}

type GetIn struct {
	Login string
}

type GetOut struct {
	Login   string
	UUID    string
	RegDate string
}

type RenameIn struct {
	OldLogin string
	NewLogin string
}

type RenameOut struct {
	Result string
}