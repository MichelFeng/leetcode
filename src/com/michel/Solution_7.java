package com.michel;

import java.math.BigInteger;

public class Solution_7 {
    public int reverse(int x) {
        String num = Integer.toString(x);
        char first = num.charAt(0);
        int start = 0;
        boolean sign = false;
        if (first == '-') {
            sign = true;
            start = 1;
        }
        StringBuilder sb = new StringBuilder(num.substring(start));
        String newNum = sb.reverse().toString();
        long res = Long.parseLong(newNum) * (sign ? -1L : 1L);
        if (res <= Integer.MAX_VALUE && res >= Integer.MIN_VALUE) {
            return (int) (res);
        }

        return 0;
    }

    public static void main(String[] args) {
        Solution_7 solution = new Solution_7();
        System.out.println(solution.reverse(120));
        System.out.println(solution.reverse(-123));
        System.out.println(solution.reverse(Integer.MAX_VALUE));
        System.out.println(solution.reverse(Integer.MIN_VALUE));
    }
}
