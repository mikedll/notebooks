
import { Aristochart } from './aristochart/Aristochart.js';
import { logging } from './logging.js';

document.addEventListener("DOMContentLoaded", () => {    
  const plotContainer = document.getElementById('plot');
  fetch('/data.json')
  .then(r => r.json())
  .then((data) => {
    const plot = new Aristochart(plotContainer, {
      style: {
        y1: {
          line: {
            stroke: 'blue'
          }
        },
        y2: {
          line: {
            stroke: 'green'
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
    // logging.info(JSON.stringify(plot.options.style.default, null, 2));
  })
  .catch(err => {
    logging.info("Error: " + err);
  });
  // logging.info(JSON.stringify(a2.options, null, 2));
});
