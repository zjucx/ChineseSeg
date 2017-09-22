/**
  * suffixtree implemented according to ukkonen
  **/
package suffixtree

import(

)

type TreeNode struct{
  Sidx  int
  Eidx  int
  Child map[string]*TreeNode
}

type SuffixTree struct {
  Root *TreeNode
  // // for search convert frist char to idx of root childs
  // Edge map[string]int

  // # = 3, active_point = (root, '\0x', 1), remainder = 1
  CurStep int // #
  // 活动点（active point），是一个三元组，包括（Activenode, Activeedge, Activelength）
  ActiveNode  *TreeNode
  ActiveEdge  string
  ActiveLen   int
  //剩余后缀数（remainder），是一个整数，代表着还需要插入多少个新的后缀
  Remainder int
}

func (st *SuffixTree) Build(str string) {
  // prepare for chisese sentence
  sentence := []rune(str)

  root := st.Root
  //fmt.Println("rune=", string(sentence[0]))
  for index, value := range sentence {
    st.CurStep = index
    word := string(value)

    // if cur word is frist word of suffix in one edge in cur root
    if v, ok := root.Child[word]; ok && st.Remainder == 1 {
        st.ActiveEdge = word
        st.ActiveLen = 1
        st.Remainder += 1
        return
    }

    if v, ok := root.Child[word]; !ok && st.Remainder == 1 {
      child := st.NewChild(index)
      root.Child[word] = child
      return
    }

    if st.Remainder > 1 && value == sentence[st.ActiveNode.Sidx+st.ActiveLen] {
      st.Remainder += 1
      st.ActiveLen += 1
      nodeLen := st.ActiveNode.Eidx - st.ActiveNode.Bidx
      if nodeLen > 0 &&  st.ActiveLen > nodeLen {
        st.ActiveNode = st.ActiveNode.Child[st.ActiveEdge]
        st.ActiveEdge = word
        st.ActiveLen = 1
      }
      return
    }

    lastSplit = nil
    for st.Remainder >= 1 {
      // insert root node child
      if st.Remainder == 1 {
        child := st.NewChild(index)
        root.Child[word] = child
        st.ActiveEdge = nil
        return
      }

      // current split node
      curSplit := st.ActiveNode[st.ActiveEdge]
      st.Split(curSplit, str)

      // update active point
      st.Remainder -= 1
      lastword := sentence[st.CurStep-st.Remainder+1]
      suffixLink := st.ActiveNode.SuffixLink
      if st.ActiveNode == root {
        /*  Rule 1:
            当向根节点插入时遵循：
            active_node 保持为 root；
            active_edge 被设置为即将被插入的新后缀的首字符；
            active_length 减 1；
        */
        st.ActiveEdge = lastword
        st.ActiveLen -= 1
      } else if suffixLink != nil{
        /*  Rule 3
            当从 active_node 不为 root 的节点分裂边时，
            我们沿着后缀连接（Suffix Link）的方向寻找节点，如果存在一个节点，
            则设置该节点为 active_noe；如果不存在，则设置 active_node 为 root。
            active_edge 和 active_length 保持不变。
        */
        st.ActiveNode = suffixLink
      } else {
        st.ActiveNode = root
      }

      /*  Rule 2
          如果我们分裂（Split）一条边并且插入（Insert）一个新的节点，
          并且如果该新节点不是当前步骤中创建的第一个节点，
          则将先前插入的节点与该新节点通过一个特殊的指针连接，称为后缀连接（Suffix Link）。
          后缀连接通过一条虚线来表示。
      */
      if lastSplit != nil {
        curSplit.SuffixLink = lastSplit
      }
      lastSplit = curSplit
    }
  }
}

func (st *SuffixTree) Split(node *TreeNode, str string) {
  node.Flag = 1
  node.Eidx = node.Bidx + st.ActiveLen
  child1 := st.NewChild(node.Eidx)
  child2 := st.NewChild(st.CurStep)

  sentence := []rune(str)
  node.Child[string(sentence[node.Eidx])] = child1
  node.Child[string(sentence[st.CurStep])] = child2
}
