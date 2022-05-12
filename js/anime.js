// 年份导航
function init() {
    var years = document.getElementById("years");
    var nowYear = new Date().getFullYear();
    for (var i = nowYear; i >= 2000; i--) {
        // li 标签
        var liChildNode = document.createElement("li");
        liChildNode.setAttribute("class", "nav-item");
        years.appendChild(liChildNode);
        // a 标签
        var aChildNode = document.createElement("a");
        aChildNode.setAttribute("class", "nav-link h4");
        aChildNode.setAttribute("data-bs-toggle", "pill");
        aChildNode.setAttribute("href", "#Year" + i.toString());
        aChildNode.setAttribute("id", "AYear" + i.toString());
        aChildNode.innerHTML = i.toString() + " 年";
        if (i == 2000) {
            aChildNode.innerHTML = i.toString() + " 年前";
        }
        liChildNode.appendChild(aChildNode);

        // 添加 tabcontent 框
        var tabContent = document.getElementById("tabContent");
        var DivChildNode = document.createElement("div");
        DivChildNode.setAttribute("id", "Year" + i.toString());
        DivChildNode.setAttribute("class", "container tab-pane");
        tabContent.appendChild(DivChildNode);
        // 子元素 1
        var Div_1ChildNode = document.createElement("div");
        Div_1ChildNode.setAttribute("class", "container p-3 my-3 border");
        DivChildNode.appendChild(Div_1ChildNode);
        // 子元素 2
        var Div_1_1ChildNode = document.createElement("div");
        Div_1_1ChildNode.setAttribute("class", "tab-content");
        Div_1_1ChildNode.setAttribute("id", "Year" + i.toString() + "everyPageSearch");
        Div_1ChildNode.appendChild(Div_1_1ChildNode);
        // 子元素 3
        var Div_2ChildNode = document.createElement("div");
        Div_2ChildNode.setAttribute("class", "container justify-content-center");
        Div_2ChildNode.setAttribute("id", "Year" + i.toString() + "waitToClear");
        DivChildNode.appendChild(Div_2ChildNode);
        // 添加脚本
        var ListenScript = document.createElement("script");
        ListenScript.setAttribute("id", "ListenScript" + i.toString());
        ListenScript.innerHTML = `document.getElementById("AYear` + i.toString() + `").onclick = function() {searchByYear(` + i + `);}`;
        document.body.appendChild(ListenScript);
    }
}

// 初始化
window.onload = function() {
    init();
    var everyPageSearch = document.getElementById("everyPageSearchLastestAnime");
    // 添加新番列表
    $.ajax({
        url: "/anime/searchNewAnime",
        type: "GET",
        success: function(data) {
            if (data["msg"] != undefined) {
                alert(data["msg"]);
                return;
            }
            // 每页显示数量，可以根据总量自动改变
            var onceAppear = 30;
            if (data["count"].length <= 100) {
                onceAppear = 20;
            } else if (data["count"].length <= 300) {
                onceAppear = 25;
            } else {
                onceAppear = 30;
            }

            // 页面数量
            var numPage = Math.floor(data["count"].length / onceAppear);
            var numPageRest = data["count"].length - onceAppear * numPage;
            if (numPageRest > 0) {
                numPage++;
            }

            // 添加每页的框架
            for (var i = 1; i <= numPage; i++) {
                var DivChildNode = document.createElement("div");
                if (i == 1) {
                    DivChildNode.setAttribute("class", "container tab-pane active");
                } else {
                    DivChildNode.setAttribute("class", "container tab-pane");
                }
                DivChildNode.setAttribute("id", "NewPage" + i.toString());
                everyPageSearch.appendChild(DivChildNode);
            }

            // 计算本页是否装满
            var isFull = onceAppear;
            var nowPage = 1;
            for (var i = 0; i < data["count"].length; i++) {
                // 每页框架作为父节点
                var fatherNode = document.getElementById("NewPage" + nowPage.toString());
                isFull--;
                if (isFull <= 0) {
                    nowPage++;
                    isFull = onceAppear;
                }
                // 1 号子节点
                var div1ChildNode = document.createElement("div");
                div1ChildNode.setAttribute("class", "container border p-3 my-3");
                div1ChildNode.setAttribute("style", "display: flex; flex-direction: row;");
                fatherNode.appendChild(div1ChildNode);
                // 1.1 号子节点
                var div1_1ChildNode = document.createElement("div");
                div1_1ChildNode.setAttribute("class", "container");
                div1_1ChildNode.setAttribute("style", "width: 250px;");
                div1ChildNode.appendChild(div1_1ChildNode);
                // 1.1.1 号子节点
                var div1_1_1ChildNode = document.createElement("img");
                div1_1_1ChildNode.setAttribute("src", data[data["count"][i]].Picurl);
                div1_1_1ChildNode.setAttribute("style", "max-width: 200px;");
                div1_1ChildNode.appendChild(div1_1_1ChildNode);
                // 1.2 号子节点
                var div1_2ChildNode = document.createElement("div");
                div1_2ChildNode.setAttribute("class", "container text-center");
                div1_2ChildNode.setAttribute("style", "width: 100% - 250px;");
                div1ChildNode.appendChild(div1_2ChildNode);
                // 1.2.1 号子节点
                var div1_2_1ChildNode = document.createElement("span");
                div1_2_1ChildNode.setAttribute("class", "h2 text-white");
                div1_2_1ChildNode.innerHTML = data[data["count"][i]].Name;
                div1_2ChildNode.appendChild(div1_2_1ChildNode);
                div1_2ChildNode.appendChild(document.createElement("br"));
                // 1.2.2 号子节点
                var div1_2_2ChildNode = document.createElement("span");
                div1_2_2ChildNode.setAttribute("class", "text-break text-white");
                div1_2_2ChildNode.innerHTML = data[data["count"][i]].Description;
                div1_2ChildNode.appendChild(div1_2_2ChildNode);
                // 1.2.3 号子节点
                var div1_2_3ChildNode = document.createElement("ul");
                div1_2_3ChildNode.setAttribute("class", "nav justify-content-center p-3 my-3");
                div1_2ChildNode.appendChild(div1_2_3ChildNode);
                if (data[data["count"][i]].Source == null) {
                    continue;
                }
                for (var index = 0; index < data[data["count"][i]].Source.length; index++) {
                    if (data[data["count"][i]].Source == null) {
                        continue;
                    }
                    // 1.2.3.x 号子节点
                    var tempChildNode = document.createElement("li");
                    tempChildNode.setAttribute("class", "nav-item border");
                    div1_2_3ChildNode.appendChild(tempChildNode);
                    // 1.2.3.x.x 号子节点
                    var temp_1ChildNode = document.createElement("a");
                    temp_1ChildNode.setAttribute("class", "nav-link");
                    temp_1ChildNode.setAttribute("target", "_blank");
                    temp_1ChildNode.setAttribute("href", data[data["count"][i]].Urls[index]);
                    temp_1ChildNode.innerHTML = data[data["count"][i]].Source[index];
                    tempChildNode.appendChild(temp_1ChildNode);
                }
            }

            // 添加分页节点
            var ulChildNode = document.createElement("ul");
            ulChildNode.setAttribute("class", "nav nav-pills p-3 container justify-content-center");
            ulChildNode.setAttribute("role", "tablist");
            document.getElementById("waitToClearLastestAnime").appendChild(ulChildNode);
            // 分页节点的子节点
            for (var i = 1; i <= numPage; i++) {
                var liChildNode = document.createElement("li");
                liChildNode.setAttribute("class", "nav-item");
                ulChildNode.appendChild(liChildNode);
                // a 节点
                var aChildNode = document.createElement("a");
                if (i == 1) {
                    aChildNode.setAttribute("class", "nav-link active");
                } else {
                    aChildNode.setAttribute("class", "nav-link");
                }
                aChildNode.setAttribute("data-bs-toggle", "pill");
                aChildNode.setAttribute("href", "#NewPage" + i.toString());
                aChildNode.innerHTML = "第 " + i.toString() + " 页";
                liChildNode.appendChild(aChildNode);
            }

            // 没有结果时
            if (everyPageSearch.children.length == 0) {
                everyPageSearch.innerHTML = "没有搜索到相关内容";
            }
        },
        fail: function() {}
    })
}

// 按年份检索，懒加载
function searchByYear(NowYear) {
    var everyPageSearch = document.getElementById("Year" + NowYear.toString() + "everyPageSearch");
    if (everyPageSearch.children.length != 0) {
        return;
    }
    $.ajax({
        url: "/anime/searchByYear",
        type: "POST",
        data: {
            "year": NowYear
        },
        success: function(data) {
            if (data["msg"] != undefined) {
                alert(data["msg"]);
                return;
            }
            // 每页显示数量，可以根据总量自动改变
            var onceAppear = 30;
            if (data["count"].length <= 100) {
                onceAppear = 20;
            } else if (data["count"].length <= 300) {
                onceAppear = 25;
            } else {
                onceAppear = 30;
            }
            // 页面数量
            var numPage = Math.floor(data["count"].length / onceAppear);
            var numPageRest = data["count"].length - onceAppear * numPage;
            if (numPageRest > 0) {
                numPage++;
            }

            // 添加每页的框架
            for (var i = 1; i <= numPage; i++) {
                var DivChildNode = document.createElement("div");
                if (i == 1) {
                    DivChildNode.setAttribute("class", "container tab-pane active");
                } else {
                    DivChildNode.setAttribute("class", "container tab-pane");
                }
                DivChildNode.setAttribute("id", "Page" + NowYear.toString() + i.toString());
                everyPageSearch.appendChild(DivChildNode);
            }

            // 计算本页是否装满
            var isFull = onceAppear;
            var nowPage = 1;
            for (var i = 0; i < data["count"].length; i++) {
                // 每页框架作为父节点
                var fatherNode = document.getElementById("Page" + NowYear.toString() + nowPage.toString());
                isFull--;
                if (isFull <= 0) {
                    nowPage++;
                    isFull = onceAppear;
                }
                // 1 号子节点
                var div1ChildNode = document.createElement("div");
                div1ChildNode.setAttribute("class", "container border p-3 my-3");
                div1ChildNode.setAttribute("style", "display: flex; flex-direction: row;");
                fatherNode.appendChild(div1ChildNode);
                // 1.1 号子节点
                var div1_1ChildNode = document.createElement("div");
                div1_1ChildNode.setAttribute("class", "container");
                div1_1ChildNode.setAttribute("style", "width: 250px;");
                div1ChildNode.appendChild(div1_1ChildNode);
                // 1.1.1 号子节点
                var div1_1_1ChildNode = document.createElement("img");
                div1_1_1ChildNode.setAttribute("src", data[data["count"][i]].Picurl);
                div1_1_1ChildNode.setAttribute("style", "max-width: 200px;");
                div1_1ChildNode.appendChild(div1_1_1ChildNode);
                // 1.2 号子节点
                var div1_2ChildNode = document.createElement("div");
                div1_2ChildNode.setAttribute("class", "container text-center");
                div1_2ChildNode.setAttribute("style", "width: 100% - 250px;");
                div1ChildNode.appendChild(div1_2ChildNode);
                // 1.2.1 号子节点
                var div1_2_1ChildNode = document.createElement("span");
                div1_2_1ChildNode.setAttribute("class", "h2 text-white");
                div1_2_1ChildNode.innerHTML = data[data["count"][i]].Name;
                div1_2ChildNode.appendChild(div1_2_1ChildNode);
                div1_2ChildNode.appendChild(document.createElement("br"));
                // 1.2.2 号子节点
                var div1_2_2ChildNode = document.createElement("span");
                div1_2_2ChildNode.setAttribute("class", "text-break text-white");
                div1_2_2ChildNode.innerHTML = data[data["count"][i]].Description;
                div1_2ChildNode.appendChild(div1_2_2ChildNode);
                // 1.2.3 号子节点
                var div1_2_3ChildNode = document.createElement("ul");
                div1_2_3ChildNode.setAttribute("class", "nav justify-content-center p-3 my-3");
                div1_2ChildNode.appendChild(div1_2_3ChildNode);
                if (data[data["count"][i]].Source == null) {
                    continue;
                }
                for (var index = 0; index < data[data["count"][i]].Source.length; index++) {
                    if (data[data["count"][i]].Source == null) {
                        continue;
                    }
                    // 1.2.3.x 号子节点
                    var tempChildNode = document.createElement("li");
                    tempChildNode.setAttribute("class", "nav-item border");
                    div1_2_3ChildNode.appendChild(tempChildNode);
                    // 1.2.3.x.x 号子节点
                    var temp_1ChildNode = document.createElement("a");
                    temp_1ChildNode.setAttribute("class", "nav-link");
                    temp_1ChildNode.setAttribute("target", "_blank");
                    temp_1ChildNode.setAttribute("href", data[data["count"][i]].Urls[index]);
                    temp_1ChildNode.innerHTML = data[data["count"][i]].Source[index];
                    tempChildNode.appendChild(temp_1ChildNode);
                }
            }

            // 添加分页节点
            var ulChildNode = document.createElement("ul");
            ulChildNode.setAttribute("class", "nav nav-pills p-3 container justify-content-center");
            ulChildNode.setAttribute("role", "tablist");
            document.getElementById("Year" + NowYear.toString() + "waitToClear").appendChild(ulChildNode);
            // 分页节点的子节点
            for (var i = 1; i <= numPage; i++) {
                var liChildNode = document.createElement("li");
                liChildNode.setAttribute("class", "nav-item");
                ulChildNode.appendChild(liChildNode);
                // a 节点
                var aChildNode = document.createElement("a");
                if (i == 1) {
                    aChildNode.setAttribute("class", "nav-link active");
                } else {
                    aChildNode.setAttribute("class", "nav-link");
                }
                aChildNode.setAttribute("data-bs-toggle", "pill");
                aChildNode.setAttribute("href", "#Page" + NowYear.toString() + i.toString());
                aChildNode.innerHTML = "第 " + i.toString() + " 页";
                liChildNode.appendChild(aChildNode);
            }
            // 没有结果时
            if (everyPageSearch.children.length == 0) {
                everyPageSearch.innerHTML = "没有搜索到相关内容";
            }
        },
        fail: function() {}
    })
}

// 搜索功能
function searchs() {
    var text = document.getElementById("keyword");
    var year = document.getElementById("years");
    var tabContent = document.getElementById("tabContent");
    var searchResult = document.getElementById("searchResult");
    var everyPageSearch = document.getElementById("everyPageSearch");

    // 清除原内容
    everyPageSearch.innerHTML = "";
    $.ajax({
        url: "/anime/search",
        type: "POST",
        data: {
            "text": text.value
        },
        success: function(data) {
            if (data["msg"] != undefined) {
                alert(data["msg"]);
                return;
            }
            // 每页显示数量，可以根据总量自动改变
            var onceAppear = 30;
            if (data["count"].length <= 100) {
                onceAppear = 20;
            } else if (data["count"].length <= 300) {
                onceAppear = 25;
            } else {
                onceAppear = 30;
            }

            // 所有页面去除 active
            for (var index = 0; index < year.children.length; index++) {
                year.children[index].children[0].setAttribute("class", "nav-link h4");
                tabContent.children[index].setAttribute("class", "container tab-pane");
            }
            // 使搜索结果页面 active
            document.getElementById("searchNav").setAttribute("class", "nav-link h4 active");
            document.getElementById("searchResult").setAttribute("class", "container tab-pane active");

            // 页面数量
            var numPage = Math.floor(data["count"].length / onceAppear);
            var numPageRest = data["count"].length - onceAppear * numPage;
            if (numPageRest > 0) {
                numPage++;
            }

            // 添加每页的框架
            for (var i = 1; i <= numPage; i++) {
                var DivChildNode = document.createElement("div");
                if (i == 1) {
                    DivChildNode.setAttribute("class", "container tab-pane active");
                } else {
                    DivChildNode.setAttribute("class", "container tab-pane");
                }
                DivChildNode.setAttribute("id", "Page" + i.toString());
                everyPageSearch.appendChild(DivChildNode);
            }

            // 计算本页是否装满
            var isFull = onceAppear;
            var nowPage = 1;
            for (var i = 0; i < data["count"].length; i++) {
                // 每页框架作为父节点
                var fatherNode = document.getElementById("Page" + nowPage.toString());
                isFull--;
                if (isFull <= 0) {
                    nowPage++;
                    isFull = onceAppear;
                }
                // 1 号子节点
                var div1ChildNode = document.createElement("div");
                div1ChildNode.setAttribute("class", "container border p-3 my-3");
                div1ChildNode.setAttribute("style", "display: flex; flex-direction: row;");
                fatherNode.appendChild(div1ChildNode);
                // 1.1 号子节点
                var div1_1ChildNode = document.createElement("div");
                div1_1ChildNode.setAttribute("class", "container");
                div1_1ChildNode.setAttribute("style", "width: 250px;");
                div1ChildNode.appendChild(div1_1ChildNode);
                // 1.1.1 号子节点
                var div1_1_1ChildNode = document.createElement("img");
                div1_1_1ChildNode.setAttribute("src", data[data["count"][i]].Picurl);
                div1_1_1ChildNode.setAttribute("style", "max-width: 200px;");
                div1_1ChildNode.appendChild(div1_1_1ChildNode);
                // 1.2 号子节点
                var div1_2ChildNode = document.createElement("div");
                div1_2ChildNode.setAttribute("class", "container text-center");
                div1_2ChildNode.setAttribute("style", "width: 100% - 250px;");
                div1ChildNode.appendChild(div1_2ChildNode);
                // 1.2.1 号子节点
                var div1_2_1ChildNode = document.createElement("span");
                div1_2_1ChildNode.setAttribute("class", "h2 text-white");
                div1_2_1ChildNode.innerHTML = data[data["count"][i]].Name;
                div1_2ChildNode.appendChild(div1_2_1ChildNode);
                div1_2ChildNode.appendChild(document.createElement("br"));
                // 1.2.2 号子节点
                var div1_2_2ChildNode = document.createElement("span");
                div1_2_2ChildNode.setAttribute("class", "text-break text-white");
                div1_2_2ChildNode.innerHTML = data[data["count"][i]].Description;
                div1_2ChildNode.appendChild(div1_2_2ChildNode);
                // 1.2.3 号子节点
                var div1_2_3ChildNode = document.createElement("ul");
                div1_2_3ChildNode.setAttribute("class", "nav justify-content-center p-3 my-3");
                div1_2ChildNode.appendChild(div1_2_3ChildNode);
                if (data[data["count"][i]].Source == null) {
                    continue;
                }
                for (var index = 0; index < data[data["count"][i]].Source.length; index++) {
                    if (data[data["count"][i]].Source == null) {
                        continue;
                    }
                    // 1.2.3.x 号子节点
                    var tempChildNode = document.createElement("li");
                    tempChildNode.setAttribute("class", "nav-item border");
                    div1_2_3ChildNode.appendChild(tempChildNode);
                    // 1.2.3.x.x 号子节点
                    var temp_1ChildNode = document.createElement("a");
                    temp_1ChildNode.setAttribute("class", "nav-link");
                    temp_1ChildNode.setAttribute("target", "_blank");
                    temp_1ChildNode.setAttribute("href", data[data["count"][i]].Urls[index]);
                    temp_1ChildNode.innerHTML = data[data["count"][i]].Source[index];
                    tempChildNode.appendChild(temp_1ChildNode);
                }
            }

            // 添加分页节点
            document.getElementById("waitToClear").innerHTML = "";
            var ulChildNode = document.createElement("ul");
            ulChildNode.setAttribute("class", "nav nav-pills p-3 container justify-content-center");
            ulChildNode.setAttribute("role", "tablist");
            document.getElementById("waitToClear").appendChild(ulChildNode);
            // 分页节点的子节点
            for (var i = 1; i <= numPage; i++) {
                var liChildNode = document.createElement("li");
                liChildNode.setAttribute("class", "nav-item");
                ulChildNode.appendChild(liChildNode);
                // a 节点
                var aChildNode = document.createElement("a");
                if (i == 1) {
                    aChildNode.setAttribute("class", "nav-link active");
                } else {
                    aChildNode.setAttribute("class", "nav-link");
                }
                aChildNode.setAttribute("data-bs-toggle", "pill");
                aChildNode.setAttribute("href", "#Page" + i.toString());
                aChildNode.innerHTML = "第 " + i.toString() + " 页";
                liChildNode.appendChild(aChildNode);
            }

            // 没有结果时
            if (everyPageSearch.children.length == 0) {
                everyPageSearch.innerHTML = "没有搜索到相关内容";
            }
        },
        fail: function() {}
    })
}

// 点击搜索时的操作
document.getElementById("confirm").onclick = searchs;

// 按下确认键搜索
document.getElementById("keyword").addEventListener("keydown", function(event) {
    if (event.key == "Enter") {
        searchs();
    }
})