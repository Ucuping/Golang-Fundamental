package main

import "fmt"

func main() {
	// pointer
	// numberA := 10
	// numberB := &numberA // numberB menginisialisasi alokasi memory yang sama dengan numberA
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB) // deliference mengembalikan nilai
	// *numberB = 20         // numberA akan berubah juga nilainya
	// fmt.Println(*numberB)
	// fmt.Println(numberA)
	var numberA int = 10
	var numberB *int = &numberA
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)
	numberA = 20
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)
	number := 10
	fmt.Println("Alamat memory:", &number)
	fmt.Println("Nilai awal:", number)
	change(&number, 50)
	fmt.Println("Nilai akhir:", number)
	fmt.Println("Alamat memory:", &number)
}

// pointer
func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory:", old)
	fmt.Println("Di dalam function:", *old)
}
