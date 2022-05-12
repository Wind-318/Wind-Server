window.onload = function() {
    var webs = document.getElementById("webs");
    $.ajax({
        url: "/collections/GetWebs",
        type: "GET",
        success: function(data) {
            if (data["msg"] != undefined) {
                alert(data["msg"]);
                return;
            }
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
        url: "/user/IsSystem",
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
                var litemps = document.createElement("li");
                var atemps = document.createElement("a");
                atemps.innerHTML = "新增图片";
                atemps.setAttribute("id", "picbtn");
                litemps.appendChild(atemps);
                temp.appendChild(litemps);

                // 添加图片
                var divs = document.createElement("div");
                divs.setAttribute("style", "margin: 0 auto;");
                var inputpic = document.createElement("input");
                inputpic.setAttribute("id", "backgroundpic");
                inputpic.setAttribute("style", "position: absolute; top: 0; right: 0; opacity: 0; width: 70px; height: 76px");
                inputpic.setAttribute("type", "file");
                inputpic.setAttribute("accept", "image/*");
                var divshow = document.createElement("div");
                var picspan = document.createElement("span");
                picspan.setAttribute("style", "font-size: 12px;");
                picspan.innerHTML = "上传背景：";
                var imgpic = document.createElement("img");
                imgpic.setAttribute("id", "upload");
                imgpic.setAttribute("style", "height: 40px; vertical-align: middle;");
                imgpic.setAttribute("src", "../picture/profile.jpg");
                divshow.appendChild(picspan);
                divshow.appendChild(imgpic);
                divs.appendChild(inputpic);
                divs.appendChild(divshow);
                document.body.appendChild(divs);

                var jsScripts = document.createElement("script");
                jsScripts.text = `var btn = document.getElementById("picbtn");
                    btn.onclick = function(ev) {
                        var pic = document.getElementById("backgroundpic");
                        if (pic.files[0] == undefined) {
                            alert('背景图不能为空！');
                            return;
                        } 
                        
                        var formData = new FormData();
                        formData.append("pic", pic.files[0]);
                        $.ajax({
                            url: "/collections/PutPic",
                            type: "POST",
                            data: formData,
                            cache: false,
                            processData: false,
                            contentType: false,
                            success: function(data) {
                                alert(data["msg"]);
                            },
                            fail: function() {
                                alert("fail");
                            }
                        })
                }`;
                document.body.appendChild(jsScripts);

                var jsScript = document.createElement("script");
                jsScript.text = `var btn = document.getElementById("newbtn");
                    btn.onclick = function(ev) {
                    var url = prompt("请输入 url：");
                    var comment = prompt("请输入备注：");
                    var picurl = prompt("请输入图片地址：");
                    if (url == "" || comment == "" || picurl == "") {
                        alert("输入不能为空！");
                        return;
                    }
                    
                    // 请求添加
                    $.ajax({
                        url: "/collections/PutWebs",
                        type: "POST",
                        data: {
                            "url": url,
                            "comment": comment,
                            "picurl": picurl
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