package tree

type NodeInterface interface {
  Search(c int)
  SearchNYT()
}

type Node struct {
  Right *Node
  Left *Node
  Parent *Node
  Weight int
  Info byte
}

func (node Node) Search(c byte) *Node {
  // If current node info matches c
  if (node.Info == c) {
    return &node
  }

  // If is leaf
  if (node.Right == nil) {
    return nil
  }

  // Search on right
  right_node := *(node.Right)
  ptr_result_right := right_node.Search(c)
  if (ptr_result_right != nil) {
    return ptr_result_right
  }

  // Search on left
  left_node := *(node.Left)
  ptr_result_left := left_node.Search(c)
  if (ptr_result_left != nil) {
    return ptr_result_left
  }

  // If not found
  return nil
}

func (node Node) SearchNYT() *Node {
  // If NYT
  if (node.Weight == 0) {
    return &node
  }

  // If leaf
  if (node.Right == nil) {
    return nil
  }

  // Search NYT on right
  right_node := *(node.Right)
  ptr_nyt_right := right_node.SearchNYT()
  if (ptr_nyt_right != nil) {
    return ptr_nyt_right
  }

  // Search NYT on left
  left_node := *(node.Left)
  ptr_nyt_left := left_node.SearchNYT()
  if (ptr_nyt_left != nil) {
    return ptr_nyt_left
  }

  // If not found
  return nil
}
