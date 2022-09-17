const svgtofont = require("svgtofont");

/*
  "svgtofont": {
    "outSVGReact": false,
    "outSVGPath": false,
    "emptyDist": true,
    "styleTemplates": "./templates",
    "fontName": "mp-icon",
    "css": {
      "filename": "mp-icon",
      "include": "\\.(css|scss|less|styl)$",
      "fontSize": "1em"
    }
  }
*/

svgtofont({
  dist: "./lib", // output path
  src: "./svg", // svg path
  styleTemplates: "./styles",
  emptyDist: true, // Clear output directory contents
  fontName: "mp-icon", // font name
  css: {
    filename: "mp-icon",
    include: "\\.css$", // "\\.(css|scss|less|styl)$",
    fontSize: "1em",
  }, // Create CSS files.
  outSVGReact: false,
  outSVGPath: false,
  svgicons2svgfont: {
    fontHeight: 1000,
    normalize: true,
  },
}).then(() => {
  console.log("done!");
});
