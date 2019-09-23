# coding=utf8
# 111. 二叉树的最小深度
# 给定一个二叉树，找出其最小深度。
# 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
# 说明: 叶子节点是指没有子节点的节点。
# 示例:
# 给定二叉树 [3, 9, 20, null, null, 15, 7],
#     3
#    / \
#   9  20
#     /  \
#    15   7
# 返回它的最小深度  2.


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def minDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        if not root.left and not root.right:
            return 1
        min_dep = float('inf')
        if root.left:
            min_dep = min([self.minDepth(root.left), min_dep])
        if root.right:
            min_dep = min([self.minDepth(root.right), min_dep])
        return min_dep + 1


s = Solution()
root = TreeNode(3)
l = TreeNode(9)
# r = TreeNode(20)
# rl = TreeNode(15)
# rr = TreeNode(7)
root.left = l
# root.right = r
# r.left = rl
# r.right = rr
print s.minDepth(root)
