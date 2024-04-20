
import { Aristochart } from './aristochart/aristochart.es6.js';
import { logging } from './logging.js';

document.addEventListener("DOMContentLoaded", () => {
  const labelRender = function(style, text, x, y, type, i){
		if(i % this.options.label[type].step == 0) {
			var label = style.label[type];
			if(type == "x") y = y + (style.tick.major + label.offsetY)*this.resolution;
			if(type == "y") x = x - (style.tick.major + label.offsetX)*this.resolution, y += label.offsetY*this.resolution;

			this.ctx.font = label.fontStyle + " " + (label.fontSize*this.resolution) + "px " + label.font;
			this.ctx.fillStyle = label.color;
			this.ctx.textAlign = label.align;
			this.ctx.textBaseline = label.baseline;

      if(text > 10000) {
  			var substr = /(\-?\d+(\.\d)?)/.exec(text) || [];
        // logging.info(type + "=" + text.toExponential(3));
  			this.ctx.fillText(text.toExponential(3), x, y);        
      } else {
  			var substr = /(\-?\d+(\.\d)?)/.exec(text) || [];
        // logging.info(type + "=" + substr[0]);
  			this.ctx.fillText(substr[0], x, y);        
      }
    }
	}
  
  const plotContainer = document.getElementById('plot');
  fetch('/data.json')
  .then(r => r.json())
  .then((data) => {
    const plot = new Aristochart(plotContainer, {
      label: {
        render: labelRender
      },
      style: {
        y: {
          line: {
            stroke: '#298281'
          }
        },
        y2: {
          line: {
            stroke: 'green'
          }
        },
        y3: {
          line: {
            stroke: 'purple'
          }
        },
        y4: {
          line: {
            stroke: 'red'
          }
        },
        default: {
          point: {
            visible: false
          },
          line: {
            fill: "rgba(0,0,0,0)",
            fillToBaseLine: false
          }
        }
      },
    	data: data
    });
    const toImageButton = document.querySelector('.to-image');
    toImageButton.addEventListener('click', (e) => {
      e.preventDefault();
      const asImageContainer = document.querySelector('.as-image');
      asImageContainer.appendChild(plot.toImage());
    });
    // logging.info(JSON.stringify(plot.options.style.default, null, 2));
  })
  .catch(err => {
    logging.info("Error: " + err);
  });
    
  // logging.info(JSON.stringify(a2.options, null, 2));
});
