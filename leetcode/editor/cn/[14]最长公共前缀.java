//编写一个函数来查找字符串数组中的最长公共前缀。 
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["flower","flow","flight"]
//输出："fl"
// 
//
// 示例 2： 
//
// 
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 
//
// 
//
// 提示： 
//
// 
// 1 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] 仅由小写英文字母组成 
// 
//
// Related Topics 字典树 字符串 👍 3180 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public String longestCommonPrefix(String[] strs) {
        String str1 = strs[0];
        for (int i = 1; i < strs.length; i++) {
            String str2 = strs[i];
            str1 = getCommonPrefix(str1, str2);
            if (str1.length() == 0) {
                return "";
            }
        }
        return str1;
    }

    private String getCommonPrefix(String str1, String str2) {
        int p = 0;
        while (p < str1.length() && p < str2.length() && str1.charAt(p) == str2.charAt(p)) {
            p++;
        }
        return str1.substring(0, p);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
