//你可以选择使用单链表或者双链表，设计并实现自己的链表。
//
// 单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
//
// 如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
//
// 实现 MyLinkedList 类：
//
//
// MyLinkedList() 初始化 MyLinkedList 对象。
// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
// void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果
//index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
//
//
//
//
// 示例：
//
//
//输入
//["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get",
//"deleteAtIndex", "get"]
//[[], [1], [3], [1, 2], [1], [1], [1]]
//输出
//[null, null, null, null, 2, null, 3]
//
//解释
//MyLinkedList myLinkedList = new MyLinkedList();
//myLinkedList.addAtHead(1);
//myLinkedList.addAtTail(3);
//myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
//myLinkedList.get(1);              // 返回 2
//myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
//myLinkedList.get(1);              // 返回 3
//
//
//
//
// 提示：
//
//
// 0 <= index, val <= 1000
// 请不要使用内置的 LinkedList 库。
// 调用 get、addAtHead、addAtTail、addAtIndex 和 deleteAtIndex 的次数不超过 2000 。
//
//
// Related Topics 设计 链表 👍 1102 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

//type ListNode struct {
//    Val  int
//    Next *ListNode
//    Pre  *ListNode
//}

type MyLinkedList struct {
	dummyHead *ListNode
}

func Constructor() MyLinkedList {
	node := new(ListNode)
	node.Next = node
	node.Pre = node
	return MyLinkedList{
		dummyHead: node,
	}
}

func (mll *MyLinkedList) Get(index int) int {
	head := mll.dummyHead.Next
	for head != mll.dummyHead && index > 0 {
		head = head.Next
		index--
	}
	if index != 0 {
		return -1
	}
	return head.Val
}

func (mll *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Next: mll.dummyHead.Next,
		Pre:  mll.dummyHead,
	}
	mll.dummyHead.Next.Pre = node
	mll.dummyHead.Next = node
}

func (mll *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val:  val,
		Next: mll.dummyHead,
		Pre:  mll.dummyHead.Pre,
	}
	mll.dummyHead.Pre.Next = node
	mll.dummyHead.Pre = node
}

func (mll *MyLinkedList) AddAtIndex(index int, val int) {
	head := mll.dummyHead.Next
	for head != mll.dummyHead && index > 0 {
		head = head.Next
		index--
	}
	if index > 0 {
		return
	}

	node := &ListNode{
		Val:  val,
		Next: head,
		Pre:  head.Pre,
	}
	head.Pre.Next = node
	head.Pre = node
}

func (mll *MyLinkedList) DeleteAtIndex(index int) {
	if mll.dummyHead.Next == mll.dummyHead {
		return
	}
	head := mll.dummyHead.Next
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index == 0 {
		head.Next.Pre = head.Pre
		head.Pre.Next = head.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
