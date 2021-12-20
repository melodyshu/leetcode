package main

import "fmt"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。.

s 和 t 仅包含小写字母

示例1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

解题思路:
使用数组
*/
func isAnagram(s string, t string) bool {
	//如果长度不相等
	if len(s) != len(t) {
		return false
	}
	var c [26]int
	for _, ch := range s {
		c[ch-'a']++
	}
	for _, ch := range t {
		c[ch-'a']--
		//如果小于0说明有不同的字母
		if c[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagramb"
	t := "nagarama"
	fmt.Println(isAnagram(s, t))
}
