package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type ListNode struct {
  Val float64
  OriPo int
  Next *ListNode
  Prev *ListNode
}

type Result struct {
  Dif float64
  Po int
}

func main() {
  var n string
  fmt.Scanf("%s", &n)
  intN, _ := strconv.Atoi(n)
  numNodes := make([]*ListNode, intN)
  // numnodes = [node_with_v_1, node_with_v_5, node_with_v_3]
  // sorted_numnodes = [node_with_v_1, node_with_v_3, node_with_v_5], middle state
  // sorted_numnodes_list = node_with_v_1 -> node_with_v_3 -> node_with_v_5
  // iterate numnodes
    // get Ai, Pi
    // 
  for key := range numNodes {
    var val string
    fmt.Scanf("%s", &val)
    valN, _ := strconv.Atoi(val)
    numNodes[key] = &ListNode{float64(valN),key,nil,nil}
  }

  sortedNumNodes := append([]*ListNode{}, numNodes...)
  sort.Slice(sortedNumNodes, func(i, j int) bool { return sortedNumNodes[i].Val < sortedNumNodes[j].Val })
  head := &ListNode{0,-1,nil,nil}
  
  for _, node := range sortedNumNodes {
    head.Next = node
    node.Prev = head
    head = node
  }
  tail := &ListNode{0,-1,nil,nil}
  head.Next = tail
  tail.Prev = head
  
  results := make([]*Result, intN)
  
  for i := intN - 1; i > 0; i-- {
    node := numNodes[i]
    prevDifVal := math.Abs(node.Val - node.Prev.Val)
    NextDifVal := math.Abs(node.Val - node.Next.Val)
    
    
    if (node.Prev.OriPo != -1 && prevDifVal <= NextDifVal) || node.Next.OriPo == -1  {
      results[i] = &Result{prevDifVal, node.Prev.OriPo}
    } else {
      results[i] = &Result{NextDifVal, node.Next.OriPo}
    }
    
    next := node.Next
    prev := node.Prev
    next.Prev = prev
    prev.Next = next
  }
  
  for i := 1; i < intN; i++ {
    fmt.Println(results[i].Dif, results[i].Po + 1)
  }

}