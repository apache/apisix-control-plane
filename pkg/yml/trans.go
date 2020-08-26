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

import (
	"encoding/json"

	"gopkg.in/yaml.v2"

	"github.com/apache/apisix-control-plane/pkg/log"
)

func Trans(b []byte, y []byte) YmlModel {
	// 1.trans with kind
	var yMap map[string]interface{}
	if err := json.Unmarshal(b, &yMap); err != nil {
		log.Errorw("unmarshal json failed", "err", err, "raw", string(b))
		return nil
	} else {
		kind := yMap["kind"]
		switch kind {
		case "Gateway":
			// trans to Gateway
			if g, err := ToGateway(y); err != nil {
				return nil
			} else {
				return g
			}
		case "Destination":
			// trans to Destination
			if g, err := ToDestination(y); err != nil {
				return nil
			} else {
				return g
			}
		case "Rule":
			// trans to Rule
			if r, err := ToRule(y); err != nil {
				return nil
			} else {
				return r
			}
		case "Plugin":
			// trans to Plugin
			if r, err := ToPlugin(y); err != nil {
				return nil
			} else {
				return r
			}
		default:
			log.Warnw("unknown kind", "kind", kind)
			return nil
		}
	}
}

func ToGateway(y []byte) (*Gateway, error) {
	var g *Gateway
	if err := yaml.Unmarshal(y, &g); err != nil {
		log.Errorw("unmarshal yaml failed", "err", err, "raw", string(y))
		return nil, err
	} else {
		return g, nil
	}
}

func ToRule(y []byte) (*Rule, error) {
	var r *Rule
	if err := yaml.Unmarshal(y, &r); err != nil {
		log.Errorw("unmarshal yaml failed", "err", err, "raw", string(y))
		return nil, err
	} else {
		return r, nil
	}
}

func ToDestination(y []byte) (*Destination, error) {
	var g *Destination
	if err := yaml.Unmarshal(y, &g); err != nil {
		log.Errorw("unmarshal yaml failed", "err", err, "raw", string(y))
		return nil, err
	} else {
		return g, nil
	}
}

func ToPlugin(y []byte) (*Plugin, error) {
	var g *Plugin
	if err := yaml.Unmarshal(y, &g); err != nil {
		log.Errorw("unmarshal yaml failed", "err", err, "raw", string(y))
		return nil, err
	} else {
		return g, nil
	}
}
