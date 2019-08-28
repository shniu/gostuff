
## Design

> The final goal is to achieve a distributed kv storage systems. Like: Cassandra, Redis, Rika

设计一个 key-value 存储引擎应该考虑哪些问题？
1. 读写性能
2. 内存占用
3. 容错恢复
4. 高效持久化
5. 支持随机读写
6. 实现尽可能简单

Target:

- 实现基于 key-value 的日志追加，value 可以是任何值
- 支持基于 key 的查询
- 根据 [bitcask intro](https://github.com/basho/bitcask/blob/develop/doc/bitcask-intro.pdf) 实现一个基本功能的 kvs, 可参考 GoBeansDB
- 增加可用性和稳定性，参考 Redis 和 Memcached

Plan:

- [x] Implement a basic kv storage system based on the bitcask intro
- [x] Add unit test
- [x] Support for server, now is http server(consider to implement memcached protocol)
- [x] Add a hash index to improve performance, the hash index is in memory
- [x] Support for multiple data files(chunks or segments), active file and old files(readonly)
- [x] Opt logging, add `logrus`
- [ ] Add gc (compress files and merge files) for `Bitcask` storage engine
- [ ] Support [memcached protocol](https://github.com/memcached/memcached/blob/master/doc/protocol.txt)
- [ ] Opt the code for v1.0, add test code
- [ ] Upgrade to SSTable & LSM for Store Engine 1.1
- [ ] Golang integration with Make
- [ ] Add viper and cobra
- [ ] ~~Async flush data to disk, e.g. add a writer buffer~~
- [ ] Golang sdk(client)
- [ ] Add config file
- [ ] Serialization and deserialization, e.g. golang's gob, golang's binary, proto buffers, Avro; binary encoding vs. json vs. xml
- [ ] Benchmark test
- [ ] Considering the cluster solutions, use gRpc
- [ ] Considering the docker solutions

Paper or Blog Todo List

- [x] [Bitcask intro](https://github.com/basho/bitcask/blob/develop/doc/bitcask-intro.pdf)
- [x] [Implementing a key value store](http://codecapsule.com/2012/11/07/ikvs-implementing-a-key-value-store-table-of-contents/)
- [ ] [LevelDB](https://github.com/google/leveldb/blob/master/doc/impl.md)
  - [ ] [LevelDB and Node: What is LevelDB Anyway?](http://web.archive.org/web/20130502222338/http://dailyjs.com/2013/04/19/leveldb-and-node-1/)
  - [ ] Google’s Snappy
  - [ ] [goleveldb 的实现解读](https://leveldb-handbook.readthedocs.io/zh/latest/basic.html), 比较不错，很详细，推荐阅读
  - [ ] [SSTable intro](https://www.igvita.com/2012/02/06/sstable-and-log-structured-storage-leveldb/)
  - [ ] [LevelDB整体结构](https://soulmachine.gitbooks.io/system-design/content/cn/key-value-store.html)
- [ ] [memcached protocol](https://github.com/memcached/memcached/blob/master/doc/protocol.txt)
- [ ] Redis Protocol
- [ ] [An Efficient Design and Implementation of LSM-Tree based Key-Value Store on Open-Channel SSD](http://ranger.uta.edu/~sjiang/pubs/papers/wang14-LSM-SDF.pdf)
- [ ] [LSM-tree存储引擎的优化研究成果总结](https://mp.weixin.qq.com/s/uUFeK2ptyG7r8Fnmsry3Sw)
- [ ] [RocksDB](https://github.com/facebook/rocksdb), CouchDB, Cassandra, HBase, Redis, Memcached, LMDB
- [ ] [TiKV](https://github.com/tikv/tikv), [TiDB](https://pingcap.com/docs-cn/)
- [ ] [Bigtable: A Distributed Storage System for Structured Data](https://ai.google/research/pubs/pub27898) [中文版](http://blog.bizcloudsoft.com/wp-content/uploads/Google-Bigtable%E4%B8%AD%E6%96%87%E7%89%88_1.0.pdf)

### KVS System Design

[shniu/kvs](https://github.com/shniu/kvs) 项目是一个 kv storage 项目，最终目的是构建一个分布式的 kv 存储系统。任何系统都是演化而来的，对于 kvs 项目本身的演化思路如下：

* 构建一个单机版的高性能 kv 存储服务
  * kvs Storage Engine v1.0 考虑基于 [bitcask-intro Paper](https://github.com/basho/bitcask/blob/develop/doc/bitcask-intro.pdf) 构建一个底层存储引擎
  * 提供客户端可访问的接口服务，将存储引擎服务化，并引入SDK，先支持 json-restful 的方式
* kvs Storage Engine v1.1 考虑使用 SSTable 和 LSM 做优化，解决范围查询和key必须全部在内存的问题, 可参考 goleveldb
* 考虑实现 memcached protocol
* 引入分布式，具体实现思路再补充 todo

#### kvs Storage Engine v1.0 Design

```
package store
type KVStore struct {
}

func (kvs *KVStore) Get(key string) []byte
func (kvs *KVStore) Put(key string, value []byte) (bool, error)
func (kvs *KVStore) Delete(key string) (bool, error)
```

#### Kvs SSTable and LSM Design

1. 实现内存写入功能，基于跳表的 MemTable
  1.1 MemTable 数据模型
  1.2 基于跳表实现 Put 功能
  1.3 MemTable 阈值设置, 转变成不可更改的 MemTable
  1.4 MemTable Dump 到磁盘
2. Write Ahead Log 实现，在数据库崩溃时恢复 MemTable
3. 实现查询功能
4. Merge 策略实现 LSM


### Design Reference

- [Design Review: Key-Value Storage](https://mozilla.github.io/firefox-browser-architecture/text/0015-rkv.html)
- [Implementing a Key-Value Store – Part 8: Architecture of KingDB](http://codecapsule.com/2015/05/25/implementing-a-key-value-store-part-8-architecture-of-kingdb/)
- [小米开源的分布式KV存储系统Pegasus](https://blog.csdn.net/pengzhouzhou/article/details/78288369)
- [NetCahe:Balancing Key-Value Stores with Fast In-Network Caching](https://www.jianshu.com/p/7a9224944118)
- [Design a Cache System](http://blog.gainlo.co/index.php/2016/05/17/design-a-cache-system/)
- [design-a-key-value-store-part-i](http://blog.gainlo.co/index.php/2016/06/14/design-a-key-value-store-part-i/)
- [design-key-value-store-part-ii](http://blog.gainlo.co/index.php/2016/06/21/design-key-value-store-part-ii/)
- [random id generator](http://blog.gainlo.co/index.php/2016/06/07/random-id-generator/)
- [The Design and Implementation of a Log-Structured File System](http://research.cs.wisc.edu/areas/os/Qual/papers/lfs.pdf)
- [存储与检索的 Note](https://github.com/shniu/notes/blob/master/reading/系统设计/构建密集型应用.md#存储与检索)
- [A Log-Structured Hash Table For Fast Key/Value Data](http://highscalability.com/blog/2011/1/10/riaks-bitcask-a-log-structured-hash-table-for-fast-keyvalue.html)
- [GoBeansDB](https://github.com/douban/gobeansdb)
- [Amazon Dynamo](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf)
- [LevelDB](http://leveldb.org)
- [使用raft算法快速构建一个分布式kv系统](https://laohanlinux.github.io/2016/04/25/%E4%BD%BF%E7%94%A8raft%E7%AE%97%E6%B3%95%E5%BF%AB%E7%86%9F%E6%9E%84%E5%BB%BA%E4%B8%80%E4%B8%AA%E5%88%86%E5%B8%83%E5%BC%8F%E7%9A%84key-value%E7%B3%BB%E7%BB%9F/)

### Project Ref

- [beecask](https://github.com/yplusplus/beecask)
- [leveldb java impl](https://github.com/dain/leveldb/)
- [golang/leveldb](https://github.com/golang/leveldb)
- [google/leveldb](https://github.com/google/leveldb)
- [google/leveldb chinese comments](https://github.com/cld378632668/leveldb_chinese_comments-Code_analysis)
- [goleveldb](https://github.com/syndtr/goleveldb)
- [Level/levelup](https://github.com/Level/levelup)
- [Level/leveldown](https://github.com/Level/leveldown)
- [Level/level](https://github.com/Level/level)

### Articles

- [x] [LSM Tree Summary](https://liudanking.com/arch/lsm-tree-summary/)
- [ ] [MemtableSSTable](https://wiki.apache.org/cassandra/MemtableSSTable)
- [ ] [Leveled Compaction](https://github.com/facebook/rocksdb/wiki/Leveled-Compaction)