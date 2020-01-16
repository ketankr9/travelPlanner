package main

import (
    "fmt"
    "strings"
    "flag"
    "./runHist"
)

func printLine(s *runHist.Station) {
    fmt.Printf("%s\t%d\t%d\t%d\t%d\t%.1f\n", s.Code, s.Rht, s.L1hr, s.G1hr, s.Can, s.Avg/60.0)
}

func contains(destFlag string, s []string) bool {
    for _, a := range s {
        if a == destFlag {
            return true
        }
    }
    return false
}

func main(){
    // parse input
    trainNo := flag.String("train", "19305", "Train No")
    last := flag.String("last", "1m", "1w, 1m, 3m, 6m, 1y")
    dest := flag.String("dest", "", "destination CODE separated by ,")
    flag.Parse()

    // print headers
    fmt.Printf("Train No: %s For: %s\n",*trainNo, *last)
    fmt.Printf("Code\tR\t<1\t>1\tC\tAvg(hrs)\n")

    // scrape data
    status := []*runHist.Station{}
    runHist.GetHistory(*trainNo, *last, &status)
    
    for _, s:= range status {
        if *dest == "" || contains(s.Code, strings.Split(*dest, ",")){
            printLine(s)
        }
    }
}
