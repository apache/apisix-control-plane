package yaml_test

import (
	"fmt"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yaml", func() {
	Describe("yaml & json", func() {
		var b []byte
		BeforeEach(func() {
			b = []byte(`{"name": "John", "age": 30}`)
			fmt.Println("BeforeEach executed")
		})
		Context("json to yaml", func() {
			It("json to yaml no error", func() {
				fmt.Println(string(b))
				y, err := yaml.JSONToYAML(b)
				fmt.Println(string(y))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("yaml to json", func() {

			It("yaml to json no error", func() {
				fmt.Println(5)
				y, _ := yaml.JSONToYAML(b)
				y2, err := yaml.YAMLToJSON(y)
				fmt.Println(string(y2))
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})
})
