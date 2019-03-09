# cache

An implement of Distributed cache

## Test case

```bash

# Get
curl -v http://localhost:5000

# Test: add cache
curl -X POST -H 'Content-Type: application/json' -d '{"key":"uid","value":"1009873"}' http://localhost:5000/api/cache
# Response should be: 
# Response 200
# {"code":0,"message":"succeed"} 

curl -X POST -H 'Content-Type: application/json' -d '{"key":"","value":"1009873"}' http://localhost:5000/api/cache
# Response should be:
# Response 200
# {"code":403,"message":"Null keys are forbidden"}
```
