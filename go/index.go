package main

import (
    "fmt"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "time"
    "log"
    "math/rand"
)

func main() {
	server := &http.Server{ ReadTimeout: 5 * time.Second, WriteTimeout:  10 * time.Second, Addr: ":8080" }
    http.HandleFunc("/golang/api-calls", ApiCallsServer)
    http.HandleFunc("/golang/hash", HashServer)
    http.HandleFunc("/golang/mix", MixServer)
    log.Fatal(server.ListenAndServe())
}

func ApiCallsServer(w http.ResponseWriter, r *http.Request) {
    client := http.Client{ Timeout: time.Second * 5 }
	response, _ := client.Get("some-url-you-can-put-load-on")
	response.Body.Close()
    fmt.Fprintf(w, "OK");
}

func HashServer(w http.ResponseWriter, r *http.Request) {
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 12)
    fmt.Fprintf(w, "{ \"hash\": " + string(hashedPassword) + "\" }");
}

func MixServer(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    if rand.Intn(100) == 0 {
        bcrypt.GenerateFromPassword([]byte("password"), 12)
    } else {
        client := http.Client{ Timeout: time.Second * 5 }
	    response, _ := client.Get("http://ec2-52-63-249-102.ap-southeast-2.compute.amazonaws.com/")
	    response.Body.Close()
    }
    fmt.Fprintf(w, "OK");
}
