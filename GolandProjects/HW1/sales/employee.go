package sales

type SalesEmployee struct {
	position string
	salary   float64
	address  string
}

func (e *SalesEmployee) GetPosition() string {
	return e.position
}

func (e *SalesEmployee) SetPosition(pos string) {
	e.position = pos
}

func (e *SalesEmployee) GetSalary() float64 {
	return e.salary
}

func (e *SalesEmployee) SetSalary(sal float64) {
	e.salary = sal
}

func (e *SalesEmployee) GetAddress() string {
	return e.address
}

func (e *SalesEmployee) SetAddress(addr string) {
	e.address = addr
}
