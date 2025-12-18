package Company

import "fmt"

type PartTimeEmployee struct {
	ID    uint64
	Name  string
	Hours float64
}

func NewPartTimeEmployee(ID uint64, name string, hours float64) *PartTimeEmployee {
	return &PartTimeEmployee{
		ID:    ID,
		Name:  name,
		Hours: hours,
	}
}

func (p *PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf(
		"Part-Time Employee ID: %d | Name: %s | Hours: %.1f",
		p.ID, p.Name, p.Hours,
	)
}
