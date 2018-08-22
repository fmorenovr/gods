package btree

import (
//  "fmt";
//  "github.com/jenazads/goutils";
  "github.com/jenazads/gods/trees";
)

func assertTreeImplementation() {
  var _ gotree.GoTree = (*BTree)(nil)
}

// Entry is the key-value pair in node
type Entry struct {
  Key   interface{}
  Value interface{}
}
