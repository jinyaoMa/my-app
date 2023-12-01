package jwt

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type RefreshToken string

func (refreshToken RefreshToken) Split() (id string, expirationTime time.Time, err error) {
	slice := strings.SplitN(string(refreshToken), ":", 2)
	nsec, err := strconv.ParseInt(slice[0], 10, 64)
	if err != nil {
		return "", time.Time{}, err
	}
	expirationTime = time.Unix(0, nsec)
	id = slice[1]
	return
}

func NewRefreshToken(id string, expirationTime time.Time) RefreshToken {
	return RefreshToken(fmt.Sprintf("%d:%s", expirationTime.UnixNano(), id))
}
