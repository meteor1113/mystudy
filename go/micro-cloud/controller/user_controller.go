package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"micro-cloud/framework"
	"micro-cloud/service"
	"micro-cloud/utils"
	"net/http"
)

type UserConterller struct {
}

var userService = new(service.UserService)

func (p *UserConterller) Router(router *framework.RouterHandler) {
	router.Router("/register", p.register)
	router.Router("/login", p.login)
	router.Router("/findAll", p.findAll)
}

// POST Content-Type=application/x-www-form-urlencoded
func (p *UserConterller) register(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// // log.Println(string(body))
	// var data map[string]interface{}
	// err = json.Unmarshal(body, &data)
	// if err != nil {
	// 	log.Fatalf("Error parsing JSON: %v", err)
	// }

	// username := data["username"].(string) // r.PostFormValue("username")
	// password := data["password"].(string) // r.PostFormValue("password")
	// fmt.Println("username:", username)
	// fmt.Println("password:", password)
	username, password := p.parseUserPass(r)

	if utils.Empty(username) || utils.Empty(password) {
		framework.ResultFail(w, "username or password can not be empty")
		return
	}
	// id := userService.Insert(username, password)
	// if id <= 0 {
	// 	framework.ResultFail(w, "register fail")
	// 	return
	// }
	framework.ResultOk(w, "register success")
}

// POST Content-Type=application/x-www-form-urlencoded
func (p *UserConterller) login(w http.ResponseWriter, r *http.Request) {
	username, password := p.parseUserPass(r)

	if utils.Empty(username) || utils.Empty(password) {
		framework.ResultFail(w, "username or password can not be empty")
		return
	}
	// users := userService.SelectUserByName(username)
	// if len(users) == 0 {
	// 	framework.ResultFail(w, "user does not exist")
	// 	return
	// }
	// if users[0].Password != password {
	// 	framework.ResultFail(w, "password error")
	// 	return
	// }

	framework.ResultOk(w, "login success")
}

// GET/POST
func (p *UserConterller) findAll(w http.ResponseWriter, r *http.Request) {
	users := userService.SelectAllUser()
	framework.ResultJsonOk(w, users)
}

func (p *UserConterller) parseUserPass(r *http.Request) (username, password string) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	username = data["username"].(string) // r.PostFormValue("username")
	password = data["password"].(string) // r.PostFormValue("password")
	fmt.Println("username:", username)
	fmt.Println("password:", password)
	return
}
