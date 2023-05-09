package main // Se llama el paquete main

import (
	"fmt"
	// Se importa el paquete fmt el cual es utilizado para imprimir
	"math"
	// Se importa el paquete math el cual implementa funciones matematicos como trigonométricas,
	// exponenciales y logarítmicas. Raices complejas
	//"math/cmplx"
	// Se importa este paquete math/cmplx el cual sirve para operaciones con numeros complejos
)

// La funcion transformada F, toma una lista o un arreglo de numeros complejos, para hacer la transformada
// de un tamaño n, se utiliza un metodo de divide y venceras. que es un metodo que lo resuelve de forma
// mas eficiente desarrollandolo por partes.

func transformadaF(x []complex64) []complex64 { // Se recibe un complex 128 numeros complejos

	n := len(x) // Calcula el tamaño de entrada n

	if n == 1 { // Si n es 1, la entrada es un solo elemento, por lo que se devuelve la entrada
		return x
	}

	// Si el algoritmo continua va a calcular las raíces complejas de la unidad para la transformada fft
	t := make([]complex64, n)
	for i := 0; i < n; i++ { //Calculo de raicez complejas
		imaginaria := 2 * math.Pi * float64(i) / float64(n)
		t[i] = complex64(complex(math.Cos(imaginaria), math.Sin(imaginaria)))

		//t[i] = complex64(cmplx.Rect(1, imaginaria))
		// utilizamos esto para construir un numero complejo es su parte real e imaginaria. real (1, theta) imaginaria
	}

	// Se divide la entrada en pares de forma recursiva
	pares := make([]complex64, n/2)
	impares := make([]complex64, n/2)
	// Con respecto a la señal de entrada el for separa los numeros pares e impares en cada arreglo
	for i := 0; i < n/2; i++ {
		pares[i] = x[2*i]
		impares[i] = x[2*i+1]
	}

	// Calcula la transformada de los dos arreglos pares e impares
	paresResultado := transformadaF(pares)
	imparesResultado := transformadaF(impares)

	// Combinar los resultados de los dos arreglos para tener un solo resultado
	resultadoFinal := make([]complex64, n) // Se declara un arreglo de resultado final.
	for i := 0; i < n/2; i++ {
		resultadoFinal[i] = paresResultado[i] + t[i]*imparesResultado[i]
		resultadoFinal[i+n/2] = paresResultado[i] - t[i]*imparesResultado[i]
	}
	return resultadoFinal
}

func main() {
	x := []complex64{complex(4, 0), complex(-4, 0), complex(8, 0), complex(-4, 0)} // Definir la entrada es un slice dinamico
	resultado := transformadaF(x)                                                  // Calcular la FFT de la entrada
	fmt.Println(resultado)                                                         // Imprime el resultado de la transformada
}
