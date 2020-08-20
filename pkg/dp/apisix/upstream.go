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
	"github.com/apache/apisix-control-plane/pkg/utils"
	"github.com/gxthrj/seven/conf"
)

func (upstream *Upstream) Add() (*UpstreamResponse, error) {
	baseUrl := conf.BaseUrl
	url := fmt.Sprintf("%s/upstreams", baseUrl)
	if b, err := json.Marshal(upstream); err != nil {
		return nil, err
	} else {
		if res, err := utils.Post(url, b); err != nil {
			return nil, fmt.Errorf("http post failed, url: %s, err: %+v", url, err)
		} else {
			var uRes UpstreamResponse
			if err = json.Unmarshal(res, &uRes); err != nil {
				return nil, err
			} else {
				if uRes.Upstream.Key != nil {
					return &uRes, nil
				} else {
					return nil, fmt.Errorf("apisix upstream not expected response")
				}
			}
		}
	}
}

func (upstream *Upstream) Update() (*UpstreamResponse, error) {
	// todo
	return nil, nil
}

func (upstream *Upstream) Delete() (*UpstreamResponse, error) {
	// todo
	return nil, nil
}

type UpstreamResponse struct {
	Action   string       `json:"action"`
	Upstream UpstreamNode `json:"node"`
}

type UpstreamNode struct {
	Key           *string       `json:"key"`
	UpstreamNodes UpstreamNodes `json:"value"`
}

type UpstreamNodes struct {
	Nodes  map[string]int64 `json:"nodes"`
	Desc   *string          `json:"desc"`
	LBType *string          `json:"type"`
}
