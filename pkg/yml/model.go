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

package yml

import "github.com/apache/apisix-control-plane/pkg/mem"

const seprator = ":"

type YmlModel interface {
	ToMem() []mem.MemModel
}

type Gateway struct {
	Kind    *string   `json:"kind"`
	Name    *string   `json:"name"`
	Servers []*Server `json:"servers"`
}

type Server struct {
	Port  *Port    `json:"port,omitempty"`
	Hosts []string `json:"host,omitempty"`
}

type Port struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
}

type Rule struct {
	Kind     *string   `json:"kind"`
	Name     *string   `json:"name"`
	Hosts    []*string `json:"hosts"`
	Gateways []*string `json:"gateways"`
	HTTP     []*HTTP   `json:"http"`
}
type RouteDestination struct {
	Port   int64  `json:"port"`
	Host   string `json:"host"`
	Subset string `json:"subset"`
	Weight int64  `json:"weight"`
}
type Route struct {
	Destination RouteDestination `json:"destination"`
}
type Label map[string]string

type Headers map[string]interface{}

type Match struct {
	Headers Headers `json:"headers"`
	Uris    []*Uri  `json:"uris"`
}

type Uri struct {
	Prefix  *string `json:"prefix,omitempty"`
	Exact   *string `json:"exact,omitempty"`
	Regular *string `json:"regular,omitempty"`
}

type HTTP struct {
	Route []*Route                 `json:"route"`
	Label *Label                   `json:"label"`
	Match []map[string]interface{} `json:"match,omitempty"`
}

type Destination struct {
	Kind    *string   `json:"kind"`
	Name    *string   `json:"name"`
	Host    *string   `json:"host"`
	Subsets []*Subset `json:"subsets"`
}

type Subset struct {
	Name     *string           `json:"name"`
	Ips      []*string         `json:"ips,omitempty"`
	Selector map[string]string `json:"selector,omitempty"`
	Weight   int64             `json:"weight"`
}

type Plugin struct {
	Kind     *string           `yaml:"kind"`
	Selector map[string]string `yaml:"selector"`
	Sets     []*PluginSet      `yaml:"sets"`
}

type PluginSet struct {
	Name *string                `yaml:"name"`
	Conf map[string]interface{} `yaml:"conf,omitempty"`
}
