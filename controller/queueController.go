package controller

import (
	"antrianWithoutDatabase/config"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Queue struct {
	Position int
	Name     string
	Credits  int
}

var mainQueue []Queue
var currentQueue []Queue
var counter int

func QueueHome(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	username := session.Values["name"]

	data := map[string]any{
		"username":     username,
		"mainQueue":    mainQueue,
		"currentQueue": currentQueue,
	}

	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/queue.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, data)
		}
	}
	// if r.Form.Get("action") == "normal" {
	// 	log.Println("Main normal")
	// 	moveToLast(1)
	// }

	if r.Method == "POST" {
		log.Println("post")
		r.ParseForm()
		action := r.Form.Get("action")
		credits := r.Form.Get("credits")
		creditsInt, _ := strconv.Atoi(credits)

		log.Println(action)

		if username == currentQueue[0].Name && creditsInt > 0 && creditsInt < 3 {
			moveToLast(creditsInt, username)

			// if action == "premium" {
			// 	log.Println("Main premium")
			// 	moveToLast(2, username)
			// }

			// if action == "normal" {
			// 	log.Println("Main normal")
			// 	moveToLast(1, username)
			// }
		}

		temp, err := template.ParseFiles("view/queue.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, data)
		}
	}

}

func GetName(name string) bool {
	nameExists := false
	indexCount := len(mainQueue)

	for i := 0; i < indexCount; i++ {
		if mainQueue[i].Name == name {
			nameExists = true
			break
		}
	}

	return nameExists
}

func InsertData(name string, creditsInt int) {
	position := len(mainQueue)

	mainQueue = append(mainQueue, Queue{position + 1, name, creditsInt})
	currentQueue = append(currentQueue, Queue{position + 1, name, creditsInt})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	name := session.Values["name"]
	log.Println("Session name: ", name)
	indexCount := len(mainQueue)

	for i := 0; i < indexCount; i++ {
		if mainQueue[i].Name == name {
			mainQueue = append(mainQueue[:i], mainQueue[i+1:]...)
			for j := i; j < indexCount-1; j++ {
				mainQueue[j].Position -= 1
			}
			break
		}
	}

	indexCount = len(currentQueue)

	for i := 0; i < indexCount; i++ {
		if currentQueue[i].Name == name {
			currentQueue = append(currentQueue[:i], currentQueue[i+1:]...)
			for j := i; j < indexCount-1; j++ {
				currentQueue[j].Position -= 1
			}
			break
		}
	}

	log.Println(mainQueue)

	//delete session
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func moveToLast(creditsReduce int, name interface{}) {
	counter = len(currentQueue)
	currentQueue[0].Credits -= creditsReduce
	topElement := currentQueue[0]
	currentQueue = currentQueue[1:] //discarding element(i think)

	currentQueue = append(currentQueue, topElement)

	//fixing positions
	for i := 0; i < counter; i++ {
		currentQueue[i].Position = i + 1
	}

	counter = len(mainQueue)
	mainQueue[0].Credits -= creditsReduce
	topElement = mainQueue[0]
	mainQueue = mainQueue[1:] //discarding element(i think)

	mainQueue = append(mainQueue, topElement)

	//fixing positions
	for i := 0; i < counter; i++ {
		mainQueue[i].Position = i + 1
	}
}
