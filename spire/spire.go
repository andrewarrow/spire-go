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
		token + "&date=2016-05-04")

	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println(err)
	}

	store := make(map[int][]float64)

	data := result["data"].([]interface{})
	for _, raw := range data {
		hash := raw.(map[string]interface{})
		ts := int64(hash["timestamp"].(float64))
		value := hash["value"].(float64)
		date := time.Unix(ts, 0)
		hour := date.Hour()
		if store[hour] == nil {
			store[hour] = make([]float64, 0)
		}
		store[hour] = append(store[hour], value)
	}

	averages := make(map[int]float64)
	for k, v := range store {
		length := len(v)
		sum := 0.0
		for _, val := range v {
			sum += val
		}
		avg := sum / float64(length)
		averages[k] = avg
	}
	hour := 0
	for {
		fmt.Println(hour, averages[hour])
		hour++
		if hour > 23 {
			break
		}
	}
}
