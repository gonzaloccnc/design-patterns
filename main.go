package main

import (
	"design-patterns/observer"
	"design-patterns/observer/chans"
	"fmt"
	"sync"
)

func main() {
	shirtItem := observer.NewItem("Nike Shirt")

	observerFirst := &observer.Customer{Id: "abc@gmail.com"}
	observerSecond := &observer.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()

	subject := &chans.Subject{}
	var wg sync.WaitGroup

	// Número de suscriptores
	numSuscriptores := 5

	// Crear y registrar 5 suscriptores
	for i := 1; i <= numSuscriptores; i++ {
		id := i
		sub := subject.Suscribir()

		wg.Go(
			func() {
				for e := range sub {
					fmt.Printf("Suscriptor %d recibió: %s\n", id, e.Mensaje)
					wg.Done()
				}
			},
		)
	}

	// Notificar a todos los suscriptores
	subject.Notificar(chans.Event{Mensaje: "¡Hola Observers!"})

	// Esperar a que los 5 suscriptores terminen de procesar el mensaje
	wg.Wait()
}
