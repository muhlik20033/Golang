package hr

type HREmployee struct {
	position string
	salary   float64
	address  string
}

func (e *HREmployee) GetPosition() string {
	return e.position
}

func (e *HREmployee) SetPosition(pos string) {
	e.position = pos
}

func (e *HREmployee) GetSalary() float64 {
	return e.salary
}

func (e *HREmployee) SetSalary(sal float64) {
	e.salary = sal
}

func (e *HREmployee) GetAddress() string {
	return e.address
}

func (e *HREmployee) SetAddress(addr string) {
	e.address = addr
}
