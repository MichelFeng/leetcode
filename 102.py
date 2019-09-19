# coding=utf8
# 102. 二叉树的层次遍历
# 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
# 例如:
# 给定二叉树: [3, 9, 20, null, null, 15, 7],
#     3
#    / \
#   9  20
#     /  \
#    15   7
# 返回其层次遍历结果：
# [
#   [3],
#   [9,20],
#   [15,7]
# ]


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def levelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []
        cur = [root]  # 当前层的节点
        nxt = []  # 下一层的节点
        res = []  # 最后的结果
        tmp = []  # 当前层的结果
        while True:
            if not cur:
                # 当前层遍历完
                cur = nxt
                if tmp:
                    # 当前层的结果插入最终结果的头部
                    res.append(tmp)
                tmp = []
                if not nxt:
                    # 没有下一层了退出循环
                    break
                nxt = []
            n = cur.pop(0)
            if n.left:
                nxt.append(n.left)
            if n.right:
                nxt.append(n.right)
            tmp.append(n.val)
        return res


s = Solution()
root = TreeNode(3)
l1 = TreeNode(9)
root.left = l1
r1 = TreeNode(20)
root.right = r1
r21 = TreeNode(15)
r1.left = r21
r22 = TreeNode(7)
r1.right = r22
print s.levelOrder(root)
