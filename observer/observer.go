package observer

type Observer interface {
	update(name string)
	getID() string
}
