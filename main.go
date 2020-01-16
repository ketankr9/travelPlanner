package main

import (
    "fmt"
    // "strings"
    "flag"
    "./runHist"
    "./paytm"

)

func printLine(s *runHist.Station) {
    // fmt.Printf("%s\t%d\t%d\t%d\t%d\t%.1f\n", s.Code, s.Rht, s.L1hr, s.G1hr, s.Can, s.Avg/60.0)
    fmt.Printf("%s:%.1f\t", s.Code, s.Avg/60.0)
}

func contains(destFlag string, s []string) bool {
    for _, a := range s {
        if a == destFlag {
            return true
        }
    }
    return false
}

func printHistory(trainNo string, last string, dest string, src string){
    // scrape data
    status := []*runHist.Station{}
    runHist.GetHistory(trainNo, last, &status)
    
    var can, tot int
    for _, s:= range status {
        if dest == s.Code || src == s.Code {
            printLine(s)
            can = s.Can
            tot = s.Can+s.Rht+s.L1hr+s.G1hr
        }
    }

    fmt.Printf("Can:%d/%d", can, tot)
}

func printClass(class string, t paytm.Train){
    if val, ok := t.Avail[class]; ok {
        fmt.Printf("%s(₹%d) %s\n", class, val.Fare, val.Seats)
    }
}

func main(){
    // parse input
    // trainNo := flag.String("train", "19305", "Train No")
    src := flag.String("src", "DDU", "starting point")
    last := flag.String("last", "1m", "1w, 1m, 3m, 6m, 1y")
    dest := flag.String("dest", "NJP", "destination")
    date := flag.String("date", "20200214", "date yyyymmdd")
    flag.Parse()


    data := paytm.Api(*src, *dest, *date)
    for _, t := range data {
        fmt.Println("____________________________________________________")
        fmt.Printf("%s\t%s (%shrs)  *%s*\n", t.No, t.Dept +" - "+t.Arri, t.Dura, t.Name)
        printClass("2S", t)
        printClass("CC", t)
        printClass("SL", t)
        printClass("3A", t)
        if len(t.Alt) > 0 {
            fmt.Println("*Alt:", t.Alt)
        }
        printHistory(t.No, *last, t.Dst, t.Src)
        fmt.Println()
        fmt.Println()
	}
}
