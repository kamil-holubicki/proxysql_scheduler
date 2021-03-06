package Global

import (
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MyWaitGroup struct {
	sync.WaitGroup
	count int
}

func (wg *MyWaitGroup) WaitTimeout(timeout time.Duration) bool {
	done := make(chan struct{})

	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
		return false

	case <-time.After(timeout):
		return true
	}
}

func (wg *MyWaitGroup) IncreaseCounter(){
	wg.count++
}
func (wg *MyWaitGroup) DecreaseCounter(){
	if wg.count >0 {
		wg.count--
	}
}

func (wg *MyWaitGroup) ReportCounter() int{
	return wg.count
}

const (
	Separator = string(os.PathSeparator)
)
//==================================

func FromStringToMAp(mystring string, separator string) map[string]string{
	myMap := make(map[string]string)
	if mystring != "" && separator != ""{
		keyValuePairArray := strings.Split(mystring, separator)
		for _, keyValuePair := range keyValuePairArray {
			//keyValuePair = strings.Trim(keyValuePair," ")
			keyValueSplit := strings.Split(keyValuePair,"=")
			if len(keyValueSplit) > 1 {
				var key = strings.TrimSpace(keyValueSplit[0])
				var value = strings.TrimSpace(keyValueSplit[1])
				if len(key) > 0 && key != "" && value != "" {
					myMap[key] = value
				}
			}
		}
	}
	return myMap
}

func ToInt(myString string) int {
	if len(myString) > 0 {
		i, err := strconv.Atoi(myString)
		if err != nil {
			pc, fn, line, _ := runtime.Caller(1)
			log.Error(pc," ",fn," ",line,": ",err)
			return -1
		} else {
			return i
		}
	}
	return 0
}

func ToBool (myString string, boolTrueString string) bool{
	myString = strings.ToLower(myString)
	boolTrueString  = strings.ToLower(boolTrueString)
	if myString !="" && myString == boolTrueString{
		return true
	}else {return false}
}

func Bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

