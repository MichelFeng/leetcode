# coding=utf8
# 我手中有一堆扑克牌， 但是观众不知道它的顺序。
# 1、第一步， 我从牌顶拿出一张牌， 放到桌子上。
# 2、第二步， 我从牌顶再拿一张牌， 放在手上牌的底部。
# 3、第三步， 重复第一步、第二步的操作， 直到我手中所有的牌都放到了桌子上。
# 最后， 观众可以看到桌子上牌的顺序是：(牌底部）1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13(牌顶部）
# 请问， 我刚开始拿在手里的牌的顺序是什么？


class Solution(object):
    def hand_to_table(self, cards):
        res = []
        if not cards:
            return res
        l = len(cards)
        for i in range(l):
            # 放到桌面上
            res.append(cards.pop(0))
            # 从牌顶放到牌底
            if len(cards) > 0:
                cards.append(cards.pop(0))
        return res

    def table_to_hand(self, cards):
        res = []
        if not cards:
            return res
        for card in reversed(cards):
            # 从牌底放到牌顶
            if len(res) > 1:
                res.insert(0, res.pop())
            # 从桌面拿牌
            res.insert(0, card)
        return res


s = Solution()
res = s.hand_to_table([1, 12, 2, 8, 3, 11, 4, 9, 5, 13, 6, 10, 7])
print res  # [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]
res = s.table_to_hand([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13])
print res  # [1, 12, 2, 8, 3, 11, 4, 9, 5, 13, 6, 10, 7]
