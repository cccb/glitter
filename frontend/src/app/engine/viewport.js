

export default class Viewport {
  constructor(canvas) {
    this.canvas = canvas;
    this.ctx = canvas.getContext("2d");

    // Handle size
    this.resizeToFit();
    window.addEventListener("resize", (e) => this.resizeToFit());
  }

  resizeToFit() {
    let w = this.canvas.parentElement.clientWidth;
    let h = this.canvas.parentElement.clientHeight;
    console.log("Resizing canvas to:", w, h);
    this.canvas.width = w;
    this.canvas.height = h;
  }

  /* Map Color RGBW space to RGB */
  colorToRgb(color) {
    return "rgb(" +
          Math.floor(255 * (color[0] + color[3])) + "," +
          Math.floor(255 * (color[1] + color[3])) + "," +
          Math.floor(255 * (color[2] + color[3])) + ")";
  }

  renderStep(col, row, color) {
    let offsetX = 20;
    let offsetY = 20;
    let height  = 30;

    let x = offsetX;
    let w = this.canvas.width - 2.0*offsetX;

    let y = offsetY + (row * 50);
    let h = 10;

    this.ctx.fillStyle = this.colorToRgb(color);

    this.ctx.fillRect(x, y, w, h);
    this.ctx.fillRect(x, y + h + 10, w, h);
  }

  render(framebuffer) {
    let cols = framebuffer.width; // Ignore this for now
    let rows = framebuffer.height;

    // Clear canvas and draw framebuffer contents
    this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);

    for (let row = 0; row < rows; row++) {
      let color = framebuffer.getPixel(0, row);
      this.renderStep(0, row, color);
    }
  }

}


