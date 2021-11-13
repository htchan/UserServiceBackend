package http

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/htchan/UserService/internal/utils"
	"github.com/htchan/UserService/pkg/users"
	"github.com/htchan/UserService/pkg/tokens"
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
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	username := req.FormValue("username")
	password := req.FormValue("password")
	user, err := users.Login(username, password)
	if err != nil {
		writeError(res, 401, err.Error())
	}
	token, err := tokens.LoadUserToken(*user, 24*60)
	if err != nil {
		writeError(res, 400, err.Error())
	}
	response(res, map[string]interface{} {
		"token": token.Token,
	})
}

func logout(res http.ResponseWriter, req *http.Request) {
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
	http.HandleFunc("api/users/login", login)
	http.HandleFunc("api/users/logout", logout)
	log.Println("http started")
	log.Fatal(http.ListenAndServe(addr, nil))
}