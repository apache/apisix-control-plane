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

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yaml", func() {
	Describe("yml & json", func() {
		var b []byte
		BeforeEach(func() {
			b = []byte(`{"name": "John", "age": 30}`)
			fmt.Println("BeforeEach executed")
		})
		Context("json to yml", func() {
			It("json to yml no error", func() {
				fmt.Println(string(b))
				y, err := yaml.JSONToYAML(b)
				fmt.Println(string(y))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("yml to json", func() {

			It("yml to json no error", func() {
				fmt.Println(5)
				y, _ := yaml.JSONToYAML(b)
				y2, err := yaml.YAMLToJSON(y)
				fmt.Println(string(y2))
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})
})
