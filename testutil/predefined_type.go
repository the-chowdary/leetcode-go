package testutil

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeNode(rawArray string) (root *TreeNode, err error) {
	if len(rawArray) < 2 || rawArray[0] != '[' || rawArray[len(rawArray)-1] != ']' {
		return nil, fmt.Errorf("invalid test data, %w", rawArray)
	}

	// Remove [] from start and end
	rawArray = rawArray[1 : len(rawArray) - 1]
	if rawArray == "" {
		return
	}

	splits := strings.Split(rawArray, ",")

	nodes := make([]*TreeNode, len(splits))

	for i, s := range splits {
		s = strings.TrimSpace(s)
		if s != "null" {
			nodes[i] = &TreeNode{}
			var err error
			nodes[i].Val, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
		}
	}

	children := nodes

	root, children = children[0], children[1:]
	for _, node := range nodes {
		if node != nil {
			if len(children) > 0 {
				node.Left, children = children[0], children[1:]
			}
			if len(children) > 0 {
				node.Right, children = children[0], children[1:]
			}
		}
	}
	return
}



