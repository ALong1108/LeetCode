/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖整个字符串 s 的，而不是部分字符串。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
*/
package leetcode

/*
解法: 回溯
如果没有 * ，我们只需要从左到右检查匹配串 s 是否能匹配模式串 p 的每一个字符。
当模式串中有 * 时，我们需要检查匹配串 s 中的不同后缀，以判断它们是否能匹配模式串剩余的部分。一个直观的解法就是用回溯的方法来体现这种关系。

如果模式串中有星号，它会出现在第二个位置，即 pattern[1] 。
这种情况下，我们可以直接忽略模式串中这一部分，或者删除匹配串的第一个字符，前提是它能够匹配模式串当前位置字符，即 pattern[0] 。
如果两种操作中有任何一种使得剩下的字符串能匹配，那么初始时，匹配串和模式串就可以被匹配。


结果: 执行用时 :16 ms 内存消耗 :2.1 MB
*/
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}