package main

import (
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/grpc"
	"github.com/htchan/UserService/backend/pkg/http"
	"sync"
	"fmt"
)

var wg sync.WaitGroup

const (
	grpcPort = 8000
	httpPort = 8080
)

func startGRPC() {
	grpc.StartServer(fmt.Sprintf("0.0.0.0:%v", grpcPort))
	wg.Done()
}

func startHTTP() {
	http.StartServer(fmt.Sprintf("0.0.0.0:%v", httpPort))
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