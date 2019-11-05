package main


import (
	"fmt"
	"binarytree"
)

type Ancestor struct {
	Node *binarytree.BinaryTree
	LeftChecked bool
}

func main() {
	tree := binarytree.BinaryTree{1, nil, nil}
	left := tree.InsertLeft(2)
	right := tree.InsertRight(3)
	leftleft := left.InsertLeft(4)
	left.InsertRight(5)
	right.InsertLeft(6)
	right.InsertRight(7)

	leftleft.InsertLeft(8)
	leftleft.InsertRight(9)
	fmt.Printf("%v\n", tree)

	fmt.Printf("PreOrderTransversal:%v\n", PreOrderTransversal(&tree))
}


//root-Left-Right
func PreOrderTransversal(root *binarytree.BinaryTree) []int {
	output := make([]int, 0)
	ancestors := make([]*Ancestor, 0)
	var ancestor *Ancestor

        for root != nil || ancestor != nil {
		if root != nil{
		        if root.GetLeft() == nil && root.GetRight() ==  nil {
				output = append(output, root.Value)
                                if len(ancestors) > 0 {
				        ancestor, ancestors, root = ancestors[len(ancestors)-1], ancestors[:len(ancestors)-1], nil
				} else {
					root = nil
				}
		        } else if root.GetLeft() != nil {
				output = append(output, root.Value)
			        ancestors = append(ancestors, &Ancestor{root, true})
			        root =  root.GetLeft()
		        } else if root.GetRight() != nil {
                                output = append(output, root.Value)
			        ancestors = append(ancestors, &Ancestor{root, false})
			        root =  root.GetRight()
			}
	        } else {
			if ancestor.Node.GetLeft() == nil && ancestor.Node.GetRight() == nil {
                                if len(ancestors) > 0 {
				        ancestor, ancestors, root = ancestors[len(ancestors)-1], ancestors[:len(ancestors)-1], nil
			        }
			} else if  ancestor.Node.GetLeft() != nil && ancestor.LeftChecked == false {
				root, ancestor = ancestor.Node.GetLeft(), nil
				output = append(output, root.Value)
				//anc = Ancestor{root, false}
				ancestors = append(ancestors, &Ancestor{root, false})

			} else if ancestor.Node.GetRight() != nil && ancestor.LeftChecked {
				root, ancestor = ancestor.Node.GetRight(), nil
				//output = append(output, root.Value)
                                //anc = Ancestor{root, false}
				//ancestors = append(ancestors, &anc)
				/*if root.Value ==  5 {
					break
				}*/
			}
		}
	}

	return output
}

