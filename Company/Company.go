package Company

type Company struct {
	employees map[uint64]Employee
}

func NewCompany() *Company {
	return &Company{
		employees: make(map[uint64]Employee),
	}
}

func (c *Company) AddEmployee(id uint64, emp Employee) {
	c.employees[id] = emp
}

func (c *Company) ListEmployees() []string {
	var list []string
	for _, emp := range c.employees {
		list = append(list, emp.GetDetails())
	}
	return list
}
