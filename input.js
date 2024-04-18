
import { Aristochart } from './aristochart/Aristochart.js';

class Logging {

  constructor() {
    this.el = document.querySelector('.logging code');
    this.output = "";
  }  
  
  info(s) {
    this.output += s + "<br/>\n";
    this.render();
  }
  
  render() {
    this.el.innerHTML = this.output;
  }
}

document.addEventListener("DOMContentLoaded", () => {
  const logging = new Logging();
  logging.info("point 1");
  const el = document.getElementById('tester');
  logging.info("point 2.1");

  try {
    logging.info("point 2.5");
    const a = Aristochart;
  } catch (e) {
    logging.info(e);
  }
  logging.info(typeof(Aristochart));
  logging.info("point 3.1");
  new Aristochart(el, {
  	data: {
  		x: 10,
  		y: [0, 1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 3, 4, 5, 6, 7, 8, 9, 10],
  		y1: [10, 9, 8, 7, 7, 6, 2, 1, 1, 1],
  		y2: [4, 4, 4, 3, 3, 2, 2, 1, 0, 0, 0, 0]
  	}
  });
  logging.info("point 3");
});
