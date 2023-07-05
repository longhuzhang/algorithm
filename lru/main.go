package main

import (
	"fmt"
)

func main() {


	// ["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]

	// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	


	lru := Constructor(2)
	fmt.Println(lru.Get(2))
	lru.Put(2, 6)
	fmt.Println(lru.Get(2))
	lru.Put(1, 5)
	lru.Put(1, 2)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))

}

type LRUCache struct {
	c *Cache
}

func Constructor(capacity int) LRUCache {
	lru := new(LRUCache)
	lru.c = NewCache(capacity)
	return *lru
}

func (this *LRUCache) Get(key int) int {
	n := this.c.GetNode(key)
	if  n == nil{
		return -1
	}
	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	this.c.PutNode(key, value)
}

type Node struct {
	key  int
	Val  int
	Pre  *Node
	Next *Node
}

type Cache struct {
	Size     int
	cap      int
	HashNode map[int]*Node
	head     *Node
	tail     *Node
}

func NewCache(s int) *Cache {
	cache := new(Cache)
	cache.Size = s
	cache.HashNode = make(map[int]*Node)
	head := new(Node)
	tail := new(Node)

	head.Next = tail
	tail.Pre = head
	cache.tail = tail
	cache.head = head
	return cache
}

func (c *Cache) GetNode(k int) *Node {
	n, ok := c.HashNode[k]
	if !ok {
		return nil
	}
	c.DeleteNodeList(n)
	c.AddNodeList(n)
	return n
}

func (c *Cache) PutNode(k, v int) {
	n, ok := c.HashNode[k]
	if ok {
		c.DeleteHash(n.key)
		c.DeleteNodeList(n)
	}else {
		if c.cap >= c.Size {
			c.removeFirst()
		}
		c.cap++
	}

	newNode := new(Node)
	newNode.key = k
	newNode.Val = v
	c.AddNodeList(newNode)
	c.AddHash(newNode)
}

func (c *Cache) AddNodeList(newNode *Node) {
	newNode.Next = c.tail
	newNode.Pre = c.tail.Pre
	c.tail.Pre.Next = newNode
	c.tail.Pre = newNode
}

func (c *Cache) DeleteNodeList(tempNode *Node) {
	tempNode.Pre.Next = tempNode.Next
	tempNode.Next.Pre = tempNode.Pre
}

func (c *Cache) DeleteHash(k int) {
	delete(c.HashNode, k)
}

func (c *Cache) AddHash(n *Node) {
	c.HashNode[n.key] = n
}

func (c *Cache) removeFirst() {
	node := c.head.Next
	c.DeleteNodeList(node)
	c.DeleteHash(node.key)
	c.cap--
}
