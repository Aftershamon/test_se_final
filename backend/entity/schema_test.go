package entity

import (
	"testing"
	//"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameCanNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	name := Video{
		Name: "", //ผิดddd
	}

	ok, err := govalidator.ValidateStruct(name)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name not blank"))

}

func TestMaxReason(t *testing.T){
	g := NewGomegaWithT(t)

	reason := Video{
		Name: "JaJar",
		Reason: "kkkkkk", //ผิด
	}
	ok, err := govalidator.ValidateStruct(reason)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Max 5"))
}

func TestEmail (t *testing.T){
	g := NewGomegaWithT(t)

	email := Video{
		Name: "Ja",
		Reason: "okk",
		Email: "kkk", //
	}
	ok,err := govalidator.ValidateStruct(email)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ไม่ถูกจ้า"))
}

// func TestTimeStart (t *testing.T){
// 	g := NewGomegaWithT(t)

// 	start := Video{
// 		Name: "Ja",
// 		Reason: "ok",
// 		Email: "kkk@gmail.com", 
// 		Start: time.Now(),//
// 	}
// 	ok,err := govalidator.ValidateStruct(start)
// 	g.Expect(ok).ToNot(BeTrue())
// 	g.Expect(err).ToNot(BeNil())
// 	g.Expect(err.Error()).To(Equal("เวลาห้ามเป็นอดีต"))
// }