import pkg from 'pkg'
import svgtofont from 'svgtofont'

const options = {
  src: './svg', // svg path
  dist: './dist', // output path
  emptyDist: true, // Clear output directory contents
  styleTemplates: './styles', // file templates path (optional)
  fontName: 'my-icon', // font name
  css: true, // Create CSS files.
  // css: {
  //   filename: "my-icon",
  //   include: "\\.(css|scss|less|styl)$",
  //   fontSize: "1em",
  // }, // Create CSS files.
  // outSVGReact: false,
  // outSVGPath: false,
  startUnicode: 0xea01, // unicode start number
  svgicons2svgfont: {
    fontHeight: 1000,
    normalize: true
  },
  // website = null, no demo html files
  website: {
    title: 'My Icons',
    // Must be a .svg format image.
    logo: './svg/jinyao-ma.svg',
    version: pkg.version,
    meta: {
      description: 'Converts SVG fonts to TTF/EOT/WOFF/WOFF2/SVG format.',
      keywords: 'svgtofont,TTF,EOT,WOFF,WOFF2,SVG'
    },
    description: ``,
    // Add a Github corner to your website
    // Like: https://github.com/uiwjs/react-github-corners
    corners: {
      url: 'https://github.com/jinyaoMa/my-app/tree/main/frontend/icons',
      width: 80, // default: 60
      height: 80, // default: 60
      bgColor: '#0099ff' // default: '#151513'
    },
    links: [
      {
        title: 'GitHub',
        url: 'https://github.com/jinyaoMa/my-app'
      },
      {
        title: 'Feedback',
        url: 'https://github.com/jinyaoMa/my-app/discussions'
      },
      {
        title: 'Font Class',
        url: 'index.html'
      },
      {
        title: 'Unicode',
        url: 'unicode.html'
      }
    ],
    footerInfo: `Licensed under MIT. (Yes it's free and <a href="https://github.com/jaywcjlove/svgtofont">open-sourced</a>`
  }
}

svgtofont(options).then(() => {
  console.log('done!')
})
