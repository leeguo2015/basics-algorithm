package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(newNode *ListNode) ListNode {
	newNode.Next = l
	return *newNode
}

func (l *ListNode) Len() int {
	if l == nil {
		return 0
	}
	length := 1
	nowList := l
	for {
		if nowList.Next == nil {
			break
		}
		nowList = nowList.Next
		length++
	}
	return length
}
