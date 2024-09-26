//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 
//
// 示例 2: 
//
// 
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 
//
// 示例 3: 
//
// 
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 5 * 10⁴ 
// s 由英文字母、数字、符号和空格组成 
// 
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10343 👎 0


import java.util.Set;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int lengthOfLongestSubstring(String s) {
        int left = 0;
        int right = 0;
        int len = 0;
        int n = s.length();
        while(right < n){

            while (isRepeat(s, left, right)){
                left++;
            }
            len = len > right -left + 1 ? len : right - left + 1;
            right++;
        }
        return len;
    }

    private boolean isRepeat(String s, int left, int right) {
        Set<Character> set = new HashSet<>();
        for (int i = left; i <= right; i++) {
            if (set.contains(s.charAt(i))) {
                return true;
            }
            set.add(s.charAt(i));
        }
        return false;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
