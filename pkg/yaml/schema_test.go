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
package yaml_test

import (
	"fmt"
	localYaml "github.com/apache/apisix-control-plane/pkg/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schema", func() {
	// gateway
	Describe("Gateway schema", func() {
		var okGateway string
		BeforeEach(func() {
			okGateway = `
kind: Gateway
name: foo-gw
servers:
 - port: 80
   name: http
   protocol: HTTP
   hosts:
   - "a.foo.com"
   - "b.foo.com"
`
			fmt.Println(okGateway)
		})
		Context("Gateway schema check ok", func() {
			It("Gateway yaml is ok", func() {
				fmt.Println(okGateway)
				fmt.Println(2)
				if b, err := localYaml.ToJson(okGateway); err != nil {
					fmt.Println(err.Error())
					panic(err.Error())
				} else {
					fmt.Println(string(b))
					result, err := localYaml.Validate(string(b))
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(Equal(true))
				}
			})
		})
	})

	// rule
	Describe("Rule schema", func() {
		var okRule string
		BeforeEach(func() {
			okRule = `
kind: Rule
name: xxx-rules
hosts:
- "foo.com"
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
`
			fmt.Println(okRule)
		})
		Context("Rule schema check ok", func() {
			It("Rule yaml is ok", func() {
				fmt.Println(okRule)
				if b, err := localYaml.ToJson(okRule); err != nil {
					fmt.Println(err.Error())
					panic(err.Error())
				} else {
					fmt.Println(string(b))
					result, err := localYaml.Validate(string(b))
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(Equal(true))
				}
			})
		})
	})

	// destination
	Describe("Destination schema", func() {
		var okTarget string
		BeforeEach(func() {
			okTarget = `
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
    labels:
      tag: v2
`
			fmt.Println(okTarget)
		})
		Context("Destination schema check ok", func() {
			It("Destination yaml is ok", func() {
				fmt.Println(okTarget)
				if b, err := localYaml.ToJson(okTarget); err != nil {
					fmt.Println(err.Error())
					panic(err.Error())
				} else {
					fmt.Println(string(b))
					result, err := localYaml.Validate(string(b))
					if err != nil {
						fmt.Println(err.Error())
					}
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(Equal(true))
				}
			})
		})
	})

	// plugin
	Describe("Plugin schema", func() {
		var okPlugin string
		BeforeEach(func() {
			okPlugin = `
kind: Plugin
selector: 
  labels:
     app: foo 
sort:
- name: proxy-rewrite
  conf:
    uri: "/test/home.html"
    scheme: "http"
    host: "baidu.com"
    headers: 
      X-Api-Version: "v1"
      X-Api-Engine: "apisix"
      X-Api-useless: ""
- name: prometheus
`
			fmt.Println(okPlugin)
		})
		Context("Plugin schema check ok", func() {
			It("Plugin yaml is ok", func() {
				fmt.Println(okPlugin)
				if b, err := localYaml.ToJson(okPlugin); err != nil {
					fmt.Println(err.Error())
					panic(err.Error())
				} else {
					fmt.Println(string(b))
					result, err := localYaml.Validate(string(b))
					if err != nil {
						fmt.Println(err.Error())
					}
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(Equal(true))
				}
			})
		})
	})
})
