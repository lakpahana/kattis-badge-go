package main

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly/v2"
)

func main() {
	ScrapeCountryTopList("LKA")
}

func ScrapeCountryTopList(country_code string) {

	// var string uni_code
	country_link := "https://open.kattis.com/countries/" + country_code
	fmt.Println(country_link)
	c := colly.NewCollector(
		colly.AllowedDomains("https://open.kattis.com/", "open.kattis.com", "www.open.kattis.com"),
	)
	//document.getElementById("top_users").getElementsByTagName('tbody')[0].getElementsByTagName('tr')[0].
	c.OnHTML("table[id=top_users] > tbody > tr ", func(h *colly.HTMLElement) {
		// fmt.Println("got")
		singleUser := h

		cssSelectorRank := "td:nth-child(1)"
		cssSelectorNameAndLink := "td:nth-child(2) > a:nth-child(1)"
		cssSelectorUni := "td:nth-child(3) > div:nth-child(1) > div:nth-child(2) > a:nth-child(1)"
		cssSelectorScore := "td:nth-child(4)"

		rank := singleUser.ChildText(cssSelectorRank)
		username := strings.Split(singleUser.ChildAttr(cssSelectorNameAndLink, "href"), "/")[2]
		name := singleUser.ChildText(cssSelectorNameAndLink)
		uni := singleUser.ChildText(cssSelectorUni)
		uni_code := ""
		// fmt.Println(len(strings.Split(singleUser.ChildAttr(cssSelectorUni, "href"), "/")))
		if len(strings.Split(singleUser.ChildAttr(cssSelectorUni, "href"), "/")) > 1 {
			uni_code = strings.Split(singleUser.ChildAttr(cssSelectorUni, "href"), "/")[2]
		}

		// _ = uni_code;

		score := singleUser.ChildText(cssSelectorScore)

		fmt.Println(name, rank, username, uni, uni_code, score)

		// fmt.Println()

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(country_link)
}
