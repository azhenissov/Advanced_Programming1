package Company

import "fmt"

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary float64
}

func NewFullTimeEmployee(ID uint64, name string, salary float64) *FullTimeEmployee {
	return &FullTimeEmployee{
		ID:     ID,
		Name:   name,
		Salary: salary,
	}
}

func (f *FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf(
		"Full-Time Employee ID: %d | Name: %s | Salary: %.2f",
		f.ID, f.Name, f.Salary,
	)
}
