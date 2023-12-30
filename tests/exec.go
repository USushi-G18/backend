package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"u-sushi/auth"
	"u-sushi/models"
	admin_routes "u-sushi/routes/admin"
	client_routes "u-sushi/routes/client"
	employee_routes "u-sushi/routes/employee"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	admin_routes.HandleAll(r)
	client_routes.HandleAll(r)
	employee_routes.HandleAll(r)
	return r
}

func executeRequest(method string, url string, body interface{}) *httptest.ResponseRecorder {
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}
	if !strings.Contains(url, "/auth/") {
		req.Header.Add("Authorization", bearerAuth(url))
	}

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	return w
}

func bearerAuth(url string) string {
	var user interface{}

	userType := strings.Split(url, "/")[1]
	if userType == "client" {
		user = auth.ClientLoginRequest{
			Password:    "u-sushi",
			TableNumber: 0,
			Menu:        models.MenuCarte,
			Seatings:    1,
		}
	} else {
		user = auth.LoginRequest{
			Password: "u-sushi",
		}
	}

	authUrl := fmt.Sprintf("/%s/auth/login", userType)
	w := executeRequest("POST", authUrl, user)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("Bearer %s", response["token"])
}
