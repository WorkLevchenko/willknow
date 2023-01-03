package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/WorkLevchenko/willknow/internal/users"
)

/*
Регистрация и авторизация. Основная работа проводится в модуле "users"

Sign Up - регистрация
Sign In - авторизация
*/

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

func main() {
	http.HandleFunc("/", userHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
