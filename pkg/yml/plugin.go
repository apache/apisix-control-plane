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

func (p *Plugin) ToMem() []mem.MemModel {
	result := make([]mem.MemModel, 0)
	plugins := make([]*mem.PluginSet, 0)
	for _, s := range p.Sets {
		plugin := s.ToMem()
		plugins = append(plugins, plugin)
	}
	plugin := &mem.Plugin{
		Kind:     p.Kind,
		Selector: p.Selector,
		Sets:     plugins,
	}
	result = append(result, plugin)
	return result
}

func (s *PluginSet) ToMem() *mem.PluginSet {
	return &mem.PluginSet{
		Name: s.Name,
		Conf: s.Conf,
	}
}
