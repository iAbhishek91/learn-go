package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func convertJsonToMap(jsonData string) map[string]interface{} {
	var result map[string]interface{}

	json.Unmarshal([]byte(jsonData), &result)

	return result
}

func main() {
	url := "https://gitlab.com/api/v4/projects/1568/repository/files/README%2Emd?ref=main" // change the URL
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", "test") // change the private token

	client := &http.Client{}
	res, _ := client.Do(req)

	respBody, _ := ioutil.ReadAll(res.Body)

	data := convertJsonToMap(string(respBody))

	stringData := fmt.Sprintf("%v", data["content"])

	datab64, _ := b64.StdEncoding.DecodeString(stringData)
	fmt.Println(string(datab64))

}
