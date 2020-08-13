package yml_test

import (
	"fmt"
	"github.com/apache/apisix-control-plane/pkg/yml"
	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trans", func() {
	Describe("trans to model", func() {
		var b []byte
		BeforeEach(func() {
			b = []byte(`
kind: Gateway
name: foo-gw
servers:
 - port:
     number: 80
     name: http
     protocol: HTTP
   hosts:
   - "a.foo.com"
   - "b.foo.com"
`)
			fmt.Println("BeforeEach executed")
		})
		Context("trans", func() {
			It("trans to gateway no error", func() {
				y, err := yaml.YAMLToJSON(b)
				fmt.Println(string(y))
				ym := yml.Trans(y, b)
				Expect(err).NotTo(HaveOccurred())
				Expect(ym.Type()).To(Equal("Gateway"))
				g, ok := ym.(*yml.Gateway)
				Expect(ok).To(Equal(true))
				Expect(len(g.Servers[0].Hosts)).To(Equal(2))
			})
		})
	})
})
