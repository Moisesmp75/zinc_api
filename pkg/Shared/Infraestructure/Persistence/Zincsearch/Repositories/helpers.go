package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RequestToString(query string, size int, from int64, sort string) string {
	if sort == "" {
		sort = "-@timestamp"
	}
	if size == 0 {
		size = 10
	}
	if from == 0 {
		from = 0
	}
	if query == "" {
		return parseRequestGetAll(size, from, sort)
	}
	return parseRequestSearch(query, size, from, sort)
}

func parseRequestGetAll(size int, from int64, sort string) string {
	return fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": [
					{
						"match_all": {}
					}
				]
			}
		},
		"sort": [
			"%v"
		],
		"from": %v,
		"size": %v,
		"aggs": {
			"histogram": {
					"auto_date_histogram": {
							"field": "@timestamp",
							"buckets": 100
					}
			}
		}
	}
	`, sort, from, size)
}

func parseRequestSearch(query string, size int, from int64, sort string) string {
	return fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": [
					{
						"query_string": {
							"query": "%v"
						}
					}
				]
			}
		},
		"sort": [
			"%v"
		],
		"from": %v,
		"size": %v,
		"aggs": {
			"histogram": {
					"auto_date_histogram": {
							"field": "date",
							"buckets": 100
					}
			}
		}
	}
	`, query, sort, from, size)
}

func performHTTPPostRequest(url string, request string, username string, password string) ([]byte, error) {
	reqs, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(request)))
	if err != nil {
		return nil, err
	}
	reqs.SetBasicAuth(username, password)
	reqs.Header.Set("Content-Type", "application/json")

	resps, err := http.DefaultClient.Do(reqs)
	if err != nil {
		return nil, err
	}
	defer resps.Body.Close()

	if resps.StatusCode < 200 || resps.StatusCode >= 300 {
		body, _ := io.ReadAll(resps.Body)
		return nil, fmt.Errorf("HTTP error: status code %d, response body: %s", resps.StatusCode, string(body))
	}

	bodyResponse, err := io.ReadAll(resps.Body)
	if err != nil {
		return nil, err
	}
	return bodyResponse, nil
}

func parseJSONtoDto(body []byte, response any) error {
	err := json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}
