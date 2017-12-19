# UserStorage
UserStorage is a simple json-RPC server for users storage. Golang implementstion based on in-memory storage or SQLite. 

## Installation
To install package:
```console
$ go get github.com/sokks/userstorage
$ cd `go env GOPATH`/src/github.com/sokks/userstorage
$ make prepare
```
To run some tests (Python3):
```console
$ make tests
```

To run and stop server:
```console
$ make run
$ make stop
```

## API
UserStorage:  
**Add**(login *string*) (*Result*, *error*)  
**Get**(login *string*) (*User*, *error*)  
**Rename**(oldLogin *string*, newLogin *string*) (*Result*, *error*)  



## Further work
- SQLite full support
- Key-Value DB support
- Docker containers
