package composition

type engine struct {
	hp int
}

type wheel struct {
	wheelDimensions int
}

type Car struct {
	CarName string
	engine
	wheel
}

func NewCar(carName string, hp, wd int) Car {
	return Car{
		CarName: carName,
		engine: engine{
			hp: hp,
		},
		wheel: wheel{
			wheelDimensions: wd,
		},
	}
}

func (e engine) Hp() int {
	return e.hp
}

func (w wheel) WheelDimensions() int {
	return w.wheelDimensions
}
