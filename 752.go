package main

import "fmt"

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}

func openLock(deadends []string, target string) int {
	root := "0000"
	if target == root {
		return 0
	}

	queue := make([]string, 0)
	visited := make(map[string]bool)

	invalid := make(map[string]bool, len(deadends))
	for _, item := range deadends {
		invalid[item] = true
	}

	if _, ok := invalid[root]; ok {
		return -1
	}

	queue = append(queue, root)
	visited[root] = true

	depth := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			cur := []byte(queue[0])
			tmp := queue[1:]

			if string(cur) == target {
				return depth
			}

			// 每一位加一和减一
			for j := 0; j < 4; j++ {
				newCur := make([]byte, 4)
				copy(newCur, cur)
				newCur[j] = (cur[j]-'0'+1+10)%10 + '0'

				if _, ok := invalid[string(newCur)]; !ok {
					if _, ok := visited[string(newCur)]; !ok {
						visited[string(newCur)] = true
						tmp = append(tmp, string(newCur))
					}
				}

				newCur = make([]byte, 4)
				copy(newCur, cur)
				newCur[j] = (cur[j]-'0'-1+10)%10 + '0'

				if _, ok := invalid[string(newCur)]; !ok {
					if _, ok := visited[string(newCur)]; !ok {
						visited[string(newCur)] = true
						tmp = append(tmp, string(newCur))
					}
				}

				queue = tmp
			}
		}

		depth++
	}

	return -1
}
