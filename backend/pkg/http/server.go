package http

import (
	"fmt"
	"log"
	"errors"
	"net/http"
	"encoding/json"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/user"
	"github.com/htchan/UserService/backend/pkg/token"
	"github.com/htchan/UserService/backend/pkg/service"
)

func StartServer(addr string) {
	router := httprouter.New()
	token.Route(router)
	
	log.Fatal(http.ListenAndServe(addr, router))
}