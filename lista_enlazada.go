package lista

const (
	_PANIC_LISTA_VACIA   string = "La lista esta vacia"
	_PANIC_FIN_ITERACION string = "El iterador termino de iterar"
)

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
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func nuevoNodo[T any](elemento T) *nodoLista[T] {
	return &nodoLista[T]{dato: elemento}
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
	if lista.EstaVacia() {
		panic(_PANIC_LISTA_VACIA)
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	return dato
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
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoAux := lista.primero
	sigamos := true
	for sigamos || nodoAux != nil {
		nodoAux = nodoAux.siguiente
		if !visitar(nodoAux.dato) {
			sigamos = false
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista: lista, actual: lista.primero}
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
	if iterador.lista.EstaVacia() {
		iterador.lista.InsertarPrimero(elemento)
	} else if iterador.actual == iterador.lista.primero {
		nuevoEnlace.siguiente = iterador.lista.primero
		iterador.lista.primero = nuevoEnlace
	} else {
		nuevoEnlace.siguiente = iterador.actual
		iterador.anterior.siguiente = nuevoEnlace
		if nuevoEnlace.siguiente == nil {
			iterador.lista.ultimo = nuevoEnlace
		}
	}
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic(_PANIC_FIN_ITERACION)
	}
	elementoBorrar := iterador.actual.dato
	if iterador.actual == iterador.lista.primero {
		iterador.lista.primero = iterador.lista.primero.siguiente
		iterador.actual = iterador.lista.primero
		if iterador.lista.primero == nil {
			iterador.lista.ultimo = nil
		}
	} else if iterador.actual == iterador.lista.ultimo {
		iterador.lista.ultimo = iterador.anterior
		iterador.actual.siguiente = nil
		iterador.actual = nil
	} else {
		iterador.actual = iterador.actual.siguiente
		iterador.anterior.siguiente = iterador.actual
	}
	iterador.lista.largo--
	return elementoBorrar
}
