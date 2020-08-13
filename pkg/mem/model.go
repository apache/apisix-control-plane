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

type Route struct {
	ID              *string   `json:"id,omitempty" yml:"id,omitempty"`
	Group           *string   `json:"group,omitempty" yml:"group,omitempty"`
	FullName        *string   `json:"full_name,omitempty" yml:"full_name,omitempty"`
	ResourceVersion *string   `json:"resource_version,omitempty" yml:"resource_version,omitempty"`
	Host            *string   `json:"host,omitempty" yml:"host,omitempty"`
	Path            *string   `json:"path,omitempty" yml:"path,omitempty"`
	Name            *string   `json:"name,omitempty" yml:"name,omitempty"`
	Methods         []*string `json:"methods,omitempty" yml:"methods,omitempty"`
	ServiceId       *string   `json:"service_id,omitempty" yml:"service_id,omitempty"`
	ServiceName     *string   `json:"service_name,omitempty" yml:"service_name,omitempty"`
	UpstreamId      *string   `json:"upstream_id,omitempty" yml:"upstream_id,omitempty"`
	UpstreamName    *string   `json:"upstream_name,omitempty" yml:"upstream_name,omitempty"`
	Plugins         []*Plugin `json:"plugins,omitempty" yml:"plugins,omitempty"`
}

type Upstream struct {
	ID              *string `json:"id,omitempty" yml:"id,omitempty"`
	FullName        *string `json:"full_name,omitempty" yml:"full_name,omitempty"`
	Group           *string `json:"group,omitempty" yml:"group,omitempty"`
	ResourceVersion *string `json:"resource_version,omitempty" yml:"resource_version,omitempty"`
	Name            *string `json:"name,omitempty" yml:"name,omitempty"`
	Type            *string `json:"type,omitempty" yml:"type,omitempty"`
	HashOn          *string `json:"hash_on,omitemtpy" yml:"hash_on,omitempty"`
	Key             *string `json:"key,omitempty" yml:"key,omitempty"`
	Nodes           []*Node `json:"nodes,omitempty" yml:"nodes,omitempty"`
	FromKind        *string `json:"from_kind,omitempty" yml:"from_kind,omitempty"`
}

type Node struct {
	IP     *string `json:"ip,omitempty" yml:"ip,omitempty"`
	Port   *int    `json:"port,omitempty" yml:"port,omitempty"`
	Weight *int    `json:"weight,omitempty" yml:"weight,omitempty"`
}

type Plugin struct {
	ID       *string        `json:"id,omitempty"`
	Selector Selector       `json:"selector"`
	Sort     []PluginSchema `json:"sort"`
}

type Selector map[string]string

type PluginSchema struct {
	Name string      `json:"name"`
	Conf interface{} `json:"conf"`
}
