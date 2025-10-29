package array

import "fmt"

// description: (https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/)
//给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
//
//考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
//
//nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
//
//判题标准:
//
//系统会用下面的代码来测试你的题解:
//
//int[] nums = [...]; // 输入数组
//int[] expectedNums = [...]; // 长度正确的期望答案
//
//int k = removeDuplicates(nums); // 调用
//
//assert k == expectedNums.length;
//for (int i = 0; i < k; i++) {
//assert nums[i] == expectedNums[i];
//}
//如果所有断言都通过，那么您的题解将被 通过。

// ----- Code Here -----

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
			j++
		}
	}
	return i + 1
}

// ----- Begin Test -----

func TestRemoveDuplicates() int {
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2}
	fmt.Println("原数组:", nums)
	d := removeDuplicates(nums)
	fmt.Println("数组:", nums)
	fmt.Println(d)
	return d
}
