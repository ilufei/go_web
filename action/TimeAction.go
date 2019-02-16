package action 

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
)

func PrintTime(writer http.ResponseWriter, request *http.Request) {
    t := time.Now()
	
	year := t.Year()
	month := int(t.Month())   //月份这里不用int处理，返回的是英文的月份比如November
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	nanosecond := t.Nanosecond()
	//weekday := t.Weekday()
	
	fmt.Fprintln(writer, t.Location())
	fmt.Fprintln(writer, t)
	fmt.Fprintln(writer, t.UTC())
	fmt.Fprintln(writer, t.UTC().Location())
	//fmt.Fprintln(writer, time.Date())
	fmt.Fprintln(writer, time.Unix(1487780010, 0))
	fmt.Fprintln(writer, t.Format(time.RFC3339))
	
	fmt.Fprintln(writer, strconv.Itoa(year))
	fmt.Fprintln(writer, strconv.Itoa(month))
	fmt.Fprintln(writer, strconv.Itoa(day))
	fmt.Fprintln(writer, strconv.Itoa(hour))
	fmt.Fprintln(writer, strconv.Itoa(minute))
	fmt.Fprintln(writer, strconv.Itoa(second))
	fmt.Fprintln(writer, strconv.Itoa(nanosecond))
	//fmt.Fprintln(writer, strconv.Itoa(weekday))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	
	time.Sleep(time.Second)
}
