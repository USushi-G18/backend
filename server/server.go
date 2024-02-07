package server

import (
	"log"
	"net/http"
	"sync"

	u_sushi "u-sushi"

	admin_routes "u-sushi/server/routes/admin"
	client_routes "u-sushi/server/routes/client"
	employee_routes "u-sushi/server/routes/employee"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartServer() {
	err := u_sushi.ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		r := mux.NewRouter()
		admin_routes.HandleAll(r)
		c := cors.AllowAll().Handler(r)
		err = http.ListenAndServe(":8081", c)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		defer waitGroup.Done()

		r := mux.NewRouter()
		client_routes.HandleAll(r)
		employee_routes.HandleAll(r)
		c := cors.AllowAll().Handler(r)
		err = http.ListenAndServe(":8082", c)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	waitGroup.Wait()
}
