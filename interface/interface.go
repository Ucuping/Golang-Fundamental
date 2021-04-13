package main

import "fmt"

// interface
type BangunDatar interface {
	HitungLuas() int
}

// interface
type Persegi struct {
	Sisi int
}

// interface
func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

// interface
type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

// interface
func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {

	persegi := Persegi{Sisi: 5}
	hitung := menghitungLuas(persegi)
	fmt.Println(hitung)

	persegiPanjang := PersegiPanjang{Panjang: 5, Lebar: 10}
	result := menghitungLuas(persegiPanjang)
	fmt.Println(result)

}

func menghitungLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
