package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body [] byte
}

func (P* Page) save() error {
	filename := P.Title + ".txt"
	return ioutil.WriteFile(filename,  P.Body, 0600)
}

func load(tittle string) (*Page, error){
	filename := tittle + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err //error handling.
	}
	return &Page{Title:tittle, Body:body}, nil
}

func main() {
	P1 := Page{Title:"number1", Body:[]byte("wow we are")}
	P1.save() //can only call like this if func is receiving pointer.
	P2, err := load("number1")

	_  = err // probably not good to do.
	fmt.Println(string(P1.Body))
	fmt.Println(string(P2.Body))
}
