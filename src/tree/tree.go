package tree

/**
 * 树节点的数据结构
 */
type Node struct {
	leftChild *Node		//左孩子
	rightChild *Node	//右孩子
	data string			//存储的数据
}

/**
 * 存储树信息
 */
type Tree struct {
	rootNode Node 	//根结点
	data []string	//树
}

/**
 * 构建一棵树
 */
func MakeTree(data []string) Tree {
	tree := Tree{
		data:data,
	}
	return tree
}

