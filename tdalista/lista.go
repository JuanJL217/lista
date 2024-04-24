package lista

type Lista[T any] interface {
	//EstaVacia: Devuelve verdadero si la lista no tiene elementos. false caso contrario
	EstaVacia() bool

	//InsertaPrimero: Inserta el elemento en la primera posicion de la lista
	InsertarPrimero(T)

	//InsertarUltimo: Inserta el dato en el ultimo lugar de la lista.
	InsertarUltimo(T)

	//BorrarPrimero: Borra el primer elemento de la lista y devuelve su elemento , Si la lista esta vacia, entrara en panico con el siguiente mensaje
	//"la lista esta vacia"
	BorrarPrimero() T

	//VerPrimero : Obtiene el primer elemento de la lista, Si la lista esta vacia, entrara en panico con el siguiente mensaje
	//"la lista esta vacia"
	VerPrimero() T

	//VerUltimo: Obtiene el ultimo elemento de la lista. Si la lista esta vacia, entrara en panico con el siguiente mensaje
	//"la lista esta vacia"
	VerUltimo() T
	//Largo: Obtiene el largo de la lista
	Largo() int

	Iterar(visitar func(T) bool)
	Iterador() IteradorLista[T]
}
