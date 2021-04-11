package basic

import "fmt"

type BinaryTreeNode struct {
	data int
	left *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(root *BinaryTreeNode) *BinaryTree {
	return &BinaryTree{
		Root: root,
	};
}

func (bt *BinaryTree) PreOrder(btNode *BinaryTreeNode) {
	if btNode == nil {
		return
	}

	fmt.Print(btNode.data)
	bt.PreOrder(btNode.left)
	bt.PreOrder(btNode.right)
}

func (bt *BinaryTree) InOrder(btNode *BinaryTreeNode) {
	if btNode == nil {
		return
	}

	bt.InOrder(btNode.left)
	fmt.Print( btNode.data)
	bt.InOrder(btNode.right)
}

func (bt *BinaryTree) PostOrder(btNode *BinaryTreeNode) {
	if btNode ==nil {
		return
	}

	bt.PostOrder(btNode.left)
	bt.PostOrder(btNode.right)
	fmt.Print(btNode.data)
}