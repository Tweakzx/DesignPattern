package composite

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepartMent_Count(t *testing.T) {
	departmentA := NewDepartment("A")
	for i := 0; i < 10; i++ {
		departmentA.AddSubOrganization(&Employee{Name: "A" + strconv.Itoa(i)})
	}
	departmentB := NewDepartment("B")
	for i := 0; i < 10; i++ {
		departmentB.AddSubOrganization(&Employee{Name: "B" + strconv.Itoa(i)})
	}
	departmentB.AddSubOrganization(departmentA)

	departmentC := NewDepartment("C")
	departmentC.AddSubOrganization(departmentB)

	assert.Equal(t, 10, departmentA.Count())
	assert.Equal(t, 20, departmentB.Count())
	assert.Equal(t, 20, departmentC.Count())
}
