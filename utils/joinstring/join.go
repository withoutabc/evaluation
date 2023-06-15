package joinstring

import (
	"fmt"
	"rpc/app/common/consts/maps"
)

const (
	origin = "origin"
	dealt  = "dealt"
	result = "result"
)

func Join(year, month, set, level int32) string {
	return fmt.Sprintf("%d-%d-%d-%s", year, month, set, maps.LevelMap[level])
}

func JoinOrigin(s string) string {
	return fmt.Sprintf("%s-%s", s, origin)
}

func JoinResult(s string) string {
	return fmt.Sprintf("%s-%s", s, result)
}
