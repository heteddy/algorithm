# goroutine pool

### 最大的buffer size
当所有的worker都在处理任务时，限定buffer的size，如果buffer已满，将会阻塞

### 并发worker数量
当worker数量为到达上限；对于新来的task将会启动一个新的worker执行

### worker空闲时间
当超过最大空闲等待时间，将会回收worker