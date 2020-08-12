<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

# apisix-control-plane

## what is apisix-control-plane?
apisix-control-plane provide APISIX with a `yaml` configuration capability,
We can use `yaml` to define the proxy behavior of APISIX

## Why do you want to do this?
1. In order to facilitate the integration of k8s, use `yaml` to define APISIX;
2. For easier synchronization across clusters;
3. Can be better adapted to multiple platforms (k8s, vm);

## DISCUSS
Here are some examples of `yaml` and we can add comments here.
https://github.com/apache/apisix-control-plane/issues/3
We can also submit a PR to modify this [document](doc/yaml_struct.md).
