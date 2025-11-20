package main

import "fmt"

func main() {

	person := map[string] string {

		"name" : "qeihan",
		"address" : "jakarta",

	}

	person["hobby"] = "combat sport"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["hobby"])
		


	var book map [string] string = make (map[string] string)
	book["title"] = "belajar go"
	book["author"] = "qeihan"
	book["ups"] = "wrong"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}