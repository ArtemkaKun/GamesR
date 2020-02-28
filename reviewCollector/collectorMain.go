package reviewCollector

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var REQ_PARAMS = PARAMS{"all", "en", "", "",
	"all", "all", "100", "836620"}

func GetReviews() {
	resp, err := http.Get(fmt.Sprintf("https://store.steampowered.com/appreviews/%v?json=1&filter=%v&language=%v&review_type=%vpurchase_type=%v&num_per_page=%v",
		REQ_PARAMS.AppID, REQ_PARAMS.Filter, REQ_PARAMS.Language, REQ_PARAMS.Review_type,
		REQ_PARAMS.Purchase_type, REQ_PARAMS.Num_per_page))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))
}
