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
        st.ActiveNode = root
        st.ActiveEdge = word
        st.ActiveLen += 1
        st.Remainder += 1
        return
    }

    activechild := st.ActiveNode[st.ActiveEdge]
    if st.Remainder > 1 && value == sentence[activechild.Sidx+st.Remainder-1] {
      st.ActiveLen += 1
      st.Remainder += 1
      return
    }

    if v, ok := root.Child[word]; !ok && st.Remainder == 1 {
      child := st.NewChild(index)
      st.Child[word] = child
      return
    }
  }
}
