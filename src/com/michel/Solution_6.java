package com.michel;

public class Solution_6 {
    public String convert(String s, int numRows) {
        if (s == null || numRows <= 0) {
            return "";
        }
        int len = s.length();
        if (numRows == 1 || numRows >= len) {
            return s;
        }
        StringBuilder sb = new StringBuilder();
        for (int row = 0; row < numRows; row++) {
            if (row == 0 || row == numRows - 1) {
                int i = 0;
                int pos = (i++) * (2 * numRows - 2) + row;
                while (pos < len) {
                    sb.append(s.charAt(pos));
                    pos = (i++) * (2 * numRows - 2) + row;
                }
            } else {
                sb.append(s.charAt(row));
                int i = 1;
                int pos = 0;
                while (true) {
                    pos = i * (2 * numRows - 2) - row;
                    if (pos < len) {
                        sb.append(s.charAt(pos));
                    } else {
                        break;
                    }
                    pos = i * (2 * numRows - 2) + row;
                    if (pos < len) {
                        sb.append(s.charAt(pos));
                    } else {
                        break;
                    }
                    i++;
                }

            }
        }

        return sb.toString();
    }

    public static void main(String[] args) {
        Solution_6 solution_6 = new Solution_6();
        solution_6.convert("A", 1);
    }
}
