package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestAgregarElementos(t *testing.T) {
	t.Log("Prueba agregando 3 elementos, iniciando insertando al principio")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	require.False(t, lista.EstaVacia(), "Devuelve False, porque ya hay un elemento en la lista")
	require.Equal(t, 5, lista.VerPrimero(), "Elemento agregado, es el primero en la lista")
	require.Equal(t, 5, lista.VerUltimo(), "Al ser el único elemento, debe ser el último elemento también")
	require.Equal(t, 1, lista.Largo(), "Hay 1 elemento en la lista")
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero(), "Numero 3 es agregado a la primera posicion")
	require.Equal(t, 5, lista.VerUltimo(), "El numero 5 se quedó con la última posicion")
	require.Equal(t, 2, lista.Largo(), "Hay 2 elementos en la lista")
	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo(), "Hay 3 elementos en la lista")
	lista.InsertarPrimero(0)
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 4, lista.Largo())
}

func TestListaInsertarPrimeroBorrarPrimero(t *testing.T) {
	//TEST BORRAR PRIMERO INSERTANDO PRIMERO
	var (
		sliceDecimal = []float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.9, 10.0}
	)
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
	require.Equal(t, 0, lista.Largo())
}

func TestListaInsertarUltimoBorrarPrimero(t *testing.T) {
	//Test insertando ultimo borrando primero
	const LISTA_VACIA = 0
	var (
		sliceVocales = []string{"a", "e", "i", "o", "u"}
	)
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

// --------------TEST ITERADOR INTERNO---------------------------//
func TestIteradorInternoVacio(t *testing.T) {
	//Test Iterador interno: Iterando una lista vacia
	const LISTA_VACIA = 0
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
	var (
		solucionTotal   int
		sliceNumeros    = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		sumaIterInterno int
	)
	lista := TDALista.CrearListaEnlazada[int]()
	for _, valor := range sliceNumeros {
		lista.InsertarPrimero(valor)
		solucionTotal += valor
	}

	lista.Iterar(func(i int) bool {
		sumaIterInterno += i
		return true
	})
	require.Equal(t, solucionTotal, sumaIterInterno)
	//La lista no debe ser modifica
	require.Equal(t, len(sliceNumeros), lista.Largo())
	require.Equal(t, sliceNumeros[lista.Largo()-1], lista.VerPrimero())
	require.Equal(t, sliceNumeros[lista.Largo()-lista.Largo()], lista.VerUltimo())
	require.False(t, lista.EstaVacia())

}

func TestIteradorInternoCorte(t *testing.T) {
	//TEST ENCONTRAR EL NUMERO IMPAR
	// Buscamos el indice del numero impar
	var (
		solucion       int
		impar          int
		arrBuscarInpar = []int{2, 4, 6, 8, 9, 10, 12, 14, 16, 18, 20}
	)
	listaCorte := TDALista.CrearListaEnlazada[int]()
	for _, valor := range arrBuscarInpar {
		listaCorte.InsertarUltimo(valor)
		if valor%2 != 0 {
			solucion = valor
		}
	}
	//Buscamos el indice donde se encuentra el primer numero impar
	listaCorte.Iterar(func(i int) bool {
		if i%2 != 0 {
			impar = i
			return false
		}
		return true
	})
	//verificamos que el indice coincida con el contador
	require.Equal(t, solucion, impar)
	// La lista no se debe modificar
	require.Equal(t, arrBuscarInpar[listaCorte.Largo()-listaCorte.Largo()], listaCorte.VerPrimero())
	require.Equal(t, arrBuscarInpar[listaCorte.Largo()-1], listaCorte.VerUltimo())
	require.False(t, listaCorte.EstaVacia())
}

func TestVolumen(t *testing.T) {

}

// -------------------TEST ITERADOR EXTERNO ----------------------------//
// De aquí para abajo serán Test unicamente para el Iterador Externo
func TestIteradorExternoListaVacia(t *testing.T) {
	t.Log("Haremos un test de iteracion con una lista vacia")
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "Entra en pánico, porque la lista esta vacia, así que entiende que ya terminó su iteracion")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() }, "Al no haber elemtnos para posicionarse, tira el panico")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() }, "Si no hay elemento posicionado, no puedo borrar nada")
	require.False(t, iter.HaySiguiente(), "Al ser vacia, no hay siguiente")
}

func TestAgregarElementoExternoListaVacia(t *testing.T) {
	t.Log("Agregar un elemento externamente en una lista vacia debe ser aceptado")
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()
	iter.Insertar("Alex")
	require.Equal(t, "Alex", lista.VerPrimero(), "Al insertar un elemento en una lista vacia, tanto tanto con un iterador externo o no, ese elemento es el primer y ultimo elemento de la lista")
	require.Equal(t, "Alex", lista.VerUltimo())
	require.Equal(t, "Alex", iter.VerActual(), "Al agregarse un elemento, ")
	require.Equal(t, 1, lista.Largo())
	require.True(t, iter.HaySiguiente())
}

func TestAgregarVariosElementosExternosListaVacia(t *testing.T) {
	t.Log("Agregar varios elementos a una lista vacia")
	decimales := []float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8}
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	for i := len(decimales) - 1; i >= 0; i-- {
		iter.Insertar(decimales[i])
	}
	require.Equal(t, 1.2, lista.VerPrimero())
	require.Equal(t, 7.8, lista.VerUltimo())
	require.Equal(t, 1.2, iter.VerActual())
	require.True(t, iter.HaySiguiente())

	t.Log("Nos 1 movemos posicion")
	iter.Siguiente()
	require.Equal(t, 2.3, iter.VerActual())
	require.True(t, iter.HaySiguiente())

	t.Log("Nos movemos 3 posiciones")
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 5.6, iter.VerActual())
	require.True(t, iter.HaySiguiente())

	t.Log("Avanzamos al elemento final")
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 7.8, iter.VerActual())
	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

	require.Equal(t, 7, lista.Largo(), "Al haber agregado externamente, el Largo debe ser igual a la cantidad de elementos que tiene la lista")
}

func TestInsertarElementosExternamente(t *testing.T) {
	numeros := []int{2, 3, 5, 6, 7}
	lista := TDALista.CrearListaEnlazada[int]()
	for i := len(numeros) - 1; i >= 0; i-- {
		lista.InsertarPrimero(numeros[i])
	}
	iter := lista.Iterador()

	t.Log("Insertamos un elemento al inicio de la lista")
	iter.Insertar(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual(), "El 1 se puso en la primera posicion y 'mueve' el resto a la derecha")

	t.Log("Avanzamos a la posicion del numero 5 e insertamos un elemento")
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 5, iter.VerActual())
	iter.Insertar(4)
	require.Equal(t, 4, iter.VerActual())

	t.Log("Avanzamos a la ultima posicion e insetamos al ultimo")
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	iter.Insertar(8)
	require.Equal(t, 8, iter.VerActual())
	require.Equal(t, 8, lista.VerUltimo())

	require.Equal(t, 8, lista.Largo(), "El largo de la lista es 8 por ser los 5 iniciales + 3 insertados")
}

func TestBorrarElementoExternamente1(t *testing.T) {
	t.Log("Borraremos un elemento de una lista con un solo elemento")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("Tengo estres")
	iter := lista.Iterador()
	require.Equal(t, "Tengo estres", iter.VerActual())
	require.Equal(t, "Tengo estres", iter.Borrar(), "Se elimina el elemento donde este el iterador, de la lista")
	iter.Insertar("Podré promocionar?")
	require.Equal(t, "Podré promocionar?", iter.VerActual())
	require.Equal(t, "Podré promocionar?", iter.Borrar())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "Al borrarse el unico elemento, no hay nada para iterar, entonces 'termino de iterar'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestBorrarElementosExternamente(t *testing.T) {
	apellidos := []string{"Juarez", "Limachi", "Lezama", "Cordero", "Messi", "Miranda", "Avalos"}
	lista := TDALista.CrearListaEnlazada[string]()
	for i := 0; i < len(apellidos); i++ {
		lista.InsertarUltimo(apellidos[i])
	}
	iter := lista.Iterador()

	t.Log("Borraremos el primer elemento")
	require.Equal(t, "Juarez", iter.Borrar(), "Al inicializarse el iterado el primer elemento de una lista, se elimina 'Juarez'")
	require.Equal(t, "Limachi", iter.VerActual(), "Se borra a 'Juarez' y el iterador cambia de posicion al nuevo primer elemento de la lista 'Limachi'")
	require.Equal(t, "Limachi", lista.VerPrimero(), "Limachi es el nuevo primero elemento de la lista")

	t.Log("Borraremos un elemento del medio")
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, "Cordero", iter.VerActual())
	require.Equal(t, "Cordero", iter.Borrar())
	require.Equal(t, "Messi", iter.VerActual())

	t.Log("Borramos el ultimo elemento de la lista")
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, "Avalos", iter.VerActual())
	require.Equal(t, "Avalos", iter.Borrar())
	require.Equal(t, "Miranda", lista.VerUltimo(), "Al borrar 'Avalos' de la lista, el elemento anterior es el nuevo ultimo elemento de la lista")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "Entra en panic porque actualmente está apuntando al nill de la lista (el final)")

	require.Equal(t, 4, lista.Largo(), "Por haber eliminado 3 elementos de una lista con 7 elementos")
}

func TestComprobarCambiosIteradorExterno(t *testing.T) {
	t.Log("Insertar y Borrar distintos elementos en la lista con iterador externo")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("B")
	lista.InsertarPrimero("A")
	iter := lista.Iterador()
	require.Equal(t, "A", iter.Borrar())
	iter.Insertar("C")
	iter.Siguiente()
	require.Equal(t, "B", iter.Borrar())
	iter.Insertar("D")
	iter.Siguiente()
	iter.Insertar("E")
	iter.Insertar("F")
	iter.Insertar("G")
	iter.Insertar("H")
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, "F", iter.Borrar())
	require.Equal(t, 5, lista.Largo())
	require.Equal(t, "C", lista.VerPrimero())
	require.Equal(t, "E", lista.VerUltimo())

	t.Log("Nuevo iterador externo para visualizar si los cambios efecutuados tras insertar y borrar son aceptados")
	iter2 := lista.Iterador()
	require.Equal(t, "C", iter2.Borrar())
	require.Equal(t, "D", iter2.Borrar())
	require.Equal(t, "H", iter2.Borrar())
	require.Equal(t, "G", iter2.Borrar())
	require.Equal(t, "E", iter2.Borrar())
	require.True(t, lista.EstaVacia())
}
