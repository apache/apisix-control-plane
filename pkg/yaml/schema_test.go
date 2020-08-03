package yaml_test

import (
	"fmt"
	localYaml "github.com/apache/apisix-control-plane/pkg/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schema", func() {
	Describe("Gateway schema", func() {
		var okGateway string
		BeforeEach(func() {
			okGateway = `kind: Gateway
name: baidu-gw
servers:
 - port: 80
   name: http
   protocol: HTTP
   hosts:
   - "a.domain.com"
   - "b.domain.com"`
			fmt.Println(okGateway)
		})
		Context("Gateway schema check ok", func() {
			It("Gateway yaml is ok", func() {
				fmt.Println(okGateway)
				fmt.Println(2)
				if b, err := localYaml.ToJson(okGateway); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(string(b))
					result, err := localYaml.Validate(string(b))
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(Equal(true))
				}
			})
		})
	})
})
