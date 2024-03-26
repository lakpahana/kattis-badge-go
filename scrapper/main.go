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
	c.OnHTML("table[id=top_users]", func(h *colly.HTMLElement) {
		// fmt.Println("got")

		hasSubdivision := false

		if h.ChildText("thead:nth-child(1) > tr:nth-child(1) > th:nth-child(3)") == "Subdivision" {
			// "td:nth-child(3) > div:nth-child(1) > div:nth-child(2) > a:nth-child(1)"
			hasSubdivision = true
		}

		h.ForEach("tbody > tr", func(i int, h *colly.HTMLElement) {

			singleUser := h
			uni_selector_child := "3"
			if hasSubdivision {
				uni_selector_child = "4"
			}
			cssSelectorUni := "td:nth-child(" + uni_selector_child + ") > div:nth-child(1) > div:nth-child(2) > a:nth-child(1)"
			
			cssSelectorRank := "td:nth-child(1)"
			cssSelectorNameAndLink := "td:nth-child(2) > a:nth-child(1)"
			cssSelectorSubdiv := "td:nth-child(3) > div:nth-child(1) > div:nth-child(2) > a:nth-child(1)"
			cssSelectorScore := "td:nth-child(" + uni_selector_child + ")"

			rank := singleUser.ChildText(cssSelectorRank)
			username := strings.Split(singleUser.ChildAttr(cssSelectorNameAndLink, "href"), "/")[2]
			name := singleUser.ChildText(cssSelectorNameAndLink)

			subdiv := ""
			subdiv_code := ""
 			if hasSubdivision {
				subdiv = singleUser.ChildText(cssSelectorSubdiv)
			
			// fmt.Println(len(strings.Split(singleUser.ChildAttr(cssSelectorUni, "href"), "/")))
			if len(strings.Split(singleUser.ChildAttr(cssSelectorSubdiv, "href"), "/")) > 1 {
				subdiv_code = strings.Split(singleUser.ChildAttr(cssSelectorSubdiv, "href"), "/")[3]
			}
			}


			uni := singleUser.ChildText(cssSelectorUni)
			uni_code := ""
			// fmt.Println(len(strings.Split(singleUser.ChildAttr(cssSelectorUni, "href"), "/")))
			if len(strings.Split(singleUser.ChildAttr(cssSelectorUni, "href"), "/")) > 1 {
				uni_code = strings.Split(singleUser.ChildAttr(cssSelectorUni, "href"), "/")[2]
			}

			// _ = uni_code;

			score := singleUser.ChildText(cssSelectorScore)

			fmt.Println(name, rank, username,subdiv,country_code,subdiv_code, uni, uni_code, score)
		})

		// fmt.Println()

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(country_link)
}
