package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/gocolly/colly/v2"
	"lakpahana.me/db"
	"lakpahana.me/db/models"
	"lakpahana.me/db/repo"
)

func main() {
	ScrapeCountryTopList("LKA")
}

func ScrapeCountryTopList(country_code string) {

	db.ConnectToDB()
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
			score_selector := "4"
			if hasSubdivision {
				uni_selector_child = "4"
				score_selector = "5"
			}
			cssSelectorUni := "td:nth-child(" + uni_selector_child + ") > div:nth-child(1) > div:nth-child(2) > a:nth-child(1)"

			cssSelectorRank := "td:nth-child(1)"
			cssSelectorNameAndLink := "td:nth-child(2) > a:nth-child(1)"
			cssSelectorSubdiv := "td:nth-child(3) > div:nth-child(1) > div:nth-child(2) > a:nth-child(1)"
			cssSelectorScore := "td:nth-child(" + score_selector + ")"

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

			rnk, er := strconv.Atoi(rank)
			scr, er := strconv.ParseFloat(score, 64)
			if er != nil {
				panic(er)
			}

			user := models.User{
				Username:         username,
				Name:             name,
				Rank:             0,
				Score:            scr,
				Country:          "Sri Lanka",
				Country_code:     country_code,
				Country_rank:     rnk,
				Subdivision:      subdiv,
				Subdivision_code: subdiv_code,
				Subdivision_rank: 0,
				University:       uni,
				University_code:  uni_code,
				University_rank:  0,
			}

			repo.CreateOrUpdateUser(db.DB, user)

			// fmt.Println(name, rnk, username, subdiv, country_code, subdiv_code, uni, uni_code, scr)
		})

		// fmt.Println()

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(country_link)
}
