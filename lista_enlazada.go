package lista

const _PANIC_LISTA_VACIA string = "La lista esta vacia"

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func nuevoNodo[T any](elemento T) *nodoLista[T] {
	return &nodoLista[T]{dato: elemento, siguiente: nil}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoEnlace := nuevoNodo(elemento)
	if lista.EstaVacia() {
		lista.ultimo = nuevoEnlace
	} else {
		nuevoEnlace.siguiente = lista.primero
	}
	lista.primero = nuevoEnlace
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoEnlace := nuevoNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevoEnlace
	} else {
		lista.ultimo.siguiente = nuevoEnlace
	}
	lista.ultimo = nuevoEnlace
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	//completar en adelante
	panic(_PANIC_LISTA_VACIA)
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(_PANIC_LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_PANIC_LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	panic("unimplemented")
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	panic("unimplemented")
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	panic("unimplemented")
}
