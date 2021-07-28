## Examples:  
- 鼠标
    ```Go
    package main

    import (
        "github.com/go-vgo/robotgo"
    )

    func main() {
        robotgo.ScrollMouse(10, "up")
        robotgo.MouseClick("left", true)
        robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
    }
    ```
- 键盘
    ```Go
    package main

    import (
        "fmt"
        "github.com/go-vgo/robotgo"
    )

    func main() {
        robotgo.TypeStr("Hello World. Winter is coming!")
        robotgo.TypeStr("だんしゃり", 1.0)
        // robotgo.TypeString("テストする")

        robotgo.TypeStr("Hi galaxy. こんにちは世界.")
        robotgo.Sleep(1)

        // ustr := uint32(robotgo.CharCodeAt("テストする", 0))
        // robotgo.UnicodeType(ustr)

        robotgo.KeyTap("enter")
        // robotgo.TypeString("en")
        robotgo.KeyTap("i", "alt", "command")

        arr := []string{"alt", "command"}
        robotgo.KeyTap("i", arr)

        robotgo.WriteAll("テストする")
        text, err := robotgo.ReadAll()
        if err == nil {
            fmt.Println(text)
        }
    }
    ```
- 屏幕
    ```Go
    package main

    import (
        "fmt"
        "github.com/go-vgo/robotgo"
    )

    func main() {
        x, y := robotgo.GetMousePos()
        fmt.Println("pos: ", x, y)

        color := robotgo.GetPixelColor(100, 200)
        fmt.Println("color----", color)
    }
    ```
- 位图
    ```Go
    package main

    import (
        "fmt"
        "github.com/go-vgo/robotgo"
    )

    func main() {
        bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
        // use `defer robotgo.FreeBitmap(bit)` to free the bitmap
        defer robotgo.FreeBitmap(bitmap)
        fmt.Println("...", bitmap)

        fx, fy := robotgo.FindBitmap(bitmap)
        fmt.Println("FindBitmap------", fx, fy)

        robotgo.SaveBitmap(bitmap, "test.png")
    }
    ```
- 事件
    ```Go
    package main

    import (
        "fmt"
        "github.com/go-vgo/robotgo"
        hook "github.com/robotn/gohook"
    )

    func main() {
        add()
        low()
        event()
    }

    func add() {
        fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
        robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
            fmt.Println("ctrl-shift-q")
            robotgo.EventEnd()
        })

        fmt.Println("--- Please press w---")
        robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
            fmt.Println("w")
        })

        s := robotgo.EventStart()
        <-robotgo.EventProcess(s)
    }

    func low() {
        EvChan := hook.Start()
        defer hook.End()

        for ev := range EvChan {
            fmt.Println("hook: ", ev)
        }
    }

    func event() {
        ok := robotgo.AddEvents("q", "ctrl", "shift")
        if ok {
            fmt.Println("add events...")
        }

        keve := robotgo.AddEvent("k")
        if keve {
            fmt.Println("you press... ", "k")
        }

        mleft := robotgo.AddEvent("mleft")
        if mleft {
            fmt.Println("you press... ", "mouse left button")
        }
    }
    ```
- 窗口句柄
    ```Go
    package main

    import (
        "fmt"
        "github.com/go-vgo/robotgo"
    )

    func main() {
        fpid, err := robotgo.FindIds("Google")
        if err == nil {
            fmt.Println("pids...", fpid)

            if len(fpid) > 0 {
                robotgo.ActivePID(fpid[0])

                robotgo.Kill(fpid[0])
            }
        }

        robotgo.ActiveName("chrome")

        isExist, err := robotgo.PidExists(100)
        if err == nil && isExist {
            fmt.Println("pid exists is", isExist)

            robotgo.Kill(100)
        }

        abool := robotgo.ShowAlert("test", "robotgo")
        if abool == 0 {
            fmt.Println("ok@@@ ", "ok")
        }

        title := robotgo.GetTitle()
        fmt.Println("title@@@ ", title)
    }
    ```
- ## 常用按键
| 按键 | 值 |
| :-: | :-: |
| 空格 | space |
| 小键盘的 0 到 9 | num0 到 num9 |
| 主键盘的 0 到 9 | 0 到 9 |
| f1 到 f24 | F1 到 F24 |
| 上箭头 | up |
| 下箭头 | down |
| 左箭头 | left |
| 右箭头 | right |
| windows 键 | cmd |
| 截图 | printscreen |
| 其余多数按键直接对应其拼写 | # |
# 来源
- #### [https://gitee.com/veni0/robotgo](https://gitee.com/veni0/robotgo)