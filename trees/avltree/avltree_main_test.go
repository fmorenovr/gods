package avltree

import (
	"fmt"
	"github.com/jenazads/gods/trees/avltree"
	"github.com/jenazads/goutils"
)

func main() {
	tree := avltree.NewAVLTree(goutils.IntComparator, goutils.IntOperator) // empty(keys are of type int)

	tree.Insert(1, "x") // 1->x
	tree.Insert(2, "b") // 1->x, 2->b (in order)
	tree.Insert(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Insert(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Insert(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Insert(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Insert(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(tree)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 2
	//          └── 1


	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      └── 3
	//          └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0
}
