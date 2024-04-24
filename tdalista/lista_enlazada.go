package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (le listaEnlazada[T]) EstaVacia() bool {
	return le.largo == 0
}

func (le *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := nodoCrear(dato, nil)
	if le.EstaVacia() {
		le.primero = nuevoNodo
		le.ultimo = nuevoNodo
	} else {
		nodoAnterior := le.primero
		le.primero = nuevoNodo
		le.primero.siguiente = nodoAnterior
	}
	le.largo++

}

func (le *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := nodoCrear(dato, nil)
	if le.EstaVacia() {
		le.ultimo = nuevoNodo
		le.primero = nuevoNodo
	} else {
		nodoUltimo := le.ultimo
		le.ultimo = nuevoNodo
		le.ultimo.siguiente = nodoUltimo
	}
	le.largo++
}
func (le listaEnlazada[T]) Largo() int {
	return le.largo
}

func (le listaEnlazada[T]) VerPrimero() T {
	if le.EstaVacia() {
		panic("la lista esta vacia")
	}
	return le.primero.dato
}

func (le listaEnlazada[T]) VerUltimo() T {
	if le.EstaVacia() {
		panic("la lista esta vacia")
	}
	return le.ultimo.dato
}

func (le *listaEnlazada[T]) BorrarPrimero() T {
	if le.EstaVacia() {
		panic("la lista esta vacia")
	}
	dato := le.primero.dato
	le.primero = le.primero.siguiente
	le.largo--
	return dato
}

func nodoCrear[T any](dato T, sig *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato, sig}
}
