package com.michel;

import java.util.HashMap;

public class Solution_13 {
    private HashMap<Character, Integer> map = null;

    public Solution_13() {
        this.map = new HashMap<>();
        this.map.put('I', 1);
        this.map.put('V', 5);
        this.map.put('X', 10);
        this.map.put('L', 50);
        this.map.put('C', 100);
        this.map.put('D', 500);
        this.map.put('M', 1000);
    }

    public int romanToInt(String s) {
        int length = s.length();
        int sum = 0;
        for (int i = 0; i < length; ) {
            char first = s.charAt(i);
            if (i + 1 < length) {
                char second = s.charAt(i + 1);
                int firstNum = this.map.get(first);
                int secondNum = this.map.get(second);
                if (firstNum >= secondNum) {
                    sum += firstNum;
                    i++;
                } else {
                    sum += secondNum - firstNum;
                    i += 2;
                }
            } else {
                sum += this.map.get(first);
                i++;
            }
        }
        return sum;
    }

    public static void main(String[] args) {
        Solution_13 solution_13 = new Solution_13();
        solution_13.romanToInt("MCMXCIV");
        solution_13.romanToInt("LVIII");
        solution_13.romanToInt("IV");
        solution_13.romanToInt("V");
    }
}
