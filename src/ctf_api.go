package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

func OutputLog() {
	// Output to ./test.log file
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
}

func Request(ip string, ports []string, path string, wg *sync.WaitGroup) {
	for _, port := range ports {
		go Connexion(ip, port, path, wg)
	}
}

func rangePort() []string {
	s := []string{}

	for i := 1024; i <= 4096; i++ {
		str := fmt.Sprint(i)
		s = append(s, str)
	}
	return s
}

var channel chan string
var client http.Client

// stock le port dans le channel
func Connexion(client *http.Client, host string, port string, path string, wg *sync.WaitGroup) {
	_, err := Get(client, host, port, path)
	if err == nil {
		channel = make(chan string)
		wg.Done()
		channel <- port
	}
}

func Get(client *http.Client, host string, port string, path string) (*http.Response, error) {
	url := fmt.Sprintf("http://%s:%s%s", host, port, path)
	log.Info("HTTP GET: ", url)
	return client.Get(url)
}

func Post(client *http.Client, body JsonStruct, host string, port string, path string) (*http.Response, error) {
	url := fmt.Sprintf("http://%s:%s%s", host, port, path)
	log.Info("HTTP POST: ", url)
	return client.Post(url, "application/json", bytes.NewBuffer(Serialize(body)))
}

func Ping(client *http.Client, host string, port string, path string) string {
	resp, err := Get(client, host, port, path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(body)
}

func Signup(client *http.Client, host string, port string, path string) string {

	// Body for POST request
	body := JsonStruct{User: "janedove"}

	resp, err := Post(client, body, host, port, path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	b, e := io.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e.Error())
	}
	return string(b)
}

func Check(client *http.Client, host string, port string, path string) string {
	body := JsonStruct{User: "janedove"}

	resp, err := Post(client, body, host, port, path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	b, e := io.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e.Error())
	}
	return string(b)
}

func GetUserSecret(client *http.Client, host string, port string, path string) *JsonStruct {
	body := JsonStruct{User: "janedove"}
	var str string
	for {
		resp, err := Post(client, body, host, port, path)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()

		b, e := io.ReadAll(resp.Body)
		if e != nil {
			log.Fatal(e.Error())
		}
		str = string(b)
		if str[0] != 'R' {
			break
		}
	}
	js := JsonStruct{User: body.User, Secret: str[12:]}
	return &js
}

type JsonStruct struct {
	User   string `json:"User"`
	Secret string `json:"secret"`
}

func parseRespo(body string) *JsonStruct {
	js := new(JsonStruct)
	err := json.Unmarshal([]byte(body), &js)
	if err != nil {
		log.Fatal(err.Error())
	}
	return js
}

func Serialize(object JsonStruct) []byte {
	b, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err.Error())
	}
	return b
}
