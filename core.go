package main
import (
    "strings"
    "strconv"
    "sort"
)
// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

timings:=[]float64{}
for _,word:=range strings.Fields(input){n,err:=strconv.ParseFloat(word,64);if err==nil&&n>=0{timings=append(timings,n)}}
sort.Float64s(timings);size:=len(timings);p95:=0.0;average:=0.0
for _,n:=range timings{average+=n};if size>0{average/=float64(size);index:=(95*size+99)/100-1;p95=timings[index]}
return map[string]interface{}{"samples":size,"meanMs":average,"p95Ms":p95}
}
