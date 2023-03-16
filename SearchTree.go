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
	if tree.value == nil {
		return &SearchTree{value, nil, nil}
	}

	if value.eval >= tree.value.eval {
		return append(tree.left, value)
	}
	return append(tree.right, value)
}

func (tree SearchTree) Preorder(f func(*MoveInfo)) {
	if tree.value != nil {
		f(tree.value)
		tree.left.Preorder(f)
		tree.right.Preorder(f)
	}
}

func (tree SearchTree) Inorder(f func(*MoveInfo)) {
	if tree.value != nil {
		tree.left.Inorder(f)
		f(tree.value)
		tree.right.Inorder(f)
	}
}

func (tree SearchTree) Postorder(f func(*MoveInfo)) {
	if tree.value != nil {
		tree.left.Postorder(f)
		tree.right.Postorder(f)
		f(tree.value)
	}
}
