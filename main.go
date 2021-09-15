package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var config *Config

func hello(w http.ResponseWriter, req *http.Request) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(config.APIURL + "/name")
	if err != nil {
		fmt.Fprintf(w, "Unable to connect to API!\n")
		return
	}
	if resp.StatusCode != 200 {
		fmt.Fprintf(w, "API returned %s.\n", resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Unable to determine name!\n")
		return
	}
	fmt.Fprintf(w, "Hello, I'm %s!\n", string(body))
}

func name(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, config.Name)
}

func status(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	var err error
	config, err = readConfig("config.json")
	if err != nil {
		fmt.Println("error occurred reading configuration")
		os.Exit(1)
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/name", name)
	http.HandleFunc("/", status)
	port := fmt.Sprintf(":%d", config.Port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
