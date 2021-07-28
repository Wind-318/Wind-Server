## wordcloud
- ### 安装
  - pip install wordcloud
- ### 操作
  - **生成一个词云对象**：w = wordcloud.WordCloud(参数（可以为空）)，参数如下
    | 参数 | 作用 |
    | :-: | :-: |
    | width | 指定图片宽度（像素） |
    | height | 指定图片高度（像素） |
    | min_font_size | 指定最小字号 |
    | max_font_size | 指定最大字号 |
    | font_step | 指定字体步进间隔 |
    | font_path | 指定字体文件 |
    | max_words | 最大单词数 |
    | stop_words | 排除的词 |
    | mask | 指定形状 |
    | background_color | 指定背景颜色 |
  - **加载文本数据（txt 为字符串类型）**：w.generate(txt)
  - **导出词云文件（filename 为文件路径）**：w.to_file(filename)
  - ### 示例(如果乱码要选择中文字体)
  - ```Python
    import wordcloud
    import jieba

    fileObj = open("1.txt", encoding = "utf-8")
    texts = fileObj.read()
    w = wordcloud.WordCloud(width = 1920, height = 1080, font_path="simsun.ttc")
    w.generate(" ".join(jieba.lcut(texts)))
    w.to_file("test.png")

    fileObj.close()
    ```
***
## 参考
- ### [word_cloud](https://github.com/amueller/word_cloud)
- ### [Python之wordcloud库使用](https://blog.csdn.net/jinsefm/article/details/80645588)