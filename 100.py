# coding=utf8
# 100. 相同的树
# 给定两个二叉树，编写一个函数来检验它们是否相同。
# 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
# 示例 1:
# 输入:       1         1
#           / / \
#          2   3     2   3
#         [1,2,3],   [1,2,3]
# 输出: true
# 示例 2:
# 输入:      1          1
#           /           \
#          2             2
#         [1,2],     [1,null,2]
# 输出: false
# 示例 3:
# 输入:       1         1
#           / \       / \
#          2   1     1   2
#         [1,2,1],   [1,1,2]
# 输出: false

# Definition for a binary tree node.


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def isSameTree(self, p, q):
        """
        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        if not p and not q:
            return True
        if not p or not q:
            return False
        if p.val != q.val:
            return False
        return self.isSameTree(p.left, q.left) and \
            self.isSameTree(p.right, q.right)
    # def isSameTree(self, p, q):
    #     """
    #     :type p: TreeNode
    #     :type q: TreeNode
    #     :rtype: bool
    #     """
    #     t1 = self.traverse(p, [])
    #     t2 = self.traverse(q, [])

    #     if len(t1) != len(t2):
    #         return False
    #     for i in range(len(t1)):
    #         if t1[i] != t2[i]:
    #             return False
    #     return True

    # def traverse(self, root, res):
    #     if not root:
    #         res.append(None)
    #     else:
    #         res.append(root.val)
    #         self.traverse(root.left, res)
    #         self.traverse(root.right, res)
    #     return res


s = Solution()
p = TreeNode(1)
p1 = TreeNode(2)
p.left = p1

q = TreeNode(1)
q1 = TreeNode(2)
q.right = q1
print s.isSameTree(p, q)
