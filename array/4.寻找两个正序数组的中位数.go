package array

import "fmt"

// description: (https://leetcode.cn/problems//description/)
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 算法的时间复杂度应该为 O(log (m+n)) 。

//

// 示例 1：

// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：

// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

// ----- Code Here -----
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 1 {
		return float64(getKthElement(nums1, nums2, (m+n)/2+1))
	} else {
		return float64(getKthElement(nums1, nums2, (m+n)/2)+getKthElement(nums1, nums2, (m+n)/2+1)) / 2.0
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		// 边界情况1: nums1已经全部被排除或者nums2已经全部被排除
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}

		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}

		// 边界情况2: k == 1 直接返回两个数组的最小值
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}

		// 正常情况
		// 1. 比较两个数组的第 k/2 个元素
		half := k / 2

		// 防止越界
		newIdx1 := min(idx1+half, len(nums1)) - 1
		newIdx2 := min(idx2+half, len(nums2)) - 1

		privot1, privot2 := nums1[newIdx1], nums2[newIdx2]
		// 2. 排除较小的部分
		if privot1 <= privot2 {
			k -= (newIdx1 - idx1 + 1)
			idx1 = newIdx1 + 1
		} else {
			k -= (newIdx2 - idx2 + 1)
			idx2 = newIdx2 + 1
		}
	}
}

// ----- Begin Test -----

func TestFindMedianSortedArrays() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println("nums1:", nums1, "nums2:", nums2, "median:", findMedianSortedArrays(nums1, nums2)) // 2.0
}

// ----- Solution Here -----
/*
本题可以使用归并查找但是复杂度为 O(m+n)，不符合要求 所以需要使用二分查找降低时间复杂度 O(log(min(m,n))) 。
主要思路是将中位数问题转变为寻找第 k 小的数的问题:

- 若总长度 m+n 为奇数，中位数是第 (m+n)/2 + 1 小的元素（比如总长度 13，找第 7 小）；
- 若总长度为偶数，中位数是第 (m+n)/2 和 (m+n)/2 + 1 小的元素的平均值。

关键在于如何去缩小查找范围

假设两个数组分别为 A 和 B，初始时我们需要找第 k 小的元素。
1. 比较 A[k/2-1] 和 B[k/2-1] (记为 pivot1 和 pivot2)。
2. 如果 pivot1 <= pivot2，说明 A 的前 k/2 个元素不可能是第 k 小的元素，因为这是最多有k-2个数字比它们小，可以将它们全部排除。更新 k 和 A 的起始索引。
3. 如果 pivot1 > pivot2，同理排除 B 的前 k/2 个元素。

特殊情况处理：
1. 数组越界
因为index代表该数组已经被排除的元素个数，所以在比较时需要考虑数组长度不足 k/2 的情况，防止越界。如果越界说明该数组的所有元素都可以被排除。直接返回另一个数组的第 k 小元素即可。
2. k == 1
直接返回两个数组当前起始位置的最小值, 因为第1小就是两个数组的最小值。
*/
