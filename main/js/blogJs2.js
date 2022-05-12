var t1 = 0,
    t2 = 0,
    timer = null;

window.onscroll = function() {
    clearTimeout(timer);
    timer = setTimeout(isScrollEnd, 300);
    t1 = t1 = document.documentElement.scrollTop || document.body.scrollTop;
}

function isScrollEnd() {
    t2 = document.documentElement.scrollTop || document.body.scrollTop;
    if (t2 == t1) {
        var scrollTop = $(this).scrollTop();　　
        var scrollHeight = $(document).height();　　
        var windowHeight = $(this).height();　　
        if (scrollTop + windowHeight + windowHeight >= scrollHeight && document.getElementById("keyword").value == "") {
            var divs = document.getElementById("container");
            var num = document.getElementById("container").children.length;
            $.ajax({
                url: "/bbs/InquirePageNums",
                type: "POST",
                data: {
                    "num": num
                },
                success: function(data) {
                    if (data["num"] < num) {
                        return;
                    }
                    var nums = data["end"] - data["start"] + 1;
                    if (nums <= 10) {
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
}

// 查找
function search() {
    var text = document.getElementById("keyword");
    $.ajax({
        url: "/bbs/Search",
        type: "POST",
        data: {
            "text": text.value
        },
        success: function(data) {
            var divs = document.getElementById("container");
            divs.innerHTML = "";
            if (text.value == "") {
                for (var i = data["num"]; i >= data["num"] - 10; i--) {
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
                return;
            }
            for (var i = data["num"]; i >= 1; i--) {
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
    })
}