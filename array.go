package main

import "fmt"

func main() {

	var name [3] string
	name[0] = "qeihan"
	name[1] = "hapis"
	name[2] = "tefun"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		50,
		100,
		21,
	}

	fmt.Println(values[2])

	var teams = [3]string{
		"padil",
		"qeihan",
		"suban",
	}

	fmt.Println(teams)

	var test [10]string

	fmt.Println(len(test))


}