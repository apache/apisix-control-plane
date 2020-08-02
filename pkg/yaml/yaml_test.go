package yaml_test

import (
	"fmt"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yaml", func() {
	var b = make([]byte, 0)
	BeforeEach(func() {
		b = []byte(`{"name": "John", "age": 30}`)
	})
	Describe("yaml & json", func() {
		b = []byte(`{"name": "John", "age": 30}`)
		Context("json to yaml", func() {
			fmt.Println(string(b))
			y, err := yaml.JSONToYAML(b)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
			fmt.Println(string(y))
			It("json to yaml no error", func() {
				Expect(err).To(nil)
			})
		})

		Context("yaml to json", func() {
			y, _ := yaml.JSONToYAML(b)
			y2, err := yaml.YAMLToJSON(y)
			fmt.Println(string(y2))
			It("yaml to json no error", func() {
				Expect(err).To(nil)
			})
		})

	})
})
