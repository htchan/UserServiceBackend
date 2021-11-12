package main

import (
	"github.com/htchan/UserService/internal/utils"
	"github.com/htchan/UserService/pkg/grpc"
	"github.com/htchan/UserService/pkg/http"
	"sync"
)

var wg sync.WaitGroup

const (
	port = 8000
)

func startGRPC() {
	grpc.StartServer(port)
	wg.Done()
}

func startHTTP() {
	http.StartServer()
	wg.Done()

}

func main() {
	utils.OpenDB("./bin/database.db")
	defer utils.CloseDB()
	wg.Add(1)
	go startGRPC()
	go startHTTP()
	wg.Wait()
}