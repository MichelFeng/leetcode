package com.michel;

public class Solution_12 {
    public String intToRoman(int num) {
        int[] arr = new int[4];
        int i = 0;
        while (num != 0) {
            arr[i++] = num % 10;
            num /= 10;
        }
        String chars = "IXCMVLD";
        StringBuilder sb = new StringBuilder();
        for (i = 0; i < 4; i++) {
            switch (arr[i]) {
                case 0:
                    break;
                case 1:
                    sb.insert(0, chars.charAt(i));
                    break;
                case 2:
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i));
                    break;
                case 3:
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i));
                    break;
                case 4:
                    sb.insert(0, chars.charAt(i + 4));
                    sb.insert(0, chars.charAt(i));
                    break;
                case 5:
                    sb.insert(0, chars.charAt(i + 4));
                    break;
                case 6:
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i + 4));
                    break;
                case 7:
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i + 4));
                    break;
                case 8:
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i));
                    sb.insert(0, chars.charAt(i + 4));
                    break;
                case 9:
                    sb.insert(0, chars.charAt(i + 1));
                    sb.insert(0, chars.charAt(i));
                    break;
            }
        }
        System.out.println(sb.toString());
        return sb.toString();
    }

    public static void main(String[] args) {
        Solution_12 solution_12 = new Solution_12();
        solution_12.intToRoman(1994);
        solution_12.intToRoman(104);
        solution_12.intToRoman(58);
    }
}
