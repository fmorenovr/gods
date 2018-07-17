package avltree

import (
  "fmt";
  "github.com/jenazads/gods/trees";
  "github.com/jenazads/goutils";
)

func assertTreeImplementation() {
  var _ gotree.GoTree = new(AVLTree)
}

func avlBalanceFactor(node *AVLNode) int8 {
  if node == nil {
    return 0;
  }
  return avlHeight(node.Children[1]) - avlHeight(node.Children[0]);
}

func avlRightRotate(y *AVLNode) *AVLNode {
  x:=y.Children[0];
  T2:=x.Children[1];
  
  x.Children[1]=y;
  y.Children[0]=T2;
  x.Parent=y.Parent;
  y.Parent=x;
  if T2!=nil {
    T2.Parent=y;
  }
  
  y.bf=avlBalanceFactor(y);
  x.bf=avlBalanceFactor(x);
  
  return x;
}

func avlLeftRotate(x *AVLNode) *AVLNode {
  y:=x.Children[1]
  T2:=y.Children[0];
  
  y.Children[0]=x;
  x.Children[1]=T2;
  y.Parent=x.Parent;
  x.Parent=y;
  if T2!=nil {
    T2.Parent=x;
  }
  
  y.bf=avlBalanceFactor(y);
  x.bf=avlBalanceFactor(x);
  
  return y;
}

func avlLeftRightRotate(root *AVLNode) *AVLNode {
  root.Children[0] = avlLeftRotate(root.Children[0]);
  return avlRightRotate(root);
}

func avlRightLeftRotate(root *AVLNode) *AVLNode {
  root.Children[1] = avlRightRotate(root.Children[1]);
  return avlLeftRotate(root);
}

func avlInsert(root *AVLNode, key interface{}, value interface{}, parent *AVLNode, comp goutils.TypeComparator) *AVLNode {
  if root==nil {
    aux:=NewAVLNode(key, value, parent)
    aux.Parent=parent;
    root=aux;
    return root
  }
  if comp(key, root.Key) == -1 {
    root.Children[0]=avlInsert(root.Children[0], key, value, root, comp);
  } else if comp(key, root.Key) == 1 {
    root.Children[1]=avlInsert(root.Children[1], key, value, root, comp);
  } else {
    root.Value = value
  }

  root.bf=avlBalanceFactor(root)
  balance:=root.bf
  
  // si esta desbalanceado, se evalua
  
  if balance < -1 { // Left Case
    if comp(key, root.Children[0].Key) == -1 { // Left Case
      return avlRightRotate(root);
    } else if comp(key, root.Children[0].Key) == 1 { // Right Case
      return avlLeftRightRotate(root)
    }
  } else if balance > 1 { // Right Case
    if comp(key, root.Children[1].Key) == 1 { // Right Case
      return avlLeftRotate(root);
    } else if comp(key, root.Children[1].Key) == -1 { // Left Case
      return avlRightLeftRotate(root)
    }
  }
  return root;
}

func avlRemove(root *AVLNode, key interface{}, comp goutils.TypeComparator) *AVLNode {
  if root == nil {
    return root;
  }
  
  if comp(key, root.Key) == -1 {
    root.Children[0]=avlRemove(root.Children[0], key, comp);
  } else if comp(key, root.Key) == 1 {
    root.Children[1]=avlRemove(root.Children[1], key, comp);
  } else {
    // cuando tiene un hijo
    if (root.Children[0] == nil) || (root.Children[1] == nil) {
      // evalua si left es !=null si es cierto tomo left o sino right
      temp:=root.Children[0]
      if temp==nil {
        temp=root.Children[1]
      }
      if temp==nil { // sin hijos
        temp=root;
        root=nil;
      } else { // un hijo
        root=temp;
      }
    } else { // 2 hijos agarra el minimo del arbol derecho
      temp:= avlFindNode(root.Children[1], 0)
      root.Key=temp.Key;
      root.Children[1]=avlRemove(root.Children[1], temp.Key, comp);
    }
  }
  // si solo tenia un nodito q era raiz
  if root==nil {
    return root;
  }
  
  root.bf=avlBalanceFactor(root)
  balance:=root.bf
  
  // si esta desbalanceado, se evalua

  if balance < -1 { // Left Case
    if avlBalanceFactor(root.Children[0]) < 0 { // Left Case
      return avlRightRotate(root);
    } else if avlBalanceFactor(root.Children[0]) >= 0 { // Right Case
      return avlLeftRightRotate(root)
    }
  } else if balance > 1 { // Right Case
    if avlBalanceFactor(root.Children[1]) > 0 { // Right Case
      return avlLeftRotate(root);
    } else if avlBalanceFactor(root.Children[1]) <= 0 { // Left Case
      return avlRightLeftRotate(root)
    }
  }
  
  return root;
}

func avlSearch(root *AVLNode, key interface{}, comp goutils.TypeComparator) (*AVLNode){
  if root==nil {
    return nil;
  } else {
    curr_node:=root;
    if comp(curr_node.Value, key) == 1 {
      return avlSearch(curr_node.Children[0], key, comp);
    } else if comp(curr_node.Value, key) == -1{
      return avlSearch(curr_node.Children[1], key, comp);
    } else {
      return curr_node;
    }
  }
}

func avlHeight(root *AVLNode) int8 {
  if root!=nil {
    var a,b int8;
    curr_node:=root;
    a=avlHeight(curr_node.Children[0]);
    b=avlHeight(curr_node.Children[1]);
    if a>b {
      return (a+1);
    } else {
      return (b+1);
    }
  } else {
    return (-1);
  }
}

func avlSize(root *AVLNode) int{
  if root==nil {
    return 0;
  } else {
    return (avlSize(root.Children[0]) + avlSize(root.Children[1]) + 1);
  }
}

func avlLeafCount(root *AVLNode) int {
  if root == nil {
    return 0;
  }
  if root.Children[0] == nil && root.Children[1] == nil {
    return 1
  } else {
    return avlLeafCount(root.Children[0]) + avlLeafCount(root.Children[1]);
  }
}

func avlFindNode(root *AVLNode, child int) *AVLNode {
  if root == nil {
    return nil;
  }
  curr_node := root;
  for curr_node.Children[child] != nil { 
    curr_node = curr_node.Children[child];
  }
  return curr_node;
}

func avlSumNodes(root *AVLNode, op goutils.TypeOperator) interface{}{
  var sum interface{} = 0
  if root!=nil {
    sumLeft  := avlSumNodes(root.Children[0], op)
    sumRight := avlSumNodes(root.Children[1], op)
    sum=op(op(sumLeft, root.Key ,"+"), sumRight, "+");
  }
  return sum
}

func avlHeightOfNode(root *AVLNode, key interface{}, comp goutils.TypeComparator) int8 {
  var height int8 =0;
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

func avlPreorder_print(root *AVLNode) {
  if root != nil {
    fmt.Printf("%v ", root.Key);
    avlPreorder_print(root.Children[0]);
    avlPreorder_print(root.Children[1]);
  }
}

func avlInorder_print(root *AVLNode) {
  if root != nil {
    avlInorder_print(root.Children[0]);
    fmt.Printf("%v ", root.Key);
    avlInorder_print(root.Children[1]);
  }
}

func avlPostorder_print(root *AVLNode) {
  if root != nil {
    avlPostorder_print(root.Children[0]);
    avlPostorder_print(root.Children[1]);
    fmt.Printf("%v ", root.Key);
  }
}

func getParentNode(root *AVLNode, key interface{}, comp goutils.TypeComparator) *AVLNode{
  curr_node:=avlSearch(root, key, comp);
  if curr_node!=nil {
    return curr_node.Parent;
  } else{
    return nil;
  }
}

func getBrotherNode(root *AVLNode, key interface{}, comp goutils.TypeComparator) *AVLNode{
  curr_node:=avlSearch(root, key, comp);
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

func avlFindNeighbourNode(root *AVLNode,child int) *AVLNode {
  if root==nil {
    return nil
  }
  if root.Children[child]!=nil {
    return (avlFindNode(root.Children[child], child^1));
  }
  curr_node:=root.Parent;
  for curr_node!=nil && root==curr_node.Children[child] {
    root=curr_node;
    curr_node=curr_node.Parent;
  }
  return curr_node;
}

func avlMirror(root *AVLNode) {
  if root !=nil {
    avlMirror(root.Children[0]);
    avlMirror(root.Children[1]);
    temp := root.Children[0];
    root.Children[0] = root.Children[1];
    root.Children[1]=temp;
  }
}

func avlIsSameAs(a,b *AVLNode) bool {
  if a==nil && b==nil {
    return true;
  } else if a!=nil && b!=nil && a.Key == b.Key {
    return (avlIsSameAs(a.Children[0], b.Children[0]) && avlIsSameAs(a.Children[1], b.Children[1]));
  } else {
    return false;
  }
}

func output(node *AVLNode, prefix string, isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Children[0], newPrefix, true, str)
	}
}
