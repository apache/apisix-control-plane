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

package mem

const GatewayKind = "Gateway"
const RouteKind = "Route"
const UpstreamKind = "Upstream"
const PluginKind = "Plugin"

type MemModel interface {
	Diff(m MemModel) bool
}

type Route struct {
	ID           *string                  `json:"id,omitempty" yml:"id,omitempty"`
	Kind         *string                  `json:"kind"`
	FullName     *string                  `json:"full_name,omitempty" yml:"full_name,omitempty"`
	Hosts        []*string                `json:"hosts,omitempty" yml:"hosts,omitempty"`
	Match        []map[string]interface{} `json:"paths,omitempty" yml:"paths,omitempty"`
	Name         *string                  `json:"name,omitempty" yml:"name,omitempty"`
	Methods      []*string                `json:"methods,omitempty" yml:"methods,omitempty"`
	ServiceId    *string                  `json:"service_id,omitempty" yml:"service_id,omitempty"`
	ServiceName  *string                  `json:"service_name,omitempty" yml:"service_name,omitempty"`
	UpstreamId   *string                  `json:"upstream_id,omitempty" yml:"upstream_id,omitempty"`
	UpstreamName *string                  `json:"upstream_name,omitempty" yml:"upstream_name,omitempty"`
	Plugins      []*Plugin                `json:"plugins,omitempty" yml:"plugins,omitempty"`
}

type Upstream struct {
	ID              *string `json:"id,omitempty" yml:"id,omitempty"`
	Kind            *string `json:"kind"`
	Host            *string `json:"host"`
	FullName        *string `json:"full_name,omitempty" yml:"full_name,omitempty"`
	Group           *string `json:"group,omitempty" yml:"group,omitempty"`
	ResourceVersion *string `json:"resource_version,omitempty" yml:"resource_version,omitempty"`
	Name            *string `json:"name,omitempty" yml:"name,omitempty"`
	Type            *string `json:"type,omitempty" yml:"type,omitempty"`
	HashOn          *string `json:"hash_on,omitempty" yml:"hash_on,omitempty"`
	Key             *string `json:"key,omitempty" yml:"key,omitempty"`
	Nodes           []*Node `json:"nodes,omitempty" yml:"nodes,omitempty"`
	Weight          int64   `json:"weight"`
	FromKind        *string `json:"from_kind,omitempty" yml:"from_kind,omitempty"`
}

type Node struct {
	IP     *string `json:"ip,omitempty" yml:"ip,omitempty"`
	Port   *int    `json:"port,omitempty" yml:"port,omitempty"`
	Weight *int    `json:"weight,omitempty" yml:"weight,omitempty"`
}

type Plugin struct {
	ID       *string           `json:"id,omitempty"`
	Kind     *string           `json:"kind"`
	Name     *string           `json:"name"`
	FullName *string           `json:"full_name"`
	Selector map[string]string `json:"selector"`
	Sets     []*PluginSet      `json:"sets"`
}

type PluginSet struct {
	Name *string                `json:"name"`
	Conf map[string]interface{} `json:"conf"`
}

type Gateway struct {
	ID       *string   `json:"id,omitempty"`
	FullName *string   `json:"full_name,omitempty" yml:"full_name,omitempty"`
	Kind     *string   `json:"kind"`
	Name     *string   `json:"name"`
	Servers  []*Server `json:"servers"`
}

type Server struct {
	Port  *Port    `json:"port,omitempty"`
	Hosts []string `json:"hosts,omitempty"`
}

type Port struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
}
