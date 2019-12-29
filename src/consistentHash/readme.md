# 一致性hash
> 引入

hash环是用于解决寻址的问题；如果使用的是普通hash(key)获取服务器地址，那么增加和删除一个节点的影响比较大

一致性hash环，采用将一个环分成很多份；每个服务器根据自己的名字或者地址+权重进行hash多个副本分布到环中

并且对环中的虚拟节点编号进行排序；每次获取的时候折半查找获取节点


1. 定义replica 复制品数量
1. 定义slot环中的插槽数量，用于取余
1. 定义一个hash ring,是一个有序slice
1. hash 服务的定义

    ``` go
    type ConsistentHashService struct {
    mux       sync.RWMutex
    nodes     map[uint32]*Node // 环中每个点对应的节点(虚拟节点id->节点)
    ring      hashRing
    replicas  int // 副本的数量，是一个基数，replicas*weight为一个节点的虚拟节点数量
    resources map[int]bool //是否已经存在于hash环中
    }
    ```
1. 增加删除一个节点都需要对ring重新排序(删除可以优化为直接获取虚拟节点id，然后直接删除，不排序)
1. 获取一个key对应的节点；首先获取key对应的hash值然后在ring中折半查找获取ring id

