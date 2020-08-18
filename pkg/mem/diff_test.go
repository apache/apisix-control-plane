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

package mem_test

import (
	"encoding/json"
	"fmt"
	"github.com/apache/apisix-control-plane/pkg/mem"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Diff", func() {
	Describe("compare with mems", func() {
		Context("gateway compare", func() {
			It("keep the same", func() {
				a := []byte(`{
						"kind": "Gateway",
						"name": "g1",
						"full_name": "Gateway:g1",
						"servers": [
							{
								"hosts": ["www.foo.com", "foo.com"],
								"port": {
									"number": 8080,
									"name": "http",
									"protocol": "HTTP"
								}
							},
							{
								"hosts": ["www.foo2.com", "foo2.com"],
								"port": {
									"number": 8082,
									"name": "http2",
									"protocol": "HTTP"
								}
							}
						]
					}`)
				b := []byte(`{
						"kind": "Gateway",
						"name": "g1",
						"full_name": "Gateway:g1",
						"servers": [
							{
								"hosts": ["www.foo.com", "foo.com"],
								"port": {
									"number": 8080,
									"name": "http",
									"protocol": "HTTP"
								}
							},
							{
								"hosts": ["www.foo2.com", "foo2.com"],
								"port": {
									"number": 8082,
									"name": "http2",
									"protocol": "HTTP"
								}
							}
						]
					}`)
				var ga *mem.Gateway
				if err := json.Unmarshal(a, &ga); err != nil {
					fmt.Println(err)
				}
				var gb *mem.Gateway
				if err := json.Unmarshal(b, &gb); err != nil {
					fmt.Println(err)
				}

				if result, err := mem.HasDiff(*ga, *gb); err != nil {
					fmt.Println(err)
				} else {
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(Equal(false))
				}
			})
		})

		Context("gateway compare", func() {
			It("different because of the sort of servers", func() {
				a := []byte(`{
						"kind": "Gateway",
						"name": "g1",
						"full_name": "Gateway:g1",
						"servers": [
							{
								"hosts": ["www.foo.com", "foo.com"],
								"port": {
									"number": 8080,
									"name": "http",
									"protocol": "HTTP"
								}
							},
							{
								"hosts": ["www.foo2.com", "foo2.com"],
								"port": {
									"number": 8082,
									"name": "http2",
									"protocol": "HTTP"
								}
							}
						]
					}`)
				b := []byte(`{
						"kind": "Gateway",
						"name": "g1",
						"full_name": "Gateway:g1",
						"servers": [
							{
								"hosts": ["www.foo2.com", "foo2.com"],
								"port": {
									"number": 8082,
									"name": "http2",
									"protocol": "HTTP"
								}
							},
							{
								"hosts": ["www.foo.com", "foo.com"],
								"port": {
									"number": 8080,
									"name": "http",
									"protocol": "HTTP"
								}
							}
						]
					}`)
				var ga *mem.Gateway
				if err := json.Unmarshal(a, &ga); err != nil {
					fmt.Println(err)
				}
				var gb *mem.Gateway
				if err := json.Unmarshal(b, &gb); err != nil {
					fmt.Println(err)
				}

				if result, err := mem.HasDiff(*ga, *gb); err != nil {
					fmt.Println(err)
				} else {
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(Equal(true))
				}
			})
		})
	})
})
