package assets

import _ "embed"

//go:embed favicon.ico
var Favicon []byte

//go:embed styles
var Styles []byte

//go:embed web-components
var WebComponents []byte
