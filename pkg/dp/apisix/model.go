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

package apisix

type Route struct {
	ID           *string     `json:"id,omitempty"`
	Hosts        []*string   `json:"hosts,omitempty"`
	Name         *string     `json:"name,omitempty"`
	Desc         *string     `json:"desc,omitempty"`
	Uris         []*string   `json:"uris"`
	Vars         [][]*string `json:"vars,omitempty"`
	Methods      []*string   `json:"methods,omitempty"`
	ServiceId    *string     `json:"service_id,omitempty"`
	ServiceName  *string     `json:"service_name,omitempty"`
	UpstreamId   *string     `json:"upstream_id,omitempty"`
	UpstreamName *string     `json:"upstream_name,omitempty"`
	Plugins      []*Plugin   `json:"plugins,omitempty"`
}

type Plugin map[string]interface{}

type Upstream struct {
	ID              *string `json:"id,omitempty"`
	Group           *string `json:"group,omitempty"`
	ResourceVersion *string `json:"resource_version,omitempty"`
	Name            *string `json:"name,omitempty"`
	Type            *string `json:"type,omitempty"`
	HashOn          *string `json:"hash_on,omitempty"`
	Key             *string `json:"key,omitempty"`
	Nodes           []*Node `json:"nodes,omitempty"`
	FromKind        *string `json:"from_kind,omitempty"`
}

type Node struct {
	IP     *string `json:"ip,omitempty"`
	Port   *int    `json:"port,omitempty"`
	Weight *int    `json:"weight,omitempty"`
}
