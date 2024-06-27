package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)

	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)

}

func Hello(w http.ResponseWriter, r *http.Request){

	name:= os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, Im %s. Im %s.", name, age)

	w.Write([]byte("<h1>Hello full cycle !!!</h1>"))
}

func ConfigMap(w http.ResponseWriter, r *http.Request){

	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil{
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My family: %s.", string(data))
}