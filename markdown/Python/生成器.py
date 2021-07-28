# 通过生成式来生成列表（list）
# 生成2n + 1的数(n从1到100)
list1 = [2 * i + 1 for i in range(1, 100)]
for i in list1:
    print(i,)

temp = []
for i in range(97,123):
    temp.append(chr(i))

# 例2：
list2 = [a + b for a in temp for b in temp if a < b]
for i in list2:
    print(i,)


# 生成器：将生成式的中括号变为小括号，生成器保存的是表达式
list1 = [2 * i + 1 for i in range(1, 100)]
for i in list1:
    print(i,)
