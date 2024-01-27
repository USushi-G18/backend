package tests

import (
	"net"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"
	u_sushi "u-sushi"
	"u-sushi/server"
)

func waitServer() {
	timeout := 10 * time.Millisecond

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	waitHelper := func(url string) {
		defer waitGroup.Done()
		for {
			time.Sleep(timeout)
			conn, err := net.DialTimeout("tcp", url, timeout)
			if err == nil {
				conn.Close()
				break
			}
		}
	}

	go waitHelper("localhost:8081")
	go waitHelper("localhost:8082")

	waitGroup.Wait()
}

func setupDB() error {
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

func TestIntegration(t *testing.T) {
	go server.StartServer()
	waitServer()
	err := setupDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	out, err := exec.Command(
		"venom",
		"run",
		"--var-from-file=suites/env/env.yml",
		"suites/*.yml",
		"--output-dir=venom",
		"--html-report",
	).Output()
	if err != nil {
		t.Fatalf("running venom: %s", string(out))
	}
}
