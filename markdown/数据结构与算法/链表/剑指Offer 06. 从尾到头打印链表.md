## [剑指 Offer 06. 从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

- 利用栈先进后出的特性，遍历链表，放入栈中，再依次弹栈，即可反向打印出。
- **时间复杂度**：$O(n)$
- **空间复杂度**：$O(n)$
    ```C++
    /**
    * Definition for singly-linked list.
    * struct ListNode {
    *     int val;
    *     ListNode *next;
    *     ListNode(int x) : val(x), next(NULL) {}
    * };
    */
    class Solution {
    public:
        vector<int> reversePrint(ListNode* head) {
            stack<int> s;
            while(head != NULL)
            {
                s.push(head->val);
                head = head->next;
            }
            vector<int> v;
            while(!s.empty())
            {
                v.push_back(s.top());
                s.pop();
            }
            return v;
        }
    };
    ```