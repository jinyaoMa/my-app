package log

import "io"

type IChainWriter interface {
	io.Writer

	Children() []IChainWriter
}
