window.onload = function() {
    var webs = document.getElementById("webs");
    $.ajax({
        url: "/collections/GetWebs",
        type: "GET",
        success: function(data) {
            var len = data["ids"].length;
            for (var i = 0; i < len; i++) {
                var litemp = document.createElement("li");
                var temp = document.createElement("a");
                var ptemp = document.createElement("p");
                var pic = document.createElement("img");
                pic.setAttribute("src", data["picurls"][i]);
                temp.setAttribute("href", data["urls"][i]);
                temp.appendChild(pic);
                temp.setAttribute("target", "_blank");
                ptemp.innerHTML = data["comments"][i];

                litemp.appendChild(temp);
                litemp.appendChild(ptemp);

                webs.appendChild(litemp);
            }
        }
    })
    $.ajax({
        url: "/collections/IsSystem",
        type: "GET",
        success: function(data) {
            if (data["msg"] == "success") {
                var temp = document.getElementById("ulcontainer");
                var litemp = document.createElement("li");
                var atemp = document.createElement("a");
                atemp.innerHTML = "新增链接";
                atemp.setAttribute("id", "newbtn");
                litemp.appendChild(atemp);
                temp.appendChild(litemp);

                var jsScript = document.createElement("script");
                jsScript.text = `var btn = document.getElementById("newbtn");
                    btn.onclick = function(ev) {
                    var url = prompt("请输入 url：");
                    var comment = prompt("请输入备注：");
                    if (url == "" || comment == "") {
                        alert("输入不能为空！");
                        return;
                    }

                    $.ajax({
                        url: "/collections/PutWebs",
                        type: "POST",
                        data: {
                            "url": url,
                            "comment": comment
                        },
                        success: function(data) {
                            location.href = "/collections";
                        }
                    })
                }`;
                document.body.appendChild(jsScript);
            }
        }
    })
}