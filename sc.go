package main

import "fmt"

func main(){


var angka1 int
	fmt.Println("angka_pertama")
	fmt.Scanln(&nama)

var angka2 int
	fmt.Println("angka_kedua")
	fmt.Scanln(&angka)

var angka3 string
	fmt.Println("masukan operator perhitungan(+,-,x,:)")
	fmt.Scanln(&hasil)

if hasil == "+" {
	fmt.Println(nama + angka)
}else if hasil =="-"{
	fmt.Println(nama - angka)
}else if hasil =="x"{
	fmt.Println(nama * angka)
}else if hasil ==":"{
	fmt.Println(nama / angka)
}else {
	fmt.Println("eror co")
}
}


