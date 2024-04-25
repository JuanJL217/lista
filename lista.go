package lista

type Lista[T any] interface {

	//EstaVacia: Devuelve verdadero si la lista no tiene elementos. false caso contrario
	EstaVacia() bool

	//InsertaPrimero: Inserta el elemento en la primera posicion de la lista
	InsertarPrimero(T)

	//InsertarUltimo: Inserta el elemento en la ultima posicion de la lista.
	InsertarUltimo(T)

	//BorrarPrimero: Borra el primer elemento de la lista y devuelve su elemento , Si la lista esta vacia,
	//entrara en panico con el siguiente mensaje "la lista esta vacia"
	BorrarPrimero() T

	//VerPrimero: Obtiene el primer elemento de la lista, Si la lista esta vacia, entrara en panico
	//con el siguiente mensaje "la lista esta vacia"
	VerPrimero() T

	//VerUltimo: Obtiene el ultimo elemento de la lista. Si la lista esta vacia, entrara en panico
	//con el siguiente mensaje "la lista esta vacia"
	VerUltimo() T

	//Largo: Obtiene el largo de la lista
	Largo() int

	//Iterar: Iterador interno, itera hasta que la funcion visitar reciba un "false" o la lista quede vacia.
	Iterar(visitar func(T) bool)

	//Iteador: Iterador externo, devuelve un IteradorLista con la cual vamos a poder iterar esta Lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//VerActual: Obtiene el elemento de la posicion actual donde este el iterador.
	VerActual() T

	//HaySiguiente: Devuelve verdadero si hay algo existente en la siguiente posicion, false caso contrario.
	HaySiguiente() bool

	//Siguiente: Avanza a la siguiente posicion de donde esté actualmente
	Siguiente()

	//Insertar: Agregará un nuevo elemento a la lista en dicha posicion del iterador
	Insertar(T)

	//Borrar: Quitará y devolverá el elemento que esté en esa posicion del iterador
	Borrar() T
}
