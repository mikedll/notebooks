(() => {
  // aristochart/Aristochart.js
  var Aristochart = function(element, options, theme) {
    if (!element || !element.DOCUMENT_NODE)
      options = element, element = document.createElement("canvas");
    if (!options || !options.data)
      throw new Error("Please provide some data to plot.");
    if (!options.data.y || !options.data.x)
      throw new Error("Please provide some data.x and data.y");
    if (options.width && !options.height)
      options.height = Math.floor(options.width * 0.67);
    this.defaults = Aristochart.themes.default;
    this.options = options;
    this.canvas = element;
    this.theme = theme;
    this.data = this.options.data;
    if (this.theme)
      this.defaults = Aristochart._deepMerge(this.defaults, this.theme);
    for (var key in this.defaults)
      this.options = Aristochart._deepMerge(this.defaults, this.options);
    for (var style in this.options.style)
      for (var key in this.options.style["default"])
        this.options.style[style] = Aristochart._deepMerge(this.options.style["default"], this.options.style[style]);
    this.indexes = [], that = this;
    ["fill", "axis", "tick", "line", "point", "label", "title"].forEach(function(feature) {
      if (that.indexes[that.options[feature].index])
        throw new Error("Conflicting indexes in Aristochart");
      else
        that.indexes[that.options[feature].index] = feature;
    });
    this.indexes = this.indexes.filter(function(val) {
      if (val)
        return true;
    });
    if (this.canvas.getContext)
      this.ctx = this.canvas.getContext("2d");
    else {
      var canvas = document.createElement("canvas");
      this.canvas.appendChild(canvas);
      this.canvas = canvas;
      this.ctx = canvas.getContext("2d");
    }
    this.canvas.height = this.options.height;
    this.canvas.width = this.options.width;
    if (window.devicePixelRatio > 1) {
      this.canvas.style.height = this.canvas.height + "px";
      this.canvas.style.width = this.canvas.width + "px";
      this.canvas.height = this.canvas.height * window.devicePixelRatio;
      this.canvas.width = this.canvas.width * window.devicePixelRatio;
    }
    this.resolution = window.devicePixelRatio || 1;
    this.update();
    if (this.options.render)
      this.render();
  };
  Aristochart._deepMerge = function(defaults, options) {
    return function recur(defaults2, options2) {
      for (var key in defaults2) {
        if (options2[key] == void 0)
          options2[key] = defaults2[key];
        else if (defaults2[key] instanceof Object)
          options2[key] = recur(defaults2[key], options2[key]);
      }
      return options2;
    }(defaults, options);
  };
  Aristochart.prototype.refreshBounds = function() {
    var yMax = -Infinity;
    var yMin = Infinity;
    for (var key in this.data) {
      if (key !== "x") {
        var max = -Infinity, min = Infinity;
        this.data[key].forEach(function(v) {
          if (v > max)
            max = v;
          if (v < min)
            min = v;
        });
        yMax = max > yMax ? max : yMax;
        yMin = min < yMin ? min : yMin;
      }
    }
    this.y = {
      //Check if manually overrided
      max: this.options.axis.y.max == void 0 ? yMax : this.options.axis.y.max,
      min: this.options.axis.y.min == void 0 ? yMin : this.options.axis.y.min
    };
    this.y.range = this.y.max - this.y.min;
    if (this.data.x.length == 1 || typeof this.data.x == "number")
      this.x = { min: 0, max: this.data.x[0] || this.data.x };
    else
      this.x = { min: this.data.x[0], max: this.data.x[this.data.x.length - 1] };
    this.x.range = this.x.max - this.x.min;
  };
  Aristochart.prototype.update = function() {
    var resolution = this.resolution;
    this.options.margin *= resolution;
    this.options.padding *= resolution;
    this.options.width *= resolution;
    this.options.height *= resolution;
    this.box = {
      x: this.options.margin,
      y: this.options.margin,
      x1: this.options.width - 2 * this.options.margin,
      y1: this.options.height - 2 * this.options.margin
    };
    this.refreshBounds();
    var data = this.getPoints();
    this.lines = data.lines;
    this.origin = data.origin;
    var padding = this.options.padding, box = this.box;
    this.axis = {
      x: {
        x: box.x - padding,
        y: box.y + box.y1 + padding,
        x1: that.box.x + box.x1 + padding,
        y1: box.y + box.y1 + padding
      },
      y: {
        x: box.x - padding,
        y: box.y - padding,
        x1: box.x - padding,
        y1: box.y + box.y1 + padding
      }
    };
  };
  Aristochart.prototype.render = function() {
    var that2 = this, lines = this.lines, origin = this.origin, axis = this.axis, defaults = that2.options.style.default;
    this.canvas.width = this.canvas.width;
    var stepX = Math.floor(this.options.axis.x.steps), stepY = Math.floor(this.options.axis.y.steps);
    var padding = this.options.padding, box = this.box, ox = origin.x, oy = origin.y;
    this.indexes.forEach(function(feature) {
      switch (feature) {
        case "point":
          for (var line in lines)
            if ((that2.options.style[line] || defaults).point.visible)
              lines[line].forEach(function(obj) {
                that2.options.point.render.call(that2, that2.options.style[line] || defaults, obj.rx, obj.ry, obj.x, obj.y, obj.graph);
              });
          break;
        case "axis":
          if (defaults.axis.visible) {
            if (defaults.axis.x.visible) {
              that2.options.axis.x.render.call(that2, defaults, axis.x.x, defaults.axis.y.fixed ? axis.x.y : oy, axis.x.x1, defaults.axis.y.fixed ? axis.x.y1 : oy, "x");
            }
            if (defaults.axis.y.visible) {
              that2.options.axis.y.render.call(that2, defaults, defaults.axis.x.fixed ? axis.y.x : ox, axis.y.y, defaults.axis.x.fixed ? axis.y.x1 : ox, axis.y.y1, "y");
            }
          }
          break;
        case "line":
          for (var line in lines) {
            var style = that2.options.style[line] || defaults;
            if (style.line.visible)
              that2.options.line.render.call(that2, style, lines[line]);
          }
          break;
        case "tick":
          if (defaults.tick.visible) {
            var disX = that2.box.x1 / stepX, disY = that2.box.y1 / stepY;
            for (var i = 0; i < stepX + 1; i++)
              that2.options.tick.render.call(that2, defaults, that2.box.x + disX * i, defaults.tick.x.fixed ? axis.x.y1 : oy, "x", i);
            for (var i = 0; i < stepY + 1; i++)
              that2.options.tick.render.call(that2, defaults, defaults.tick.y.fixed ? axis.y.x1 : ox, that2.box.y + disY * i, "y", i);
          }
          break;
        case "label":
          var disX = that2.box.x1 / stepX, disY = that2.box.y1 / stepY;
          if (defaults.label.x.visible)
            for (var i = 0; i < stepX + 1; i++)
              that2.options.label.render.call(that2, defaults, that2.x.min + (that2.x.max - that2.x.min) / stepX * i, that2.box.x + disX * i, defaults.label.x.fixed ? axis.x.y1 : oy, "x", i);
          if (defaults.label.y.visible)
            for (var i = 0; i < stepY + 1; i++) {
              var pos = stepY - i, label = that2.y.min + (that2.y.max - that2.y.min) / stepY * pos;
              that2.options.label.render.call(that2, defaults, label, defaults.label.y.fixed ? axis.y.x1 : ox, that2.box.y + disY * i, "y", i);
            }
          break;
        case "fill":
          for (var line in lines) {
            var style = that2.options.style[line] || defaults;
            if (style.line.fill)
              that2.options.fill.render.call(that2, style, lines[line]);
          }
          break;
        case "title":
          if (defaults.title.visible) {
            var xLabel = that2.options.title.x, yLabel = that2.options.title.y;
            if (defaults.title.x.visible)
              that2.options.title.render.call(that2, defaults, xLabel, (that2.box.x * 2 + that2.box.x1) / 2, that2.box.y + that2.box.y1, "x");
            if (defaults.title.y.visible)
              that2.options.title.render.call(that2, defaults, yLabel, that2.box.x, (that2.box.y * 2 + that2.box.y1) / 2, "y");
          }
          break;
      }
    });
  };
  Aristochart.prototype.getPoints = function(callback) {
    var lines = {}, Xmax = this.x.max, Xmin = this.x.min, Xrange = this.x.range, Ymax = this.y.max, Ymin = this.y.min, Yrange = this.y.range, bx = this.box.x, by = this.box.y, bx1 = this.box.x1, by1 = this.box.y1, Yorigin = by + by1 / Yrange * Ymax, Xorigin = bx + bx1 / Xrange * Math.abs(Xmin);
    for (var key in this.data) {
      if (key == "x")
        continue;
      lines[key] = [];
      var currArr = this.data[key], length = currArr.length, factor = 1;
      if (length > 1e3)
        factor = 5;
      if (length > 1e4)
        factor = 50;
      if (length > 1e5)
        factor = 5e3;
      var count = length / factor;
      for (var i = 0; i < count; i++) {
        var x = Xrange / (count - 1) * i + Xmin, y = currArr[i], rx = Xorigin + bx1 / Xrange * x, ry = Yorigin - by1 / Yrange * y;
        lines[key].push({ x, y, rx, ry });
        if (callback)
          callback(rx, ry, x, y, key);
      }
    }
    return {
      lines,
      origin: {
        x: Xorigin,
        y: Yorigin
      }
    };
  };
  Aristochart.prototype.toImage = function() {
    var img = new Image();
    img.src = this.canvas.toDataURL("image/png");
    return img;
  };
  Aristochart.point = {
    circle: function(style, rx, ry, x, y, graph) {
      this.ctx.save();
      this.ctx.strokeStyle = style.point.stroke;
      this.ctx.lineWidth = style.point.width * this.resolution;
      this.ctx.fillStyle = style.point.fill;
      this.ctx.beginPath();
      this.ctx.arc(rx, ry, style.point.radius * this.resolution, 0, Math.PI * 2, true);
      this.ctx.fill();
      this.ctx.stroke();
      this.ctx.restore();
    }
  };
  Aristochart.line = {
    line: function(style, points) {
      this.ctx.save();
      this.ctx.strokeStyle = style.line.stroke;
      this.ctx.lineWidth = style.line.width * this.resolution;
      this.ctx.beginPath();
      this.ctx.moveTo(points[0].rx, points[0].ry);
      var that2 = this;
      points.forEach(function(point) {
        that2.ctx.lineTo(point.rx, point.ry);
      });
      this.ctx.stroke();
      this.ctx.restore();
    },
    fill: function(style, points) {
      this.ctx.save();
      this.ctx.fillStyle = style.line.fill;
      this.ctx.beginPath();
      this.ctx.moveTo(points[0].rx, points[0].ry);
      var that2 = this;
      points.forEach(function(point) {
        that2.ctx.lineTo(point.rx, point.ry);
      });
      this.ctx.lineTo(points[points.length - 1].rx, this.box.y + this.box.y1 + (style.line.fillToBaseLine ? this.options.padding : 0));
      this.ctx.lineTo(points[0].rx, this.box.y + this.box.y1 + (style.line.fillToBaseLine ? this.options.padding : 0));
      this.ctx.closePath();
      this.ctx.fill();
      this.ctx.restore();
    }
  };
  Aristochart.tick = {
    line: function(style, x, y, type, i) {
      this.ctx.save();
      this.ctx.strokeStyle = style.tick.stroke;
      this.ctx.lineWidth = style.tick.width * this.resolution;
      this.ctx.beginPath();
      var length = i % 2 == 0 ? style.tick.major : style.tick.minor;
      length *= this.resolution;
      var mx = x, my = y;
      switch (style.tick.align) {
        case "middle":
          if (type == "x")
            my = y - length / 2;
          if (type == "y")
            mx = x - length / 2;
          break;
        case "inside":
          if (type == "x")
            my = y - length;
          mx = x;
          break;
        case "outside":
          if (type == "x")
            my = y;
          if (type == "y")
            mx = x - length;
          break;
      }
      this.ctx.moveTo(mx, my);
      if (type == "x")
        this.ctx.lineTo(mx, my + length);
      else
        this.ctx.lineTo(mx + length, my);
      this.ctx.stroke();
      this.ctx.restore();
    }
  };
  Aristochart.axis = {
    line: function(style, x, y, x1, y1, type) {
      this.ctx.save();
      this.ctx.strokeStyle = style.axis.stroke;
      this.ctx.lineWidth = style.axis.width * this.resolution;
      this.ctx.beginPath();
      this.ctx.moveTo(x, y);
      this.ctx.lineTo(x1, y1);
      this.ctx.stroke();
      this.ctx.restore();
    }
  };
  Aristochart.label = {
    text: function(style, text, x, y, type, i) {
      if (i % this.options.label[type].step == 0) {
        var label = style.label[type];
        if (type == "x")
          y = y + (style.tick.major + label.offsetY) * this.resolution;
        if (type == "y")
          x = x - (style.tick.major + label.offsetX) * this.resolution, y += label.offsetY * this.resolution;
        this.ctx.font = label.fontStyle + " " + label.fontSize * this.resolution + "px " + label.font;
        this.ctx.fillStyle = label.color;
        this.ctx.textAlign = label.align;
        this.ctx.textBaseline = label.baseline;
        var substr = /(\-?\d+(\.\d)?)/.exec(text) || [];
        this.ctx.fillText(substr[0], x, y);
      }
    }
  };
  Aristochart.title = {
    text: function(style, text, x, y, type) {
      this.ctx.save();
      if (type == "x")
        y += style.title.x.offsetY, x += style.title.x.offsetX;
      if (type == "y")
        y += style.title.y.offsetY, x += style.title.y.offsetX;
      this.ctx.font = style.title.fontStyle + " " + style.title.fontSize * this.resolution + "px " + style.title.font;
      this.ctx.fillStyle = style.title.color;
      this.ctx.translate(x, y);
      if (type == "y")
        this.ctx.rotate(Math.PI / 2);
      this.ctx.fillText(text, 0, 0);
      this.ctx.restore();
    }
  };
  if (window.jQuery)
    jQuery.fn.aristochart = function(options, theme) {
      if (this.length > 1)
        this.each(function(elem) {
          new Aristochart(this[0], options, theme);
        });
      else
        return new Aristochart(this[0], options, theme);
    };
  Aristochart.themes = {};
  Aristochart.themes.default = {
    width: 640,
    height: 400,
    margin: 70,
    padding: 20,
    render: true,
    //Automatically render
    fill: {
      index: 0,
      render: Aristochart.line.fill
    },
    axis: {
      index: 1,
      render: Aristochart.axis.line,
      x: {
        steps: 5,
        render: Aristochart.axis.line
      },
      y: {
        steps: 10,
        render: Aristochart.axis.line
      }
    },
    tick: {
      index: 2,
      render: Aristochart.tick.line
    },
    line: {
      index: 3,
      render: Aristochart.line.line
    },
    point: {
      index: 4,
      render: Aristochart.point.circle
    },
    label: {
      index: 5,
      render: Aristochart.label.text,
      x: {
        step: 1
      },
      y: {
        step: 1
      }
    },
    title: {
      index: 6,
      render: Aristochart.title.text,
      x: "x",
      y: "y"
    },
    style: {
      default: {
        point: {
          stroke: "#000",
          fill: "#fff",
          radius: 4,
          width: 3,
          visible: true
        },
        line: {
          stroke: "#298281",
          width: 3,
          fill: "rgba(150, 215, 226, 0.4)",
          fillToBaseLine: true,
          visible: true
        },
        axis: {
          stroke: "#ddd",
          width: 3,
          visible: true,
          x: {
            visible: true,
            fixed: true
          },
          y: {
            visible: true,
            fixed: true
          }
        },
        tick: {
          align: "middle",
          //"outside", "inside",
          stroke: "#ddd",
          width: 2,
          minor: 10,
          major: 15,
          visible: true,
          x: {
            fixed: true
          },
          y: {
            fixed: true
          }
        },
        label: {
          x: {
            font: "Helvetica",
            fontSize: 14,
            fontStyle: "normal",
            color: "#000",
            align: "center",
            baseline: "bottom",
            offsetY: 8,
            offsetX: 3,
            visible: true,
            fixed: true
          },
          y: {
            font: "Helvetica",
            fontSize: 10,
            fontStyle: "normal",
            color: "#000",
            align: "center",
            baseline: "bottom",
            offsetY: 8,
            offsetX: 8,
            visible: true,
            fixed: true
          }
        },
        title: {
          color: "#777",
          font: "georgia",
          fontSize: "16",
          fontStyle: "italic",
          visible: true,
          x: {
            offsetX: 0,
            offsetY: 120,
            visible: true
          },
          y: {
            offsetX: -135,
            offsetY: 10,
            visible: true
          }
        }
      }
    }
  };

  // input.js
  var Logging = class {
    constructor() {
      this.el = document.querySelector(".logging code");
      this.output = "";
    }
    info(s) {
      this.output += s + "\n";
      this.render();
    }
    render() {
      this.el.innerHTML = this.output;
    }
  };
  document.addEventListener("DOMContentLoaded", () => {
    const logging = new Logging();
    const el = document.getElementById("plot");
    const a2 = new Aristochart(el, {
      style: {
        default: {
          point: {
            visible: false,
            radius: 2
          }
        }
      },
      data: {
        x: 10,
        y: [0, 1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 3, 4, 5, 6, 7, 8, 9, 10],
        y1: [10, 9, 8, 7, 7, 6, 2, 1, 1, 1],
        y2: [4, 4, 4, 3, 3, 2, 2, 1, 0, 0, 0, 0]
      }
    });
  });
})();
/**
 * Aristochart.js
 *
 * http://dunxrion.github.com/aristochart
 * 
 * @version 0.2
 * @author Adrian Cooney <cooney.adrian@gmail.com> (http://adriancooney.ie)
 * @license http://opensource.org/licenses/MIT
 */
