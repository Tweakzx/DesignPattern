package composite

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e *Employee) Count() int {
	return 1
}

type Department struct {
	Name             string
	SubOrganizations []IOrganization
}

func (d *Department) Count() int {
	count := 0
	for _, org := range d.SubOrganizations {
		count += org.Count()
	}
	return count
}

func (d *Department) AddSubOrganization(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewDepartment(name string) *Department {
	return &Department{Name: name}
}
