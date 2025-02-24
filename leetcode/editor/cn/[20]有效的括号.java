//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 每个右括号都有一个对应的相同类型的左括号。 
// 
//
// 
//
// 示例 1： 
//
// 
// 输入：s = "()" 
// 
//
// 输出：true 
//
// 示例 2： 
//
// 
// 输入：s = "()[]{}" 
// 
//
// 输出：true 
//
// 示例 3： 
//
// 
// 输入：s = "(]" 
// 
//
// 输出：false 
//
// 示例 4： 
//
// 
// 输入：s = "([])" 
// 
//
// 输出：true 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁴ 
// s 仅由括号 '()[]{}' 组成 
// 
//
// Related Topics 栈 字符串 👍 4565 👎 0


import java.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean isValid(String s) {
        int len = s.length();
        if (len % 2 == 1) return false;

        Deque<Character> stack = new LinkedList<Character>();
        Map<Character, Character> map = new HashMap<Character, Character>();
        map.put(')', '(');
        map.put(']', '[');
        map.put('}', '{');

        for (int i = 0; i < len; i++) {

            char ch = s.charAt(i);
            if (map.containsValue(ch)) {
                stack.push(ch);
            }else if(map.containsKey(ch)){
                if(stack.isEmpty() || map.get(ch) != stack.peek()) return false;
                stack.pop();
            }else {
                return false;
            }
        }

        return stack.isEmpty();
    }
}
//leetcode submit region end(Prohibit modification and deletion)
