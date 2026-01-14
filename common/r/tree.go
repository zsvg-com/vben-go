package r

type Stree struct {
	Id       string  `json:"id"`                //主键ID
	Name     string  `json:"name"`              //名称
	Pid      string  `json:"pid"`               //父ID
	Type     int64   `json:"type"`              //类别
	Children []Stree `json:"children" gorm:"-"` //子元素
}

// TreeBuild 正确的树构建方法
func TreeBuild(nodes []Stree) []Stree {
	// 特殊情况处理
	if len(nodes) == 0 {
		return []Stree{}
	}

	// 创建两个映射
	idToNode := make(map[string]*Stree)        // ID到节点指针的映射
	pidToChildren := make(map[string][]*Stree) // 父ID到子节点指针列表的映射

	// 第一遍遍历：建立映射关系
	for i := range nodes {
		// 注意：这里要取地址，但后续会创建新的节点
		node := &nodes[i]
		idToNode[node.Id] = node

		// 收集每个父ID下的子节点
		if node.Pid != "" && node.Pid != node.Id {
			pidToChildren[node.Pid] = append(pidToChildren[node.Pid], node)
		}
	}

	// 第二遍遍历：构建树
	var result []Stree

	for i := range nodes {
		node := &nodes[i]

		// 判断是否为根节点
		isRoot := false
		if node.Pid == "" || node.Pid == "0" || node.Pid == node.Id {
			isRoot = true
		} else if _, parentExists := idToNode[node.Pid]; !parentExists {
			// 父节点不存在，也作为根节点
			isRoot = true
		}

		if isRoot {
			// 深度拷贝构建树
			treeRoot := buildSubTree(node, pidToChildren)
			result = append(result, treeRoot)
		}
	}

	return result
}

// buildSubTree 递归构建子树
func buildSubTree(node *Stree, pidToChildren map[string][]*Stree) Stree {
	// 创建当前节点的副本
	newNode := Stree{
		Id:   node.Id,
		Name: node.Name,
		Pid:  node.Pid,
		Type: node.Type,
	}

	// 如果有子节点，递归构建
	if children, exists := pidToChildren[node.Id]; exists {
		newNode.Children = make([]Stree, 0, len(children))
		for _, child := range children {
			subTree := buildSubTree(child, pidToChildren)
			newNode.Children = append(newNode.Children, subTree)
		}
	}

	return newNode
}

// 优化版本：避免递归过深
func TreeBuildIterative(nodes []Stree) []Stree {
	if len(nodes) == 0 {
		return []Stree{}
	}

	// 创建映射
	nodeMap := make(map[string]*Stree)
	childrenMap := make(map[string][]string)

	// 第一遍：建立关系
	for i := range nodes {
		node := &nodes[i]
		nodeMap[node.Id] = node

		if node.Pid != "" && node.Pid != node.Id {
			childrenMap[node.Pid] = append(childrenMap[node.Pid], node.Id)
		}
	}

	// 第二遍：构建结果
	var result []Stree

	for i := range nodes {
		node := &nodes[i]

		// 检查是否为根节点
		if node.Pid == "" || node.Pid == "0" || node.Pid == node.Id {
			result = append(result, *node)
			continue
		}

		// 检查父节点是否存在
		if _, exists := nodeMap[node.Pid]; !exists {
			result = append(result, *node)
		}
	}

	// 第三遍：为每个节点附加子节点
	for i := range result {
		attachChildren(&result[i], childrenMap, nodeMap)
	}

	return result
}

// attachChildren 为节点附加子节点
func attachChildren(node *Stree, childrenMap map[string][]string, nodeMap map[string]*Stree) {
	if childIds, exists := childrenMap[node.Id]; exists {
		node.Children = make([]Stree, 0, len(childIds))

		for _, childId := range childIds {
			if childNode, exists := nodeMap[childId]; exists {
				newChild := Stree{
					Id:   childNode.Id,
					Name: childNode.Name,
					Pid:  childNode.Pid,
					Type: childNode.Type,
				}
				// 递归附加子节点
				attachChildren(&newChild, childrenMap, nodeMap)
				node.Children = append(node.Children, newChild)
			}
		}
	}
}

// 正确的简化版本
func BuildTreeCorrect(items []Stree) []Stree {
	// 复制数据，避免修改原始数据
	nodes := make([]Stree, len(items))
	copy(nodes, items)

	// 创建映射
	nodeMap := make(map[string]*Stree)
	for i := range nodes {
		nodeMap[nodes[i].Id] = &nodes[i]
		// 确保Children切片初始化
		nodes[i].Children = []Stree{}
	}

	// 构建树
	var roots []Stree
	for i := range nodes {
		node := &nodes[i]

		// 判断根节点条件
		if node.Pid == "" || node.Pid == "0" || node.Pid == node.Id {
			roots = append(roots, *node)
		} else if parent, exists := nodeMap[node.Pid]; exists {
			// 添加到父节点的Children
			parent.Children = append(parent.Children, *node)
		} else {
			// 父节点不存在，作为根节点
			roots = append(roots, *node)
		}
	}

	return roots
}
