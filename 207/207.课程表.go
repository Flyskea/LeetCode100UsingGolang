/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */
package main

import (
	"fmt"
)

/*
// DFS
// 0表示没访问
// 1表示当前dfs访问过
// 2表示其他dfs分支访问过

func canFinish(numCourses int, prerequisites [][]int) bool {
    var (
        edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        result []int
        valid = true
        dfs func(u int)
    )

    dfs = func(u int) {
        visited[u] = 1
        for _, v := range edges[u] {
            if visited[v] == 0 {
                dfs(v)
                if !valid {
                    return
                }
            } else if visited[v] == 1 {
                valid = false
                return
            }
        }
        visited[u] = 2
        result = append(result, u)
    }

    for _, info := range prerequisites {
        edges[info[1]] = append(edges[info[1]], info[0])
    }

    for i := 0; i < numCourses && valid; i++ {
        if visited[i] == 0 {
            dfs(i)
        }
    }
    return valid
}
*/

// 拓扑排序
// 则没有环
// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	graph := make(map[int][]int)
	queue := make([]int, 0)
	count := 0
	for _, v := range prerequisites {
		indegree[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		class := queue[0]
		queue = queue[1:]
		// 学习课程数
		count++
		need := graph[class]
		for _, v := range need {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	// 如果不相等则说明 课程未完成
	return count == numCourses
}

// @lc code=end
func main() {
	fmt.Println(canFinish(2, [][]int{{0, 1}, {1, 0}}))
}
