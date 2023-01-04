package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/WorkLevchenko/willknow/internal/app/apiserver"
)

/*
Регистрация и авторизация. Основная работа проводится в модуле "users"

Sign Up - регистрация
Sign In - авторизация


func getSighInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-in.html", nil)
}

func getSighUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-up.html", nil)
}

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := users.DefaultUserService.VerifyUser(newUser)
	if !ok {
		fileName := "sign-in.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, "New user sign in failure.")
		return
	}
	fileName := "sign-in.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, "New user sign in succesfuly.")
	return
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	if err != nil {
		fileName := "sign-up.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, "New user sign up failure.")
		return
	}
	fileName := "sign-up.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, "New user sign up succesful.")
	return
}

func getUser(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User{
		Email:    email,
		Password: password,
	}
}

//Роутинг и главная функция.

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSighInPage(w, r)
	case "/sign-up-form":
		getSighUpPage(w, r)
	}
}
*/

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig() // Инициализируем конфиг
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config) // Передаём порт в качестве аргумента
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	//http.HandleFunc("/", userHandler)
}
