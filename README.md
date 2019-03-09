# cache

An implement of Distributed cache

## Test case

```bash

# Get
curl -v http://localhost:5000

# Test: add cache
curl -v -X PUT -d "bbbbbbbbbbbwwww" http://localhost:5000/api/cache/u_1234
# Response should be: 
# Response 200
# {"code":0,"message":"succeed"} 

curl -v -X PUT -d '{"key":"","value":"1009873"}' http://localhost:5000/api/cache/
# Response should be:
# Response 400
```
