PROJECT_PATH=`pwd`
GO_PATH=/mnt/d/KSENIYA/go
DB_FILE=/home/ksenia/server_db.db
default: prepare tests run
prepare:
	chmod u+x ./server/install_dependencies.sh
	sudo ./server/install_dependencies.sh
tests:
	./server/make_tests.sh
run:
	cd $(GO_PATH)
	go install ./src/github.com/sokks/userstorage/server
	./bin/server --port 9080 --dbFile $(DB_FILE) >> ./log/server.log &
stop:
	kill %1
