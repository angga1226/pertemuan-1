package main

import "fmt"

func main() {
	fmt.Println("tes")


// kondisi
//1.if
//kondisinya harus true

// if 1 < 10 {
// 	fmt.Println("Hola kecil")
// }else if a == 16{	fmt.Println("hai belasan")
// }else {	fmt.Println(ini apa)
// }

//2. switch
// nama := "angga"
// switch nama {
// case "Djohan":
// 		fmt.Println ("hai Djo")
// case "septia":
// 		fmt.Println("hai septia")
// default:
// 		fmt.Println(ini siapa)
// }

//memasukan input
var nama string
fmt.Print("masukan nama : ")
fmt.Scanln(&nama)

fmt.Println("Hola nama saya", nama)
}



