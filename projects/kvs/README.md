# kvs

An implementation of kvs using Golang.

## Project

There are many ways to implement the KVS, each have advantages and disadvantages, and is a process of continuous optimization, 
such as bitcask, leveldb, rocksdb, couchdb, Cassandra, HBase, Redis, Memcached, LMDB, Dynamo, badger, boltdb, Riak, etc.

So consider the project is divided into different categories, to store different implementations, such easy to manage:

```
- kvs
  |- bitcask/    # bitcask storage engine implementation
  |- leveldb/    # SSTable and LSM Tree implementation
  |- and so on...
```

## Docs

Document for Design, look [here](docs/design.md).

## Tests

```shell

# Set key
curl -H "Content-Type:application/json" -X POST -d "{\"key\":\"uid:1\",\"value\":\"red\"}" http://127.0.0.1:3000/set
curl -v -H "Content-Type:application/json" -X POST -d "{\"key\":\"uid:2\",\"value\":\"black\"}" http://127.0.0.1:3000/set
curl -H "Content-Type:application/json" -X POST -d "{\"key\":\"uid:3\",\"value\":\"yellow\"}" http://127.0.0.1:3000/set
curl -H "Content-Type:application/json" -X POST -d "{\"key\":\"uid:4\",\"value\":\"gray\"}" http://127.0.0.1:3000/set

# Get key
curl -X GET http://127.0.0.1:3000/get?key=uid:1
curl -X GET http://127.0.0.1:3000/get?key=unknown

# Delte key
curl -X GET http://127.0.0.1:3000/get?key=uid:2
curl -v -X DELETE http://127.0.0.1:3000/delete?key=uid:2
curl -X GET http://127.0.0.1:3000/get?key=uid:2


```

## Reference

Look [here](docs/design.md#design-reference)

## Credits

Original idea, design and implementation: [shniu](https://github.com/shniu)
