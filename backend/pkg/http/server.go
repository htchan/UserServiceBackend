package http

import (
	"log"
	"errors"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

var InvalidParamsError = errors.New("invalid_params")

func writeError(res http.ResponseWriter, statusCode int, err error) {
	res.WriteHeader(statusCode)
	messages := map[error]string{
	}
	message, ok := messages[err]
	if !ok { message = "" }
	json.NewEncoder(res).Encode(map[string]string{ "error": message })
}

func optionsHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "*")
	res.WriteHeader(http.StatusOK)
	return
}

func StartServer(addr string) {
	router := httprouter.New()
	UserRoutes(router)
	
	log.Fatal(http.ListenAndServe(addr, router))
}