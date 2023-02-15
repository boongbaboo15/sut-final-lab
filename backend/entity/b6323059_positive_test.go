package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeePass(t *testing.T) { //Pass
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

func TestNameNull(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "",
		Email:      "test@gmail.com",
		EmployeeID: "J12345678",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))
}
