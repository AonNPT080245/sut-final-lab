package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check EmployeeID invalid format", func(t *testing.T) {

		id := []string{
			"R123456789", // ผิดตัวหน้า
			"J123456789", // เกินมา 1
			"J1234567",   // ขาด 1
			"JH234567",   // หลังตัวหน้ามีสตริง
			"J1234V678",  // หลังตัวหน้ามีสตริง

			"R123456789", // ผิดตัวหน้า
			"M123456789", // เกินมา 1
			"M1234567",   // ขาด 1
			"MH234567",   // หลังตัวหน้ามีสตริง
			"M1234V678",  // หลังตัวหน้ามีสตริง

			"R123456789", // ผิดตัวหน้า
			"S123456789", // เกินมา 1
			"S1234567",   // ขาด 1
			"SH234567",   // หลังตัวหน้ามีสตริง
			"S1234V678",  // หลังตัวหน้ามีสตริง
		}

		for _, val := range id {
			emp := Employee{
				Name:       "Aon",
				Email:      "Aon12@mail.com",
				EmployeeID: val,
			}
			ok, err := govalidator.ValidateStruct(emp)
			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("EmployeeID invalid format"))
		}

	})

}
