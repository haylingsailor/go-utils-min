package datetime

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	provider = NewNowTimeProvider()
	t        time.Time
	err      error
)

var _ = Describe("NowTimeProvider", func() {

	It("implements interface TimeProvider", func() {
		// This will not compile if it doesn't!
		var _ TimeProvider = provider
	})

	Describe("Now", func() {

		Context("SetNow() has not been called previously", func() {

			It("returns the live system time", func() {
				tpNow := provider.Now()
				Expect(tpNow).Should(BeTemporally("~", time.Now(), time.Second))
			})
		})

		Context("SetNow() has been called previously", func() {

			BeforeEach(func() {
				t, err = time.Parse(time.RFC3339, "2015-01-01T00:00:00Z")
				if err != nil {
					Fail("Invalid Time set in test!")
				}

				provider.SetNow(t)
			})

			It("returns the fixed time", func() {

				tpNow := provider.Now()
				Expect(tpNow).Should(BeTemporally("==", t))
			})
		})
	})
	Describe("Until", func() {
		It("returns the correct duration", func() {
			tNow, _ := time.Parse(time.RFC3339, "2015-01-01T00:00:00Z")
			tFuture, _ := time.Parse(time.RFC3339, "2015-01-01T00:00:01Z")
			provider.SetNow(tNow)
			d := provider.Until(tFuture)
			Expect(d).To(Equal(time.Second))
		})
	})
	Describe("Since", func() {
		It("returns the correct duration", func() {
			tNow, _ := time.Parse(time.RFC3339, "2015-01-01T00:00:01Z")
			tPast, _ := time.Parse(time.RFC3339, "2015-01-01T00:00:00Z")
			provider.SetNow(tNow)
			d := provider.Since(tPast)
			Expect(d).To(Equal(time.Second))
		})
	})
})
