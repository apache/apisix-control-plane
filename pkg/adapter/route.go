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

package adapter

import (
	"reflect"

	"github.com/apache/apisix-control-plane/pkg/dp/apisix"
	"github.com/apache/apisix-control-plane/pkg/mem"
)

func ToRoute(r *mem.Route) *apisix.Route {
	return &apisix.Route{
		Name:    r.FullName,
		Hosts:   r.Hosts,
		Desc:    r.Name,
		Uris:    toUris(r),
		Vars:    toVars(r),
		Methods: r.Methods,
	}
}

func toUris(r *mem.Route) []*string {
	result := make([]*string, 0)
	for _, m := range r.Match {
		// uris
		matchUris := m["uris"]
		switch reflect.TypeOf(matchUris).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(matchUris)
			if slice, ok := s.Interface().([]map[string]string); ok {
				for _, s := range slice {
					for k, v := range s {
						switch k {
						case "prefix":
							uri := v + "*"
							result = append(result, &uri)
						case "exact":
							result = append(result, &v)
						}
					}
				}
			}
		}
	}
	return result
}

func toVars(r *mem.Route) [][]*string {
	result := make([][]*string, 0)
	for _, m := range r.Match {
		// args
		matchArgs := m["args"]
		switch reflect.TypeOf(matchArgs).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(matchArgs)
			if slice, ok := s.Interface().([]map[string]map[string]string); ok {
				for _, s := range slice {
					tmp := make([]*string, 0)
					for p, value := range s {
						tmp = append(tmp, &p)
						for k, v := range value {
							switch k {
							case "greater":
								greater := ">"
								tmp = append(tmp, &greater)
								tmp = append(tmp, &v)
							case "exact":
								equals := "=="
								tmp = append(tmp, &equals)
								tmp = append(tmp, &v)
							}
						}
						result = append(result, tmp)
					}

				}
			}
		}
	}
	return result
}
