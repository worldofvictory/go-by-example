package main

import (
	"fmt"
	"math"
)

func main() {

	println("__Calcolatore del indice del peso__")
	userKg, userHeight := inputUser()

	var IMT = userKg / math.Pow(userHeight, 2)
	outputResult(IMT)
	println(interpretaIMT(IMT))
	fmt.Printf("Vuoi fare un altro calcolo? (s/n): ")

}

func interpretaIMT(IMT float64) string {
	switch {
	case IMT < 16:
		return "Sei gravemente sottopeso."
	case IMT >= 16 && IMT < 18.5:
		return "Sei sottopeso."
	case IMT >= 18.5 && IMT < 25:
		return "Hai un peso normale."
	case IMT >= 25 && IMT < 30:
		return "Sei in sovrappeso."
	case IMT >= 30:
		return "Obesità."
	default:
		return "Valore non valido."
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Il tuo IMT  %.0f%%\n", imt)
	fmt.Print(result)
}

func inputUser() (float64, float64) {
	var userHeight float64
	var userKg float64 //
	println("Insersci la tua altezza in centimetri")
	fmt.Scan(&userHeight)

	println("Insersci il tuo peso")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

// se indice è da 1 a 3 è normale, facciamo o if else oppure proviamo sitch con i varianti . dopo se voglono ripetere tutto un altra volta, qui dobbiamo fare un ciclo ? dopo che ho finito mi deve dare possibilità di fare un altro
