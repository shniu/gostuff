class LineStyle {
    constructor(width, color) {
        this.width = width;
        this.color = color;
    }
}

// 线段模型，图形的一个子类，实现了 Shape 接口
class Line {
    constructor(point1, point2, lineStyle) {
        this.p1 = point1;
        this.p2 = point2;
        this.lineStyle = lineStyle;
    }

    onpaint(ctx) {
        let lineStyle = this.lineStyle;
        ctx.lineWidth = lineStyle.width;
        ctx.strokeStyle = lineStyle.color;
        ctx.beginPath();
        ctx.moveTo(this.p1.x, this.p1.y);
        ctx.lineTo(this.p2.x, this.p2.y);
        ctx.stroke();
    }
}

// 矩形模型
class Rect {
    constructor(r, lineStyle) {
        this.x = r.x;
        this.y = r.y;
        this.width = r.width;
        this.height = r.height;
        this.lineStyle = lineStyle;
    }

    onpaint(ctx) {
        let lineStyle = this.lineStyle;
        ctx.lineWidth = lineStyle.width;
        ctx.strokeStyle = lineStyle.color;

        ctx.beginPath();
        ctx.rect(this.x, this.y, this.width, this.height);
        ctx.stroke();
    }
}

// 椭圆模型
class Ellipse {
    constructor(x, y, radiusX, radiusY, lineStyle) {
        this.x = x;
        this.y = y;
        this.radiusX = radiusX;
        this.radiusY = radiusY;
        this.lineStyle = lineStyle;
    }

    onpaint(ctx) {
        let lineStyle = this.lineStyle;
        ctx.lineWidth = lineStyle.width;
        ctx.strokeStyle = lineStyle.color;
        ctx.beginPath();
        ctx.ellipse(this.x, this.y, this.radiusX, this.radiusY, 0, 0, 2 * Math.PI);
        ctx.stroke();
    }
}

// 路径
class Path {
    constructor(points, close, lineStyle) {
        this.points = points;
        this.close = close;
        this.lineStyle = lineStyle;
    }

    onpaint(ctx) {
        let n = this.points.length;
        if (n < 1) {
            return;
        }

        let points = this.points;
        let lineStyle = this.lineStyle;
        ctx.lineWidth = lineStyle.width
        ctx.strokeStyle = lineStyle.color
        ctx.beginPath();
        ctx.moveTo(points[0].x, points[0].y);
        for (let i = 1; i < n; i++) {
            ctx.lineTo(points[i].x, points[i].y);
        }
        if (this.close) {
            ctx.closePath()
        }
        ctx.stroke();
    }
}

// 画图
class PaintDoc {
    constructor() {
        // 跟踪画布上所有已经添加的图形
        this.shapes = [];
    }

    addShape(shape) {
        if (shape != null) {
            this.shapes.push(shape);
        }
    }

    // 绘制
    onpaint(ctx) {
        // 实际的画图逻辑委托给具体的图形实现，只有每个具体的实现自身才直到该如何画图
        let shapes = this.shapes;
        for (let s in shapes) {
            shapes[s].onpaint(ctx);
        }
    }
}