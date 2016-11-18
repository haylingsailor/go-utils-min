package datetime

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ITime", func() {

	var (
		nanoStr = "2012-10-31T16:13:58.292387Z"
		t       time.Time
		sut     ITime
		err     error
	)
	BeforeEach(func() {
		t, err = time.Parse(time.RFC3339Nano, nanoStr)
		Expect(err).To(BeNil())

	})
	Describe("MarshalJSON", func() {
		It("Creates the expected JSON", func() {
			sut = ITime(t)
			result, err := sut.MarshalJSON()
			Expect(err).To(BeNil())
			Expect(string(result)).To(Equal(`"2012-10-31T16:13:58Z"`))
		})
	})
})
