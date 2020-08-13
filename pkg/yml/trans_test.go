/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
		Context("trans", func() {
			It("trans to gateway no error", func() {
				b := []byte(`
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
				y, err := yaml.YAMLToJSON(b)
				fmt.Println(string(y))
				ym := yml.Trans(y, b)
				Expect(err).NotTo(HaveOccurred())
				v := typeof(ym)
				Expect(v).To(Equal("*yml.Gateway"))
				g, ok := ym.(*yml.Gateway)
				Expect(ok).To(Equal(true))
				Expect(len(g.Servers[0].Hosts)).To(Equal(2))
			})

			It("trans to rule no error", func() {
				b := []byte(`
kind: Rule
name: xxx-rules
hosts:
- "a.foo.com"
gateways:
- foo-gw
http:
- route:
  - destination:
     port: 28002
     host: foo-server
     subset: foo-v1
     weight: 10
  label:
    app: foo
    version: v1
  match:
  - headers:
     product_id:
       exact: v1
- route:
  - destination:
       port: 28002
       host: foo-server
       subset: v2
  label:
    app: foo
    version: v2
`)

				y, err := yaml.YAMLToJSON(b)
				fmt.Println(string(y))
				ym := yml.Trans(y, b)
				Expect(err).NotTo(HaveOccurred())
				v := typeof(ym)
				Expect(v).To(Equal("*yml.Rule"))
				r, ok := ym.(*yml.Rule)
				Expect(ok).To(Equal(true))
				Expect(r.Kind).To(Equal("Rule"))
				Expect(r.Kind).To(Equal("Rule"))
			})

			It("trans to destination no error", func() {
				b := []byte(`
kind: Destination
name: foo-dest
host: foo-server
subsets:
- name: foo-v1
  ips:
  - 127.0.0.1
  - 127.0.0.2
- name: v2
  selector:
    tag: v2
`)
				y, err := yaml.YAMLToJSON(b)
				fmt.Println(string(y))
				ym := yml.Trans(y, b)
				Expect(err).NotTo(HaveOccurred())
				v := typeof(ym)
				Expect(v).To(Equal("*yml.Destination"))
				g, ok := ym.(*yml.Destination)
				Expect(ok).To(Equal(true))
				Expect(g.Kind).To(Equal("Destination"))
				Expect(g.Host).To(Equal("foo-server"))
			})
		})
	})
})

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
