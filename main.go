package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/satori/go.uuid"
)

var a *Agent

var HWID string

const HOME_URL string = "http://git.leeingram.com:8765"

func main() {
	HWID = generateHWID()
	a = createBlankAgent()
	if checkFirstRun() {
		firstStartup()
	} else {
		fmt.Println(a)
	}
}

func checkFirstRun() bool {
	url := HOME_URL + "/agents/" + HWID
	fmt.Println(url)
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(a)
	if a.ID != 0 {
		return false
	}
	return true
}

func firstStartup() {
	fmt.Println("First time")
	a.OS = gleanOS()
	a.Location = gleanLocation()
	a.UUID = generateHWID()
	registerAgent()
}

func registerAgent() {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(a)
	r, err := http.Post(HOME_URL+"/agents", "application/json; charset=utf-8", b)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(a)
	fmt.Println("Agent registered:", a.UUID)
}

func gleanLocation() string {
	return "china"
}

func gleanOS() string {
	return runtime.GOOS
}

func generateHWID() string {
	return uuid.NewV1().String()
}

func createBlankAgent() *Agent {
	return new(Agent)
}

//Agent data structure
type Agent struct {
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	OS          string    `json:"os"`
	Location    string    `json:"location"`
	CreatedDate time.Time `json:"createddate"`
}
