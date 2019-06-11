package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type pageVariables struct {
	Message string
}

// HomePage renders index.html template with variables
func HomePage(w http.ResponseWriter, r *http.Request) {
	log.SetOutput(os.Stdout)
	now := time.Now()
	msg := ""
	weekDay := now.Weekday()
	if checkFriday(weekDay) {
		msg = "Today is God bless Friday"
	} else {
		msg = "Today is " + now.Weekday().String()
	}
	HomePageVars := pageVariables{
		Message: msg,
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Print("failed to parse template: ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("failed to execute template: ", err)
	}
}

func checkFriday(weekDay time.Weekday) bool {
	if weekDay == 5 {
		return true
	}
	return false
}
