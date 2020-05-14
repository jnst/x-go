package main

import (
	"fmt"
	"sort"
	"time"
)

type Hero struct {
	ID   string
	Name string
	Born int64
}

type YearMonth struct {
	Year  int
	Month int
}

func printSortedHeroes() {
	t1 := time.Date(2001, 5, 31, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(1999, 4, 1, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2015, 2, 21, 0, 0, 0, 0, time.UTC)
	t4 := time.Date(2008, 1, 10, 0, 0, 0, 0, time.UTC)
	t5 := time.Date(2002, 5, 9, 0, 0, 0, 0, time.UTC)

	heroes := []Hero{
		{ID: "A001", Name: "Hero1", Born: t1.Unix()},
		{ID: "A002", Name: "Hero2", Born: t2.Unix()},
		{ID: "A003", Name: "Hero3", Born: t3.Unix()},
		{ID: "A004", Name: "Hero4", Born: t4.Unix()},
		{ID: "A005", Name: "Hero5", Born: t5.Unix()},
	}

	sort.Slice(heroes, func(i, j int) bool {
		return heroes[i].Born < heroes[j].Born
	})

	for _, hero := range heroes {
		fmt.Printf("%+v\n", hero)
	}
}

func printSortedYearMonths() {
	yms := []YearMonth{
		{Year: 2020, Month: 8},
		{Year: 2020, Month: 4},
		{Year: 1999, Month: 7},
		{Year: 2032, Month: 1},
		{Year: 2005, Month: 1},
		{Year: 2020, Month: 6},
	}

	sort.Slice(yms, func(i, j int) bool {
		if yms[i].Year < yms[j].Year {
			return true
		}
		if yms[i].Year == yms[j].Year && yms[i].Month < yms[j].Month {
			return true
		}
		return false
	})

	for _, ym := range yms {
		fmt.Printf("%+v\n", ym)
	}
}

func main() {
	printSortedHeroes()
	printSortedYearMonths()
}
