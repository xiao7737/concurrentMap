# concurrentMap

### Features
 - thread safe and high performance read and write
 - sharding mechanism to reduce lock range
 - read-write lock to improve concurrency
 - murmurHash to make the distribution more reasonable
 - support user program to customize the number of shard
 - provide operations including create, get, set, del, count, exists 
 
### Import
```
go get "github.com/xiao7737/concurrentMap"
```
### Usage
```
    // Init map (recommend to take an alias)
    cm := concurrentMap.CreateConcurrentMap(32)

    // Add or update
    cm.Set(ConvertStr("hello"), "go")

    // Get kv
    res, ok := cm.Get(ConvertStr("hello"))   

    // Del kv
    cm.Del(ConvertStr("hello"))
```


### Performance comparison with sync.Map
![image](https://github.com/xiao7737/concurrentMap/blob/master/bench.png)

### TODO
 - Batch setting operation
 
### More details
 -  See more details through test file: [concurrent_map_test.go](https://github.com/xiao7737/concurrentMap/blob/master/concurrent_map_test.go)   
 -  Get performance comparison:   
 ```go test -bench=.```

