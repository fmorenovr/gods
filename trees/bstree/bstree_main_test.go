package bstree_test

import (
  "fmt"
  "github.com/jenazads/gods/trees/bstree"
  "github.com/jenazads/goutils"
)

func Example_bsTree() {
  tree := bstree.NewBSTree(goutils.IntComparator, goutils.IntOperator)


  tree.Insert(5, "e")
  tree.Insert(2, "b")
  tree.Insert(3, "c")
  tree.Insert(1, "a")
  tree.Insert(4, "d")
  tree.Insert(6, "f")
  tree.Insert(7, "g")
  tree.Insert(8, "h")

  tree.PrintPreOrder()
  tree.PrintInOrder()
  tree.PrintPostOrder()

  fmt.Println(tree.Ceiling(4))
  fmt.Println(tree.Floor(5))
  tree.Print()

  fmt.Println(tree.Values())
  fmt.Println(tree.Keys())

  fmt.Print(tree)
  tree.Remove(3)
  fmt.Println(tree)
  tree.Remove(2)
  fmt.Println(tree)
  tree.Remove(7)
  fmt.Println(tree)
  tree.Remove(5)
  fmt.Println(tree)
  tree.Remove(4)
  fmt.Println(tree)
  tree.Remove(1)
  fmt.Println(tree)
  tree.Remove(6)
  fmt.Println(tree)
  tree.Remove(8)
  fmt.Println(tree)

  tree.Clear()
  tree.IsEmpty()
  tree.Size()
}
