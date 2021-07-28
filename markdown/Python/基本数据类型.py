# List中可以装不同类型的元素

list1 = [1,1]
# 复制
list2 = list1.copy()
print(list2)
# 追加
list1.append("asd")
print(list1)
# 统计数量
print(list1.count(1))
# 搜索指定区间，返回第一个出现值的下标
print(list1.index(1, 0, 3))
# 插入
list1.insert(2, "qwe")
# 删除1索引的元素
list1.pop(1)
# 删除最后一个元素
list1.pop()
# 删除
list1.remove("asd")
# 翻转
list1.reverse()
print(list1)
# 清空
list1.clear()
print(list1)
# 排序(list中的元素本身要能比较)
list2.sort()



# tuple与List相似，不过元素不能修改，所以没有append和pop等方法
# 如果tuple中包含list，当list中的元素变化时，tuple仍然指向此list，所以tuple没有改变，改变的是list指向的元素

# 示例：
tuple = (1,2, [3,4])
tuple[2][0] = 5
tuple[2][1] = 6
print(tuple)

# Python中的dict在某些语言中叫map，使用键值对存储(key-value)
map = {"name":"龙套乙"}
# 复制一个map
temp = map.copy()
# 返回key对应的值，若没有则返回默认设置的值
map.get("asd",13)
# 查看map的所有key
print(map.keys())
# 删除key
map.pop("name")
# 插入key，若已经有则返回对应值，若没有则返回设置的默认值
print(map.setdefault("age", 18))
# map中的值
print(map.values())
# 清空
map.clear()

# set存储key，而且不能重复
# 添加一个元素
set.add(value)

# 显示一个集合和其它集合的不同，返回这个新的集合
set.difference(self, other)

# 从这个集合中移除其它集合中的key
set.difference_update(self, *other)

# 从集合中移除一个元素
set.remove(value)