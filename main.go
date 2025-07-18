package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64
	var userKg float64 //float64 = 0.0
	println("__Calcolatore del indice del peso__")

	println("Insersci la tua altezza in centimetri")
	fmt.Scan(&userHeight)

	println("Insersci il tuo peso")
	fmt.Scan(&userKg)

	var IMT = userKg / math.Pow(userHeight, 2)

	result := fmt.Printf("Il tuo indice è: %.0f%%\n", IMT)

	switch {
	case result < 16%:
		fmt.Scan("sei sottopeso")

	case result == 16%:
		fmt.Scan("sei peso norma")

	}

	// se indice è da 1 a 3 è normale, facciamo o if else oppure proviamo sitch con i varianti . dopo se voglono ripetere tutto un altra volta, qui dobbiamo fare un ciclo ? dopo che ho finito mi deve dare possibilità di fare un altro
	fmt.Printf("Vuoi fare un'altro calcolo?:")
}
