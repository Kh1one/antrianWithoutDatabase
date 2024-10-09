package main

import (
	"antrianWithoutDatabase/controller"
	"log"
	"net/http"
)

func main() {
	//home
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/queue", controller.QueueHome)
	http.HandleFunc("/logout", controller.Logout)

	//http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// 	log.Printf("defaulting to port %s", port)
	// }

	// log.Printf("listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println(("server running"))
	http.ListenAndServe(":5050", nil)
}
