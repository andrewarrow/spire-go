package spire

import "os"
import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"
import "time"

func GetDate(date string) {
	token := os.Getenv("SPIRE_TOKEN")
	response, err := http.Get("https://app.spire.io/api/events/br?access_token=" +
		token + "&date=2016-05-05")

	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
	result := make(map[string]interface{})
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println(err)
	}

	data := result["data"].([]interface{})
	for _, raw := range data {
		hash := raw.(map[string]interface{})
		ts := int64(hash["timestamp"].(float64))
		value := hash["value"].(float64)
		date := time.Unix(ts, 0)
		fmt.Println(date, value)
	}

}
