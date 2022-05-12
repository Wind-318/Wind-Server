window.onload = function() {
    var container = document.getElementById("container");
    $.ajax({
        url: "/bbs/InquireClassification",
        type: "GET",
        success: function(data) {
            for (var i = 0; i < data["num"]; i++) {
                // 新建 a 标签
                var temp = document.createElement("a");
                temp.setAttribute("id", data["types"][i]);

                // a 标签的子元素
                var att1 = document.createElement("img");
                var att2 = document.createElement("span");
                att1.setAttribute("src", data["pic"][i]);
                att2.setAttribute("id", temp.id + "Span");
                att2.setAttribute("style", "left: 40%; top: 40%; font-size: 30px");
                att2.innerHTML = data["types"][i];
                temp.appendChild(att1);
                temp.appendChild(att2);

                // 将 a 标签加到 div 中
                container.appendChild(temp);

                // 增加监控 js
                var jsScript = document.createElement("script");
                jsScript.type = "text/javascript";
                jsScript.text = `var btn = document.getElementById("` + temp.id + `");
                    var divs = document.getElementById("container");
                    var Addr = "` + data["Addr"] + `";

                    btn.onclick = function(ev) {
                        // 刷新分类下的文章
                        document.getElementById("container").innerHTML = "";

                        $.ajax({
                            url: "/bbs/InquireText",
                            type: "POST",
                            data: {
                                "types": "` + temp.id + `"
                            },
                            success: function(data) {
                                for (var i = data["num"] - 1; i >= 0; i--) {  
                                    var temp = document.createElement('a');
                                    temp.setAttribute("href", data["urls"][i]);
                                    temp.setAttribute("target", "_blank");
                                    temp.setAttribute("id", data["id"][i]);

                                    // a 标签的子元素
                                    var att = document.createElement("input");
                                    att.setAttribute("type", "checkbox");
                                    att.setAttribute("id", data["id"][i] + "checkbox");
                                    att.setAttribute("class", "deletebtn");
                                    temp.appendChild(att);
                                    var att1 = document.createElement("img");
                                    var att2 = document.createElement("span");
                                    var att3 = document.createElement("p");
                                    var att4 = document.createElement("p");
                                    att1.setAttribute("src", data["picurl"][i]);
                                    att2.setAttribute("id", temp.id + "Span");
                                    att3.setAttribute("id", temp.id + "Spans");
                                    att4.setAttribute("id", temp.id + "Spanss");
                                    att3.setAttribute("class", "Spans");
                                    att4.setAttribute("class", "Spanss");
                                    att4.innerHTML = data["titles"][i];
                                    att2.innerHTML = data["description"][i];
                                    att3.innerHTML = "作者：" + data["author"][i] + "\t\t\t发布时间：" + data["create_time"][i] + "\t\t\t修改于：" + data["update_time"][i];
                                    temp.appendChild(att1);
                                    temp.appendChild(att2);
                                    temp.appendChild(att3);
                                    temp.appendChild(att4);

                                    divs.appendChild(temp);
                                }
                            },
                            fail: function() {
                                location.href = "/serverError";
                            }
                        })
                    }`;
                document.body.appendChild(jsScript);
            }

        },
        fail: function(data) {
            location.href = "/serverError";
        }
    })
}