package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func OutputLog() {
	// Output to ./test.log file
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error!")
		log.Fatal(err)
	}
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
}

func raw_connect(host string, ports []string, path string) {
	p := []string{}
	for _, port := range ports {
		timeout := time.Millisecond * 10
		address := fmt.Sprintf("%s%s", net.JoinHostPort(host, port), path)
		fmt.Println("address: ", address)
		conn, err := net.DialTimeout("tcp", address, timeout)

		if err != nil {
			fmt.Println("Connecting error:--", err.Error())
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", address)
			p = append(p, address)
			//log.Fatal("STOP program")
		}
	}
	fmt.Println("ports ouverts: ", p)
}

func Request(ip string, ports []string, path string, wg *sync.WaitGroup) {
	for _, port := range ports {
		go testconnexion(ip, port, path, wg)
	}
	fmt.Println("before channel")
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

// stock le port dans le channel
func testconnexion(host string, port string, path string, wg *sync.WaitGroup) {

	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s%s", host, port, path)
	_, err := client.Get(url)

	if err == nil {
		fmt.Println("port: ", port)
		channel = make(chan string)
		wg.Done()
		fmt.Println("url: ", url)
		fmt.Println("err == nil ")
		channel <- port
	}
}

func Ping(host string, port string, path string) string {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s%s", host, port, path)
	log.Info("URL: ", url)
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(body)
}

func Signup(host string, port string, path string) string {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s%s", host, port, path)
	log.Info("URL Signup: ", url)
	body := JsonStruct{User: "janedove"}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(Serialize(body)))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	b, e := io.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e.Error())
	}
	//fmt.Println(string(b))
	return string(b)
}

func parseRespo(body string) *JsonStruct {
	js := new(JsonStruct)
	err := json.Unmarshal([]byte(body), &js)
	if err != nil {
		log.Fatal(err.Error())
	}
	return js
}

type JsonStruct struct {
	User   string `json:"User"`
	Secret string `json:"secret"`
}

func Serialize(object JsonStruct) []byte {
	b, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err.Error())
	}
	return b
}

func Check(host string, port string, path string) string {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s%s", host, port, path)
	log.Info("URL Check: ", url)
	body := JsonStruct{User: "janedove"}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(Serialize(body)))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	b, e := io.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e.Error())
	}
	//fmt.Println(string(b))
	return string(b)
}

func GetUserSecret(host string, port string, path string) *JsonStruct {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s%s", host, port, path)
	log.Info("URL Check: ", url)
	body := JsonStruct{User: "janedove"}
	var str string
	for {
		resp, err := client.Post(url, "application/json", bytes.NewBuffer(Serialize(body)))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()
		b, e := io.ReadAll(resp.Body)
		if e != nil {
			log.Fatal(e.Error())
		}
		str = string(b)
		fmt.Println(str)
		if str[0] != 'R' {
			break
		}
	}
	log.Info("URL getUserSecret: ", url)
	js := JsonStruct{User: body.User, Secret: str[12:]}
	return &js
}
