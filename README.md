## Protocol Definition Exploration

This repository contains some exploratory code for 

```bash
$ curl http://localhost:8080/api/v1/job -X POST -d '{"job": {"name": "myjob", "size": 10, "type": "CPU"}}' -H "Content-Type: application/json"
{"job":{"name":"myjob","size":10}}
```