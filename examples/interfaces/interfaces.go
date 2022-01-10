package interfaces

import "fmt"

type IF interface {
	GetName() (string, error)
}

type Car struct {
	Name string
}

func (c *Car) GetName() (string, error) {
	return c.Name, nil
}

type Boat struct {
	Name string
}

func (b *Boat) GetName() (string, error) {
	return b.Name, nil
}

func echo_interface() {
	interfaces := []IF{}

	car := new(Car)
	car.Name = "car"
	interfaces = append(interfaces, car)

	boat := new(Boat)
	boat.Name = "boat"
	interfaces = append(interfaces, boat)

	for index, _ := range interfaces {
		fmt.Println(interfaces[index].GetName())
	}

}
