package main

import (
	"fmt"
	"time"
)

func main() {
	var min, sec int
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	const now = 1589570165
	addtime := ((1000000000 * 60) * min) + (1000000000 * sec)
	curTime := time.Unix(now, int64(addtime))
	resultTime := curTime.Format(time.UnixDate)
	fmt.Println(resultTime)	
}
