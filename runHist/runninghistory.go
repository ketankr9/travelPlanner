package runHist

import (
	"log"
    "strconv"
    "strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
    "regexp"
)


type Station struct {
    Code string // station code
    Avg float32 // average delay in minutes
    Rht int // right time
    L1hr int // less than 1 hour delay
    G1hr int // greater than 1 hour delay
    Can int // no of times cancelled
}

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func toInt(valStr string) int {
    valInt, _ := strconv.Atoi(valStr)
    return valInt
}

func toIntFloat(valStr string) float32 {
    valFlt, _ := strconv.ParseFloat(valStr, 32)
    return float32(valFlt)
}

func convertToStruct(data string, Status *[]*Station){
    vm := otto.New()

    _, _ = vm.Run("data="+data)
    len1, _ := vm.Run("data.length")

    for i := 1; i < toInt(len1.String()); i += 1 {

    	val, _ := vm.Run("data["+strconv.Itoa(i)+"];")
        arr := strings.Split(val.String(), ",")
        
        *Status = append(*Status, &Station{
            arr[0],
            toIntFloat(arr[1]),
            toInt(arr[2]),
            toInt(arr[3]),
            toInt(arr[4]),
            toInt(arr[5]),
        })
    }
}

func GetHistory(trainNumber string, last string, Status *[]*Station){
    doc, err := goquery.NewDocument("https://etrain.info/in?PAGE=runningHistory--" + trainNumber + "--" + last)
    check(err)
    doc.Find("script").Each(func (i int, s *goquery.Selection)  {
        if i == 12 {
            match := regexp.MustCompile("(.*);").FindStringSubmatch(s.Text()[450:])
            convertToStruct(match[0], Status)
        }
    })
}


