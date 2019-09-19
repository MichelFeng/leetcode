# coding=utf8
# 101. 对称二叉树

# 给定一个二叉树，检查它是否是镜像对称的。
# 例如，二叉树 [1, 2, 2, 3, 4, 4, 3] 是对称的。
#     1
#    / \
#   2   2
#  / \ / \
# 3  4 4  3
# 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
#     1
#    / \
#   2   2
#    \   \
#    3    3


class TreeNode(object):
    '''
    Definition for a binary tree node.
    '''

    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


# class Solution(object):
#     def isSymmetric(self, root):
#         """
#         :type root: TreeNode
#         :rtype: bool
#         """
#         if not root:
#             return True
#         if not root.left and not root.right:
#             return True
#         if not root.left or not root.right:
#             return False
#         ls = self.dfs(root.left, True, [])
#         rs = self.dfs(root.right, False, [])
#         print ls, rs
#         if len(ls) != len(rs):
#             return False
#         for i in range(len(ls)):
#             if ls[i] != rs[i]:
#                 return False
#         return True

#     def dfs(self, node, switch, res):
#         if not node:
#             res.append(None)
#         else:
#             res.append(node.val)
#             if switch:
#                 self.dfs(node.left, switch, res)
#                 self.dfs(node.right, switch, res)
#             else:
#                 self.dfs(node.right, switch, res)
#                 self.dfs(node.left, switch, res)
#         return res


class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.isMirror(root, root)

    def isMirror(self, t1, t2):
        if not t1 and not t2:
            return True
        if not t1 or not t2:
            return False
        if t1.val != t2.val:
            return False
        return self.isMirror(t1.left, t2.right) and self.isMirror(t1.right, t2.left)


s = Solution()
root = TreeNode(1)
l1 = TreeNode(2)
r1 = TreeNode(2)
l2 = TreeNode(3)
r2 = TreeNode(3)
root.left = l1
root.right = r1
l1.left = l2
r1.left = r2
print s.isSymmetric(root)
