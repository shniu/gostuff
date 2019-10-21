package cn
//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例： 
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
// 
// Related Topics 链表

type ListNode struct {
	Val int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil { return nil }
	if l1 == nil && l2 != nil { return l2 }
	if l1 != nil && l2 == nil { return l1 }

	head := l1
	if l1.Val > l2.Val {
		head = l2
	}

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			// 在 l1 中找到一个合适的位置插入 l2 的元素
			for l1.Next != nil && l1.Next.Val < l2.Val {
				l1 = l1.Next
			}

			// move ptr
			l1, l2 = movePtr(l1, l2)
		} else {
			// 在 l2 中找到一个合适的位置插入 l1 的元素
			for l2.Next != nil && l2.Next.Val < l1.Val {
				l2 = l2.Next
			}

			// move ptr
			//tmp2 := l2.Next
			//l2.Next = l1
			//l1 = l1.Next
			//l2.Next.Next = tmp2
			l2, l1 = movePtr(l2, l1)
		}
	}

	return head
}

func movePtr(l1, l2 *ListNode) (*ListNode, *ListNode) {
	tmp1 := l1.Next
	l1.Next = l2
	l2 = l2.Next
	l1.Next.Next = tmp1
	return l1, l2
}

// 找到可以插入的位置
func findPosition(l1 *ListNode, l2 *ListNode) {
	tmp := l1
	for tmp.Next != nil && tmp.Next.Val < l2.Val {
		tmp = tmp.Next
	}
	l1 = tmp
}
//leetcode submit region end(Prohibit modification and deletion)
