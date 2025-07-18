package main

import "fmt"

func main() {
	var altezza uint
	var peso uint
	var indice uint
	println("__Calcolatore del indice del peso__")

	println("Insersci la tua altezza")
	fmt.Scan(&altezza)

	println("Insersci il tuo peso")
	fmt.Scan(&peso)

	indice = altezza / peso

	fmt.Println("Il tuo indice è:", indice)

	// se indice è da 1 a 3 è normale, facciamo o if else oppure proviamo sitch con i varianti . dopo se voglono ripetere tutto un altra volta, qui dobbiamo fare un ciclo ? dopo che ho finito mi deve dare possibilità di fare un altro
	fmt.Println("Vuoi fare un'altro calcolo?", indice)
}
