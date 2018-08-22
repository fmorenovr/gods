package btree

import (
  "bytes"
  "fmt"
  "github.com/jenazads/goutils"
  "strings"
)

// B-Tree object
type BTree struct {
  Root       *BNode                  // Root node
  comparator goutils.TypeComparator  // Key comparator
  size       int                     // Total number of keys in the tree
  m          int                     // order (maximum number of children)
}

// Node
type BNode struct {
  Parent   *BNode
  Entries  []*Entry  // Contained keys in node
  Children []*BNode   // Children nodes
}

// New B Tree
func NewBTree(order int, comp goutils.TypeComparator) *BTree {
  if order < 3 {
    panic("Invalid order, should be at least 3")
  }
  return &BTree{m: order, comparator: comp}
}

// New BTree Node
func NewBNode(p *BNode, entry []*Entry, child []*BNode) (*BNode) {
  return &BNode{Parent: p, Entries: entry, Children: child}
}

// IsEmpty, true if tree doesnt have nodes
func (t *BTree) IsEmpty() bool {
  return t.size == 0
}

// Return true if the node is leaf
func IsLeaf(node *BNode) bool {
  return len(node.Children) == 0
}

// Removes all nodes
func (t *BTree) Clear() {
  t.Root = nil
  t.size = 0
}

// Put inserts key-value pair node into the tree
func (t *BTree) Put(key interface{}, value interface{}) {
  entry := &Entry{Key: key, Value: value}

  if t.Root == nil {
    t.Root = NewBNode(nil, []*Entry{entry}, []*BNode{})
    t.size++
    return
  }

  if t.insert(t.Root, entry) {
    t.size++
  }
}

// Remove Node by key
func (t *BTree) Remove(key interface{}) {
  node, index, found := t.searchRecursively(t.Root, key)
  if found {
    t.delete(node, index)
    t.size--
  }
}

// Get Value
func (t *BTree) Get(key interface{}) (interface{}) {
  node, index, found := t.searchRecursively(t.Root, key)
  if found {
    return node.Entries[index].Value
  }
  return nil
}

// Returns the height
func (t *BTree) Height() int {
  return t.Root.height()
}

// Return Size of tree
func (t *BTree) Size() int {
  return t.size
}

// Return minimum element
func (t *BTree) Left() *BNode {
  return t.left(t.Root)
}

// Return maximum element
func (t *BTree) Right() *BNode {
  return t.right(t.Root)
}

// Keys returns all keys in-order
func (t *BTree) Keys() []interface{} {
  keys := make([]interface{}, t.size)
  it := t.Iterator()
  for i := 0; it.Next(); i++ {
    keys[i] = it.Key()
  }
  return keys
}

// Values returns all values in-order based on the key.
func (t *BTree) Values() []interface{} {
  values := make([]interface{}, t.size)
  it := t.Iterator()
  for i := 0; it.Next(); i++ {
    values[i] = it.Value()
  }
  return values
}

// LeftKey returns the left-most (min) key or nil if tree is empty.
func (t *BTree) LeftKey() interface{} {
  if left := t.Left(); left != nil {
    return left.Entries[0].Key
  }
  return nil
}

// LeftValue returns the left-most value or nil if tree is empty.
func (t *BTree) LeftValue() interface{} {
  if left := t.Left(); left != nil {
    return left.Entries[0].Value
  }
  return nil
}

// RightKey returns the right-most (max) key or nil if tree is empty.
func (t *BTree) RightKey() interface{} {
  if right := t.Right(); right != nil {
    return right.Entries[len(right.Entries)-1].Key
  }
  return nil
}

// RightValue returns the right-most value or nil if tree is empty.
func (t *BTree) RightValue() interface{} {
  if right := t.Right(); right != nil {
    return right.Entries[len(right.Entries)-1].Value
  }
  return nil
}

// String returns a string representation of container (for debugging purposes)
func (t *BTree) String() string {
  var buffer bytes.Buffer
  if _, err := buffer.WriteString("BTree\n"); err != nil {
  }
  if !t.IsEmpty() {
    t.output(&buffer, t.Root, 0, true)
  }
  return buffer.String()
}

func (entry *Entry) String() string {
  return fmt.Sprintf("%v", entry.Key)
}

func (t *BTree) output(buffer *bytes.Buffer, node *BNode, level int, isTail bool) {
  for e := 0; e < len(node.Entries)+1; e++ {
    if e < len(node.Children) {
      t.output(buffer, node.Children[e], level+1, true)
    }
    if e < len(node.Entries) {
      if _, err := buffer.WriteString(strings.Repeat("    ", level)); err != nil {
      }
      if _, err := buffer.WriteString(fmt.Sprintf("%v", node.Entries[e].Key) + "\n"); err != nil {
      }
    }
  }
}

func (node *BNode) height() int {
  height := 0
  for ; node != nil; node = node.Children[0] {
    height++
    if len(node.Children) == 0 {
      break
    }
  }
  return height
}

func (t *BTree) isFull(node *BNode) bool {
  return len(node.Entries) == t.maxEntries()
}

func (t *BTree) shouldSplit(node *BNode) bool {
  return len(node.Entries) > t.maxEntries()
}

func (t *BTree) maxChildren() int {
  return t.m
}

func (t *BTree) minChildren() int {
  return (t.m + 1) / 2 // ceil(m/2)
}

func (t *BTree) maxEntries() int {
  return t.maxChildren() - 1
}

func (t *BTree) minEntries() int {
  return t.minChildren() - 1
}

func (t *BTree) middle() int {
  return (t.m - 1) / 2 // "-1" to favor right nodes to have more keys when splitting
}

// search searches only within the single node among its entries
func (t *BTree) search(node *BNode, key interface{}) (index int, found bool) {
  low, high := 0, len(node.Entries)-1
  var mid int
  for low <= high {
    mid = (high + low) / 2
    compare := t.comparator(key, node.Entries[mid].Key)
    switch {
    case compare > 0:
      low = mid + 1
    case compare < 0:
      high = mid - 1
    case compare == 0:
      return mid, true
    }
  }
  return low, false
}

// searchRecursively searches recursively down the tree starting at the startNode
func (t *BTree) searchRecursively(startNode *BNode, key interface{}) (node *BNode, index int, found bool) {
  if t.IsEmpty() {
    return nil, -1, false
  }
  node = startNode
  for {
    index, found = t.search(node, key)
    if found {
      return node, index, true
    }
    if IsLeaf(node) {
      return nil, -1, false
    }
    node = node.Children[index]
  }
}

func (t *BTree) insert(node *BNode, entry *Entry) (inserted bool) {
  if IsLeaf(node) {
    return t.insertIntoLeaf(node, entry)
  }
  return t.insertIntoInternal(node, entry)
}

func (t *BTree) insertIntoLeaf(node *BNode, entry *Entry) (inserted bool) {
  insertPosition, found := t.search(node, entry.Key)
  if found {
    node.Entries[insertPosition] = entry
    return false
  }
  // Insert entry's key in the middle of the node
  node.Entries = append(node.Entries, nil)
  copy(node.Entries[insertPosition+1:], node.Entries[insertPosition:])
  node.Entries[insertPosition] = entry
  t.split(node)
  return true
}

func (t *BTree) insertIntoInternal(node *BNode, entry *Entry) (inserted bool) {
  insertPosition, found := t.search(node, entry.Key)
  if found {
    node.Entries[insertPosition] = entry
    return false
  }
  return t.insert(node.Children[insertPosition], entry)
}

func (t *BTree) split(node *BNode) {
  if !t.shouldSplit(node) {
    return
  }

  if node == t.Root {
    t.splitRoot()
    return
  }

  t.splitNonRoot(node)
}

func (t *BTree) splitNonRoot(node *BNode) {
  middle := t.middle()
  parent := node.Parent

  left := NewBNode(parent, append([]*Entry(nil), node.Entries[:middle]...), nil)
  right := NewBNode(parent, append([]*Entry(nil), node.Entries[middle+1:]...), nil)

  // Move children from the node to be split into left and right nodes
  if !IsLeaf(node) {
    left.Children = append([]*BNode(nil), node.Children[:middle+1]...)
    right.Children = append([]*BNode(nil), node.Children[middle+1:]...)
    setParent(left.Children, left)
    setParent(right.Children, right)
  }

  insertPosition, _ := t.search(parent, node.Entries[middle].Key)

  // Insert middle key into parent
  parent.Entries = append(parent.Entries, nil)
  copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
  parent.Entries[insertPosition] = node.Entries[middle]

  // Set child left of inserted key in parent to the created left node
  parent.Children[insertPosition] = left

  // Set child right of inserted key in parent to the created right node
  parent.Children = append(parent.Children, nil)
  copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
  parent.Children[insertPosition+1] = right

  t.split(parent)
}

func (t *BTree) splitRoot() {
  middle := t.middle()

  left := NewBNode(nil, append([]*Entry(nil), t.Root.Entries[:middle]...), nil)
  right := NewBNode(nil, append([]*Entry(nil), t.Root.Entries[middle+1:]...), nil)

  // Move children from the node to be split into left and right nodes
  if !IsLeaf(t.Root) {
    left.Children = append([]*BNode(nil), t.Root.Children[:middle+1]...)
    right.Children = append([]*BNode(nil), t.Root.Children[middle+1:]...)
    setParent(left.Children, left)
    setParent(right.Children, right)
  }

  // Root is a node with one entry and two children (left and right)
  newRoot := NewBNode(nil, []*Entry{t.Root.Entries[middle]}, []*BNode{left, right})

  left.Parent = newRoot
  right.Parent = newRoot
  t.Root = newRoot
}

func setParent(nodes []*BNode, parent *BNode) {
  for _, node := range nodes {
    node.Parent = parent
  }
}

func (t *BTree) left(node *BNode) *BNode {
  if t.IsEmpty() {
    return nil
  }
  current := node
  for {
    if IsLeaf(current) {
      return current
    }
    current = current.Children[0]
  }
}

func (t *BTree) right(node *BNode) *BNode {
  if t.IsEmpty() {
    return nil
  }
  current := node
  for {
    if IsLeaf(current) {
      return current
    }
    current = current.Children[len(current.Children)-1]
  }
}

// leftSibling returns the node's left sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
func (t *BTree) leftSibling(node *BNode, key interface{}) (*BNode, int) {
  if node.Parent != nil {
    index, _ := t.search(node.Parent, key)
    index--
    if index >= 0 && index < len(node.Parent.Children) {
      return node.Parent.Children[index], index
    }
  }
  return nil, -1
}

// rightSibling returns the node's right sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
func (t *BTree) rightSibling(node *BNode, key interface{}) (*BNode, int) {
  if node.Parent != nil {
    index, _ := t.search(node.Parent, key)
    index++
    if index < len(node.Parent.Children) {
      return node.Parent.Children[index], index
    }
  }
  return nil, -1
}

// delete deletes an entry in node at entries' index
// ref.: https://en.wikipedia.org/wiki/B-tree#Deletion
func (t *BTree) delete(node *BNode, index int) {
  // deleting from a leaf node
  if IsLeaf(node) {
    deletedKey := node.Entries[index].Key
    t.deleteEntry(node, index)
    t.rebalance(node, deletedKey)
    if len(t.Root.Entries) == 0 {
      t.Root = nil
    }
    return
  }

  // deleting from an internal node
  leftLargestNode := t.right(node.Children[index]) // largest node in the left sub-tree (assumed to exist)
  leftLargestEntryIndex := len(leftLargestNode.Entries) - 1
  node.Entries[index] = leftLargestNode.Entries[leftLargestEntryIndex]
  deletedKey := leftLargestNode.Entries[leftLargestEntryIndex].Key
  t.deleteEntry(leftLargestNode, leftLargestEntryIndex)
  t.rebalance(leftLargestNode, deletedKey)
}

// rebalance rebalances the tree after deletion if necessary and returns true, otherwise false.
// Note that we first delete the entry and then call rebalance, thus the passed deleted key as reference.
func (t *BTree) rebalance(node *BNode, deletedKey interface{}) {
  // check if rebalancing is needed
  if node == nil || len(node.Entries) >= t.minEntries() {
    return
  }

  // try to borrow from left sibling
  leftSibling, leftSiblingIndex := t.leftSibling(node, deletedKey)
  if leftSibling != nil && len(leftSibling.Entries) > t.minEntries() {
    // rotate right
    node.Entries = append([]*Entry{node.Parent.Entries[leftSiblingIndex]}, node.Entries...) // prepend parent's separator entry to node's entries
    node.Parent.Entries[leftSiblingIndex] = leftSibling.Entries[len(leftSibling.Entries)-1]
    t.deleteEntry(leftSibling, len(leftSibling.Entries)-1)
    if !IsLeaf(leftSibling) {
      leftSiblingRightMostChild := leftSibling.Children[len(leftSibling.Children)-1]
      leftSiblingRightMostChild.Parent = node
      node.Children = append([]*BNode{leftSiblingRightMostChild}, node.Children...)
      t.deleteChild(leftSibling, len(leftSibling.Children)-1)
    }
    return
  }

  // try to borrow from right sibling
  rightSibling, rightSiblingIndex := t.rightSibling(node, deletedKey)
  if rightSibling != nil && len(rightSibling.Entries) > t.minEntries() {
    // rotate left
    node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1]) // append parent's separator entry to node's entries
    node.Parent.Entries[rightSiblingIndex-1] = rightSibling.Entries[0]
    t.deleteEntry(rightSibling, 0)
    if !IsLeaf(rightSibling) {
      rightSiblingLeftMostChild := rightSibling.Children[0]
      rightSiblingLeftMostChild.Parent = node
      node.Children = append(node.Children, rightSiblingLeftMostChild)
      t.deleteChild(rightSibling, 0)
    }
    return
  }

  // merge with siblings
  if rightSibling != nil {
    // merge with right sibling
    node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1])
    node.Entries = append(node.Entries, rightSibling.Entries...)
    deletedKey = node.Parent.Entries[rightSiblingIndex-1].Key
    t.deleteEntry(node.Parent, rightSiblingIndex-1)
    t.appendChildren(node.Parent.Children[rightSiblingIndex], node)
    t.deleteChild(node.Parent, rightSiblingIndex)
  } else if leftSibling != nil {
    // merge with left sibling
    entries := append([]*Entry(nil), leftSibling.Entries...)
    entries = append(entries, node.Parent.Entries[leftSiblingIndex])
    node.Entries = append(entries, node.Entries...)
    deletedKey = node.Parent.Entries[leftSiblingIndex].Key
    t.deleteEntry(node.Parent, leftSiblingIndex)
    t.prependChildren(node.Parent.Children[leftSiblingIndex], node)
    t.deleteChild(node.Parent, leftSiblingIndex)
  }

  // make the merged node the root if its parent was the root and the root is empty
  if node.Parent == t.Root && len(t.Root.Entries) == 0 {
    t.Root = node
    node.Parent = nil
    return
  }

  // parent might underflow, so try to rebalance if necessary
  t.rebalance(node.Parent, deletedKey)
}

func (t *BTree) prependChildren(fromNode *BNode, toNode *BNode) {
  children := append([]*BNode(nil), fromNode.Children...)
  toNode.Children = append(children, toNode.Children...)
  setParent(fromNode.Children, toNode)
}

func (t *BTree) appendChildren(fromNode *BNode, toNode *BNode) {
  toNode.Children = append(toNode.Children, fromNode.Children...)
  setParent(fromNode.Children, toNode)
}

func (t *BTree) deleteEntry(node *BNode, index int) {
  copy(node.Entries[index:], node.Entries[index+1:])
  node.Entries[len(node.Entries)-1] = nil
  node.Entries = node.Entries[:len(node.Entries)-1]
}

func (t *BTree) deleteChild(node *BNode, index int) {
  if index >= len(node.Children) {
    return
  }
  copy(node.Children[index:], node.Children[index+1:])
  node.Children[len(node.Children)-1] = nil
  node.Children = node.Children[:len(node.Children)-1]
}
