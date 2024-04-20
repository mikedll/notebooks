
class Logging {

  constructor() {
    this.el = null;
    this.output = "";
  }  
  
  info(s) {
    this.output += s + "\n";
    this.render();      
  }
  
  render() {
    if(this.el === null) {
      this.el = document.querySelector('.logging code');      
    }
    this.el.innerHTML = this.output;
  }
}

const logging = new Logging();

export { logging };