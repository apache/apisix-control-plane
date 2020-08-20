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

func Map(raws []MemModel) map[string][]MemModel {
	result := make(map[string][]MemModel)
	result[GatewayKind] = make([]MemModel, 0)
	result[RouteKind] = make([]MemModel, 0)
	result[UpstreamKind] = make([]MemModel, 0)
	result[PluginKind] = make([]MemModel, 0)
	for _, r := range raws {
		switch v := r.(type) {
		case *Gateway:
			result[*v.Kind] = append(result[*v.Kind], v)
		case *Route:
			result[*v.Kind] = append(result[*v.Kind], v)
		case *Upstream:
			result[*v.Kind] = append(result[*v.Kind], v)
		case *Plugin:
			result[*v.Kind] = append(result[*v.Kind], v)
		}
	}
	return result
}
