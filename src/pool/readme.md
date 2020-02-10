# goroutine pool

+ 最大的buffer size

当所有的worker都在处理任务时，限定buffer的size，如果buffer已满，将会阻塞

+ 并发worker数量

当worker数量为到达上限；对于新来的task将会启动一个新的worker执行

+ worker空闲时间

当超过最大空闲等待时间，将会回收worker

+ 优雅的退出，等待所有的任务执行完成

    1. 支持两种退出方式，暴力退出，直接退出；
    1. 优雅退出；不接收新的任务并等待已经添加的任务退出