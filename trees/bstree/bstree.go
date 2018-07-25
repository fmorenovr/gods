package bstree

import (
  "fmt";
  "github.com/jenazads/goutils";
)

// Tree holds elements of the BS tree.
type BSTree struct {
  Root       *BSTNode                // Root node
  comparator goutils.TypeComparator  // Key comparator
  operator   goutils.TypeOperator    // Type Operator
}

// Node 
type BSTNode struct {
  Key      interface{}
  Value    interface{}
  Parent   *BSTNode    // Parent node
  Children [2]*BSTNode // Children nodes, 0-> left, 1-> right
}

// New BS Tree
func NewBSTree(comp goutils.TypeComparator, op goutils.TypeOperator) (*BSTree) {
  return &BSTree{comparator: comp, operator: op}
}

// New BS Node
func NewBSTNode(key interface{}, value interface{}, p *BSTNode) (*BSTNode) {
  return &BSTNode{Key: key, Value: value, Parent: p}
}

// IsEmpty, true if tree doesnt have nodes
func (t *BSTree) IsEmpty() (bool) {
  return (t.Root == nil)
}

// Return true if the node is leaf
func IsLeaf(node *BSTNode) (bool) {
  return (node.Children[0]==nil && node.Children[1]==nil);
}

// Removes all nodes
func (t *BSTree) Clear() {
  t.Root = nil
}

// Insert New Node by Key
func (t *BSTree) Insert(key interface{}, value interface{}) {
  t.Root = bstInsert(t.Root, key, value, nil, t.comparator)
}

// Remove Node by key
func (t *BSTree) Remove(key interface{}) {
  t.Root = bstRemove(t.Root, key, t.comparator)
}

// Search Value, return the node
func (t *BSTree) Search(key interface{}) (*BSTNode) {
  return bstSearch(t.Root, key, t.comparator)
}

// return BS Tree Height
func (t *BSTree) Height() (int) {
  return bstHeight(t.Root)
}

// Return Size of tree
func (t *BSTree) Size() (int) {
  return bstSize(t.Root)
}

// Return number of Leaf
func (t *BSTree) LeafCount() (int) {
  return bstLeafCount(t.Root)
}

// Return minimum element
func (t *BSTree) Left() (*BSTNode) {
  return bstFindNode(t.Root, 0)
}

// Return maximum element
func (t *BSTree) Right() (*BSTNode) {
  return bstFindNode(t.Root, 1)
}

// Return sum of all nodes
func (t *BSTree) SumNodes() (interface{}) {
  return bstSumNodes(t.Root, t.operator)
}

// Return the height of a specific node
func (t *BSTree) HeightOfNode(key interface{}) (int) {
  return bstHeightOfNode(t.Root, key, t.comparator)
}

// print preorder
func (t *BSTree) PrintPreOrder() {
  bstPrintPreorder(t.Root)
  fmt.Println()
}

// print inorder
func (t *BSTree) PrintInOrder() {
  bstPrintInorder(t.Root)
  fmt.Println()
}

// print postorder
func (t *BSTree) PrintPostOrder() {
  bstPrintPostorder(t.Root)
  fmt.Println()
}

// Return the parent node of a specific value
func (t *BSTree) Parent(key interface{}) (*BSTNode) {
  return getParentNode(t.Root, key, t.comparator)
}

// Return the brother of a specific value
func (t *BSTree) Brother(key interface{}) (*BSTNode) {
  return getBrotherNode(t.Root, key, t.comparator)
}

// Return previous or predecessor node
func (node *BSTNode) Prev() (*BSTNode) {
  return bstFindNeighbourNode(node, 0)
}

// Return next or sucessor node
func (node *BSTNode) Next() (*BSTNode) {
  return bstFindNeighbourNode(node, 1)
}

// Return bst mirror
func (t *BSTree) Mirror() {
  bstMirror(t.Root)
}

// Compare 2 BSTree
func (t *BSTree) IsSameAs(otherTree *BSTree) (bool) {
  return bstIsSameAs(t.Root, otherTree.Root)
}

// Keys returns all keys in-order
func (t *BSTree) Keys() ([]interface{}) {
  size:=t.Size()
  keys := make([]interface{}, size)
  it := t.Iterator()
  for i := 0; it.Next(); i++ {
    keys[i] = it.Key()
  }
  return keys
}

// Values returns all values in-order based on the key.
func (t *BSTree) Values() ([]interface{}) {
  size:=t.Size()
  values := make([]interface{}, size)
  it := t.Iterator()
  for i := 0; it.Next(); i++ {
    values[i] = it.Value()
  }
  return values
}

// Print the BSTree
func (t *BSTree) Print() {
  bstTreePrintLevels(t.Root)
}

// Print Tree with fmt.Print*
func (t *BSTree) String() (string) {
  str := ""
	if !t.IsEmpty() {
		bstTreePrint(t.Root, "", false, &str)
	}
	return str
}

// Print Node with fmt.Print*
func (node *BSTNode) String() (string) {
  return fmt.Sprintf("%v", node.Key)
}

// Return the largest node that is smaller than or equal to the given node.
func (t *BSTree) Floor(key interface{}) (*BSTNode) {
  node := t.Search(key)
  return node.Next()
}

// Return the smallest node that is larger than or equal to the given node.
func (t *BSTree) Ceiling(key interface{}) (*BSTNode) {
  node := t.Search(key)
  return node.Prev()
}
