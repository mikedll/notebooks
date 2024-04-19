
import { Aristochart } from './aristochart/Aristochart.js';


class Logging {

  constructor() {
    this.el = document.querySelector('.logging code');
    this.output = "";
  }  
  
  info(s) {
    this.output += s + "\n";
    this.render();
  }
  
  render() {
    this.el.innerHTML = this.output;
  }
}

document.addEventListener("DOMContentLoaded", () => {
  const logging = new Logging();
  const el = document.getElementById('plot');
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
  // logging.info(JSON.stringify(a2.options, null, 2));
});
