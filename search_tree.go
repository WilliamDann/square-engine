package main

import "github.com/notnil/chess"

type SearchTree struct {
	value       *MoveInfo
	left, right *SearchTree
}

type MoveInfo struct {
	move *chess.Move
	eval int
}

func append(tree *SearchTree, value *MoveInfo) *SearchTree {
	if tree == nil || tree.value == nil {
		return &SearchTree{value, nil, nil}
	}

	if value.eval <= tree.value.eval {
		return &SearchTree{tree.value, tree.left, append(tree.right, value)}
	}
	return &SearchTree{tree.value, append(tree.left, value), tree.right}
}

func (tree *SearchTree) Preorder(f func(*MoveInfo)) {
	if tree != nil && tree.value != nil {
		f(tree.value)
		tree.left.Preorder(f)
		tree.right.Preorder(f)
	}
}

func (tree *SearchTree) Inorder(f func(*MoveInfo)) {
	if tree != nil && tree.value != nil {
		tree.left.Inorder(f)
		f(tree.value)
		tree.right.Inorder(f)
	}
}

func (tree *SearchTree) Postorder(f func(*MoveInfo)) {
	if tree != nil && tree.value != nil {
		tree.left.Postorder(f)
		tree.right.Postorder(f)
		f(tree.value)
	}
}

func (tree *SearchTree) getRightmost() *MoveInfo {
	if tree.right != nil {
		return tree.right.getRightmost()
	}
	return tree.value
}

func (tree *SearchTree) getLeftmost() *MoveInfo {
	if tree.left != nil {
		return tree.left.getLeftmost()
	}
	return tree.value
}
