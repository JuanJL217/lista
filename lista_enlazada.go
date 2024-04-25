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
	listaIterar *listaEnlazada[T]
	actual      *nodoLista[T]
	anterior    *nodoLista[T]
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
	for nodoAux != nil {
		elemento := nodoAux.dato
		nodoAux = nodoAux.siguiente
		if !visitar(elemento) {
			break
		}
	}

}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{listaIterar: lista}
}

func (lista *listaEnlazada[T]) VerActual() T {

}

func (lista *listaEnlazada[T]) HaySiguiente() bool {

}

func (lista *listaEnlazada[T]) Siguiente() {

}

func (lista *listaEnlazada[T]) Insertar(elemento T) {

}

func (lista *listaEnlazada[T]) Borrar() T {

}
