package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	LISTA_VACIA = 0
)

var (
	sliceNumeros   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceVocales   = []string{"a", "e", "i", "o", "u"}
	sliceDecimal   = []float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.9, 10.0}
	arrBuscarInpar = []int{2, 4, 6, 8, 9, 10, 12, 14, 16, 18, 20}
)

func TestListaVacia(t *testing.T) {
	t.Log("Pruebas con una lista crecien creada")
	lista := TDALista.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia(), "Devuelve True, por ser una lista vacia")
	require.Equal(t, 0, lista.Largo(), "Al ser una lista vacia, tiene largo 0")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Lista vacia, no hay elementos para borrar")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Si no hay elemento, no hay nada en la primera posicion")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Si no hay un inicio, no hay un final")
}

func TestAgregarElementosParte1(t *testing.T) {
	t.Log("Prueba agregando 3 elementos, iniciando insertando al principio")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	require.False(t, lista.EstaVacia(), "Devuelve False, porque ya hay un elemento en la lista")
	require.Equal(t, 5, lista.VerPrimero(), "Elemento agregado, es el primero en la lista")
	require.Equal(t, 5, lista.VerUltimo(), "Al ser el único elemento, debe ser el último elemento también")
	require.Equal(t, 1, lista.Largo(), "Hay 1 elemento en la lista")
	lista.InsertarPrimero(7)
	require.Equal(t, 7, lista.VerPrimero(), "Numero 3 es agregado a la primera posicion")
	require.Equal(t, 5, lista.VerUltimo(), "El numero 5 se quedó con la última posicion")
	require.Equal(t, 2, lista.Largo(), "Hay 2 elementos en la lista")
	lista.InsertarUltimo(10)
	require.Equal(t, 7, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo(), "Hay 3 elementos en la lista")

}

func TestAgregarElementosParte2(t *testing.T) {
	t.Log("Prueba agregando 3 elementos, iniciando insertando al principio")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("Juan")
	require.False(t, lista.EstaVacia(), "Devuelve False, porque ya hay un elemento en la lista")
	require.Equal(t, "Juan", lista.VerPrimero(), "")
	require.Equal(t, "Juan", lista.VerUltimo(), "")
	require.Equal(t, 1, lista.Largo(), "Hay 1 elemento en la lista")
	// Juan: Yo termino con esta prueba
	// Alex: weno pa
}

func TestListaInsertarPrimeroBorrarPrimero(t *testing.T) {
	//TEST BORRAR PRIMERO INSERTANDO PRIMERO
	lista := TDALista.CrearListaEnlazada[float64]()
	for _, valor := range sliceDecimal {
		lista.InsertarPrimero(valor)
	}
	require.Equal(t, sliceDecimal[lista.Largo()-1], lista.VerPrimero())
	for i := 0; i < len(sliceDecimal); i++ {
		require.Equal(t, sliceDecimal[lista.Largo()-1], lista.VerPrimero())
		lista.BorrarPrimero()

	}
	//Verificamos que la lista quede vacia una vez borrado los primeros elementos
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() {
		lista.VerUltimo()
		lista.VerPrimero()
		lista.BorrarPrimero()
	})
	require.Equal(t, LISTA_VACIA, lista.Largo())
}

func TestListaInsertarUltimoBorrarPrimero(t *testing.T) {
	//Test insertando ultimo borrando primero
	lista := TDALista.CrearListaEnlazada[string]()
	for _, valor := range sliceVocales {
		lista.InsertarUltimo(valor)
	}
	require.Equal(t, sliceVocales[lista.Largo()-1], lista.VerUltimo())
	//Vamos borrando los primeros datos , el ultimo no se modifica
	for i := 0; i < len(sliceVocales); i++ {
		require.Equal(t, sliceVocales[len(sliceVocales)-1], lista.VerUltimo())
		lista.BorrarPrimero()
	}
	//Verificamos que la lista quede vacia una vez borrado los primeros elementos
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() {
		lista.VerUltimo()
		lista.VerPrimero()
		lista.BorrarPrimero()
	})
	require.Equal(t, LISTA_VACIA, lista.Largo())
}

// De aquí para abajo serán Test unicamente para el Iterador Externo
func TestIteradorExternoListaVacia(t *testing.T) {
	t.Log("Haremos un test de iteracion con una lista vacia")
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "Entra en pánico, porque la lista esta vacia, así que entiende que ya terminó su iteracion")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "Al no haber elemtnos para posicionarse, tira el panico")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() }, "Si no hay elemento posicionado, no puedo borrar nada")
	require.False(t, iter.HaySiguiente(), "Al ser vacia, no hay siguiente")
}

func TestAgregarElementosExternos(t *testing.T) {
	t.Log("Agregar un elemento en una lista vacia debe ser aceptado")
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()
	iter.Insertar("Hola")
	require.Equal(t, "Hola", lista.VerPrimero())
	require.Equal(t, "Hola", lista.VerUltimo())
}

// --------------TEST ITERADOR INTERNO---------------------------//
func TestIteradorInternoVacio(t *testing.T) {
	//Test Iterador interno: Iterando una lista vacia
	lista := TDALista.CrearListaEnlazada[int]()
	lista.Iterar(func(i int) bool {
		require.True(t, lista.EstaVacia())
		return true
	})
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() {
		lista.VerPrimero()
		lista.VerUltimo()
		lista.BorrarPrimero()
	})
	require.Equal(t, LISTA_VACIA, lista.Largo())
}

func TestIteradorInternoCompleto(t *testing.T) {
	//Test SUMA TODOS LOS ELEMENTOS
	const TOTAL_SUMA = 55
	lista := TDALista.CrearListaEnlazada[int]()
	for _, valor := range sliceNumeros {
		lista.InsertarPrimero(valor)
	}
	contador := 0
	lista.Iterar(func(i int) bool {
		contador += i
		return true
	})
	require.Equal(t, contador, TOTAL_SUMA)
	//La lista no debe ser modifica
	require.Equal(t, len(sliceNumeros), lista.Largo())
	require.Equal(t, sliceNumeros[lista.Largo()-1], lista.VerPrimero())
	require.Equal(t, sliceNumeros[lista.Largo()-lista.Largo()], lista.VerUltimo())
	require.False(t, lista.EstaVacia())

}

func TestIteradorInternoCorte(t *testing.T) {
	//TEST ENCONTRAR EL NUMERO IMPAR
	// Buscamos el indice del numero impar
	const INDICE_NUMERO_IMPAR = 4
	listaCorte := TDALista.CrearListaEnlazada[int]()
	for _, valor := range arrBuscarInpar {
		listaCorte.InsertarUltimo(valor)
	}
	contador := 0
	//Buscamos el indice donde se encuentra el primer numero impar
	listaCorte.Iterar(func(i int) bool {
		if i%2 != 0 {
			return false
		}
		contador++
		return true
	})
	//verificamos que el indice coincida con el contador
	require.Equal(t, contador, INDICE_NUMERO_IMPAR)
	// La lista no se debe modificar
	require.Equal(t, arrBuscarInpar[listaCorte.Largo()-listaCorte.Largo()], listaCorte.VerPrimero())
	require.Equal(t, arrBuscarInpar[listaCorte.Largo()-1], listaCorte.VerUltimo())
	require.False(t, listaCorte.EstaVacia())

}
