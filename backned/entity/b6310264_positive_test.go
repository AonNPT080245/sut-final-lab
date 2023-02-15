package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployee(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check input correct", func(t *testing.T) {

		emp := Employee{
			Name:       "Aon",
			Email:      "Aon12@mail.com",
			EmployeeID: "J12345678",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})

}
