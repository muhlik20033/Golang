package engineering

type EngineeringEmployee struct {
	position string
	salary   float64
	address  string
}

func (e *EngineeringEmployee) GetPosition() string {
	return e.position
}

func (e *EngineeringEmployee) SetPosition(pos string) {
	e.position = pos
}

func (e *EngineeringEmployee) GetSalary() float64 {
	return e.salary
}

func (e *EngineeringEmployee) SetSalary(sal float64) {
	e.salary = sal
}

func (e *EngineeringEmployee) GetAddress() string {
	return e.address
}

func (e *EngineeringEmployee) SetAddress(addr string) {
	e.address = addr
}
