# LRU 算法
最近最少使用的算法，使用一个list+Map的方式实现

双向链表保存value, map[string] *Node; map中的value指向list中的节点；
这样可以保证



## redis中LRU的实现

Redis支持和LRU相关淘汰策略包括，

+ volatile-lru 设置了过期时间的key参与近似的lru淘汰策略
+ allkeys-lru 所有的key均参与近似的lru淘汰策略

Redis采用了一个近似的做法，就是随机取出若干个key，然后按照访问时间排序后，淘汰掉最不经常使用的，具体分析如下：

为了支持LRU，Redis 2.8.19中使用了一个全局的LRU时钟，server.lruclock，定义如下，
```
#define REDIS_LRU_BITS 24
unsigned lruclock:REDIS_LRU_BITS; /* Clock for LRU eviction */
```

默认的LRU时钟的分辨率是1秒，可以通过改变REDIS_LRU_CLOCK_RESOLUTION宏的值来改变，Redis会在serverCron()中调用updateLRUClock定期的更新LRU时钟，更新的频率和hz参数有关，默认为100ms一次，如下，
```
#define REDIS_LRU_CLOCK_MAX ((1<<REDIS_LRU_BITS)-1) /* Max value of obj->lru */
#define REDIS_LRU_CLOCK_RESOLUTION 1 /* LRU clock resolution in seconds */

void updateLRUClock(void) {
    server.lruclock = (server.unixtime / REDIS_LRU_CLOCK_RESOLUTION) &
                                                REDIS_LRU_CLOCK_MAX;
}
```

server.unixtime是系统当前的unix时间戳，当 lruclock 的值超出REDIS_LRU_CLOCK_MAX时，会从头开始计算，所以在计算一个key的最长没有访问时间时，可能key本身保存的lru访问时间会比当前的lrulock还要大，这个时候需要计算额外时间，如下，
```c
/* Given an object returns the min number of seconds the object was never
 * requested, using an approximated LRU algorithm. */
unsigned long estimateObjectIdleTime(robj *o) {
    if (server.lruclock >= o->lru) {
        return (server.lruclock - o->lru) * REDIS_LRU_CLOCK_RESOLUTION;
    } else {
        return ((REDIS_LRU_CLOCK_MAX - o->lru) + server.lruclock) *
                    REDIS_LRU_CLOCK_RESOLUTION;
    }
}
```


在redis中有相应的时间，redis考虑了综合性能，在设计的时候不是每次访问都检查
而是在一个时间周期内去检查一次


Redis会基于server.maxmemory_samples配置选取固定数目的key，然后比较它们的lru访问时间，然后淘汰最近最久没有访问的key，maxmemory_samples的值越大，Redis的近似LRU算法就越接近于严格LRU算法，但是相应消耗也变高，对性能有一定影响，样本值默认为5。