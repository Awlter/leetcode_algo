func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  result := &ListNode{0, nil}
  r := result

  for list1 != nil && list2 != nil {
    if list1.Val <= list2.Val {
      nextNode := list1.Next
      list1.Next = nil
      result.Next = list1
      list1 = nextNode
    } else {
      nextNode := list2.Next
      list2.Next = nil
      result.Next = list2
      list2 = nextNode
    }
    result = result.Next
  }

  if list1 != nil {
    result.Next = list1
  } else {
    result.Next = list2
  }
}