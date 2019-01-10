package binaryTree

import "testing"

func TestBinaryTreePreOrder(t *testing.T) {
	b0 := &BinaryTree{0, nil, nil}
	b1 := &BinaryTree{1, nil, nil}
	b2 := &BinaryTree{2, nil, nil}
	b3 := &BinaryTree{3, nil, nil}
	b4 := &BinaryTree{4, nil, nil}

	b5 := &BinaryTree{5, b0, b1}
	b6 := &BinaryTree{6, b2, b3}
	b7 := &BinaryTree{7, b4, nil}

	b8 := &BinaryTree{8, b5, b6}
	b9 := &BinaryTree{9, b7, nil}

	b10 := &BinaryTree{10, b8, b9}

	b10.preOrder()
}

func TestBinaryTreeInOrder(t *testing.T) {
	b0 := &BinaryTree{0, nil, nil}
	b1 := &BinaryTree{1, nil, nil}
	b2 := &BinaryTree{2, nil, nil}
	b3 := &BinaryTree{3, nil, nil}
	b4 := &BinaryTree{4, nil, nil}

	b5 := &BinaryTree{5, b0, b1}
	b6 := &BinaryTree{6, b2, b3}
	b7 := &BinaryTree{7, b4, nil}

	b8 := &BinaryTree{8, b5, b6}
	b9 := &BinaryTree{9, b7, nil}

	b10 := &BinaryTree{10, b8, b9}

	b10.inOrder()
}

func TestBinaryTreePostOrder(t *testing.T) {
	b0 := &BinaryTree{0, nil, nil}
	b1 := &BinaryTree{1, nil, nil}
	b2 := &BinaryTree{2, nil, nil}
	b3 := &BinaryTree{3, nil, nil}
	b4 := &BinaryTree{4, nil, nil}

	b5 := &BinaryTree{5, b0, b1}
	b6 := &BinaryTree{6, b2, b3}
	b7 := &BinaryTree{7, b4, nil}

	b8 := &BinaryTree{8, b5, b6}
	b9 := &BinaryTree{9, b7, nil}

	b10 := &BinaryTree{10, b8, b9}

	b10.postOrder()
}

func TestBinaryTreeBF(t *testing.T) {
	b0 := &BinaryTree{0, nil, nil}
	b1 := &BinaryTree{1, nil, nil}
	b2 := &BinaryTree{2, nil, nil}
	b3 := &BinaryTree{3, nil, nil}
	b4 := &BinaryTree{4, nil, nil}

	b5 := &BinaryTree{5, b0, b1}
	b6 := &BinaryTree{6, b2, b3}
	b7 := &BinaryTree{7, b4, nil}

	b8 := &BinaryTree{8, b5, b6}
	b9 := &BinaryTree{9, b7, nil}

	b10 := &BinaryTree{10, b8, b9}

	b10.breadthFirst()
}
