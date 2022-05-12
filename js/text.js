// 获取文章内容
$.ajax({
    url: "/bbs/GetUserText",
    type: "POST",
    data: {
        "ids": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        document.getElementById('contentText').innerHTML = marked(data["content"]);
        document.getElementById("titles").innerHTML = data["title"];
    },
})

// 获取头像
$.ajax({
    url: "/bbs/GetProfile",
    type: "POST",
    data: {
        "id": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        document.getElementById("profile").setAttribute("src", data["pic"]);
    }
})

// 获取文章评论
$.ajax({
    url: "/bbs/TextComment",
    type: "POST",
    data: {
        "id": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        var commentroot = document.getElementById("root");
        for (var i = 0; i < data["nums"]; i++) {
            // 评论
            var tempsr = document.createElement("div");
            tempsr.setAttribute("id", data["ids"][i] + "tempsr");
            tempsr.setAttribute("class", "tempsr");
            var temps = document.createElement("div");
            temps.setAttribute("class", "temps");
            var tempscon = document.createElement("div");
            tempscon.setAttribute("class", "tempscon");
            var tempstitle = document.createElement("div");
            tempstitle.setAttribute("class", "tempstitle");
            // 头像
            var imgs = document.createElement("img");
            imgs.setAttribute("style", "border-radius: 50%; width: auto");
            imgs.setAttribute("src", data["pics"][i]);
            // 回复按钮
            var btns = document.createElement("button");
            btns.setAttribute("id", data["ids"][i] + "childReply");
            btns.setAttribute("style", "color: white; cursor: pointer; right: 0; position: absolute; bottom: 0; width: 60px; height: 60px; border-radius: 50%; background-color: rgb(50, 75, 150);")
            btns.innerHTML = "回复";
            // 回复人
            var authorss = document.createElement("span");
            authorss.innerHTML = data["authors"][i];
            authorss.setAttribute("style", "color: white; margin: 0 auto;");
            // 评论内容
            var contents = document.createElement("span");
            if (data["parents"][i] == "") {
                contents.innerHTML = marked(data["contents"][i]);
            } else {
                contents.innerHTML = "回复 " + data["parents"][i] + "：" + marked(data["contents"][i]);
            }
            // 修改时间
            var update_times = document.createElement("span");
            update_times.innerHTML = "最后修改于：" + data["update_time"][i];
            update_times.setAttribute("style", "color: white; margin-bottom: 0;");

            // 添加到子评论中
            temps.appendChild(imgs);
            temps.appendChild(authorss);
            temps.setAttribute("style", "flex-direction: column; display: flex; margin-left: 100px; width: 200px; flex-direction: column; background-color: rgba(122, 122, 122, 0.6);");
            tempscon.appendChild(contents);
            tempscon.setAttribute("style", "flex-direction: column; width: auto; background-color: rgba(255, 255, 255, 0.8);");
            tempscon.setAttribute("id", data["ids"][i] + "tempscon");
            var tempconsplits = document.createElement("div");
            tempconsplits.setAttribute("style", "width: 100%; height: 70px;");
            tempstitle.appendChild(update_times);
            tempstitle.appendChild(btns);
            tempstitle.setAttribute("style", "height: 30px; width: 100%; display: flex; flex-direction: row; bottom: 0; position: absolute;");
            tempsr.appendChild(temps);
            var contentandtitles = document.createElement("div");
            contentandtitles.setAttribute("style", "width: 1000px; height: auto; flex-direction: column; display: flex; position: relative;");
            contentandtitles.appendChild(tempscon);
            contentandtitles.appendChild(tempconsplits);
            contentandtitles.appendChild(tempstitle);
            tempsr.appendChild(contentandtitles);
            tempsr.setAttribute("style", "flex-direction: row; display: flex; height: auto;");
            commentroot.appendChild(tempsr);

            // 分割线
            var splitlines = document.createElement("div");
            splitlines.setAttribute("style", "margin-left: 100px; width: 1200px; height: 10px; background-color: rgb(47, 30, 92);");
            commentroot.appendChild(splitlines);
        }
    }
})

// 检测键盘抬起事件
function mdSwitch() {
    var mdValue = document.getElementById("texts").value;
    var html = marked(mdValue);
    document.getElementById("show-area").innerHTML = html;
}

// 赞
document.getElementById("praise").onclick = function(ev) {
    $.ajax({
        url: "/bbs/Parise",
        type: "POST",
        data: {
            "id": document.getElementsByName("main")[0].id
        },
        success: function(data) {
            $.ajax({
                url: "/bbs/PariseNum",
                type: "POST",
                data: {
                    "id": document.getElementsByName("main")[0].id
                },
                success: function(data) {
                    document.getElementById("praiseNum").innerHTML = data["num"];
                }
            })
        }
    })
}

// 获取赞数
$.ajax({
    url: "/bbs/PariseNum",
    type: "POST",
    data: {
        "id": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        document.getElementById("praiseNum").innerHTML = data["num"];
    }
})

// 获取点击数
$.ajax({
    url: "/bbs/Views",
    type: "POST",
    data: {
        "id": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        document.getElementById("views").innerHTML = "浏览量：" + data["num"];
    }
})


// 获取最后修改时间
$.ajax({
    url: "/bbs/GetLastModify",
    type: "POST",
    data: {
        "id": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        document.getElementById("lastmodify").innerHTML = "最后修改于：" + data["lastmodify"];
    }
})

// 获取作者
$.ajax({
    url: "/bbs/Author",
    type: "POST",
    data: {
        "id": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        document.getElementById("author").innerHTML = "作者：" + data["author"];
    }
})

// 回复文章
document.getElementById("reply").onclick = function(ev) {
    if (document.getElementById("replycomment") != undefined) {
        document.getElementById("replycomment").remove();
    }
    // 评论框
    var divs = document.createElement("div");
    divs.setAttribute("id", "replycomment");
    divs.setAttribute("style", "z-index:10; background-color: rgba(222, 222, 222, 0.8); position: fixed; bottom: 0; flex-direction: column; display: flex; width: 100%; height: 50%;");
    // 评论和展示框
    var commentdiv = document.createElement("div");
    commentdiv.setAttribute("style", "flex-direction: row; display: flex; resize: none; width: 100%; height: 85%");
    // 写评论
    var textarea = document.createElement("textarea");
    textarea.setAttribute("id", "texts");
    textarea.setAttribute("onkeyup", "mdSwitch()");
    textarea.setAttribute("maxlength", "1000");
    textarea.setAttribute("style", "display: flex; resize: none; width: 50%; height: 100%; background-color: rgba(222, 222, 222, 0.8);");
    // 展示框
    var show = document.createElement("div");
    show.setAttribute("id", "show-area");
    show.setAttribute("style", "width: 50%; height: 100%;");
    // 选项框
    var titlearea = document.createElement("div");
    titlearea.setAttribute("style", "display: flex; width: 100%; height: 15%;");
    var btn1 = document.createElement("button");
    btn1.setAttribute("id", "cancel");
    btn1.setAttribute("style", "position: absolute; left: 0; width: 60px; height: 15%")
    var btn2 = document.createElement("button");
    btn2.setAttribute("id", "confirm");
    btn2.setAttribute("style", "position: absolute; right: 0; width: 60px; height: 15%");
    btn1.innerHTML = "取消";
    btn2.innerHTML = "回复";
    titlearea.appendChild(btn1);
    titlearea.appendChild(btn2);

    commentdiv.appendChild(textarea);
    commentdiv.appendChild(show);
    divs.appendChild(titlearea);
    divs.appendChild(commentdiv);
    document.body.appendChild(divs);

    // 添加脚本
    var js1 = document.createElement("script");
    var js2 = document.createElement("script");
    js1.innerHTML = `// 取消评论
    document.getElementById("cancel").onclick = function(ev) {
        document.getElementById("replycomment").remove();
    }`;
    js2.innerHTML = `// 回复文章
    document.getElementById("confirm").onclick = function(ev) {
        $.ajax({
            url: "/bbs/AddComment",
            type: "POST",
            data: {
                "id": document.getElementsByName("main")[0].id,
                "parent": -1,
                "content": document.getElementById("texts").value
            },
            success: function(data) {
                location.reload();
            },
            fail: function() {
                alert("请先登录");
            }
        })
    }`;
    document.body.appendChild(js1);
    document.body.appendChild(js2);
}

// 回复父评论
function childreply(childid) {
    if (document.getElementById("replycomment") != undefined) {
        document.getElementById("replycomment").remove();
    }
    // 评论框
    var divs = document.createElement("div");
    divs.setAttribute("id", "replycomment");
    divs.setAttribute("style", "z-index:10; background-color: rgba(222, 222, 222, 0.8); position: fixed; bottom: 0; flex-direction: column; display: flex; width: 100%; height: 50%;");
    // 评论和展示框
    var commentdiv = document.createElement("div");
    commentdiv.setAttribute("style", "flex-direction: row; display: flex; resize: none; width: 100%; height: 85%");
    // 写评论
    var textarea = document.createElement("textarea");
    textarea.setAttribute("id", "texts");
    textarea.setAttribute("onkeyup", "mdSwitch()");
    textarea.setAttribute("maxlength", "1000");
    textarea.setAttribute("style", "display: flex; resize: none; width: 50%; height: 100%; background-color: rgba(222, 222, 222, 0.8);");
    // 展示框
    var show = document.createElement("div");
    show.setAttribute("id", "show-area");
    show.setAttribute("style", "width: 50%; height: 100%;");
    // 选项框
    var titlearea = document.createElement("div");
    titlearea.setAttribute("style", "display: flex; width: 100%; height: 15%;");
    var btn1 = document.createElement("button");
    btn1.setAttribute("id", "cancel");
    btn1.setAttribute("style", "position: absolute; left: 0; width: 60px; height: 15%")
    var btn2 = document.createElement("button");
    btn2.setAttribute("id", childid + "confirm");
    btn2.setAttribute("style", "position: absolute; right: 0; width: 60px; height: 15%");
    btn1.innerHTML = "取消";
    btn2.innerHTML = "回复";
    titlearea.appendChild(btn1);
    titlearea.appendChild(btn2);

    commentdiv.appendChild(textarea);
    commentdiv.appendChild(show);
    divs.appendChild(titlearea);
    divs.appendChild(commentdiv);
    document.body.appendChild(divs);

    // 添加脚本
    var js1 = document.createElement("script");
    var js2 = document.createElement("script");
    js1.innerHTML = `// 取消评论
    document.getElementById("cancel").onclick = function(ev) {
        document.getElementById("replycomment").remove();
    }`;
    js2.innerHTML = `// 回复文章
    document.getElementById("` + childid + `" + "confirm").onclick = function(ev) {
        $.ajax({
            url: "/bbs/AddComment",
            type: "POST",
            data: {
                "id": document.getElementsByName("main")[0].id,
                "parent": ` + childid + `,
                "content": document.getElementById("texts").value
            },
            success: function(data) {
                location.reload();
            },
            fail: function() {
                alert("请先登录");
            }
        })
    }`;
    document.body.appendChild(js1);
    document.body.appendChild(js2);
}

// 删除文章
function deletecomment(cid) {
    $.ajax({
        url: "/bbs/DeleteComment",
        type: "POST",
        data: {
            "id": cid
        },
        success: function(data) {
            location.reload();
        }
    })
}

// 添加回复脚本
function replyjs() {
    $.ajax({
        url: "/bbs/GetCommentsID",
        type: "POST",
        data: {
            "id": document.getElementsByName("main")[0].id
        },
        success: function(data) {
            for (var i = 0; i < data["ids"].length; i++) {
                // 添加回复 btn
                var btnjs = document.createElement("script");
                btnjs.innerHTML = `document.getElementById("` + data["ids"][i] + `" + "childReply").onclick = function(ev) {
                    childreply(` + data["ids"][i] + `);
                }`;
                document.body.appendChild(btnjs);
            }
        }
    })
}

// 添加删除按钮和修改按钮
function adddelete() {
    $.ajax({
        url: "/user/IsSystems",
        type: "POST",
        data: {
            "id": document.getElementsByName("main")[0].id
        },
        success: function(data) {
            if (data["msg"] == "success") {
                $.ajax({
                    url: "/bbs/GetCommentsID",
                    type: "POST",
                    data: {
                        "id": document.getElementsByName("main")[0].id
                    },
                    success: function(data) {
                        for (var i = 0; i < data["ids"].length; i++) {
                            var temp = document.getElementById(data["ids"][i] + "tempscon");
                            var att = document.createElement("input");
                            att.setAttribute("type", "checkbox");
                            att.setAttribute("id", data["ids"][i] + "checkbox");
                            att.setAttribute("class", "deletebtn");
                            att.setAttribute("style", "z-index: 5; width: 20px; height: 20px; position: absolute; right: 0; top: 0;")
                            temp.appendChild(att);

                            var btnjs = document.createElement("script");
                            btnjs.innerHTML = `document.getElementById("` + data["ids"][i] + `" + "checkbox").onclick = function(ev) {
                                deletecomment(` + data["ids"][i] + `);
                            }`;
                            document.body.appendChild(btnjs);
                        }
                        var modifytext = document.createElement("button");
                        modifytext.setAttribute("id", "modifybutton");
                        modifytext.innerHTML = "修改";
                        modifytext.setAttribute("style", "height: auto; width: 100px;")
                        document.body.appendChild(modifytext);
                        var modifysrc = document.createElement("script");
                        modifysrc.innerHTML = `document.getElementById("modifybutton").onclick = function(ev) {
                            modifyfunc();
                        }`;
                        document.body.appendChild(modifysrc);
                    }
                })
            }
        }
    })
}

// 添加修改逻辑
function modifyfunc() {
    // 修改框
    var divs = document.createElement("div");
    divs.setAttribute("id", "replycomment");
    divs.setAttribute("style", "z-index:10; background-color: rgba(222, 222, 222, 0.8); position: fixed; bottom: 0; flex-direction: column; display: flex; width: 100%; height: 100%;");
    // 修改和展示框
    var commentdiv = document.createElement("div");
    commentdiv.setAttribute("style", "flex-direction: row; display: flex; resize: none; width: 100%; height: 85%");
    // 修改
    var textarea = document.createElement("textarea");
    textarea.setAttribute("id", "texts");
    textarea.setAttribute("onkeyup", "mdSwitch()");
    textarea.setAttribute("maxlength", "500000");
    textarea.setAttribute("style", "display: flex; resize: none; width: 50%; height: 100%; background-color: rgba(222, 222, 222, 0.8);");
    // 展示框
    var show = document.createElement("div");
    show.setAttribute("id", "show-area");
    show.setAttribute("style", "width: 50%; height: 100%; word-wrap: break-word; overflow-y: auto; overflow-x: hidden;");

    // 选项框
    var titlearea = document.createElement("div");
    titlearea.setAttribute("style", "display: flex; width: 100%; height: 15%;");
    var btn1 = document.createElement("button");
    btn1.setAttribute("id", "cancel");
    btn1.setAttribute("style", "position: absolute; left: 0; width: 60px; height: 15%; background-color: rgba(255, 255, 255, 0.8);");
    var btn2 = document.createElement("button");
    btn2.setAttribute("id", "confirm");
    btn2.setAttribute("style", "position: absolute; right: 0; width: 60px; height: 15%; background-color: rgba(255, 255, 255, 0.8);");
    btn1.innerHTML = "取消";
    btn2.innerHTML = "修改";
    // 选项框中加入图片
    var backgroundpic = document.createElement("input");
    backgroundpic.setAttribute("id", "backgroundPic");
    backgroundpic.setAttribute("style", "width: 138px; left: 80px; position: absolute;");
    backgroundpic.setAttribute("type", "file");
    backgroundpic.setAttribute("accept", "image/*");
    backgroundpic.innerHTML = "背景";
    // 附加文件
    var attfile = document.createElement("input");
    attfile.setAttribute("id", "attfile");
    attfile.setAttribute("style", "width: 138px; left: 250px; position: absolute;");
    attfile.setAttribute("type", "file");
    attfile.setAttribute("multiple", "multiple");
    attfile.innerHTML = "文件";
    // 标题
    titleinput = document.createElement("input");
    titleinput.setAttribute("id", "titlesss");
    titleinput.setAttribute("placeholder", "标题");
    titleinput.setAttribute("maxlength", "100");
    titleinput.setAttribute("style", "margin-left: 500px;");
    // 分类
    typeinput = document.createElement("input");
    typeinput.setAttribute("id", "types");
    typeinput.setAttribute("placeholder", "分类");
    typeinput.setAttribute("maxlength", "100");
    // 简介
    descriptioninput = document.createElement("input");
    descriptioninput.setAttribute("id", "description");
    descriptioninput.setAttribute("placeholder", "简介");
    descriptioninput.setAttribute("maxlength", "250");
    titlearea.appendChild(btn1);
    titlearea.appendChild(backgroundpic);
    titlearea.appendChild(attfile);
    titlearea.appendChild(titleinput);
    titlearea.appendChild(typeinput);
    titlearea.appendChild(descriptioninput);
    titlearea.appendChild(btn2);

    commentdiv.appendChild(textarea);
    commentdiv.appendChild(show);
    divs.appendChild(titlearea);
    divs.appendChild(commentdiv);
    document.body.appendChild(divs);

    // 给正文赋值
    $.ajax({
            url: "/bbs/GetModifyBlog",
            type: "POST",
            data: {
                "id": document.getElementsByName("main")[0].id
            },
            success: function(data) {
                document.getElementById("titlesss").innerHTML = data["title"];
                document.getElementById("description").innerHTML = data["description"];
                document.getElementById("texts").innerHTML = data["content"];
            }
        })
        // 添加脚本
    var js1 = document.createElement("script");
    var js2 = document.createElement("script");
    js1.innerHTML = `// 取消修改
    document.getElementById("cancel").onclick = function(ev) {
        document.getElementById("replycomment").remove();
    }`;
    js2.innerHTML = `// 修改文章
    document.getElementById("confirm").onclick = function(ev) {
        var texts = document.getElementById("texts").value;
        var titles = document.getElementById("titlesss").value;
        var description = document.getElementById("description").value;
        var types = document.getElementById("types").value;
        var pic = document.getElementById('backgroundPic');
        var attFile = document.getElementById("attfile");

        var formData = new FormData();
        var formDatas = new FormData();
        if (pic.files[0] != undefined) {
            if (pic.files[0].size > 5120000) {
                alert('文件大小最大为 5 MB');
                return;
            }
            formData.append("pic", pic.files[0]);
            var pictypes = pic.files[0].type;
            var index = pictypes.lastIndexOf("/");
            formData.append("picType", pictypes.substr(index + 1));
        }

        formData.append("id", "` + document.getElementsByName("main")[0].id + `");
        formData.append("texts", texts);
        formData.append("titles", titles);
        formData.append("types", types);
        formData.append("description", description);
        formData.append("authority", 0);

        if (attFile.files[0] != undefined) {
            for (var i = 0; i < attFile.files.length; i++) {
                formData.append("attFiles", attFile.files[i]);
            }
        }
        
        $.ajax({
            url: "/bbs/ModifyBlog",
            type: "POST",
            data: formData,
            cache: false,
            processData: false,
            contentType: false,
            success: function(data) {
                location.reload();
            }
        })
    }`;
    document.body.appendChild(js1);
    document.body.appendChild(js2);
}

$.ajax({
    url: "/bbs/Getpicurl",
    type: "POST",
    data: {
        "id": document.getElementsByName("main")[0].id
    },
    success: function(data) {
        document.getElementById("root").setAttribute("style", `flex-direction: column; display: flex; width: 100%; background-image: url('` + data["picurl"] + `'); background-size:cover; background-repeat: no-repeat; background-attachment: fixed;`)
    }
})