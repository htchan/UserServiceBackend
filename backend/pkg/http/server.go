package http

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/users"
	"github.com/htchan/UserService/backend/pkg/tokens"
)


func response(res http.ResponseWriter, data map[string]interface{}) {
	dataByte, err := json.Marshal(data)
	utils.CheckError(err)
	fmt.Fprintln(res, string(dataByte))
}

func writeError(res http.ResponseWriter, code int, msg string) {
	res.WriteHeader(code)
	response(res, map[string]interface{} {
		"code": code,
		"message": msg,
	})
}

func login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	username := req.Form.Get("username")
	password := req.Form.Get("password")
	serviceToken := req.Form.Get("serviceToken")
	fmt.Println("hi", req.Form, username, " ", password)
	user, err := users.Login(username, password)
	if err != nil {
		writeError(res, 401, err.Error())
		return
	}
	service, err := services.FindServiceByTokenStr(serviceToken)
	if err != nil {
		writeError(res, 401, err.Error())
		return
	}
	token, err := tokens.LoadUserToken(user, 24*60)
	if err != nil {
		writeError(res, 400, err.Error())
	}
	response(res, map[string]interface{} {
		"token": token.Token,
		"url": token.Url + "user_service/login",
	})
}

func logout(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	tokenStr := req.Header.Get("Authorization")
	token, err := tokens.FindUserTokenByTokenStr(tokenStr)
	if err != nil {
		writeError(res, 404, err.Error())
	}
	token.Expire()
	response(res, map[string]interface{} {
		"result": "success",
	})
}

func StartServer(addr string) {
	http.HandleFunc("/api/users/login", login)
	http.HandleFunc("/api/users/logout", logout)
	log.Println("http started")
	log.Fatal(http.ListenAndServe(addr, nil))
}