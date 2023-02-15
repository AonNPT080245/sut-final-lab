package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameIsblank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check Name cannot blank", func(t *testing.T) {

		emp := Employee{
			Name:       "",
			Email:      "Aon12@mail.com",
			EmployeeID: "J12345678",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot blank"))

	})

}
