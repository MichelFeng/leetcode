# coding=utf8
# 155. 最小栈
# 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
# push(x) - - 将元素 x 推入栈中。
# pop() - - 删除栈顶的元素。
# top() - - 获取栈顶元素。
# getMin() - - 检索栈中的最小元素。
# 示例:
# MinStack minStack = new MinStack()
# minStack.push(-2)
# minStack.push(0)
# minStack.push(-3)
# minStack.getMin() - -> 返回 - 3.
# minStack.pop()
# minStack.top() - -> 返回 0.
# minStack.getMin() - -> 返回 - 2.


class MinStack(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.s1 = []
        self.s2 = []

    def push(self, x):
        """
        :type x: int
        :rtype: None
        """
        self.s1.append(x)
        if not len(self.s2):
            self.s2.append(x)
        else:
            if x < self.s2[-1]:
                self.s2.append(x)
            else:
                self.s2.append(self.s2[-1])

    def pop(self):
        """
        :rtype: None
        """
        if len(self.s1) == 0:
            return None
        self.s1 = self.s1[:-1]
        self.s2 = self.s2[:-1]

    def top(self):
        """
        :rtype: int
        """

        if len(self.s1) == 0:
            return None
        d = self.s1[-1]
        return d

    def getMin(self):
        """
        :rtype: int
        """
        if len(self.s2) == 0:
            return None
        return self.s2[-1]


minStack = MinStack()
minStack.push(-2)
minStack.push(0)
minStack.push(-3)
print minStack.getMin()  # - -> 返回 - 3.
print minStack.pop()
print minStack.top()  # - -> 返回 0.
print minStack.getMin()  # - -> 返回 - 2.
