package main

import "fmt"

func main () {

	 type NoKTP string
	 
	 var ktpQeihan NoKTP = "1111111"

	 var contoh string = "222222"

	 var contohKtp NoKTP = NoKTP(contoh)

	 fmt.Println(ktpQeihan)
	 fmt.Println(contohKtp)			
}