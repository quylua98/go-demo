package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	res := ""

	//file, err := ioutil.ReadFile("go-docs.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//file2, err := ioutil.ReadFile("go-docs.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//res = string(file) + string(file2)

	a, b := callApi(), callApi()
	res += <-a + <-b


	fmt.Fprintf(w, res)
}

func callApi() <-chan string {
	r := make(chan string)
	go func() {
		response, err := http.Get("http://localhost:9069/test")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			r <- string(data)
		}
	}()
	return r
}

func main() {
	http.HandleFunc("/test", test) // set router
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
