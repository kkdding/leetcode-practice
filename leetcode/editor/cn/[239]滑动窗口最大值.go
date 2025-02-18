//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2961 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

type moQueue struct {
    queue []int
}

func (m *moQueue) front() int {
    return m.queue[0]
}

func (m *moQueue) back() int {
    return m.queue[len(m.queue)-1]
}

func (m *moQueue) empty() bool {
    return len(m.queue) == 0
}

func (m *moQueue) push(elem int) {
    for !m.empty() && elem > m.back() {
        m.queue = m.queue[:len(m.queue)-1]
    }
    m.queue = append(m.queue, elem)
}

func (m *moQueue) pop(elem int) {
	if m.front() == elem{
        m.queue = m.queue[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
    mo := new(moQueue)
    for i := 0; i < k; i++ {
        mo.push(nums[i])
    }
    res = append(res, mo.front())
    for i := 0; i < len(nums)-k; i++ {
        mo.pop(nums[i])
        mo.push(nums[i+k])
        res = append(res, mo.front())
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
