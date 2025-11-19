package main

import "fmt"

func main() {

	var names = [...]string {
		"eko0",
		"eko1",
		"eko2",
		"eko3",
		"eko4",
		"eko5",
		"eko6",
		"eko7",
		"eko8",
		"eko9",
	}

	var slice1 = names[4:6]

	fmt.Println(slice1[0])
	fmt.Println(slice1[1])
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	names[5] = "diubah"
	fmt.Println(slice1)
}

