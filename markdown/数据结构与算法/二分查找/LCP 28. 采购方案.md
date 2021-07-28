- ## [Leetcode.LCP 28. 采购方案](https://leetcode-cn.com/problems/4xy4Wx/)
- 对数组中每个元素，查找第一个大于 target - nums[i] 的元素的位置 t，这个元素的组合数就为 t 减这个元素开始的位置再减一。
- 时间复杂度：$O(nlogn)$
- 空间复杂度：$O(logn)$
    ```C++
    class Solution {
    public:
        int purchasePlans(vector<int>& nums, int target) {
            sort(nums.begin(), nums.end());
            int ans = 0;
            int n = nums.size();
            for (int i = 0; i < n; i++) {
                auto temp = target - nums[i];
                
                auto t = upper_bound(nums.begin() + i, nums.end(), temp) - (nums.begin() + i);
                if (t > 0) {
                    t = t - 1;
                }
                ans += t;
                ans %= (int)(1e9 + 7);
            }
            
            return ans;
        }
    };
    ```