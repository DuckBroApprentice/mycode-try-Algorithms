package classtwo

func (l *ListNode) Enqueue(target *ListNode) {
	//不允許插隊的List
	//insert只能加在最後面
	tail := l
	for tail.Next != nil {
		tail = tail.Next
	}
	//指在最後一個節點
	tail.Next = target
}

func (l *ListNode) Dequeue() {
	//pop掉的數據可能會存在別的變數中或進行操作
	//只實現delete
	l.Next = l.Next.Next
}
