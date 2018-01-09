package com.michel;

public class Solution_717 {
    public boolean isOneBitCharacter(int[] bits) {
        int len = bits.length;
        int i = 0;
        while (i < len - 1) {
            if (bits[i] == 1) {
                i += 2;
            } else {
                i++;
            }
        }
        return i == len - 1;
    }
}
