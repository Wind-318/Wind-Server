window.onscroll = function() {
    var scrollTop = $(this).scrollTop();　　
    var scrollHeight = $(document).height();　　
    var windowHeight = $(this).height();　　
    if (scrollTop + windowHeight == scrollHeight) {
        var divs = document.getElementById("container");
        var num = document.getElementById("container").children.length;
        $.ajax({
            url: "/blogs/InquirePageNums",
            type: "POST",
            data: {
                "num": num
            },
            success: function(data) {
                if (data["num"] < num) {
                    retrun;
                }
                var nums = data["end"] - data["start"] + 1;
                if (nums <= 20) {
                    for (var i = data["start"]; i < data["end"]; i++) {
                        var temp = document.createElement('a');
                        temp.setAttribute("href", data["urls"][i - 1]);
                        temp.setAttribute("target", "_blank");
                        temp.setAttribute("id", data["id"][i - 1]);

                        // a 标签的子元素
                        if (data["isSystem"] == 1) {
                            var att = document.createElement("input");
                            att.setAttribute("type", "checkbox");
                            att.setAttribute("id", data["id"][i - 1] + "checkbox");
                            att.setAttribute("class", "deletebtn");
                            temp.appendChild(att);
                        }
                        var att1 = document.createElement("img");
                        var att2 = document.createElement("span");
                        var att3 = document.createElement("p");
                        var att4 = document.createElement("p");
                        att1.setAttribute("src", data["picurl"][i - 1]);
                        att2.setAttribute("id", temp.id + "Span");
                        att3.setAttribute("id", temp.id + "Spans");
                        att4.setAttribute("id", temp.id + "Spanss");
                        att3.setAttribute("class", "Spans");
                        att4.setAttribute("class", "Spanss");
                        att4.innerHTML = data["titles"][i - 1];
                        att2.innerHTML = data["description"][i - 1];
                        att3.innerHTML = "作者：" + data["author"][i - 1] + "\t\t\t发布时间：" + data["create_time"][i - 1] + "\t\t\t修改于：" + data["update_time"][i - 1];
                        temp.appendChild(att1);
                        temp.appendChild(att2);
                        temp.appendChild(att3);
                        temp.appendChild(att4);

                        divs.appendChild(temp);
                    }
                } else {
                    for (var i = data["start"]; i < data["end"]; i++) {
                        var temp = document.createElement('a');
                        temp.setAttribute("href", data["urls"][i - 1]);
                        temp.setAttribute("target", "_blank");
                        temp.setAttribute("id", data["id"][i - 1]);

                        // a 标签的子元素
                        if (data["isSystem"] == 1) {
                            var att = document.createElement("input");
                            att.setAttribute("type", "checkbox");
                            att.setAttribute("id", data["id"][i - 1] + "checkbox");
                            att.setAttribute("class", "deletebtn");
                            temp.appendChild(att);
                        }

                        var att1 = document.createElement("img");
                        var att2 = document.createElement("span");
                        var att3 = document.createElement("p");
                        var att4 = document.createElement("p");
                        att1.setAttribute("src", data["picurl"][i - 1]);
                        att2.setAttribute("id", temp.id + "Span");
                        att3.setAttribute("id", temp.id + "Spans");
                        att4.setAttribute("id", temp.id + "Spanss");
                        att3.setAttribute("class", "Spans");
                        att4.setAttribute("class", "Spanss");
                        att4.innerHTML = data["titles"][i - 1];
                        att2.innerHTML = data["description"][i - 1];
                        att3.innerHTML = "作者：" + data["author"][i - 1] + "\t\t\t发布时间：" + data["create_time"][i - 1] + "\t\t\t修改于：" + data["update_time"][i - 1];
                        temp.appendChild(att1);
                        temp.appendChild(att2);
                        temp.appendChild(att3);
                        temp.appendChild(att4);

                        divs.appendChild(temp);
                    }
                }
            },
            fail: function() {
                location.href = "/serverError";
            }
        })
    }
}