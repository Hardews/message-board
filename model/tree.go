package model

type TreeNode struct {
	ID      int
	Txt     string
	Time    string
	Name    string
	LikeNum int
	Branch  *TreeNode
}
