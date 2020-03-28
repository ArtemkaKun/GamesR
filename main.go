package main

import (
	"fmt"
	"github.com/artemkakun/GamesR/reviewCollector"
	"strconv"
)

func main() {
	var reviews []string
	for i := 0; i < 100; i++ {
		var REQ_PARAMS = reviewCollector.PARAMS{"recent", "", "", "*",
		"all", "all", "100", ""}
		REQ_PARAMS.AppID = strconv.Itoa(i)
		buff_reviews := reviewCollector.GetReviews(REQ_PARAMS)

		for _, one_review := range buff_reviews {
			reviews = append(reviews, one_review.Review)
		}
	}

	fmt.Println(len(reviews))
}
