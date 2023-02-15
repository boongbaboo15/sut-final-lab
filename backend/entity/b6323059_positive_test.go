package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeePass(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "test",
		Email:      "test@gmail.com",
		EmployeeID: "J12345678",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
