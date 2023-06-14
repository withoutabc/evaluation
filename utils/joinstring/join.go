package joinstring

import "fmt"

const (
	origin = "origin"
	dealt  = "dealt"
	result = "result"
)

var levelMap = map[int32]string{
	4: "四级",
	6: "六级",
}

func Join(year, month, set, level int32) string {
	return fmt.Sprintf("%d-%d-%d-%s", year, month, set, levelMap[level])
}

func JoinOrigin(s string) string {
	return fmt.Sprintf("%s-%s", s, origin)
}

func JoinDealt(s string) string {
	return fmt.Sprintf("%s-%s", s, dealt)
}

func JoinResult(s string) string {
	return fmt.Sprintf("%s-%s", s, result)
}
