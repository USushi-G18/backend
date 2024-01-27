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
		err = http.ListenAndServe(":8081", r)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		defer waitGroup.Done()

		r := mux.NewRouter()
		client_routes.HandleAll(r)
		employee_routes.HandleAll(r)
		err = http.ListenAndServe(":8082", r)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	waitGroup.Wait()
}
