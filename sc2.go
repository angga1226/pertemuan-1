package main

import "fmt"

func main() {
	 sc := "";
	for kondisi := 0; kondisi <= 3;kondisi++{
		for kondisi1 := 0 ; kondisi1 <= kondisi;kondisi1++{
			sc += "*";
		}
		sc += "\n";
	}
	fmt.Println(sc)
}
