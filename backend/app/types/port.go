package types

import (
	"fmt"
	"regexp"
	"strconv"
)

var PortMatcher = regexp.MustCompile(`^:([0-9]{1,5})$`)

type Port uint

func ParsePort(port string) Port {
	matches := PortMatcher.FindStringSubmatch(port)
	if len(matches) == 2 {
		if p, err := strconv.ParseUint(matches[1], 10, 0); err == nil {
			return Port(p)
		}
	}
	return 0
}

func (p Port) ToUint() uint {
	return uint(p)
}

func (p Port) ToString() string {
	return fmt.Sprintf(":%d", p)
}
