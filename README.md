# concurrentMap

### Import
```
go get "github.com/xiao7737/concurrentMap"
```
### Usage
```
    cm := CreateConcurrentMap(32)

    // Add or Update for map
    cm.Set(ConvertStr("hello"), "go")

    // get kv from map
    res, ok := cm.Get(ConvertStr("hello"))   

    // del kv
    cm.Del(ConvertStr("hello"))
```


### Performance comparison with sync.Map
![image](https://github.com/xiao7737/concurrentMap/blob/master/bench.png) 

### More details
> See more details through test file: [concurrent_map_test.go](https://github.com/xiao7737/concurrentMap/blob/master/concurrent_map_test.go)
