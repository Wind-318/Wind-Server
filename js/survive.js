document.getElementById("start").onclick = function(ev) {
    document.getElementById("options").innerHTML = "";
    var forward = document.createElement("li");
    forward.innerHTML = '前进';
    forward.setAttribute('id', "forward");
    document.getElementById("options").appendChild(forward);

    $.ajax({
        url: "/survive/getproperty",
        type: "GET",
        success: function(data) {
            var display = document.getElementById("display");
            display.innerHTML = "姓名：" + data["name"] + "<br>" + "性别：" + data["gender"] + "<br>" + "HP：" + data["hp"] + "<br>" + "饥饿度：" + data["hunger"] + "<br>" + "渴度：" + data["thirst"] + "<br>" + "温度：" + data["temperature"] + "<br>" + "体力：" + data["stamina"] + "<br>" + "睡眠值：" + data["sleep"] + "<br>" + "效果：" + data["effect"] + "<br>" + "状态：" + data["status"];
        }
    })
}

document.getElementById("forward").onclick = function(ev) {
    $.ajax({
        url: "/survive/add",
        type: "GET",
        success: function(data) {

        }
    })
}

document.getElementById("goon").onclick = function(ev) {

}