package avltree

import (
  "fmt";
  "github.com/jenazads/goutils";
)

// AVLTree object
type AVLTree struct {
  Root       *AVLNode                // Root node
  comparator goutils.TypeComparator  // Key comparator
  operator   goutils.TypeOperator    // Type Operator
}

// Node 
type AVLNode struct {
  Key      interface{}
  Value    interface{}
  Parent   *AVLNode    // Parent node
  Children [2]*AVLNode // Children nodes, 0-> left, 1-> right
  bf       int         // balance factor
}

// New AVL Tree
func NewAVLTree(comp goutils.TypeComparator, op goutils.TypeOperator) (*AVLTree) {
  return &AVLTree{comparator: comp, operator: op}
}

// New AVL Node
func NewAVLNode(key interface{}, value interface{}, p *AVLNode) (*AVLNode) {
  return &AVLNode{Key: key, Value: value, Parent: p, bf: 0}
}

// IsEmpty, true if tree doesnt have nodes
func (t *AVLTree) IsEmpty() (bool) {
  return (t.Root == nil)
}

// Return true if the node is leaf
func IsLeaf(node *AVLNode) (bool) {
  return (node.Children[0]==nil && node.Children[1]==nil);
}

// Removes all nodes
func (t *AVLTree) Clear() {
  t.Root = nil
}

// Insert New Node by Key
func (t *AVLTree) Insert(key interface{}, value interface{}) {
  t.Root = avlInsert(t.Root, key, value, nil, t.comparator)
}

// Remove Node by key
func (t *AVLTree) Remove(key interface{}) {
  t.Root = avlRemove(t.Root, key, t.comparator)
}

// Search Value, return the node
func (t *AVLTree) Search(key interface{}) (*AVLNode) {
  return avlSearch(t.Root, key, t.comparator)
}

// return AVL Tree Height
func (t *AVLTree) Height() (int) {
  return avlHeight(t.Root)
}

// Return Size of tree
func (t *AVLTree) Size() (int) {
  return avlSize(t.Root)
}

// Return number of Leaf
func (t *AVLTree) LeafCount() (int) {
  return avlLeafCount(t.Root)
}

// Return minimum element
func (t *AVLTree) Left() (*AVLNode) {
  return avlFindNode(t.Root, 0)
}

// Return maximum element
func (t *AVLTree) Right() (*AVLNode) {
  return avlFindNode(t.Root, 1)
}

// Return sum of all nodes
func (t *AVLTree) SumNodes() (interface{}) {
  return avlSumNodes(t.Root, t.operator)
}

// Return the height of a specific node
func (t *AVLTree) HeightOfNode(key interface{}) (int) {
  return avlHeightOfNode(t.Root, key, t.comparator)
}

// print preorder
func (t *AVLTree) PrintPreOrder() {
  avlPrintPreorder(t.Root)
  fmt.Println()
}

// print inorder
func (t *AVLTree) PrintInOrder() {
  avlPrintInorder(t.Root)
  fmt.Println()
}

// print postorder
func (t *AVLTree) PrintPostOrder() {
  avlPrintPostorder(t.Root)
  fmt.Println()
}

// Return the parent node of a specific value
func (t *AVLTree) Parent(key interface{}) (*AVLNode) {
  return getParentNode(t.Root, key, t.comparator)
}

// Return the brother of a specific value
func (t *AVLTree) Brother(key interface{}) (*AVLNode) {
  return getBrotherNode(t.Root, key, t.comparator)
}

// Return previous or predecessor node
func (node *AVLNode) Prev() (*AVLNode) {
  return avlFindNeighbourNode(node, 0)
}

// Return next or sucessor node
func (node *AVLNode) Next() (*AVLNode) {
  return avlFindNeighbourNode(node, 1)
}

// Return avl mirror
func (t *AVLTree) Mirror() {
  avlMirror(t.Root)
}

// Compare 2 AVLTree
func (t *AVLTree) IsSameAs(otherTree *AVLTree) (bool) {
  return avlIsSameAs(t.Root, otherTree.Root)
}

// Keys returns all keys in-order
func (t *AVLTree) Keys() ([]interface{}) {
  size:=t.Size()
  keys := make([]interface{}, size)
  it := t.Iterator()
  for i := 0; it.Next(); i++ {
    keys[i] = it.Key()
  }
  return keys
}

// Values returns all values in-order based on the key.
func (t *AVLTree) Values() ([]interface{}) {
  size:=t.Size()
  values := make([]interface{}, size)
  it := t.Iterator()
  for i := 0; it.Next(); i++ {
    values[i] = it.Value()
  }
  return values
}

// Print the AVLTree
func (t *AVLTree) Print() {
  avlTreePrintLevels(t.Root)
}

// Print Tree with fmt.Print*
func (t *AVLTree) String() (string) {
  str := ""
	if !t.IsEmpty() {
		avlTreePrint(t.Root, "", false, &str)
	}
	return str
}

// Print Node with fmt.Print*
func (node *AVLNode) String() (string) {
  return fmt.Sprintf("%v(%v)", node.Key, node.bf)
}

// Return the largest node that is smaller than or equal to the given node.
func (t *AVLTree) Floor(key interface{}) (*AVLNode) {
  node := t.Search(key)
  return node.Next()
}

// Return the smallest node that is larger than or equal to the given node.
func (t *AVLTree) Ceiling(key interface{}) (*AVLNode) {
  node := t.Search(key)
  return node.Prev()
}
