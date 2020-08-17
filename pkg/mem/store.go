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

import "github.com/hashicorp/go-memdb"

var DB *memdb.MemDB

const (
	GatewayTable  = "Gateway"
	RouteTable    = "Route"
	UpstreamTable = "Upstream"
	PluginTable   = "Plugin"
)

func init() {
	if db, err := NewDB(); err != nil {
		panic(err)
	} else {
		DB = db
	}
}

func NewDB() (*memdb.MemDB, error) {
	var schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			GatewayTable:  gatewaySchema,
			RouteTable:    routeSchema,
			UpstreamTable: upstreamSchema,
			PluginTable:   pluginSchema,
		},
	}
	if memDB, err := memdb.NewMemDB(schema); err != nil {
		return nil, err
	} else {
		return memDB, nil
	}
}

var gatewaySchema = &memdb.TableSchema{
	Name: GatewayTable,
	Indexes: map[string]*memdb.IndexSchema{
		"id": {
			Name:    "id",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "FullName"},
		},
		"name": {
			Name:         "name",
			Unique:       true,
			Indexer:      &memdb.StringFieldIndex{Field: "Name"},
			AllowMissing: true,
		},
	},
}

var routeSchema = &memdb.TableSchema{
	Name: RouteTable,
	Indexes: map[string]*memdb.IndexSchema{
		"id": {
			Name:    "id",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "FullName"},
		},
		"name": {
			Name:         "name",
			Unique:       true,
			Indexer:      &memdb.StringFieldIndex{Field: "Name"},
			AllowMissing: true,
		},
	},
}

var upstreamSchema = &memdb.TableSchema{
	Name: UpstreamTable,
	Indexes: map[string]*memdb.IndexSchema{
		"id": {
			Name:    "id",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "FullName"},
		},
		"name": {
			Name:         "name",
			Unique:       true,
			Indexer:      &memdb.StringFieldIndex{Field: "Name"},
			AllowMissing: true,
		},
	},
}

var pluginSchema = &memdb.TableSchema{
	Name: PluginTable,
	Indexes: map[string]*memdb.IndexSchema{
		"id": {
			Name:    "id",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "FullName"},
		},
		"name": {
			Name:         "name",
			Unique:       true,
			Indexer:      &memdb.StringFieldIndex{Field: "Name"},
			AllowMissing: true,
		},
	},
}
