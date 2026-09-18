package chans

type Event struct {
	Mensaje string
}

type Subject struct {
	suscriptores []chan Event
}

func (s *Subject) Suscribir() <-chan Event {
	ch := make(chan Event)
	s.suscriptores = append(s.suscriptores, ch)
	return ch
}

func (s *Subject) Notificar(e Event) {
	for _, ch := range s.suscriptores {
		ch <- e // Envía la notificación
	}
}
