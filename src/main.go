package main

import (
  "fmt"
  "bufio"
  "os"
  "./tree"
)

func main()  {
  ptr_root := new(tree.Node)
  root := *ptr_root

  // Get message
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Message: ")
  scanner.Scan()
  msg := scanner.Text()

  // Add characters of message
  for i:= range msg {
    // Search character on tree
    ptr_current := root.Search(msg[i])

    // If hasn't occurred
    if (ptr_current == nil) {
      ptr_nyt := root.SearchNYT()
      nyt := *ptr_nyt
      nyt.Weight = 1

      // Parent
      nyt.Parent = ptr_nyt
      // Create NYT
      nyt.Left = new(tree.Node)
      // Add value
      nyt.Right = new(tree.Node)
      right_node := *(nyt.Right)
      right_node.Weight = 1

      // Increment parents weight
      ptr_current := nyt.Parent
      _ = ptr_current
    }

    // Update parent weights
    for ptr_current != nil {
      current := *(ptr_current)
      ptr_parent := current.Parent
      parent := *(ptr_parent)
      ptr_right_node := parent.Right
      right_node := *(ptr_right_node)
      ptr_left_node := parent.Left
      left_node := *(ptr_left_node)

      // If isn't line max, swap
      if (left_node.Weight <= right_node.Weight) {
        parent.Right = ptr_left_node
        parent.Left = ptr_right_node
      }

      current.Weight += 1

      ptr_current := ptr_parent
      _ = ptr_current
    }

    fmt.Printf("Added %c\n", msg[i]);
  }
}
