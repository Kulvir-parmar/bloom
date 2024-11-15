## Bloom Filter

Bloom Filters are space efficient probabilistic data structure that specifically used for membership queries. A hash coded space is used which is very less than conventional hashing mehthods. The reduction in space comes at cost of small fraction of tolerable errors(false positives).
Suitable for applications where most of the queries will result in false. If query returned false, it is definitely false but if it returns true then there is very slight possibility of a false positive. So you have do a disk lookup in case it returns true. In such applications the early rejection at cost of few false disk lookups can greatly increase the performance.

Hash area of M bits is used i.e. bit adressable. All the bits of hash area is set to 0 at first. Each message to be stored is hash coded into distinct bit addresses (a1, a2, ..., ad) and all *d* bits are set to 1.
To test a message a sequence of *d* bit addresses is generated in same way. If all *d* bits are set to 1 then new message is accepted. If any of the bit is 0 the message is rejected.

More about bloom filters: [wikipedia](https://en.wikipedia.org/wiki/Bloom_filter)


## Installation

```bash
cd /path/to/project
go get github.com/kulvirdotgg/bloomfilter
```

## Usage
- **Creating a bloom filter instance**
```go
import (
    bf "github.com/Kulvir-parmar/bloomfilter"
)

// adjust these numbers according to the need of your application
const (
    falsePositiveRate := 0.01
    expectedInsertion := 10000
)

bloomFilter, err := bf.NewBloomFilter(falsePosRate, expectedInsertions)
```

- **Insertion to bloom filter**

```go
bloomFilter.Put(message)
```

- **Checking for membership in bloom filter**

```go
bloomFilter.MightContain(message)
```

---

- **My implementation was highly inspired from [Guava](https://github.com/google/guava/blob/master/guava/src/com/google/common/hash/BloomFilterStrategies.java) library's implementation of a bloom filter**

- **RocksDB implementation of [Bloom Filters](https://github.com/rockset/rocksdb-cloud/blob/master/util/bloom_impl.h) is also very cool**
- **More [details](https://github.com/facebook/rocksdb/wiki/RocksDB-Bloom-Filter#the-math) about the RocksDB implementation**
