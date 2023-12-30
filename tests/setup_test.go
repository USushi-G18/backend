package tests

import (
	"log"
	"os"
	"testing"
	u_sushi "u-sushi"
)

// this function will be called before every test
func TestMain(m *testing.M) {
	err := setupTests()
	if err != nil {
		log.Fatalln(err)
	}
	code := m.Run()

	os.Exit(code)
}

func setupTests() error {
	err := u_sushi.ConnectToDB()
	if err != nil {
		return err
	}

	file, err := os.ReadFile("setup.sql")
	if err != nil {
		return err
	}

	_, err = u_sushi.GetDB().Exec(string(file))
	if err != nil {
		return err
	}

	return nil
}
