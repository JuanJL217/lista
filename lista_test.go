package lista_test

import (
	"fmt"
	TDALista "tdas/lista"
	"testing"
)

func TestCrearListaEnlazada(t *testing.T) {
	slicePrueba := []int{1, 2, 3, 4, 5}
	lista := TDALista.CrearListaEnlazada[int]()
	//for _, valor := range slicePrueba {
	//	lista.InsertarPrimero(valor)
	//}
	//require.Equal(t, slicePrueba[len(slicePrueba)-1], lista.VerPrimero())
	//for i := 0; i < len(slicePrueba); i++ {
	//	lista.BorrarPrimero()
	//}
	//require.True(t, lista.EstaVacia())
	//for _, valor := range slicePrueba {
	//	lista.InsertarUltimo(valor)
	//}

	lista.Iterar(func(i int) bool {
		fmt.Println(i)
		return true
	})

}
