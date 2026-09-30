package r

type Ltree struct {
	Id       int64   `json:"id"`                //主键ID
	Name     string  `json:"name"`              //名称
	Pid      int64   `json:"pid"`               //父ID
	Type     int64   `json:"type"`              //类别
	Children []Ltree `json:"children" gorm:"-"` //子元素
}

// TreeBuild 正确的树构建方法
func BuildLtree(nodes []Ltree) []Ltree {
	// 特殊情况处理
	if len(nodes) == 0 {
		return []Ltree{}
	}

	// 创建两个映射
	idToNode := make(map[int64]*Ltree)        // ID到节点指针的映射
	pidToChildren := make(map[int64][]*Ltree) // 父ID到子节点指针列表的映射

	// 第一遍遍历：建立映射关系
	for i := range nodes {
		// 注意：这里要取地址，但后续会创建新的节点
		node := &nodes[i]
		idToNode[node.Id] = node

		// 收集每个父ID下的子节点
		if node.Pid != 0 && node.Pid != node.Id {
			pidToChildren[node.Pid] = append(pidToChildren[node.Pid], node)
		}
	}

	// 第二遍遍历：构建树
	var result []Ltree

	for i := range nodes {
		node := &nodes[i]

		// 判断是否为根节点
		isRoot := false
		if node.Pid == 0 || node.Pid == node.Id {
			isRoot = true
		} else if _, parentExists := idToNode[node.Pid]; !parentExists {
			// 父节点不存在，也作为根节点
			isRoot = true
		}

		if isRoot {
			// 深度拷贝构建树
			treeRoot := buildSubLtree(node, pidToChildren)
			result = append(result, treeRoot)
		}
	}

	return result
}

// buildSubTree 递归构建子树
func buildSubLtree(node *Ltree, pidToChildren map[int64][]*Ltree) Ltree {
	// 创建当前节点的副本
	newNode := Ltree{
		Id:   node.Id,
		Name: node.Name,
		Pid:  node.Pid,
		Type: node.Type,
	}

	// 如果有子节点，递归构建
	if children, exists := pidToChildren[node.Id]; exists {
		newNode.Children = make([]Ltree, 0, len(children))
		for _, child := range children {
			subTree := buildSubLtree(child, pidToChildren)
			newNode.Children = append(newNode.Children, subTree)
		}
	}

	return newNode
}
