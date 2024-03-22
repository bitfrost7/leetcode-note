# 题目描述：
# 围棋棋盘由纵横各19条线垂直相交组成，棋盘上一共19x19-361个交点，对弈双方一方执白棋，一方执
# 黑棋，落子时只能将棋子置于交点上。
# “气”是围棋中很重要的一个概念，某个棋子有几口气，是指其上下左右方向四个相邻的交叉点中，有几个
# 交叉点没有棋子，由此可知：
# 1、在棋盘的边缘上的棋子最多有3口气（黑1），在棋盘角点的棋子最多有2口气（黑2），其它情况最
# 多有4口气（白1）
# 2、所有同色棋子的气之和叫作该色棋子的气，需要注意的是，同色棋子重合的气点，对于该颜色棋子来
# 说，只能计算一次气，比如下图中，黑棋一共4口气，而不是5口气，因为黑1和黑2中间红色三角标出的
# 气是两个黑棋共有的，对于黑棋整体来说只能算一个气。
# 3、本题目只计算气，对于眼也按气计算，如果您不清楚：眼"的概念，可忽略，按照前面描述的规则计算
# 即可。
# 现在，请根据输入的黑棋和白棋的坐标位置，计算黑棋和白起一共各有多少气？
# 假设输入：
# 0 5 8 9 9 10
# 5 0 9 9 9 8
# 输出：
# 8 7

def neighbor(x, y, board_size):
    neighbors = []
    if x + 1 < board_size:
        neighbors.append((x + 1, y))
    if x - 1 >= 0:
        neighbors.append((x - 1, y))
    if y + 1 < board_size:
        neighbors.append((x, y + 1))
    if y - 1 >= 0:
        neighbors.append((x, y - 1))

    return neighbors


def calc_liberty(black_co_list, white_co_list):
    board_size = 19
    black_cnt, white_cnt = 0, 0
    for i in range(0, board_size - 1):
        for j in range(0, board_size - 1):
            co = (i, j)
            if co not in black_co_list and co not in white_co_list:
                has_black, has_white = False, False
                for nei_co in neighbor(i, j, 19):
                    if not has_black and nei_co in black_co_list:
                        black_cnt += 1
                        has_black = True
                    if not has_white and nei_co in white_co_list:
                        white_cnt += 1
                        has_white = True
    return black_cnt, white_cnt


def main():
    line1 = list(map(int, input().split()))
    black_co_list = [tuple(line1[i:i + 2]) for i in range(0, len(line1), 2)]
    line2 = list(map(int, input().split()))
    white_co_list = [tuple(line2[i:i + 2]) for i in range(0, len(line2), 2)]
    black_cnt, white_cnt = calc_liberty(black_co_list, white_co_list)
    print(black_cnt, white_cnt)


if __name__ == '__main__':
    main()
