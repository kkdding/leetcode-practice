//给定一个非空的字符串
// s ，检查是否可以通过由它的一个子串重复多次构成。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abab"
//输出: true
//解释: 可由子串 "ab" 重复两次构成。
// 
//
// 示例 2: 
//
// 
//输入: s = "aba"
//输出: false
// 
//
// 示例 3: 
//
// 
//输入: s = "abcabcabcabc"
//输出: true
//解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
// 
//
// 
//
// 提示： 
//
// 
// 
//
// 
// 1 <= s.length <= 10⁴ 
// s 由小写英文字母组成 
// 
//
// Related Topics 字符串 字符串匹配 👍 1255 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func repeatedSubstringPattern(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        sub := s[:i+1]
        if len(s)%len(sub) != 0 {
            continue
        }
        n := len(s) / len(sub)
        for n > 1 {
            sub += sub
            n--
        }
        if sub == s {
            return true
        }
    }
    return false
}

//leetcode submit region end(Prohibit modification and deletion)
