package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/health", Health)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	fmt.Fprintf(w, "Hello, I'm %s. v1", name)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/family/pets.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My pets: %s.", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s. Password: %s", user, password)
}

func Health(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() < 10 { // liveness test: > 25
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}