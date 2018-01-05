package com.michel;

public class Solution_747 {
    public int dominantIndex(int[] nums) {
        int length = nums.length;

        int secondMax, max, index;
        secondMax = 0;
        max = 0;
        index = -1;
        for(int i = 0; i < length; i++){
            if(nums[i]>max){
                secondMax = max;
                max = nums[i];
                index = i;
            }else if(nums[i] > secondMax){
                secondMax = nums[i];
            }
        }

        if(max >= secondMax * 2){
            return index;
        }
        return -1;
    }

    public int dominantIndexOfficial(int[] nums) {
        int maxIndex = 0;
        for (int i = 0; i < nums.length; ++i) {
            if (nums[i] > nums[maxIndex])
                maxIndex = i;
        }
        for (int i = 0; i < nums.length; ++i) {
            if (maxIndex != i && nums[maxIndex] < 2 * nums[i])
                return -1;
        }
        return maxIndex;
    }
}
