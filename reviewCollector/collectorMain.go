package reviewCollector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetReviews(params PARAMS) []Review{
	var collected_reviews []Review

	for true {
		var buffer_info Info

		resp, err := http.Get(fmt.Sprintf("https://store.steampowered.com/appreviews/%v?json=1&filter=%v&language=%v&review_type=%v&purchase_type=%v&num_per_page=%v&cursor=%v",
			params.AppID, params.Filter, params.Language, params.Review_type,
			params.Purchase_type, params.Num_per_page, params.Cursor))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = json.Unmarshal(body, &buffer_info)
		if err != nil {
			fmt.Println(err.Error())
		}

		if len(buffer_info.Reviews) == 0 {
			break
		}

		collected_reviews = append(collected_reviews, buffer_info.Reviews...)
		params.Cursor = buffer_info.Cursor
	}

	return collected_reviews
}
