#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def divide(self, dividend, divisor):
        """
        超时
        :type dividend: int
        :type divisor: int
        :rtype: int
        """
        if not dividend:
            return 0
        if not divisor:
            return 0
        if dividend < 0 and divisor < 0:
            left = dividend - divisor
            idx = 1
            while left < 0:
                dividend = left
                left = dividend - divisor
                idx += 1
            return idx - 1
        elif dividend > 0 and divisor > 0:
            left = dividend - divisor
            idx = 1
            while left > 0:
                dividend = left
                left = dividend - divisor
                idx += 1
            return idx - 1
        elif dividend < 0 and divisor > 0:
            left = dividend + divisor
            idx = 1
            while left < 0:
                dividend = left
                left = dividend + divisor
                idx += 1
            return -idx + 1
        else:
            left = dividend + divisor
            idx = 1
            while left > 0:
                dividend = left
                left = dividend + divisor
                idx += 1
            return -idx + 1

    def divide2(self, dividend, divisor):
        if not dividend:
            return 0
        if not divisor:
            return 0
        sign = 1 if (dividend ^ divisor) < 0 else -1

        dividend = abs(dividend)
        divisor = abs(divisor)
        idx = 1
        while dividend >= divisor << 1:
            divisor <<= 1
            idx <<= 1

        res = 0
        while idx > 0 and dividend > 0:
            if dividend >= divisor:
                dividend -= divisor
                res += 1
            idx >>= 1
            divisor >>= 1

        return res if sign else -res


if __name__ == "__main__":
    s = Solution()
    dividend = 10
    divisor = 3
    print(s.divide(dividend, divisor))

    dividend = 7
    divisor = -3
    print(s.divide(dividend, divisor))

    dividend = -10
    divisor = -3
    print(s.divide(dividend, divisor))

    dividend = -7
    divisor = 3
    print(s.divide(dividend, divisor))

    dividend = -2147483648
    divisor = -1
    print(s.divide(dividend, divisor))
    print(-2147483648 / -1)
