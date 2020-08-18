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

import "fmt"

type UpstreamDB struct {
	Upstreams []*Upstream
}

// insert Upstream to memDB
func (db *UpstreamDB) Insert() error {
	txn := DB.Txn(true)
	defer txn.Abort()
	for _, r := range db.Upstreams {
		if err := txn.Insert(UpstreamTable, r); err != nil {
			return err
		}
	}
	txn.Commit()
	return nil
}

func (g *Upstream) FindByFullName() (*Upstream, error) {
	txn := DB.Txn(false)
	defer txn.Abort()
	if raw, err := txn.First(UpstreamTable, "id", *g.FullName); err != nil {
		return nil, err
	} else {
		if raw != nil {
			current := raw.(*Upstream)
			return current, nil
		}
		return nil, fmt.Errorf("NOT FOUND")
	}
}

func (g *Upstream) Diff(t MemModel) bool {
	return true
}
