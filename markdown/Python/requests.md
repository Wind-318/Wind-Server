# requests
| 方法 | 作用 |
| :-: | :-: |
| requests.get(url) | 以 GET 方式请求 url，返回一个响应对象 res |
| requests.post(url) | 以 POST 方式请求 url，返回一个响应对象 res |

- ## res 响应对象的方法：
  - **res.status_code**：返回响应状态码
  - **res.encoding**：设定编码格式
  - **res.content**：返回请求内容二进制形式
  - **res.text**：返回请求内容的字符串形式