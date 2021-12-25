package http

import (
	"fmt"
	"log"
	"errors"
	"net/http"
	"encoding/json"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/users"
	"github.com/htchan/UserService/backend/pkg/tokens"
	"github.com/htchan/UserService/backend/pkg/services"
)

func response(res http.ResponseWriter, data map[string]interface{}) {
	dataByte, err := json.Marshal(data)
	utils.CheckError(err)
	fmt.Fprintln(res, string(dataByte))
}

func writeError(res http.ResponseWriter, err error) {
	errorMap := map[string]int{
		"unauthorized": 401,
	}
	code, ok := errorMap[err.Error()]
	if !ok { code = 400 }
	
	res.WriteHeader(code)
	response(res, map[string]interface{} {
		"code": code,
		"error": err.Error(),
	})
}

func usernameLogin(username, password string) (authToken *tokens.UserToken, err error) {
	user, err := users.Login(username, password)
	if err != nil {
		return nil, err
	}

	token, err := tokens.GenerateUserToken(user, services.UserService(), -1)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func tokenLogin(userTokenStr, serviceName string) (serviceAccessToken *tokens.UserToken, err error) {
	user, err := tokens.FindUserByTokenStr(userTokenStr)
	if err != nil {
		return nil, err
	}

	service, err := services.FindServiceByName(serviceName)
	if err != nil {
		return nil, err
	}

	return tokens.LoadUserToken(user, service, -1)	
}

func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	username := req.Form.Get("username")
	password := req.Form.Get("password")
	if username != "" && password != "" {
		userToken, err := usernameLogin(username, password)
		if err != nil {
			writeError(res, err)
		} else {
			response(res, map[string]interface{} {
				"token": userToken.Token,
			})
		}
		return
	}
	authToken := req.Header.Get("authorization")
	serviceName := req.Form.Get("service")
	if authToken != "" && serviceName != "" {
		service, err := services.FindServiceByName(serviceName)
		if err != nil {
			writeError(res, err)
			return
		}
		accessToken, err := tokenLogin(authToken, serviceName)
		if err != nil {
			writeError(res, err)
		} else {
			response(res, map[string]interface{} {
				"token": accessToken.Token,
				"url": service.Url + "user_service/login",
			})
		}
		return
	}
	writeError(res, errors.New("unauthorized"))
}

func Logout(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	tokenStr := req.Header.Get("authorization")
	token, err := tokens.FindUserTokenByTokenStr(tokenStr)
	if err != nil {
		writeError(res, err)
	}
	token.Expire()
	response(res, map[string]interface{} {
		"result": "success",
	})
}

func StartServer(addr string) {
	http.HandleFunc("/api/users/login", Login)
	http.HandleFunc("/api/users/logout", Logout)
	log.Println("http started")
	log.Fatal(http.ListenAndServe(addr, nil))
}