/**
有的时候我们希望排序不是仅仅按照自然顺序排序。
例如，我们希望按照字符串的长度来对一个字符串数组排序而不是按照字母顺序来排序。

// 我们实现了sort接口的Len，Less和Swap方法
// 这样我们就可以使用sort包的通用方法Sort
// Len和Swap方法的实现在不同的类型之间大致
// 都是相同的，只有Less方法包含了自定义的排序
// 逻辑，这里我们希望以字符串长度升序排序
 */

package main

import (
	"sort"
	"fmt"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	//
	fruits := []string{"peach", "banana", "kiwi"}

	sort.Sort(ByLength(fruits))

	fmt.Println(fruits)
}
