function installControllers() {

    document.getElementById("menu").innerHTML = `
    <input type="button" id="PathCreator" value="Create Path" style="visibility:hidden">
    <input type="button" id="FreePathCreator" value="Create FreePath" style="visibility:hidden">
    <input type="button" id="LineCreator" value="Create Line" style="visibility:hidden">
    <input type="button" id="RectCreator" value="Create Rect" style="visibility:hidden">
    <input type="button" id="EllipseCreator" value="Create Ellipse" style="visibility:hidden">
    <input type="button" id="CircleCreator" value="Create Circle" style="visibility:hidden">
    `;

    for (let gkey in qview.controllers) {
        let key = gkey;
        let elem = document.getElementById(key);
        elem.style.visibility = "visible";
        // 为每个图形控制器元素注册点击事件，提供给用户使用
        elem.onclick = function () {
            if (qview.currentKey != "") {
                document.getElementById(qview.currentKey).removeAttribute("style");
            }
            elem.style.borderColor = "blue";
            elem.blur();
            // 执行该控制器，作用是设置当前激活的 Controller，并创建新的 Controller
            qview.invokeController(key);
        }
    }
}

function onLineWidthChanged() {
    let elem = document.getElementById("LineWidth");
    elem.blur();
    let val = parseInt(elem.value);
    if (val > 0) {
        qview.properties.lineWidth = val;
    }
}

function onLineColorChanged() {
    let elem = document.getElementById("LineColor");
    elem.blur();
    qview.properties.lineColor = elem.value;
}

function installPropSelectors() {
    document.getElementById("menu").insertAdjacentHTML("afterend", `<br><div id="properties">
    <label for="LineWidth">LineWidth: </label>
    <select id="LineWidth" onchange="onLineWidthChanged()">
        <option value="1">1</option>
        <option value="3">3</option>
        <option value="5">5</option>
        <option value="7">7</option>
        <option value="9">9</option>
        <option value="11">11</option>
    </select>&nbsp;
    <label for="LineColor">LineColor: </label>
    <select id="LineColor" onchange="onLineColorChanged()">
        <option value="black">black</option>
        <option value="red">red</option>
        <option value="blue">blue</option>
        <option value="green">green</option>
        <option value="yellow">yellow</option>
        <option value="gray">gray</option>
    </select>
    </div>`);
}

function installMousePos() {
    document.getElementById("properties")
        .insertAdjacentHTML("beforeend", `&nbsp;<span id="mousepos"></span>`);

    let old = qview.drawing.onmousemove;
    let mousepos = document.getElementById("mousepos");
    qview.drawing.onmousemove = function (event) {
        let pos = qview.getMousePos(event);
        mousepos.innerText = "MousePos: " + pos.x + ", " + pos.y;
        old(event);
    }
}

// 初始化所有的 Controller
installControllers();
// 初始化全局属性设置选择器
installPropSelectors();
// 跟踪鼠标位置并显示
installMousePos();