package finance

type FinanceEmployee struct {
	position string
	salary   float64
	address  string
}

func (e *FinanceEmployee) GetPosition() string {
	return e.position
}

func (e *FinanceEmployee) SetPosition(pos string) {
	e.position = pos
}

func (e *FinanceEmployee) GetSalary() float64 {
	return e.salary
}

func (e *FinanceEmployee) SetSalary(sal float64) {
	e.salary = sal
}

func (e *FinanceEmployee) GetAddress() string {
	return e.address
}

func (e *FinanceEmployee) SetAddress(addr string) {
	e.address = addr
}
