package classtwo

//leetcode Q82 Remove Duplicates from Sorted List II

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Del(target int) *ListNode {
	curr := l
	free := &ListNode{} //Go有GC，該Pointer沒有被利用就會釋放，free是C的用法(講師用的是C語言)
	for curr != nil {
		if curr.Next.Val == target {
			free = curr.Next //保留這個Pointer
			curr.Next = curr.Next.Next
		}
	}
	free.Next = nil
	return l
}

func deleteDuplicates(head *ListNode) *ListNode {
	head_cur := &ListNode{Val: 0}
	res := head_cur
	i := 0
	for head != nil {
		if head.Next == nil {
			if i == 0 {
				head_cur.Next = head
				head_cur = head_cur.Next
			} else {
				head_cur.Next = nil
			}
			break
		}
		if head.Val == head.Next.Val {
			i++
		} else if head.Val != head.Next.Val {
			if i == 0 {
				head_cur.Next = head
				head_cur = head_cur.Next
			} else if i != 0 {
				i = 0
			}
		}
		head = head.Next
	}
	head_cur.Next = nil
	return res.Next
}

//Runtime 0ms Beats 100.00%
//Memory 4.69MB Beats 99.49%
