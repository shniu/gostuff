
## cache

An implement of Distributed cache.

这个子项目提供灵活的缓存实现

### Test case

```bash

# Get
curl -v http://localhost:5000/api/cache/u_1234

# Test: add cache
curl -v -X PUT -d "bbbbbbbbbbbwwww" http://localhost:5000/api/cache/u_1234
# Response should be:
# Response 200
# {"code":0,"message":"succeed"}

curl -v -X PUT -d "{\"name\":\"xx\",\"age\":\"100\"}" http://localhost:5000/api/cache/u:999
# Response should be:
# Response 400
```