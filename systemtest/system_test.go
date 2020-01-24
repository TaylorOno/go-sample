package system_test

import (
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the webserver", func() {
	It("responds to Post with 200", func() {
		//pkgPath := "github.com/rosenhouse/counter-demo"
		url := "http://localhost:8000"

		resp, err := http.Post(url, "application/json", strings.NewReader("{}"))
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(200))
	})
})
