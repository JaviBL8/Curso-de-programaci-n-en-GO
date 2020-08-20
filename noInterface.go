package main

import "fmt"

type perro struct{}

type pez struct{}

type pajaro struct{}

func (perro) caminar() string {
	return "Soy un perro y camino"
}

func (pez) nadar() string {
	return "Soy un pez y estoy nadando"
}

func (pajaro) volar() string {
	return "Soy un pajaro y estoy volando"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(pe pez) {
	fmt.Println(pe.nadar())
}

func moverPajaro(pa pajaro) {
	fmt.Println(pa.volar())
}

func main() {
	p := perro{}
	moverPerro(p)
	pe := pez{}
	moverPez(pe)
	pa := pajaro{}
	moverPajaro(pa)
}
