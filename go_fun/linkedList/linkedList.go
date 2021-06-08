package linkedList

type Object interface{}

type Node struct {
	Data Object //定义数据域
	Next *Node  //定义地址域(指向下一个表的地址)
}

type List struct {
	HeadNode *Node //头节点
}

func (list *List) Length() int {
	//获取链表头结点
	cur := list.HeadNode
	//定义一个计数器，初始值为0
	count := 0

	for cur != nil {
		//如果头节点不为空，则count++
		count++
		//对地址进行逐个位移
		cur = cur.Next
	}
	return count
}

func (l *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = l.HeadNode
	l.HeadNode = node
	return node
}

func (this *List) len() int {
	//获取链表头结点
	cur := this.HeadNode
	//定义一个计数器，初始值为0
	count := 0

	for cur != nil {
		//如果头节点不为空，则count++
		count++
		//对地址进行逐个位移
		cur = cur.Next
	}
	return count
}
