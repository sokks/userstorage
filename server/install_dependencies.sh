#!/bin/bash
PROJECT_PATH=`pwd`
echo $PROJECT_PATH
GO_PATH=/mnt/d/KSENIYA/go
cd $HOME
# db settings
sqlite3 server_db.db < $PROJECT_PATH/db/create_db.sql
# packages
export GOPATH=$GO_PATH
cd $GOPATH
go get github.com/satori/go.uuid
go get github.com/gorilla/rpc
go get github.com/mattn/go-sqlite3
go get github.com/sokks/userstorage
#go get github.com/sokks/userstorage/server
cd $PROJECT_PATH
pip3 install requests
