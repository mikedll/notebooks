
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
  logging.info("fetching");
  fetch('/data.json')
  .then(r => r.json())
  .then((data) => {
    const a2 = new Aristochart(el, {
      style: {
        default: {
          point: {
            visible: false
          }
        }
      },
    	data: data
    });
  })
  .catch(err => {
    logging.info("Error: " + err);
  });
  // logging.info(JSON.stringify(a2.options, null, 2));
});
