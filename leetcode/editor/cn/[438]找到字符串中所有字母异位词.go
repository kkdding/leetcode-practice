//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
// 
//
// 示例 2: 
//
// 
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
// 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length, p.length <= 3 * 10⁴ 
// s 和 p 仅包含小写字母 
// 
//
// Related Topics 哈希表 字符串 滑动窗口 👍 1550 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func findAnagrams(s string, p string) []int {
    var ans []int
    sLen, pLen := len(s), len(p)
    if sLen < pLen {
        return ans
    }

    var sCount, pCount [26]int
    for i, ch := range p {
        sCount[s[i]-'a']++
        pCount[ch-'a']++
    }
    if sCount == pCount {
        ans = append(ans, 0)
    }

    for i, ch := range s[:sLen-pLen] {
        sCount[ch-'a']--
        sCount[s[i+pLen]-'a']++
        if sCount == pCount {
            ans = append(ans, i+1)
        }
    }
    return ans
}

//leetcode submit region end(Prohibit modification and deletion)
