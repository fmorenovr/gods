package bstree

import (
  "fmt";
  "github.com/jenazads/goutils";
  "github.com/jenazads/gods/trees";
)

func assertTreeImplementation() {
  var _ gotree.GoTree = new(BSTree)
}

func bstInsert(root *BSTNode, key interface{}, value interface{}, parent *BSTNode, comp goutils.TypeComparator) *BSTNode {
  if root==nil {
    aux:=NewBSTNode(key, parent)
    aux.Value = append(aux.Value, value)
    aux.Parent=parent;
    aux.Count = 1
    root=aux;
    return root
  }
  if comp(key, root.Key) == -1 {
    root.Children[0]=bstInsert(root.Children[0], key, value, root, comp);
  } else if comp(key, root.Key) == 1 {
    root.Children[1]=bstInsert(root.Children[1], key, value, root, comp);
  } else if comp(key, root.Key) == 0 {
    root.Value = append(root.Value, value)
    root.Count = root.Count + 1
  }
  return root;
}

func bstTransplant(u, v *BSTNode) (*BSTNode) {
  v.Parent = u.Parent
  u=v
  return u;
}

func bstRemove(root *BSTNode, key interface{}, comp goutils.TypeComparator) *BSTNode {
  if root == nil {
    return root;
  }
  
  if comp(key, root.Key) == -1 {
    root.Children[0]=bstRemove(root.Children[0], key, comp);
  } else if comp(key, root.Key) == 1 {
    root.Children[1]=bstRemove(root.Children[1], key, comp);
  } else if comp(key, root.Key) == 0 {
    if root.Count > 1 {
      root.Value = root.Value[:len(root.Value)-1]
      root.Count = root.Count - 1
    } else if root.Count == 1 {
      // sin hijos
      if (root.Children[0] == nil) && (root.Children[1] == nil) {
        root = nil
      } else if (root.Children[0] == nil) && (root.Children[1] != nil) {
        root=bstTransplant(root, root.Children[1])
      } else if (root.Children[0] != nil) && (root.Children[1] == nil) {
        root=bstTransplant(root, root.Children[0])
      }else { // 2 hijos agarra el minimo del arbol derecho
        temp:= bstFindNode(root.Children[1], 0)
        root.Key=temp.Key;
        root.Children[1]=bstRemove(root.Children[1], temp.Key, comp);
      }
    }
  }
  // si solo tenia un nodito q era raiz
  return root;
}

func bstSearch(root *BSTNode, key interface{}, comp goutils.TypeComparator) (*BSTNode){
  if root==nil {
    return nil;
  } else {
    curr_node:=root;
    if comp(curr_node.Key, key) == 1 {
      return bstSearch(curr_node.Children[0], key, comp);
    } else if comp(curr_node.Key, key) == -1{
      return bstSearch(curr_node.Children[1], key, comp);
    } else{
      return curr_node;
    }
  }
}

func bstHeight(root *BSTNode) int {
  if root!=nil {
    var a,b int;
    curr_node:=root;
    a=bstHeight(curr_node.Children[0]);
    b=bstHeight(curr_node.Children[1]);
    if a>b {
      return (a+1);
    } else {
      return (b+1);
    }
  } else {
    return (-1);
  }
}

func bstSize(root *BSTNode) int{
  if root==nil {
    return 0;
  } else {
    return (bstSize(root.Children[0]) + bstSize(root.Children[1]) + 1);
  }
}

func bstLeafCount(root *BSTNode) int {
  if root == nil {
    return 0;
  }
  if root.Children[0] == nil && root.Children[1] == nil {
    return 1
  } else {
    return bstLeafCount(root.Children[0]) + bstLeafCount(root.Children[1]);
  }
}

func bstFindNode(root *BSTNode, child int) *BSTNode {
  if root == nil {
    return nil;
  }
  curr_node := root;
  for curr_node.Children[child] != nil { 
    curr_node = curr_node.Children[child];
  }
  return curr_node;
}

func bstSumNodes(root *BSTNode, op goutils.TypeOperator) interface{}{
  var sum interface{} = 0
  if root!=nil {
    sumLeft  := bstSumNodes(root.Children[0], op)
    sumRight := bstSumNodes(root.Children[1], op)
    sum=op(op(sumLeft, root.Key ,"+"), sumRight, "+");
  }
  return sum
}

func bstHeightOfNode(root *BSTNode, key interface{}, comp goutils.TypeComparator) int {
  height:=0;
  curr_node:=root;
  for curr_node!=nil {
    if comp(key, curr_node.Key) == 0 {
      return height;
    } else {
      height++;
      if comp(key, curr_node.Key) == -1 {
        curr_node = curr_node.Children[0];
      } else if comp(key, curr_node.Key) == 1 {
        curr_node = curr_node.Children[1];
      }
    }
  }
  return -1;
}

func bstPrintPreorder(root *BSTNode) {
  if root != nil {
    fmt.Printf("%v ", root.Key);
    bstPrintPreorder(root.Children[0]);
    bstPrintPreorder(root.Children[1]);
  }
}

func bstPrintInorder(root *BSTNode) {
  if root != nil {
    bstPrintInorder(root.Children[0]);
    fmt.Printf("%v ", root.Key);
    bstPrintInorder(root.Children[1]);
  }
}

func bstPrintPostorder(root *BSTNode) {
  if root != nil {
    bstPrintPostorder(root.Children[0]);
    bstPrintPostorder(root.Children[1]);
    fmt.Printf("%v ", root.Key);
  }
}

func getParentNode(root *BSTNode, key interface{}, comp goutils.TypeComparator) *BSTNode{
  curr_node:=bstSearch(root, key, comp);
  if curr_node!=nil {
    return curr_node.Parent;
  } else{
    return nil;
  }
}

func getBrotherNode(root *BSTNode, key interface{}, comp goutils.TypeComparator) *BSTNode{
  curr_node:=bstSearch(root, key, comp);
  if curr_node!=nil && curr_node.Parent!=nil && curr_node.Parent.Children[1]==curr_node {
    return curr_node.Parent.Children[0]
  } else if curr_node!=nil && curr_node.Parent!=nil && curr_node.Parent.Children[0]==curr_node {
    return curr_node.Parent.Children[1]
  } else if curr_node!=nil && curr_node.Parent==nil {
    return nil;
  } else{
    return nil;
  }
}

func bstFindNeighbourNode(root *BSTNode,child int) *BSTNode {
  if root==nil {
    return nil
  }
  if root.Children[child]!=nil {
    return (bstFindNode(root.Children[child], child^1));
  }
  curr_node:=root.Parent;
  for curr_node!=nil && root==curr_node.Children[child] {
    root=curr_node;
    curr_node=curr_node.Parent;
  }
  return curr_node;
}

func bstMirror(root *BSTNode) {
  if root !=nil {
    bstMirror(root.Children[0]);
    bstMirror(root.Children[1]);
    temp := root.Children[0];
    root.Children[0] = root.Children[1];
    root.Children[1]=temp;
  }
}

func bstIsSameAs(a,b *BSTNode) bool {
  if a==nil && b==nil {
    return true;
  } else if a!=nil && b!=nil && a.Key == b.Key {
    return (bstIsSameAs(a.Children[0], b.Children[0]) && bstIsSameAs(a.Children[1], b.Children[1]));
  } else {
    return false;
  }
}

func bstTreePrint(node *BSTNode, prefix string, isTail bool, str *string) {
  if node.Children[1] != nil {
    newPrefix := prefix
    if node.Parent == nil {
      newPrefix += "     "
    } else {
      if isTail {
        newPrefix += "│    "
      } else {
        newPrefix += "     "
      }
    }
    bstTreePrint(node.Children[1], newPrefix, false, str)
  }
  *str += prefix
  if node.Parent == nil {
    *str += "nil──"
  } else {
    if isTail {
      *str += "└────"
    } else {
      *str += "┌────"
    }
  }
  *str += node.String() + "\n"
  if node.Children[0] != nil {
    newPrefix := prefix
    if node.Parent == nil {
      newPrefix += "     "
    } else {
      if isTail {
        newPrefix += "     "
      } else {
        newPrefix += "│    "
      }
    }
    bstTreePrint(node.Children[0], newPrefix, true, str)
  }
}

func bstTreePrintLevels(root *BSTNode) {
  height:=bstHeight(root)+1;
  for i := 0;i < height;i++ {
    fmt.Printf("\nNIVEL %v  :", i);
    bstPrintNodeAtLevel(root, i, 0);
    fmt.Println();
  }
}

func bstPrintNodeAtLevel(root *BSTNode, height, level int) {
  if root!=nil {
    if(height==level) {
      fmt.Printf("\t%v", root.Key)
    } else {
      bstPrintNodeAtLevel(root.Children[0], height, level+1);
      bstPrintNodeAtLevel(root.Children[1], height, level+1);
    }
  }
}
