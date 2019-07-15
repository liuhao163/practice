package lru

type LRUNode struct {
	Prev *LRUNode //Lru双向链表
	Next *LRUNode //Lru双向链表

	Data interface{}

	Hnext *LRUNode //Lru散列
}
