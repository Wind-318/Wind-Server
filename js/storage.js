window.onload = function() {
    // users 节点
    var ulNode = document.getElementById("users");
    // 添加 user
    $.ajax({
        url: "/user/getUsersName",
        type: "GET",
        success: function(data) {
            for (var i = 0; i < data["names"].length; i++) {
                var names = data["names"][i];
                // 添加子节点
                var liNode = document.createElement("li");
                liNode.setAttribute("class", "nav-item");
                ulNode.appendChild(liNode);
                // 添加子节点
                var aNode = document.createElement("a");
                aNode.setAttribute("class", "nav-link h4");
                aNode.setAttribute("data-bs-toggle", "pill");
                aNode.setAttribute("href", "#" + names);
                aNode.setAttribute("id", "user" + names);
                aNode.innerHTML = names;
                liNode.appendChild(aNode);
                // 添加脚本
                var ListenScript = document.createElement("script");
                ListenScript.innerHTML = `document.getElementById("user` + names + `").onclick = function() { loadFirst("` + names + `", 1); }`;
                document.body.appendChild(ListenScript);
                // 添加框架
                var root = document.getElementById("tabContent");
                var divNode = document.createElement("div");
                divNode.setAttribute("class", "container tab-pane");
                divNode.setAttribute("id", names);
                root.appendChild(divNode);
                // 框架子节点
                var div1Node = document.createElement("div");
                div1Node.setAttribute("class", "container border");
                divNode.appendChild(div1Node);
                // 子节点的子节点 1
                var div1_1Node = document.createElement("div");
                div1_1Node.setAttribute("class", "tab-content");
                div1Node.appendChild(div1_1Node);
                // 子节点的子节点 2
                var div1_2Node = document.createElement("div");
                div1_2Node.setAttribute("class", "container justify-content-center");
                div1Node.appendChild(div1_2Node);
                // 添加 ul 节点
                var ul1Node = document.createElement("ul");
                ul1Node.setAttribute("class", "nav nav-pills");
                ul1Node.setAttribute("role", "tablist");
                div1_2Node.appendChild(ul1Node);

                $.ajax({
                    url: "/storage/getUserStoragePicturePage",
                    type: "POST",
                    async: false,
                    data: {
                        "name": names,
                        "pageNum": 50
                    },
                    success: function(data) {
                        for (var i = 1; i <= data["page"]; i++) {
                            // 子 1 子节点
                            var div1_1_1Node = document.createElement("div");
                            if (i == 1) {
                                div1_1_1Node.setAttribute("class", "container tab-pane active");
                            } else {
                                div1_1_1Node.setAttribute("class", "container tab-pane");
                            }
                            div1_1_1Node.setAttribute("id", names + "Page" + i);
                            div1_1Node.appendChild(div1_1_1Node);
                            // 子 2 子节点
                            var li1Node = document.createElement("li");
                            li1Node.setAttribute("class", "nav-item");
                            li1Node.setAttribute("style", "list-style: none");
                            ul1Node.appendChild(li1Node);
                            var a1Node = document.createElement("a");
                            if (i == 1) {
                                a1Node.setAttribute("class", "nav-link active");
                            } else {
                                a1Node.setAttribute("class", "nav-link");
                            }
                            a1Node.setAttribute("data-bs-toggle", "pill");
                            a1Node.setAttribute("href", "#" + names + "Page" + i);
                            a1Node.setAttribute("id", names + i);
                            a1Node.innerHTML = "第 " + i + " 页";
                            li1Node.appendChild(a1Node);
                            // 添加脚本
                            var ListenScripts = document.createElement("script");
                            ListenScripts.innerHTML = `document.getElementById("` + names + i + `").onclick = function() { loadFirst("` + names + `", ` + i + `); }`;
                            document.body.appendChild(ListenScripts);
                        }
                    },
                    fail: function() {}
                })
            }
        },
        fail: function() {}
    })
}


// 加载第 n 页
function loadFirst(names, page) {
    var root = document.getElementById(names + "Page" + page);
    root.innerHTML = "";
    // 每页加载数量
    var pageNum = 50;
    // 偏移量
    var moveNum = pageNum * (page - 1);
    $.ajax({
        url: "/storage/getUserStoragePicture",
        type: "POST",
        data: {
            "name": names,
            "num": moveNum,
            "onceChoose": pageNum
        },
        success: function(data) {
            // 加载
            for (var i = 0; i < data["num"]; i++) {
                var aNode = document.createElement("a");
                aNode.setAttribute("href", data["picPath"][i]);
                aNode.setAttribute("target", "_blank");
                root.appendChild(aNode);
                var imgNode = document.createElement("img");
                imgNode.setAttribute("src", data["smallPicPath"][i]);
                imgNode.setAttribute("style", "width: 400px;");
                imgNode.setAttribute("class", "p-2");
                aNode.appendChild(imgNode);
            }
        },
        fail: function() {}
    })
}