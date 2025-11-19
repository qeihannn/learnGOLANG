package main

import "fmt"

func main() {

	var nilaiAkhir = 90
	var Absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 90
	var lulusAbsensi bool = Absensi >= 80 

	var lulus bool = lulusNilaiAkhir && lulusAbsensi 

	fmt.Println(lulus)


}

