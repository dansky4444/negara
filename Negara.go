package main

import (
	"fmt"
)

var aseanCountry = [10]string{
	"Brunei",
	"Cambodia",
	"Indonesia",
	"Laos",
	"Malaysia",
	"Myanmar",
	"Philippines",
	"Singapore",
	"Thailand",
	"Vietnam"}

type country struct {
	name      string
	year, gdp int
}

type aseanCountryTab struct {
	country  [100]country
	cCountry int
}

type tabShort struct {
	shortAsenCountry [10]shortAsenCountry
	n                int
}

type shortAsenCountry struct {
	name   string
	arrgdp [3]int
	len    int
}

func validate_AseanCountry_by_name_1302220048(search string) bool {
	fmt.Println(search)
	var output bool = false
	i := 0
	for i < len(aseanCountry) && !output {
		if aseanCountry[i] == search {
			output = true
		}
		i++
	}
	return output
}

func inputCountry_1302220048(astab *aseanCountryTab) {
	var (
		name      string
		year, gdp int
	)
	fmt.Scanln(&name, &year, &gdp)
	for validate_AseanCountry_by_name_1302220048(name) {
		astab.country[astab.cCountry].name = name
		astab.country[astab.cCountry].year = year
		astab.country[astab.cCountry].gdp = gdp
		astab.cCountry++
		fmt.Scanln(&name, &year, &gdp)
	}
}

func GroupByAseanGDP_1302220048(tab *tabShort, astab *aseanCountryTab) {
	for i := 0; i < astab.cCountry; i++ {
		var found = false
		var j = 0
		for j < tab.n && !found {
			if tab.shortAsenCountry[j].name == astab.country[i].name {
				found = true
				tab.shortAsenCountry[j].arrgdp[indexByYears_1302220048(astab.country[i].year)] = astab.country[i].gdp
				tab.shortAsenCountry[j].len++
			}
			j++
		}
		if !found {
			tab.shortAsenCountry[tab.n].name = astab.country[i].name
			tab.shortAsenCountry[tab.n].arrgdp[indexByYears_1302220048(astab.country[i].year)] = astab.country[i].gdp
			tab.shortAsenCountry[tab.n].len++
			tab.n++
		}
	}
}

func indexByYears_1302220048(years int) int {
	if years == 2021 {
		return 0
	} else if years == 2022 {
		return 1
	} else if years == 2023 {
		return 2
	}

	return 0
}

func ShortByLast_1302220048(s *tabShort) {
	for i := 0; i < s.n-1; i++ {
		var minIndex = i
		var j = minIndex + 1
		for j < s.n {
			if s.shortAsenCountry[minIndex].arrgdp[2] < s.shortAsenCountry[j].arrgdp[2] {
				minIndex = j
			}
			j++
		}
		var temp = s.shortAsenCountry[minIndex]
		s.shortAsenCountry[minIndex] = s.shortAsenCountry[i]
		s.shortAsenCountry[i] = temp
	}
}

func printTabs_1302220048(s *tabShort) {
	for i := 0; i < s.n; i++ {
		fmt.Println(s.shortAsenCountry[i].name, s.shortAsenCountry[i].arrgdp[0], s.shortAsenCountry[i].arrgdp[1], s.shortAsenCountry[i].arrgdp[0])
	}
}

func main() {
	var astab aseanCountryTab
	var s tabShort
	inputCountry_1302220048(&astab)
	GroupByAseanGDP_1302220048(&s, &astab)
	ShortByLast_1302220048(&s)
	printTabs_1302220048(&s)

}
