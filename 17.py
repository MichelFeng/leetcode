#!/usr/bin/env python
# -*- coding: utf-8 -*-
class Solution(object):
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        num2char = {
            '2': 'abc',
            '3': 'def',
            '4': 'ghi',
            '5': 'jkl',
            '6': 'mno',
            '7': 'pqrs',
            '8': 'tuv',
            '9': 'wxyz'
        }
        
        res = [""]
        for d in digits:
            t = []
            for c in num2char[d]:
                for r in res:
                    t.append(r + c)
            res = t
                    
        return res


if __name__ == "__main__":
    digits = "23"
    print Solution().letterCombinations(digits)
