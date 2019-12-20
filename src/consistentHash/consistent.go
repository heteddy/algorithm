package consistentHash

import (
	"fmt"
	"hash/crc32" // 也可以使用murmur的hash算法，比较快
	"sort"
	"strconv"
	"sync"
)

const (
	DEFAULT_REPLICA = 100    // 每个node复制品数量
	MAX_SLOTS       = 163234 // 总共的slot数量，用来取余
)

type hashRing []uint32 //sortable

func (c hashRing) Len() int {
	return len(c)
}

func (c hashRing) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c hashRing) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type Node struct {
	id     int
	name   string
	ip     string
	port   int
	weight int // 设置节点的权重，
}

func (n *Node) GetURL() string {
	return n.ip + ":" + strconv.Itoa(n.port)
}

func (n *Node) String() string {
	return fmt.Sprintf("%s_id%d_w%d", n.name, n.id, n.weight)
}

func NewNode(id int, ip string, port int, name string, weight int) *Node {
	return &Node{
		id:     id,
		name:   name,
		ip:     ip,
		port:   port,
		weight: weight,
	}
}

type ConsistentHashService struct {
	mux       sync.RWMutex
	nodes     map[uint32]*Node // 环中每个点对应的节点
	ring      hashRing
	replicas  int // 副本的数量，是一个基数，replicas*weight为一个节点的实际数量
	resources map[int]bool
}

func NewConsistentService(replicas int) *ConsistentHashService {
	var replicaNumber int
	if replicas > 0 {
		replicaNumber = replicas
	} else {
		replicaNumber = DEFAULT_REPLICA
	}

	return &ConsistentHashService{
		mux:       sync.RWMutex{},
		nodes:     make(map[uint32]*Node),
		ring:      hashRing{},
		replicas:  replicaNumber,
		resources: make(map[int]bool),
	}
}

// 增加一个节点
func (c *ConsistentHashService) Add(node *Node) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.resources[node.id]; ok {
		return false
	}

	count := c.replicas * node.weight
	for i := 0; i < count; i++ {
		// 获取node的uuid--
		nodeKey := c.generateNodeSlotString(i, node)
		// 生成node的slot id；
		hashCode := c.hashNode(nodeKey) % MAX_SLOTS
		c.nodes[hashCode] = node
	}
	c.resources[node.id] = true
	c.sortRing()
	return true
}

// 删除一个节点
func (c *ConsistentHashService) Remove(n *Node) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	if _, ok := c.resources[n.id]; !ok {
		return false
	}
	delete(c.resources, n.id)
	count := c.replicas * n.weight
	for i := 0; i < count; i++ {
		str := c.generateNodeSlotString(i, n)
		key := c.hashNode(str)
		delete(c.nodes, key)
	}
	c.sortRing()
	return true
}

// 给定一个key，获取node节点，将存放到指定的节点中
func (c *ConsistentHashService) Get(key string) *Node {
	c.mux.RLock()
	defer c.mux.RUnlock()
	keySlotID := c.hashNode(key) % MAX_SLOTS
	i := c.search(keySlotID)
	// 返回节点
	return c.nodes[c.ring[i]]
}

func (c *ConsistentHashService) generateNodeSlotString(i int, n *Node) string {
	return n.name + "-" + n.ip + "-" + strconv.Itoa(i) + "-" + strconv.Itoa(n.weight) + "-" + strconv.Itoa(n.id)
}

func (c *ConsistentHashService) hashNode(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// 将所有的hash值排序，
func (c *ConsistentHashService) sortRing() {
	c.ring = hashRing{}
	// 使用node的key，组成一个环
	for i, _ := range c.nodes {
		//
		// log.Printf("nodes_key:%d;node:%s\n", i, node.String())
		c.ring = append(c.ring, i)
	}
	sort.Sort(c.ring)
}

// 所
func (c *ConsistentHashService) search(slotID uint32) int {
	i := sort.Search(len(c.ring), func(i int) bool {
		return c.ring[i] >= slotID
	})
	if i < len(c.ring) {
		// 找到前面的一个节点
		if i == len(c.ring)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.ring) - 1
	}
}
