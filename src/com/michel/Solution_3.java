package com.michel;

public class Solution_3 {
    public int lengthOfLongestSubstring(String s) {
        int max = 0;
        int count = 0;
        int idx = 0;
        int len = s.length();
        String tmp = "";
        while (idx != len) {
            String sub = s.substring(idx, idx + 1);
            int sIdx = tmp.indexOf(sub);
            if (sIdx == -1) {
                tmp += sub;
                count++;
            } else {
                tmp = tmp.substring(sIdx + 1) + sub;
                if (count > max) {
                    max = count;
                }
                count = tmp.length();
            }
            idx++;
        }
        if (count > max) {
            max = count;
        }
        return max;
    }

    public static void main(String[] args) {
        Solution_3 solution = new Solution_3();
        System.out.println(solution.lengthOfLongestSubstring("abcabcbb"));
        System.out.println(solution.lengthOfLongestSubstring("bbbbb"));
        System.out.println(solution.lengthOfLongestSubstring("pwwkew"));
        System.out.println(solution.lengthOfLongestSubstring("dvdf"));
        System.out.println(solution.lengthOfLongestSubstring("anviaj"));
    }
}