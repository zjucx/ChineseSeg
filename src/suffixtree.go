/**
  * suffixtree implemented according to ukkonen
  **/
package suffixtree

import(

)

type TreeNode struct{
  Sidx  int
  Eidx  int
  Child []*TreeNode
}

type SuffixTree struct {
  Root *TreeNode
  // for search convert frist char to idx of root childs
  Edge map[string]int

  // # = 3, active_point = (root, '\0x', 1), remainder = 1

  // 活动点（active point），是一个三元组，包括（Activenode, Activeedge, Activelength）
  ActiveNode  *TreeNode
  ActiveEdge  int
  ActiveLen   int
  //剩余后缀数（remainder），是一个整数，代表着还需要插入多少个新的后缀
  Remainder int

}
