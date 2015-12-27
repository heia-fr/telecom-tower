var towerSocket = new WebSocket("ws://localhost:8100/stream");

towerSocket.onmessage = function (event) {
  tower.update_canvas(JSON.parse(event.data))
}

var tower = {
  radius: 2,
  margin: 3,
  sep: 1,
  columns: 128,
  rows: 8,
  init: function() {
    this.width = this.columns * (this.radius * 2 + this.sep) - this.sep + 2 * this.margin;
    this.height = this.rows * (this.radius * 2 + this.sep) - this.sep + 2 * this.margin;
    var c = document.getElementById("towerCanvas");
    c.width = this.width;
    c.height = this.height;
    var ctx = c.getContext("2d");

    ctx.beginPath();
    ctx.rect(0,0,this.width,this.height);
    ctx.fillStyle="#333333";
    ctx.fill();

    for (var y=0; y < this.rows; y++) {
      for (var x=0; x < this.columns; x++) {
        ctx.beginPath();
        ctx.arc(
          (this.sep + 2 * this.radius) * x + this.margin + this.radius,
          (this.sep + 2 * this.radius) * y + this.margin + this.radius,
          this.radius, 0, 2*Math.PI);
        ctx.strokeStyle="white";
        ctx.stroke();     
      }
    }
  },
  update_canvas: function(data) {
    var c = document.getElementById("towerCanvas");
    var ctx = c.getContext("2d");

    for (var y=0; y < this.rows; y++) {
      for (var x=0; x < this.columns; x+=2) {
        ctx.beginPath();
        ctx.arc(
          (this.sep + 2 * this.radius) * x + this.margin + this.radius,
          (this.sep + 2 * this.radius) * y + this.margin + this.radius,
          this.radius, 0, 2*Math.PI);
        var rgb = data[x*this.rows + y]
        ctx.fillStyle = '#' + (0x1000000 + rgb).toString(16).slice(1)
        ctx.fill();     
      }
      for (var x=1; x < this.columns; x+=2) {
        ctx.beginPath();
        ctx.arc(
          (this.sep + 2 * this.radius) * x + this.margin + this.radius,
          (this.sep + 2 * this.radius) * y + this.margin + this.radius,
          this.radius, 0, 2*Math.PI);
        var rgb = data[x*this.rows + (this.rows - 1 - y)]
        ctx.fillStyle = '#' + (0x1000000 + rgb).toString(16).slice(1)
        ctx.fill();     
      }

    }
  }
}
