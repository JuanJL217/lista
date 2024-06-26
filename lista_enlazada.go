package lista

const (
	PANIC_LISTA_VACIA   string = "La lista esta vacia"
	PANIC_FIN_ITERACION string = "El iterador termino de iterar"
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

func crearNodo[T any](elemento T) *nodoLista[T] {
	return &nodoLista[T]{dato: elemento}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// --------------ITERADOR INTERNO ----------------------------//
func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoAux := lista.primero
	estado := true
	for nodoAux != nil && estado {
		elemento := nodoAux.dato
		nodoAux = nodoAux.siguiente
		if !visitar(elemento) {
			estado = false
		}
	}
}

//-----------------ITERADOR EXTERNO -----------------//

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista: lista, actual: lista.primero}
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic(PANIC_FIN_ITERACION)
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic(PANIC_FIN_ITERACION)
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if iterador.lista.EstaVacia() {
		iterador.lista.primero = nuevoNodo
		iterador.lista.ultimo = nuevoNodo
	} else if iterador.actual == iterador.lista.primero {
		nuevoNodo.siguiente = iterador.lista.primero
		iterador.lista.primero = nuevoNodo
	} else if iterador.actual == iterador.lista.ultimo.siguiente {
		nuevoNodo.siguiente = nil
		iterador.anterior.siguiente = nuevoNodo
		iterador.lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = iterador.actual
		iterador.anterior.siguiente = nuevoNodo
	}
	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic(PANIC_FIN_ITERACION)
	}
	elementoBorrardo := iterador.actual.dato
	if iterador.actual == iterador.lista.primero {
		iterador.lista.primero = iterador.lista.primero.siguiente
		iterador.actual = iterador.lista.primero
		iterador.anterior = nil
	} else if iterador.actual == iterador.lista.ultimo {
		iterador.anterior.siguiente = iterador.lista.ultimo.siguiente
		iterador.lista.ultimo = iterador.anterior
		iterador.actual = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
		iterador.actual = iterador.actual.siguiente
	}
	iterador.lista.largo--
	return elementoBorrardo
}
