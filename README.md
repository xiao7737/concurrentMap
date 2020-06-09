# concurrentMap

### Import
```
go get "github.com/xiao7737/concurrentMap"
```
### Usage
```
    // Init map (recommend to take a alias)
    cm := concurrentMap.CreateConcurrentMap(32)

    // Add or update
    cm.Set(ConvertStr("hello"), "go")

    // Get kv
    res, ok := cm.Get(ConvertStr("hello"))   

    // Del kv
    cm.Del(ConvertStr("hello"))
```


### Performance comparison with sync.Map
![image](https://github.com/xiao7737/concurrentMap/blob/master/benchmark.png) 

### More details
 -  See more details through test file: [concurrent_map_test.go](https://github.com/xiao7737/concurrentMap/blob/master/concurrent_map_test.go)   
 -  Get performance comparison:   
 ```go test -bench=.```
