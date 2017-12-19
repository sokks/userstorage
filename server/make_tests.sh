#!/bin/bash
PROJECT_PATH=`pwd`
GO_PATH=/mnt/d/KSENIYA/go
DB_FILE=/home/ksenia/server_db.db

export GOPATH=$GO_PATH
cd $GOPATH
go install ./src/github.com/sokks/userstorage/server

# in-memory
./bin/server --port 9081 --inmemory true >> $PROJECT_PATH/log/debug.log &
cd $GOPATH/src/github.com/sokks/userstorage
sleep 5
python3 ./test/userstorage_test.py > $PROJECT_PATH/log/test.log
kill %1

# DB
./bin/server --port 9081 --dbFile $DB_FILE >> $PROJECT_PATH/log/debug.log &
cd $GOPATH/src/github.com/sokks/userstorage
sleep 5
python3 ./test/userstorage_test.py > $PROJECT_PATH/log/testDB.log
kill %1

go clean -i -r github.com/sokks/userstorage/server
