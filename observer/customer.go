package observer

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) update(name string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, name)
}

func (c *Customer) getID() string {
	return c.Id
}
