# 题目描述：
# 小明今年升学到小学一年级，来到新班级后发现其他小朋友们身高参差不齐，然后就想基于各小朋友
# 和自己的身高差对他们进行排序，请帮他实现排序。
# 输入描述：
# 第一行为正整数H和N，0<H<200，为小明的身高，O<N<50，为新班级其他小朋友个数。
# 第二行为N个正整数H1-HN，分别是其他小朋友的身高，取值范围0<Hi<200 (1<=i<=N)，且N个正整数
# 各不相同。
# 输出描述：
# 输出排序结果，各正整数以空格分割。和小明 身高差Q绝对值最小的小朋友排在前面，和小明身高差绝
# 对值最大的小朋友排在最后，如果两个小朋友和小明身高差一样，则个子较小的小朋友排在前面。


def main():
    H, N = map(int, input().split())
    heights = list(map(int, input().split()))
    result = sort_height(H, N, heights)
    print(" ".join(map(str, result)))


def sort_height(H: int, N: int, heights: list):
    heights.sort(key=lambda x: (abs(x - H), x))
    return heights


if __name__ == '__main__':
    main()
