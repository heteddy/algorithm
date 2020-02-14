/*
@Copyright:
*/
/*
@Time : 2020/2/11 19:18
@Author : teddy
@File : trie.go
*/

package trie

type Node struct {
	isWord  bool
	counter int
	char    string
	child   map[string]*Node
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{
		isWord:  false,
		counter: 0,
		char:    "",
		child:   make(map[string]*Node),
	}}
}

func (t *Trie) Add(word string) {
	current := t.root
	length := len([]byte(word))
	// 从第一个字符开始匹配
	for idx, c := range []byte(word) {
		// 如果存在这个树的分支，当前的
		if node, ok := current.child[string(c)]; ok {
			//
			node.counter++
			current = node
			// 最后一个字符 设置为一个word
			if idx == length-1 {
				node.isWord = true
			}
		} else { // 创建这个字符的分支
			_node := &Node{
				isWord:  false,
				counter: 1,
				char:    string(c),
				child:   make(map[string]*Node),
			}
			current.child[string(c)] = _node
			current = _node
			if idx == length-1 {
				_node.isWord = true
			}
		}
	}
}
func (t *Trie) Remove(word string) {
	current := t.root
	parent := t.root
	for _, c := range []byte(word) {
		if node, ok := current.child[string(c)]; ok {
			parent, current = current, node

		} else {
			// 没有这个word
			return
		}
	}
	if current.isWord {
		current.counter -= 1
	}
	if current.counter == 0 {
		delete(parent.child, current.char)
	}
}

func (t *Trie) ContainWord(word string) bool {
	current := t.root
	for _, c := range []byte(word) {
		if node, ok := current.child[string(c)]; ok {
			current = node
		} else {
			return false
		}
	}
	return current.isWord
}
func (t *Trie) ContainPrefix(word string) int {
	current := t.root
	for _, c := range []byte(word) {
		if node, ok := current.child[string(c)]; ok {
			current = node
		} else {
			return 0
		}
	}
	return current.counter
}
