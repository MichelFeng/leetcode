# coding=utf8
# 112. 路径总和
# 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

# 说明: 叶子节点是指没有子节点的节点。

# 示例: 给定如下二叉树，以及目标和 sum = 22，

#               5
#              / \
#             4   8
#            /   / \
#           11  13  4
#          /  \      \
#         7    2      1
# 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def hasPathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: bool
        """
        if not root:
            return False
        if not root.left and not root.right:
            return root.val == sum
        return self.hasPathSum(root.left, sum-root.val) or self.hasPathSum(root.right, sum-root.val)


s = Solution()
root = TreeNode(5)
l1 = TreeNode(4)
r1 = TreeNode(8)
root.left = l1
root.right = r1
l21 = TreeNode(11)
l1.left = l21
r21 = TreeNode(13)
r22 = TreeNode(4)
r1.left = r21
r1.right = r22
l31 = TreeNode(7)
l32 = TreeNode(2)
l21.left = l31
l21.right = l32
r32 = TreeNode(1)
r22.right = r32
print s.hasPathSum(root, 22)
print s.hasPathSum(root, 26)
print s.hasPathSum(root, 13)
