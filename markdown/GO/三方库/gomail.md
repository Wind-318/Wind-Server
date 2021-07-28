## gomail
- ### 获取包
  - go get gopkg.in/gomail.v2
  - ### 示例代码
    ```Go
    package main

    import (
        "github.com/go-gomail/gomail"
    )

    func main() {
        // 发送者邮箱
        senderAccount := "xxx@gmail.com"

        // 发送者邮箱密码（或授权码）
        senderPassword := "passwd"

        // 接收者邮箱
        receiver := "yyy@outlook.com"

        title := "这里是标题"
        text := "这里是正文"

        // 附件地址
        attachs := "c:\\xxx\\yyy.jpg"

        // 服务器地址，outlook 是 smtp.office365.com
        serverAddr := "smtp.office365.com"

        // 服务器端口，outlook 是 587
        serverPort := 587

        // 生成一条新消息
        mess := gomail.NewMessage()
        // 设置发送方
        mess.SetHeader("From", senderAccount)
        // 设置接收方
        mess.SetHeader("To", receiver)
        // 设置标题
        mess.SetHeader("Subject", title)
        // 设置正文
        mess.SetBody("text/html", text)
        // 添加附件
        mess.Attach(attachs)
        // 添加图片
        /*
          m := gomail.NewMessage()
          m.Embed(`图片地址`)
          m.SetBody("text/html", `<img src="图片名" alt="My image" />`)
        */

        // 发送
        dial := gomail.NewDialer(serverAddr, serverPort, senderAccount, senderPassword)
        dial.DialAndSend(mess)
    }
    ```