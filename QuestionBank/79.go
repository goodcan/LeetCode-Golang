/*
 * @Time     : 2020/9/13 01:04
 * @Author   : cancan
 * @File     : 79.go
 * @Function : 单词搜索
 */

/*
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 * 示例:
 * board =
 * [
 *   ['A','B','C','E'],
 *   ['S','F','C','S'],
 *   ['A','D','E','E']
 * ]
 *
 * 给定 word = "ABCCED", 返回 true
 * 给定 word = "SEE", 返回 true
 * 给定 word = "ABCB", 返回 false
 *
 * 提示：
 * board 和 word 中只包含大写和小写英文字母。
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 */

package QuestionBank

import "fmt"

func exist(board [][]byte, word string) bool {
	maxH := len(board)
	maxW := len(board[0])
	maxLen := len(word)

	for i := 0; i < maxH; i++ {
		for j := 0; j < maxW; j++ {
			if board[i][j] == word[0] {
				if existDfs(board, word, 0, i, j, maxH, maxW, maxLen, map[string]bool{fmt.Sprintf("%v-%v", i, j): true}) {
					return true
				}
			}
		}
	}

	return false
}

func existDfs(board [][]byte, word string, idx, i, j, maxH, maxW, maxLen int, check map[string]bool) bool {
	if i < 0 || j < 0 || i >= maxH || j >= maxW || board[i][j] != word[idx] {
		return false
	}

	idx++
	if idx >= maxLen {
		return true
	}

	var key string
	for _, item := range [][2]int{{i, j + 1}, {i, j - 1}, {i + 1, j}, {i - 1, j}} {
		key = fmt.Sprintf("%v-%v", item[0], item[1])
		if !check[key] {
			check[key] = true
			if existDfs(board, word, idx, item[0], item[1], maxH, maxW, maxLen, check) {
				return true
			}
			delete(check, key)
		}
	}

	return false
}
