package controller

import (
	"antrianWithoutDatabase/config"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type UserInput struct {
	name    string `validate:"required"`
	credits string `validate:"required"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("home")

	if r.Method == "GET" {
		temp, _ := template.ParseFiles("view/home.html")
		log.Println("showing home page")
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		r.ParseForm()

		log.Println("post")
		userInput := &UserInput{
			name:    r.Form.Get("name"),
			credits: r.Form.Get("credits"),
		}
		var message error

		if userInput.name != "" && userInput.credits != "" {
			var nameExists = GetName(userInput.name)

			if nameExists == false { //username available
				log.Println("Username available")
				creditsInt, _ := strconv.Atoi(userInput.credits)

				session, _ := config.Store.Get(r, config.SESSION_ID)

				session.Values["loggedIn"] = true
				session.Values["name"] = userInput.name
				session.Values["credits"] = creditsInt

				session.Save(r, w)

				InsertData(userInput.name, creditsInt)
				http.Redirect(w, r, "/queue", http.StatusSeeOther)

			} else {
				message = errors.New("Username sudah terpakai :(")

				tempData := map[string]interface{}{
					"error": message,
				}

				temp, err := template.ParseFiles("views/home.html")
				if err != nil {
					panic(err.Error())
				} else {
					temp.Execute(w, tempData)
				}
			}
		}

	}

}
