window.onload = function() {
    var divs = document.getElementById("container");
    var num = document.getElementById("container").children.length;
    var userStatus = document.getElementById("userStatus")
    $.ajax({
        url: "/bbs/InquirePageNums",
        type: "POST",
        data: {
            "num": num
        },
        success: function(data) {
            if (data["status"] == 1) {
                userStatus.innerText = "文章管理";
                userStatus.href = "/bbsFile/" + data["ids"] + "/user.html";

                if (data["isSystem"] == 1) {
                    var litemp = document.createElement("li");
                    var temp = document.createElement("a");
                    var ulcontainer = document.getElementById("ulcontainer");
                    temp.setAttribute("id", "deletebtn");
                    temp.innerText = "删除"
                    litemp.appendChild(temp);
                    ulcontainer.appendChild(litemp);

                    var jsScript = document.createElement("script");
                    jsScript.text = `var btn = document.getElementById("deletebtn");
                            btn.onclick = function(ev) {
                                var arr = document.getElementById("container").children;
                                var datas = new Array();
                                for (var i = 0; i < arr.length; i++) {
                                    if (document.getElementById(arr[i].id + "checkbox").checked) {
                                        datas.push(arr[i].id);
                                    }
                                }
                                $.ajax({
                                    url: "/bbs/DeleteBlog",
                                    type: "POST",
                                    data: {
                                        "checked": JSON.stringify(datas)
                                    },
                                    success: function(data) {
                                        location.href = "/bbs";
                                    }
                                })
                            }`;
                    document.body.appendChild(jsScript);
                }


                var litemp = document.createElement("li");
                var temp = document.createElement("a");
                var ulcontainer = document.getElementById("ulcontainer");
                temp.setAttribute("href", "/user/Exit")
                temp.innerText = "退出登录"
                litemp.appendChild(temp);
                ulcontainer.appendChild(litemp);
            }
            if (data["num"] < num) {
                retrun;
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
                    att3.innerHTML = "作者:" + data["author"][i - 1] + "\t\t\t发布时间:" + data["create_time"][i - 1] + "\t\t\t修改于:" + data["update_time"][i - 1];
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
                    att3.innerHTML = "作者:" + data["author"][i - 1] + "\t\t\t发布时间:" + data["create_time"][i - 1] + "\t\t\t修改于:" + data["update_time"][i - 1];
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