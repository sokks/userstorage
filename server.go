package userstorage

import (
	"log"
	"net"
	"net/http"
	"os"
	
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

var (
	logger *log.Logger
)

func initLogger() (*log.Logger) {
	return log.New(os.Stdout, "userstorage:", log.Ltime)
}


func StartServer(inmemory bool, port, dbFile string) {
	logger = initLogger()
	
	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer listener.Close()

	var storage *Storage
	if inmemory {
		storage = &Storage{ NewTmpStorage() }
	} else {
		storage = &Storage{ NewSqliteStorage(dbFile) }
	}

	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	err = server.RegisterService(storage, "UserStorage")
	if err != nil {
		log.Fatal("service registration error:", err)
	}
	http.Handle("/userstorage", server)
	http.Serve(listener, nil)
}
