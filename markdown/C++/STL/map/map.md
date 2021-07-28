# map
- ## unordered_map 哈希冲突
  - c++ 采用开链法解决哈希冲突，冲突时会将值以头插法插入链表头部，实现 $O(1)$ 时间复杂度的插入。
  - 负载因子：哈希表中元素个数与桶数的比值
  - unordered_map 在负载因子超过上限后发生 rehash，增加桶的数量，然后将原数据迁移到新桶中。
- ## map
  - map 底层由红黑树实现，红黑树可以自动排序，插入、删除与查找时间复杂度为 $O(logn)$。
---
# 参考
- #### [深入了解C++（1）：hash冲突、退化](https://mp.weixin.qq.com/s?__biz=MzkyMjIxMzIxNA==&mid=2247483656&idx=1&sn=a204fedfbf2cf7f2023979c56b756c8a&chksm=c1f68f39f681062fe0e31f2a07fd576ff8fcb3b3016bf024afa8703520390ca7fa3b257997da&token=377973187&lang=zh_CN#rd)
- #### [深入了解C++（2）：unordered_map的insert函数](https://mp.weixin.qq.com/s?__biz=MzkyMjIxMzIxNA==&mid=2247483848&idx=1&sn=d459a04730a4e56653452eae9f71d424&chksm=c1f68ff9f68106ef0c606d105f8a25d9e0e7e241faf3bbf4b92b7d0484e711c10597596535a6&token=377973187&lang=zh_CN#rd)