## Turtle 库
- ### 随机位置画出随机颜色的五角星 Demo
    ```Python
    import turtle

    import random



    turtle.pensize(4)

    turtle.screensize(1500, 800)

    turtle.colormode(255)



    for _ in range(50):

        x, y = random.randint(-750, 750), random.randint(-400, 400)

        turtle.up()

        turtle.goto(x, y)

        turtle.down()

        r, g, b = random.randint(0, 255), random.randint(0, 255), random.randint(0, 255)

        turtle.pencolor(r, g, b)

        turtle.fillcolor(r, g, b)

        turtle.begin_fill()

        for _ in range(5):

            turtle.forward(100)

            turtle.right(144) 

        turtle.end_fill()

        

    turtle.mainloop()
    ```