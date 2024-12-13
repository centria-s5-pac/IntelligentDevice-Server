package lightbrain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetJson() []map[string]interface{} {
	url := "http://127.0.0.1:8080/sensor"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.SetBasicAuth("admin", "password")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	// * The response body is a JSON array of objects
	// [{"id":2000001200,"type":1,"value":1.2,"timestamp":"2021-01-01T00:00:00Z"},{"id":20000012000,"type":1,"value":1.2,"timestamp":"2021-01-01T00:00:00Z"}]

	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	return data
}

func GetLightLevel() int {
	data := GetJson()
	if data == nil {
		return 0
	}

	return 1
}
