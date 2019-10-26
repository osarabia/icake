package binarytree


type BinaryTree struct {
	Value int
	Left *BinaryTree
	Right *BinaryTree
}


func (b BinaryTree) InsertLeft(value int) BinaryTree {
	tree := BinaryTree{value, nil, nil}
	b.Left = &tree

	return tree
}

func (b BinaryTree) InsertRight(value int) BinaryTree {
	tree := BinaryTree{value, nil, nil}
	b.Right = &tree

	return tree
}

func (b BinaryTree) GetLeft() BinaryTree {
	return *b.Left
}

func (b BinaryTree) GetRight() BinaryTree {
	return *b.Right
}
