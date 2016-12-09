package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

var a *Agent

func main() {
	firstStartup()
}

func firstStartup() {
	fmt.Println("First time")
	a = createBlankAgent()
	a.OS = gleanOS()
	a.Location = gleanLocation()
	a.UUID = generateHWID()
	phoneHome()
}

func phoneHome() {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(a)
	r, err := http.Post("http://git.leeingram.com:8765/agents", "application/json; charset=utf-8", b)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(a)
	fmt.Println(a.UUID)
}

func gleanLocation() string {
	return "china"
}

func gleanOS() string {
	return "Windows"
}

func generateHWID() uuid.UUID {
	return uuid.NewV1()
}

func createBlankAgent() *Agent {
	return new(Agent)
}

//Agent data structure
type Agent struct {
	ID          int       `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	OS          string    `json:"os"`
	Location    string    `json:"location"`
	CreatedDate time.Time `json:"createddate"`
}
