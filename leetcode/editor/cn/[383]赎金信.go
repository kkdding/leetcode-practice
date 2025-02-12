//给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//
// 如果可以，返回 true ；否则返回 false 。
//
// magazine 中的每个字符只能在 ransomNote 中使用一次。
//
//
//
// 示例 1：
//
//
//输入：ransomNote = "a", magazine = "b"
//输出：false
//
//
// 示例 2：
//
//
//输入：ransomNote = "aa", magazine = "ab"
//输出：false
//
//
// 示例 3：
//
//
//输入：ransomNote = "aa", magazine = "aab"
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= ransomNote.length, magazine.length <= 10⁵
// ransomNote 和 magazine 由小写英文字母组成
//
//
// Related Topics 哈希表 字符串 计数 👍 949 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func canConstruct(ransomNote string, magazine string) bool {
	set := make(map[rune]int)
	for _, c := range ransomNote {
		set[c]++
	}

	for _, c := range magazine {
		if _, ok := set[c]; ok {
			set[c]--
			if set[c] == 0 {
				delete(set, c)
			}
		}
	}

	return len(set) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
