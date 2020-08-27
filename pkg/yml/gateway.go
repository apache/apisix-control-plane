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

func (g *Gateway) ToMem() []mem.MemModel {
	result := make([]mem.MemModel, 0)
	fullName := *g.Kind + separator + *g.Name
	servers := make([]*mem.Server, 0)
	for _, e := range g.Servers {
		server := e.ToMem()
		servers = append(servers, server)
	}
	gateway := &mem.Gateway{
		FullName: &fullName,
		Kind:     g.Kind,
		Name:     g.Name,
		Servers:  servers,
	}
	result = append(result, gateway)
	return result
}

func (s *Server) ToMem() *mem.Server {
	return &mem.Server{
		Port:  s.Port.ToMem(),
		Hosts: s.Hosts,
	}
}

func (p *Port) ToMem() *mem.Port {
	return &mem.Port{
		Number:   p.Number,
		Name:     p.Name,
		Protocol: p.Protocol,
	}
}
