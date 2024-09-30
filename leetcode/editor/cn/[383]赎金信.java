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
// Related Topics 哈希表 字符串 计数 👍 917 👎 0


import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        /**
         * 双指针
         */
//        char[] ransomNoteCharArray = ransomNote.toCharArray();
//        char[] magazineCharArray = magazine.toCharArray();
//        Arrays.sort(ransomNoteCharArray);
//        Arrays.sort(magazineCharArray);
//        int ransomNoteIndex = 0;
//        int magazineIndex = 0;
//        while (ransomNoteIndex < ransomNoteCharArray.length && magazineIndex < magazineCharArray.length){
//            if (ransomNoteCharArray[ransomNoteIndex] == magazineCharArray[magazineIndex]){
//                ransomNoteIndex++;
//                magazineIndex++;
//            }else {
//                magazineIndex++;
//            }
//        }
//        if(ransomNoteIndex == ransomNoteCharArray.length) return true;
//        return false;
        /**
         * 哈希表
         */
        Map<Character, Integer> map = new HashMap<>();
        for (Character c : magazine.toCharArray()) {
            if (map.containsKey(c)) {
                map.put(c, map.get(c) + 1);
            } else {
                map.put(c, 1);
            }
        }
        for (Character c : ransomNote.toCharArray()) {
            if (map.containsKey(c)) {
                map.put(c, map.get(c) - 1);
                if (map.get(c) < 0) return false;
            } else {
                return false;
            }
        }
        return true;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
