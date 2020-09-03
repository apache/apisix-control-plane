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

import (
	"encoding/json"
	"fmt"
	"github.com/apache/apisix-control-plane/pkg/conf"
	"github.com/apache/apisix-control-plane/pkg/utils"
)

func (route *Route) Add() (*RouteResponse, error) {
	baseUrl := conf.BaseUrl
	url := fmt.Sprintf("%s/routes", baseUrl)
	if b, err := json.Marshal(route); err != nil {
		return nil, err
	} else {
		if res, err := utils.Post(url, b); err != nil {
			return nil, err
		} else {
			var routeResp RouteResponse
			if err = json.Unmarshal(res, &routeResp); err != nil {
				return nil, err
			} else {
				if routeResp.Route.Key != nil {
					return &routeResp, nil
				} else {
					return nil, fmt.Errorf("apisix route not expected response")
				}

			}
		}
	}
}

func (route *Route) Update() (*RouteResponse, error) {
	// todo
	return nil, nil
}

func (route *Route) Delete() (*RouteResponse, error) {
	// todo
	return nil, nil
}

type RouteResponse struct {
	Action string    `json:"action"`
	Route  RouteNode `json:"node"`
}

type RouteNode struct {
	Key   *string `json:"key"`
	Value Value   `json:"value"`
}

type Value struct {
	UpstreamId *string                `json:"upstream_id"`
	ServiceId  *string                `json:"service_id"`
	Plugins    map[string]interface{} `json:"plugins"`
	Host       *string                `json:"host,omitempty"`
	Uri        *string                `json:"uri"`
	Desc       *string                `json:"desc"`
	Methods    []*string              `json:"methods,omitempty"`
}
