package main

import (
	"fmt"
	"strings"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 示例 1：
//
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
//
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	// 初始化公告字符串。默认将第一个字符串当成公共字符串
	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		// 使用for类似于while循环，判断公共前缀是否和字符串一致
		for strings.Index(strs[i], prefix) != 0 {
			// 如果不一致，就从公共字符串减去一位
			prefix = prefix[:len(prefix)-1]

			// 如果公共字符串为空，说明没有公共前缀
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

func main() {

	var strs = []string{"flower", "flow", "fight"}

	commonPrefix := longestCommonPrefix(strs)
	fmt.Println(commonPrefix)
}
