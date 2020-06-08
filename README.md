# concurrentMap

### Import
```
go get "github.com/xiao7737/concurrentMap"
```
### Usage
```
    cm := CreateConcurrentMap(32)
    cm.Set(ConvertStr("hello"), "go")
    res, ok := cm.Get(ConvertStr("hello"))   
    cm.Del(ConvertStr("hello"))
```


### Performance comparison with sync.Map
![image](https://github.com/xiao7737/concurrentMap/blob/master/bench.png) 

### More details
> See more details through test file: [concurrent_map_test.go](https://github.com/xiao7737/concurrentMap/blob/master/concurrent_map_test.go)
