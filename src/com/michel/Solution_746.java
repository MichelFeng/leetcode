package com.michel;

public class Solution_746 {
    Integer[] cache = null;
    public int minCostClimbingStairs(int[] cost) {
        this.cache = new Integer[cost.length];
        return Math.min(calculateCost(cost, 0), calculateCost(cost, 1));
    }

    private int calculateCost(int[] cost, int idx) {
        if (idx >= cost.length) {
            return 0;
        }
        if (this.cache[idx]!=null){
            return this.cache[idx];
        }

        int currentCost =  cost[idx] + Math.min(calculateCost(cost, idx + 1), calculateCost(cost, idx + 2));
        this.cache[idx] = currentCost;
        return currentCost;
    }

    /**
     * 官方答案，思路是类似求解斐波那契数列
     * @param cost
     * @return
     */
    public int minCostClimbingStairsOfficial(int[] cost) {
        int f1 = 0, f2 = 0;
        for (int i = cost.length - 1; i >= 0; --i) {
            int f0 = cost[i] + Math.min(f1, f2);
            f2 = f1;
            f1 = f0;
        }
        return Math.min(f1, f2);
    }
}
