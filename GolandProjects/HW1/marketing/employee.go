package marketing

type MarketingEmployee struct {
	position string
	salary   float64
	address  string
}

func (e *MarketingEmployee) GetPosition() string {
	return e.position
}

func (e *MarketingEmployee) SetPosition(pos string) {
	e.position = pos
}

func (e *MarketingEmployee) GetSalary() float64 {
	return e.salary
}

func (e *MarketingEmployee) SetSalary(sal float64) {
	e.salary = sal
}

func (e *MarketingEmployee) GetAddress() string {
	return e.address
}

func (e *MarketingEmployee) SetAddress(addr string) {
	e.address = addr
}
