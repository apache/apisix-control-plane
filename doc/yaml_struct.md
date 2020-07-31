## Currently, 4 types are defined, namely
1.Gateway
2.Traffic distribution rules
3.Define the target service
4.Plug-in extension

### Gateway
For edge traffic, extract the host separately and define a Gateway type

```yaml
kind:Gateway
name: baidu-gw
servers:
 - port:
     number: 80
     name: http
     protocol: HTTP
   hosts:
   - "a.domain.com"
   - "b.domain.com"
```
|  object/field   | describition |
|  ----  | ----  |
| Gateway  | the type for Edge traffic  |
| Gateway.servers  | Define edge traffic service list  |
| Gateway.servers.port  | the port for service  |
| Gateway.servers.port.protocol  | Specify protocol  |
| Gateway.servers.hosts  | List of domain names that a certain service can accept  |

### Define traffic distribution rules

```yaml
kind: Rule
name: xxx-rules
hosts:
- "domain.com"
gateways:
- baidu-gw
http:
- route:
 - destination:
     port: 28002
     host: baidu-server
     subset: baidu-v1
     weight: 10
 label：
    app: baidu
    version: v1
  match:
  - headers:
     product_id:
       exact: v1
- route:
  - destination:
       port: 28002
       host: baidu-server
       subset: v2
  label:
    app: baidu
    version: v2

```
|  object/field   | describition |
|  ----  | ----  |
|  Rule  | the rule type for traffic |
|  Rule.hosts  | Specify the list of accepted hosts |
|  Rule.gateways  | Specify which gateway it belongs to, receive traffic from the edge gateway |
|  Rule.http  | Specify the HTTP protocol (similar to other protocols) |
|  Rule.http.route  | define route|
|  Rule.http.match  | Conditions for hitting the route|
|  Rule.http.route.destination  | Target service definition|

### Define the target service

```yaml
kind: destinations
name: baidu-dest
host: baidu-server
subsets:
- name: baidu-v1
  ips：
  - 127.0.0.1
  - 127.0.0.2
- name: v2
  selector:
    labels:
      tag: v2

```
|  object/field   | describition |
|  ----  | ----  |
|  destinations  | Target service |
| destinations.host | Target service internal host |
| destinations.subsets | Target service collection |

### Plug-in extension
The logic of the plugin itself is not defined here, only known plugins will be referenced here

```yaml
kind：Plugin
selector：
  labels：
     app: baidu
sort:
- name: limit-count
  conf:
      max：100
- name: prometheus
  conf:
     ...schema..

```

### A complete demo
1. Use APISIX Admin API
```shell
curl -XPUT http://127.0.0.1:9080/apisix/admin/routes/1 -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -d '
{
    "desc": "test_route",
    "uris": [
        "/*"
    ],
    vars:[
        ["arg_name", "==", "json"],
        ["arg_age", ">", "18"]
    ],
    "hosts": [
        "baidu.com"
    ],
    "upstream": {
        "type": "roundrobin",
        "nodes": {
            "127.0.0.1:8080": 10
        },
        "timeout": {
            "connect": 6000,
            "send": 6000,
            "read": 6000
        },
        "enable_websocket": false
    },
    "plugins":{
    	"prometheus":{},
    	"proxy-rewrite":{
    		"uri": "/test/home.html",
            "scheme": "http",
            "host": "baidu.com",
            "headers": {
                "X-Api-Version": "v1",
                "X-Api-Engine": "apisix",
                "X-Api-useless": ""
            }
    	}
    }
}'

```
2. Use YAML
```yaml
kind:Gateway
name: baidu-gw
servers:
 - port:
     number: 80
     name: http
     protocol: HTTP
   hosts:
   - "baidu.com"

-----------------

kind: Rule
name: baidu-rules
hosts:
- "baidu.com"
gateways:
- baidu-gw
http:
- route:
  - destination:
       port: 8080
       host: baidu-server
       subset: baidu-v1
    label：
      app: baidu
      version: v1
  match:
  - args:
     name:
       exact: "json"
     age:
       Greater: 18
  - uri:
      prefix: "/"
-------------------

kind: destinations
name: baidu-dest
host: baidu-server
subsets:
- name: baidu-v1
  ips：
  - 127.0.0.1

-------------------

kind：Plugin
selector：
  labels：
     app: baidu
sort:
- name: proxy-rewrite
  conf:
    uri: "/test/home.html"
    scheme: "http"
    host: "baidu.com"
    headers:
      X-Api-Version: "v1"
      X-Api-Engine: "apisix"
      X-Api-useless: ""
- name: prometheus

```