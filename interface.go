package main

import (
	"fmt"
	"math"
)

func main() {
	p := persegiPanjang{panjang: 5, lebar: 7}
	l := lingkaran{jariJari: 7}

	penghitungan(p)
	penghitungan(l)

}

// Apa yang akan di-print
func penghitungan(g geometri) {
	fmt.Println(g)
	fmt.Println("Luas bidang : ", g.luas())
	fmt.Println("Keliling bidang : ", g.keliling())
}

// Begin Golang Interface
type geometri interface {
	luas() float64
	keliling() float64
}

// Bidang Geometri yang ingin dicari
type persegiPanjang struct {
	panjang, lebar float64
}

type lingkaran struct {
	jariJari float64
}

// Implementasi persegiPanjang ke interface geometri
func (p persegiPanjang) luas() float64 {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() float64 {
	return 2 * p.panjang + 2 * p.lebar
}

// Implementasi lingkaran ke interface geometri
func (l lingkaran) luas() float64 {
	return math.Pi * l.jariJari * l.jariJari
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jariJari
}

