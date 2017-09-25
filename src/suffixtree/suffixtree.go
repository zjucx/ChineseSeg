/**
  * suffixtree implemented according to ukkonen
  **/
package suffixtree

import(
  "fmt"
)

type TreeNode struct{
  Sidx  int
  Eidx  int
  SuffixLink *TreeNode
  Child map[string]*TreeNode
}

type SuffixTree struct {
  Root *TreeNode

  // raw words
  Words string
  // # = 3, active_point = (root, '\0x', 1), remainder = 1
  CurStep int // #
  // 活动点（active point），是一个三元组，包括（Activenode, Activeedge, Activelength）
  ActiveNode  *TreeNode
  ActiveEdge  string
  ActiveLen   int
  //剩余后缀数（remainder），是一个整数，代表着还需要插入多少个新的后缀
  Remainder int
}

func New() *SuffixTree{
  return &SuffixTree{
    Root: &TreeNode{Sidx: 0, Eidx: 0, Child: make(map[string]*TreeNode)},
    CurStep: -1,
    ActiveNode: nil,
    ActiveEdge: "",
    ActiveLen: 0,
    //剩余后缀数（remainder），是一个整数，代表着还需要插入多少个新的后缀
    Remainder: 1,
    Words: "",
  }
}


func (st *SuffixTree) Build(str string) {
  st.Reset()
  st.Words += str
  // prepare for chisese sentence
  sentence := []rune(str)

  root := st.Root
  //fmt.Println("rune=", string(sentence[0]))
  for _, value := range sentence {
    st.CurStep += 1
    word := string(value)

    // if cur word is frist word of suffix in one edge in cur root
    if _, ok := root.Child[word]; ok && st.Remainder == 1 {
        st.ActiveNode = root
        st.ActiveEdge = word
        st.ActiveLen = 1
        st.Remainder += 1
        fmt.Printf("#:%d ActiveEdge:%s ActiveLen:%d Remainder:%d\n", st.CurStep, st.ActiveEdge, st.ActiveLen, st.Remainder)
        continue
    }

    // if cur word is not in the root child and is also not a suffix, then add a child
    if _, ok := root.Child[word]; !ok && st.Remainder == 1 {
      child := st.NewChild(st.CurStep, 0)
      root.Child[word] = child
      fmt.Println("add new child " + word)
      continue
    }

    // if cur word is included by a child then add modify the variables
    //nodeLen := len(st.ActiveNode.Child[st.ActiveEdge].Words)
    if st.Remainder > 1 && value == []rune(st.Words)[st.ActiveNode.Child[st.ActiveEdge].Sidx+st.ActiveLen] {
      st.Remainder += 1
      st.ActiveLen += 1
      nodeLen := st.ActiveNode.Child[st.ActiveEdge].Eidx - st.ActiveNode.Child[st.ActiveEdge].Sidx
      if nodeLen > 0 &&  st.ActiveLen > nodeLen && st.ActiveNode.Child[st.ActiveEdge] != nil {
        st.ActiveNode = st.ActiveNode.Child[st.ActiveEdge]
        st.ActiveEdge = word
        st.ActiveLen = 1
      }
      fmt.Printf("#:%d ActiveEdge:%s ActiveLen:%d Remainder:%d\n", st.CurStep, st.ActiveEdge, st.ActiveLen, st.Remainder)
      continue
    }

    // reminder is larger than 1 and cur word is not match the suffix tnen split the suffix tree
    var lastSplit *TreeNode
    for st.Remainder >= 1 {
      // insert root node child
      if st.Remainder == 1 {
        st.ActiveEdge = ""
        if root.Child[word] == nil {
          child := st.NewChild(st.CurStep, 0)
          root.Child[word] = child
          fmt.Println("add new child " + word)
          fmt.Printf("#:%d ActiveEdge:%s ActiveLen:%d Remainder:%d\n", st.CurStep, st.ActiveEdge, st.ActiveLen, st.Remainder)
        }
        break
      }

      // current split node
      curSplit := st.ActiveNode.Child[st.ActiveEdge]
      st.Split(curSplit, st.Words)

      // update active point
      st.Remainder -= 1
      lastword := []rune(st.Words)[st.CurStep-st.Remainder+1]
      suffixLink := st.ActiveNode.SuffixLink
      if st.ActiveNode == root {
        /*  Rule 1:
            当向根节点插入时遵循：
            active_node 保持为 root；
            active_edge 被设置为即将被插入的新后缀的首字符；
            active_length 减 1；
        */
        st.ActiveEdge = string(lastword)
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
        lastSplit.SuffixLink = curSplit
      }
      lastSplit = curSplit
      fmt.Printf("#:%d ActiveEdge:%s ActiveLen:%d Remainder:%d\n", st.CurStep, st.ActiveEdge, st.ActiveLen, st.Remainder)
    }
  }
  st.Print(st.Words)
}

func (st *SuffixTree) Split(node *TreeNode, str string) {
  sentence := []rune(str)

  if node.Sidx + st.ActiveLen != node.Eidx {
    child1 := st.NewChild(node.Sidx + st.ActiveLen, node.Eidx)
    node.Eidx = node.Sidx + st.ActiveLen
    tmp := child1.Child
    child1.Child = node.Child
    node.Child = tmp
    node.Child[string(sentence[node.Eidx])] = child1
    fmt.Println("add new child " + string(sentence[node.Eidx]))
  }
  child2 := st.NewChild(st.CurStep, 0)
  node.Child[string(sentence[st.CurStep])] = child2

  fmt.Println("add new child " + string(sentence[st.CurStep]))
}

func (st *SuffixTree) NewChild(s, e int) *TreeNode{
  return &TreeNode{
    Sidx: s,
    Eidx: e,
    SuffixLink: nil,
    Child: make(map[string]*TreeNode),
  }
}

func (st *SuffixTree) Reset(){
  st.ActiveNode = nil
  st.ActiveEdge = ""
  st.ActiveLen = 0
  //剩余后缀数（remainder），是一个整数，代表着还需要插入多少个新的后缀
  st.Remainder = 1
}

func (st *SuffixTree) Print(str string) {
  sentence := []rune(str)
  print(st.Root, sentence, 0, st.CurStep+1)
  fmt.Printf("len%*s->%d %d\n", 1*4, "", len(st.Root.Child), len(str))
}

func print(root *TreeNode, sentence []rune, ph, idx int) {
  for _, node := range root.Child {
    if node.Eidx == 0 {
      node.Eidx = idx
    }
    fmt.Printf("%*s->%s\n", ph*4, "", string(sentence[node.Sidx:node.Eidx]))
    print(node, sentence, ph+1, idx)
  }
}
