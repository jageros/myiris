package errcode

import "strconv"

type ErrCode int32

func (e ErrCode) Error() string {
	return strconv.Itoa(int(e))
}

const (
	Success      ErrCode = 0
	InternalErr  ErrCode = -1
)