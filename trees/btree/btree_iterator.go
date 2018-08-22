package btree

import(
  "github.com/jenazads/goutils";
)

func assertIteratorImplementation() {
  var _ goutils.ReverseIteratorKey = (*Iterator)(nil)
}

// Iterator holding the iterator's state
type Iterator struct {
  tree     *BTree
  node     *BNode
  entry    *Entry
  position position
}

type position byte

const (
  begin, between, end position = 0, 1, 2
)

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (t *BTree) Iterator() Iterator {
  return Iterator{tree: t, node: nil, position: begin}
}

// Moves to the next element
func (iterator *Iterator) Next() bool {
  if iterator.position == end {
    goto end
  }
  if iterator.position == begin {
    left := iterator.tree.Left()
    if left == nil {
      goto end
    }
    iterator.node = left
    iterator.entry = left.Entries[0]
    goto between
  }
  {
    e, _ := iterator.tree.search(iterator.node, iterator.entry.Key)
    if e+1 < len(iterator.node.Children) {
      iterator.node = iterator.node.Children[e+1]
      for len(iterator.node.Children) > 0 {
        iterator.node = iterator.node.Children[0]
      }
      iterator.entry = iterator.node.Entries[0]
      goto between
    }
    if e+1 < len(iterator.node.Entries) {
      iterator.entry = iterator.node.Entries[e+1]
      goto between
    }
  }
  for iterator.node.Parent != nil {
    iterator.node = iterator.node.Parent
    e, _ := iterator.tree.search(iterator.node, iterator.entry.Key)
    if e < len(iterator.node.Entries) {
      iterator.entry = iterator.node.Entries[e]
      goto between
    }
  }

end:
  iterator.End()
  return false

between:
  iterator.position = between
  return true
}

// Move to the Prev element
func (iterator *Iterator) Prev() bool {
  if iterator.position == begin {
    goto begin
  }
  if iterator.position == end {
    right := iterator.tree.Right()
    if right == nil {
      goto begin
    }
    iterator.node = right
    iterator.entry = right.Entries[len(right.Entries)-1]
    goto between
  }
  {
    e, _ := iterator.tree.search(iterator.node, iterator.entry.Key)
    if e < len(iterator.node.Children) {
      iterator.node = iterator.node.Children[e]
      for len(iterator.node.Children) > 0 {
        iterator.node = iterator.node.Children[len(iterator.node.Children)-1]
      }
      iterator.entry = iterator.node.Entries[len(iterator.node.Entries)-1]
      goto between
    }
    if e-1 >= 0 {
      iterator.entry = iterator.node.Entries[e-1]
      goto between
    }
  }
  for iterator.node.Parent != nil {
    iterator.node = iterator.node.Parent
    e, _ := iterator.tree.search(iterator.node, iterator.entry.Key)
    if e-1 >= 0 {
      iterator.entry = iterator.node.Entries[e-1]
      goto between
    }
  }

begin:
  iterator.Begin()
  return false

between:
  iterator.position = between
  return true
}

// Return current value
func (iterator *Iterator) Value() interface{} {
  return iterator.entry.Value
}

// Return current Key
func (iterator *Iterator) Key() interface{} {
  return iterator.entry.Key
}

// Set pointer to Begin state
func (iterator *Iterator) Begin() {
  iterator.node = nil
  iterator.position = begin
  iterator.entry = nil
}

// Set pointer to last state
func (iterator *Iterator) End() {
  iterator.node = nil
  iterator.position = end
  iterator.entry = nil
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
