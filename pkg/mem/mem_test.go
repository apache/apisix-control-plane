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

var _ = Describe("Mem", func() {
	Describe("mem object store", func() {
		Context("gateway store", func() {
			It("inert and find", func() {
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
						}
					]
				}`)
				var g *mem.Gateway
				if err := json.Unmarshal(b, &g); err != nil {
					fmt.Println(err)
				}
				gatewayDB := &mem.GatewayDB{Gateways: []*mem.Gateway{g}}
				if err := gatewayDB.Insert(); err != nil {
					fmt.Println(err)
				}
				if gFromDB, err := g.FindByFullName(); err != nil {
					fmt.Println(err)
				} else {
					Expect(gFromDB).NotTo(Equal(nil))
					Expect(*gFromDB.FullName).To(Equal("Gateway:g1"))
				}
			})
		})
		Context("route store", func() {
			It("inert and find", func() {
				b := []byte(`{
					"kind": "Route",
					"name": "r1",
					"full_name": "Route:r1",
					"hosts": ["foo-server", "foo2-com"],
					"match": [
					{
						"args": {
							"name": {
								"exact": "user"
							},
							"age": {
								"greater": 18
							}
						}
					},{
						"uris": [{"prefix": "/"}]
					}]
				}`)
				var g *mem.Route
				if err := json.Unmarshal(b, &g); err != nil {
					fmt.Println(err)
				}
				routeDB := &mem.RouteDB{Routes: []*mem.Route{g}}
				if err := routeDB.Insert(); err != nil {
					fmt.Println(err)
				}
				if gFromDB, err := g.FindByFullName(); err != nil {
					fmt.Println(err)
				} else {
					Expect(gFromDB).NotTo(Equal(nil))
					Expect(*gFromDB.FullName).To(Equal("Route:r1"))
				}
			})
		})

		Context("upstream store", func() {
			It("inert and find", func() {
				b := []byte(`{
					"kind": "Upstream",
					"name": "u1",
					"full_name": "Upstream:u1",
					"host": "foo-server",
					"group": "foo",
					"type": "Roundrobin",
					"nodes": [
						{"ip": "127.0.0.1", "port": 80, "weight": 100},
						{"ip": "127.0.0.2", "port": 80, "weight": 100}
					]
				}`)
				var g *mem.Upstream
				if err := json.Unmarshal(b, &g); err != nil {
					fmt.Println(err)
				}
				upstreamDB := &mem.UpstreamDB{Upstreams: []*mem.Upstream{g}}
				if err := upstreamDB.Insert(); err != nil {
					fmt.Println(err)
				}
				if gFromDB, err := g.FindByFullName(); err != nil {
					fmt.Println(err)
				} else {
					Expect(gFromDB).NotTo(Equal(nil))
					Expect(*gFromDB.FullName).To(Equal("Upstream:u1"))
					Expect(len(gFromDB.Nodes)).To(Equal(2))
				}
			})
		})

		Context("Plugin store", func() {
			It("inert and find", func() {
				b := []byte(`{
					"kind": "Plugin",
					"name": "p1",
					"full_name": "Plugin:p1",
					"selector": {
						"app": "foo-server"
					},
					"sets": [
						{
							"name": "proxy-rewrite",
							"conf": {
								"uri": "/test/home.html",
								"scheme": "http",
								"host": "foo.com",
								"headers": {
									"X-Api-Version:": "v1",
									"X-Api-Engine": "apisix",
									"X-Api-useless": ""
								}
							}
						},
						{
							"name": "prometheus"
						}
					]
				}`)
				var g *mem.Plugin
				if err := json.Unmarshal(b, &g); err != nil {
					fmt.Println(err)
				}
				pluginDB := &mem.PluginDB{Plugins: []*mem.Plugin{g}}
				if err := pluginDB.Insert(); err != nil {
					fmt.Println(err)
				}
				if gFromDB, err := g.FindByFullName(); err != nil {
					fmt.Println(err)
				} else {
					Expect(gFromDB).NotTo(Equal(nil))
					Expect(*gFromDB.FullName).To(Equal("Plugin:p1"))
					Expect(len(gFromDB.Sets)).To(Equal(2))
					Expect(gFromDB.Selector["app"]).To(Equal("foo-server"))
				}
			})
		})
	})
})
