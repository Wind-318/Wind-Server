function init() {
    var years = document.getElementById("years");
    $.ajax({
        url: "/anime/getYear",
        type: "GET",
        success: function(datas) {
            for (var i = datas["year"]; i >= 1999; i--) {
                var liLable = document.createElement("li");
                liLable.setAttribute("class", "nav-item");
                var aLable = document.createElement("a");
                if (i == 0) {
                    aLable.setAttribute("class", "nav-link active");
                } else {
                    aLable.setAttribute("class", "nav-link");
                }
                aLable.setAttribute("data-bs-toggle", "pill");
            }
        },
        fail: function() {}
    })
}

window.onload = function() {

}

function searchs() {
    var text = document.getElementById("keyword");
    var year = document.getElementById("year");
    var tabContent = document.getElementById("tabContent");
    $.ajax({
        url: "/anime/search",
        type: "POST",
        data: {
            "text": text.value
        },
        success: function(data) {
            for (var i = 0; i < data["count"]; i++) {
                // 清除所有 active 状态
                for (var index = 0; index < year.childNodes.length; index++) {
                    year.childNodes[index].setAttribute("class", "nav-link h4");
                    tabContent.childNodes[index].setAttribute("class", "container tab-pane");
                }
                // 使搜索结果页面 active
                document.getElementById("searchNav").setAttribute("class", "nav-link h4 active");
                document.getElementById("searchResult").setAttribute("class", "container tab-pane active");
            }

            for (var i = 0; i < data["count"]; i++) {
                // 获取搜索结果
                var searchResult = document.getElementById("searchResult");
                // 1 号子节点
                var div1ChildNode = document.createElement("div");
                div1ChildNode.setAttribute("class", "container border p-3 my-3");
                div1ChildNode.setAttribute("style", "display: flex; flex-direction: row;");
                searchResult.appendChild(div1ChildNode);
                // 1.1 号子节点
                var div1_1ChildNode = document.createElement("div");
                div1_1ChildNode.setAttribute("class", "container");
                div1_1ChildNode.setAttribute("style", "width: 250px;");
                div1ChildNode.appendChild(div1_1ChildNode);
                // 1.1.1 号子节点
                var div1_1_1ChildNode = document.createElement("img");
                div1_1_1ChildNode.setAttribute("style", "width: 150px; height: 180px;");
                div1_1_1ChildNode.setAttribute("src", data[i.toString()].Picurl);
                div1_1ChildNode.appendChild(div1_1_1ChildNode);
                // 1.2 号子节点
                var div1_2ChildNode = document.createElement("div");
                div1_2ChildNode.setAttribute("class", "container text-center");
                div1_2ChildNode.setAttribute("style", "width: 100% - 250px;");
                div1ChildNode.appendChild(div1_2ChildNode);
                // 1.2.1 号子节点
                var div1_2_1ChildNode = document.createElement("span");
                div1_2_1ChildNode.setAttribute("class", "h2 text-white");
                div1_2_1ChildNode.innerHTML = data[i.toString()].Name;
                div1_2ChildNode.appendChild(div1_2_1ChildNode);
                div1_2ChildNode.appendChild(document.createElement("br"));
                // 1.2.2 号子节点
                var div1_2_2ChildNode = document.createElement("span");
                div1_2_2ChildNode.setAttribute("class", "text-break text-white");
                div1_2_2ChildNode.innerHTML = data[i.toString()].Description;
                div1_2ChildNode.appendChild(div1_2_2ChildNode);
                // 1.2.3 号子节点
                var div1_2_3ChildNode = document.createElement("ul");
                div1_2_3ChildNode.setAttribute("class", "nav justify-content-center p-3 my-3");
                div1_2ChildNode.appendChild(div1_2_3ChildNode);
                // 1.2.3.1 号子节点
                var div1_2_3_1ChildNode = document.createElement("li");
                div1_2_3_1ChildNode.setAttribute("class", "nav-item border");
                div1_2_3.appendChild(div1_2_3_1ChildNode);
                // 1.2.3.1.1 号子节点
                var div_1_2_3_1_1ChildNode = document.createElement("a");
                div_1_2_3_1_1ChildNode.setAttribute("class", "nav-link");
                div_1_2_3_1_1ChildNode.setAttribute("target", "_blank");
                div_1_2_3_1_1ChildNode.setAttribute("href", data[i.toString()].Url);
                div_1_2_3_1_1ChildNode.innerHTML = "bangumi";
                div1_2_3_1ChildNode.appendChild(div_1_2_3_1_1ChildNode);
                for (var index = 0; index < data[i.toString()].Source.length; index++) {
                    // 1.2.3.x 号子节点
                    var tempChildNode = document.createElement("li");
                    tempChildNode.setAttribute("class", "nav-item border");
                    div1_2_3.appendChild(tempChildNode);
                    // 1.2.3.x.x 号子节点
                    var temp_1ChildNode = document.createElement("a");
                    temp_1ChildNode.setAttribute("class", "nav-link");
                    temp_1ChildNode.setAttribute("target", "_blank");
                    temp_1ChildNode.setAttribute("href", data[i.toString()].Urls[index]);
                    temp_1ChildNode.innerHTML = data[i.toString()].Source[index];
                    tempChildNode.appendChild(temp_1ChildNode);
                }
            }
        },
        fail: function() {}
    })
}

document.getElementById("confirm").onclick = searchs;

document.getElementById("keyword").addEventListener("keydown", function(event) {
    if (event.key == "Enter") {
        searchs();
    }
})