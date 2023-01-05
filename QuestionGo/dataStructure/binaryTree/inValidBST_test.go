package binaryTree

import "testing"

type BST *TreeNode

func Test_isValidBST(t *testing.T) {
	type Detail struct {
		want  bool
		mater *TreeNode
	}
	BST1 := &TreeNode{
		Val: 2,

		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	BST2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	BST3 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	var TestList = []Detail{
		{want: true, mater: BST1},
		{want: false, mater: BST2},
		{want: false, mater: BST3},

	}

	for _, detail := range TestList {
		if got := isValidBST(detail.mater); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v\n", detail.want, got, *detail.mater)
			return
		}
	}
	t.Log("pass all test case")
}
