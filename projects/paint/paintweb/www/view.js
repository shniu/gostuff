class PaintView {
    constructor() {
        // view 自身的属性，属于 view 的全局设置，由用户触发控制
        this.properties = {
            lineWidth: 1,
            lineColor: "black"
        };

        // 跟踪当前激活的图形是哪一个，属于 view 的全局设置，由用户触发控制
        this._currentKey = "";
        this._current = null;

        // 维护每个图形对应的 controller，view 通过 controller 的名称来获取其实例
        // 在页面初始化时注册每一个图形对应的 controller 处理程序
        this.controllers = {};

        // view 能响应的事件列表，主要用来跟踪 Canvas 画布的对应事件，然后委托给 Model 和 Controller 进行处理
        this.onmousedown = null;
        this.onmousemove = null;
        this.onmouseup = null;
        this.ondblclick = null;
        this.onkeydown = null;

        // 获取 Canvas 的 DOMElement
        let drawing = document.getElementById("drawing");
        let view = this;

        // 在 Canvas 的 DOM 上注册事件，用来响应用户操作
        drawing.onmousedown = function (event) {
            event.preventDefault();
            if (view.onmousedown != null) {
                view.onmousedown(event);
            }
        };
        drawing.onmousemove = function(event) {
            if (view.onmousemove != null) {
                view.onmousemove(event)
            }
        }
        drawing.onmouseup = function(event) {
            if (view.onmouseup != null) {
                view.onmouseup(event)
            }
        }
        drawing.ondblclick = function(event) {
            event.preventDefault()
            if (view.ondblclick != null) {
                view.ondblclick(event)
            }
        }

        // 在 document 上注册 onkeydown
        document.onkeydown = function (event) {
            switch (event.keyCode) {
                case 9:
                case 13:
                case 27:
                    event.preventDefault();
            }

            if (view.onkeydown != null) {
                view.onkeydown(event);
            }
        };

        // 其实是 Canvas 的代理对象，是访问 Canvas API 的入口
        this.drawing = drawing;

        // 实际画图的 Model，最终由 doc 来完成图形的绘制，借助 Canvas 的 API
        this.doc = new PaintDoc();
    }

    // 获取当前的 controller 注册名称
    get currentKey() {
        return this._currentKey;
    }

    // 返回 line 的样式设置
    get lineStyle() {
        let props = this.properties;
        return new LineStyle(props.lineWidth, props.lineColor);
    }

    // 绘制图形
    onpaint(ctx) {
        // 委托给 doc 进行实际的绘制，重绘之前所有的元素
        this.doc.onpaint(ctx);
        if (this._current != null) {
            // 绘制当前正在画的元素
            this._current.onpaint(ctx);
        }
    }

    // 返回鼠标的当前位置
    getMousePos(event) {
        return {
            x: event.offsetX,
            y: event.offsetY
        }
    }

    invalidateRect(reserved) {
        // ctx 是 HTMLCanvasElement 的绘制上下文
        // 2d 是一个二维的渲染上下文，是 CanvasRenderingContext2D 的实例
        let ctx = this.drawing.getContext("2d");
        let bound = this.drawing.getBoundingClientRect();
        // 擦除之前画的所有元素
        ctx.clearRect(0, 0, bound.width, bound.height);
        // 重绘所有元素
        this.onpaint(ctx);
    }

    // 注册 Controller
    registerController(name, controller) {
        if (name in this.controllers) {
            alert("Controller exists: " + name);
        } else {
            this.controllers[name] = controller;
        }
    }

    // 执行某个 Controller
    invokeController(name) {
        // 停止上一个 Controller，清空设置
        this.stopController();
        if (name in this.controllers) {
            // 找到当前要操作的 Controller，并设置
            let controller = this.controllers[name];
            // this.controllers[name] 其实是一个函数对象，调用后产生一个 Controller 实例
            // 因为绘制的每个图形都是独一无二的
            this._setCurrent(name, controller());
        }
    }

    // 停止一个 Controller
    stopController() {
        if (this._current != null) {
            this._current.stop()
            this._setCurrent("", null)
        }
    }

    // 设置当前的 Controller 对象
    _setCurrent(name, ctrl) {
        this._current = ctrl
        this._currentKey = name
    }
}

var qview = new PaintView();

function invalidate(reserved) {
    qview.invalidateRect(null);
}
