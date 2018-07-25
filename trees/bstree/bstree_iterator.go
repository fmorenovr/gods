package bstree

import(
  "github.com/jenazads/goutils";
)

func assertIteratorImplementation() {
  var _ goutils.ReverseIteratorKey = (*Iterator)(nil)
}

// Iterator
type Iterator struct {
  tree     *BSTree
  node     *BSTNode
  position position
}

type position byte

const (
  begin, between, end position = 0, 1, 2
)

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (t *BSTree) Iterator() goutils.ReverseIteratorKey {
  return &Iterator{tree: t, node: nil, position: begin}
}

// Moves to the next element
func (iterator *Iterator) Next() bool {
  switch iterator.position {
  case begin:
    iterator.position = between
    iterator.node = iterator.tree.Left()
  case between:
    iterator.node = iterator.node.Next()
  }

  if iterator.node == nil {
    iterator.position = end
    return false
  }
  return true
}

// Move to the Prev element
func (iterator *Iterator) Prev() bool {
  switch iterator.position {
  case end:
    iterator.position = between
    iterator.node = iterator.tree.Right()
  case between:
    iterator.node = iterator.node.Prev()
  }

  if iterator.node == nil {
    iterator.position = begin
    return false
  }
  return true
}

// Return current value
func (iterator *Iterator) Value() interface{} {
  if iterator.node == nil {
    return nil
  }
  return iterator.node.Value
}

// Return current Key
func (iterator *Iterator) Key() interface{} {
  if iterator.node == nil {
    return nil
  }
  return iterator.node.Key
}

// Set pointer to Begin state
func (iterator *Iterator) Begin() {
  iterator.node = nil
  iterator.position = begin
}

// Set pointer to last state
func (iterator *Iterator) End() {
  iterator.node = nil
  iterator.position = end
}

// Moves to the first element
func (iterator *Iterator) First() bool {
  iterator.Begin()
  return iterator.Next()
}

// Moves to the Last element
func (iterator *Iterator) Last() bool {
  iterator.End()
  return iterator.Prev()
}
