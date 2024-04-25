package lista

const _PANIC_FIN_ITERACION string = "El iterador termino de iterar"

type iterListaEnlazada[T any] struct {
	listaIterar *listaEnlazada[T]
	actual      *nodoLista[T]
	anterior    *nodoLista[T]
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{listaIterar: lista, actual: lista.primero}
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic(_PANIC_FIN_ITERACION)
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic(_PANIC_FIN_ITERACION)
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoEnlace := nuevoNodo(elemento)
	if iterador.listaIterar.EstaVacia() {
		iterador.listaIterar.InsertarPrimero(elemento)
	} else if iterador.actual == iterador.listaIterar.primero {
		nuevoEnlace.siguiente = iterador.listaIterar.primero
		iterador.listaIterar.primero = nuevoEnlace
	} else {
		nuevoEnlace.siguiente = iterador.actual
		iterador.anterior.siguiente = nuevoEnlace
		if nuevoEnlace.siguiente == nil {
			iterador.listaIterar.ultimo = nuevoEnlace
		}
	}
}

func (lisiteradorta *iterListaEnlazada[T]) Borrar() T {
	//Completar
}
