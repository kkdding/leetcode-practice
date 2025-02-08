//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
// 注意：
//
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在
//o(m+n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 3111 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func minWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }
    if s == t {
        return s
    }

    // 记录 t 中每个字符的出现次数
    target := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        target[t[i]]++
    }

    // 记录当前窗口中每个字符的出现次数
    window := make(map[byte]int)
    // 记录当前窗口中满足 t 字符频率要求的字符种类数
    count := 0
    // 最小窗口的左右边界和最小长度
    left, right := 0, 0
    minLeft, minRight, minLen := 0, 0, len(s)+1

    for right < len(s) {
        // 移入窗口的字符
        c := s[right]
        right++
        // 更新窗口内字符的出现次数
        window[c]++
        // 如果当前字符的出现次数达到了 t 中该字符的出现次数，则满足要求的字符种类数加 1
        if target[c] > 0 && window[c] == target[c] {
            count++
        }

        // 判断左侧窗口是否需要收缩
        for count == len(target) {
            // 更新最小窗口
            if right-left < minLen {
                minLen = right - left
                minLeft = left
                minRight = right
            }
            // 移出窗口的字符
            d := s[left]
            left++
            // 更新窗口内字符的出现次数
            window[d]--
            // 如果当前字符的出现次数小于 t 中该字符的出现次数，则满足要求的字符种类数减 1
            if target[d] > 0 && window[d] < target[d] {
                count--
            }
        }
    }

    if minLen > len(s) {
        return ""
    }
    return s[minLeft:minRight]
}

//leetcode submit region end(Prohibit modification and deletion)
