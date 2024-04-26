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

// func TestCrearListaEnlazada(t *testing.T) {
// 	slicePrueba := []int{1, 2, 3, 4, 5}
// 	lista := TDALista.CrearListaEnlazada[int]()
// 	//for _, valor := range slicePrueba {
// 	//	lista.InsertarPrimero(valor)
// 	//}
// 	//require.Equal(t, slicePrueba[len(slicePrueba)-1], lista.VerPrimero())
// 	//for i := 0; i < len(slicePrueba); i++ {
// 	//	lista.BorrarPrimero()
// 	//}
// 	//require.True(t, lista.EstaVacia())
// 	//for _, valor := range slicePrueba {
// 	//	lista.InsertarUltimo(valor)
// 	//}

// 	lista.Iterar(func(i int) bool {
// 		fmt.Println(i)
// 		return true
// 	})

// }
