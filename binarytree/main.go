package binarytree


type BinaryTree struct {
	Value int
	left *BinaryTree
	right *BinaryTree
}


func (b BinaryTree) InsertLeft(value int) BinaryTree {
	tree := BinaryTree{value, nil, nil}
	b.left = &tree

	return tree
}

func (b BinaryTree) InsertRight(value int) BinaryTree {
	tree := BinaryTree{value, nil, nil}
	b.right = &tree

	return tree
}

func (b BinaryTree) GetLeft() BinaryTree {
	return b.left
}

func (b BinaryTree) GetRight() BinaryTree {
	return b.right
}
