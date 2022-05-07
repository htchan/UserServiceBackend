package http

import {
	"net/http"
	"github.com/htchan/UserService/pkg/user"
	"github.com/htchan/UserService/pkg/token"
}


func getUserLoginBody(req *http.Request) (username, password string, err error) {
	if err := req.ParseForm(); err != nil {
		return "", "", InvalidParamsError
	}
	return req.Form.Get("username"), req.Form.Get("password"), nil
}

func userLoginHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	// extract params from request
	username, password, err := getUserLoginBody(req)
	if err != nil { writeError(res, http.StatusBadRequest, err); return }

	tkn, _, err := token.UserNameLogin(username, password, service.DefaultUserService().UUID)
	if err != nil { writeError(res, http.StatusBadRequest, err); return }

	json.NewEncoder(res).Encode(map[string]string{
		"token": tkn.Token,
	})
}

func userLogoutHandler(res http.ResponseWriter, req *http.Request) {
	tokenString := req.Header.Get("authorization")
	err := token.UserLogout(tokenString)
	if err != nil {
		writeError(res, http.StatusUnauthorized, nil)
		return
	}
	json.NewEncoder(res).Encode(map[string]interface{} {
		"message": "logout_success",
	})
}

func getServiceLoginBody(req *http.Request) (authTokenStr, serviceUUID string, err error) {
	if err := req.ParseForm(); err != nil {
		return "", "", InvalidParamsError
	}
	return req.Header.Get("authorization"), req.Form.Get("service"), nil
}

func serviceLoginHandler(res http.ResponseWriter, req *http.Request) {
	userServiceTokenString, serviceUUID, err := getServiceLoginBody(req)
	if err != nil { writeError(res, http.StatusBadRequest, err); return }

	tkn, err := token.UserTokenLogin(userServiceTokenString, serviceUUID)
	if err != nil { writeError(res, http.StatusBadRequest, err); return }

	s, _ := tkn.Service()

	http.Redirect(res, req, s.RedirectURL(tkn.Token), http.StatusFound)
}

func UserRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodOptions, "/api/users/login", optionsHandler)
	router.HandlerFunc(http.MethodOptions, "/api/users/logout", optionsHandler)
	router.HandlerFunc(http.MethodOptions, "/api/users/service/login", optionsHandler)
	
	router.HandlerFunc(http.MethodPost, "/api/users/login", userLoginHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/logout", userLogoutHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/service/login", serviceLoginHandler)

}