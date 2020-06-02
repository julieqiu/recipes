package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// http.Get = function. Input = string, Output = two variables
	//
	// func Get(url string) (resp *Response, err error) {

	url := "https://www.delish.com/cooking/recipe-ideas/recipes/a50630/perfect-mashed-potatoes-recipe"

	// Either resp or err == nil
	// (1) Call func
	// (2) Check if error is nil
	// (2a) If error is nil, something bad happened. Fail.
	// (3) Error is not nil. This was successful. Look at resp.
	resp, err := http.Get(url)
	if err != nil {
		// resp is nil here
		log.Fatal(err)
	}
	// resp will NOT be nil, because err == nil
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("http.Get: %d", resp.Status)
	}

	// url == delish
	// err == nil
	// resp != nil
	// resp.Status == 200

	// anytime you want to read what is in resp.Body, you gotta open it, then
	// you gotta close it
	//
	// resp.Body == file

	// A defer statement defers the execution of a function until the
	// surrounding function returns.
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
