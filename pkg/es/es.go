package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Elastic struct {
	Client *elasticsearch.Client
}

func New(address string) *Elastic {
	var client *elasticsearch.Client
	var err error

	cfg := elasticsearch.Config{
		Addresses: []string{address},
	}

	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	return &Elastic{Client: client}
}

func (e *Elastic) Search(index, key, value string) (ret []interface{}, err error) {
	// index 索引
	// key   需要查询的键
	// value 需要查找对应的值
	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"wildcard": map[string]interface{}{
				key: map[string]interface{}{
					"value": value,
				},
				// key: value,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return ret, fmt.Errorf("Error encoding query: %s", err)
	}
	opts := []func(*esapi.SearchRequest){
		e.Client.Search.WithSort("@timestamp:asc"),
		e.Client.Search.WithSize(10000),
		e.Client.Search.WithContext(context.Background()),
		e.Client.Search.WithIndex(index),
		e.Client.Search.WithBody(&buf),
		e.Client.Search.WithTrackTotalHits(true),
		e.Client.Search.WithPretty(),
		e.Client.Search.WithScroll(time.Minute),
	}

	res, err := e.Client.Search(opts...)

	if err != nil {
		return ret, fmt.Errorf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return ret, fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			return ret, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return ret, fmt.Errorf("Error parsing the response body: %s", err)
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

		ret = append(ret, hit.(map[string]interface{})["_source"].(map[string]interface{})["message"])
	}
	return ret, nil
	// 暂时注释掉scroll，后面补充
	// for {

	// 	var scrollID string
	// 	scrollID = r["_scroll_id"].(string)
	// 	response, err := e.Client.Scroll(
	// 		e.Client.Scroll.WithContext(context.Background()),
	// 		e.Client.Scroll.WithScrollID(scrollID),
	// 		e.Client.Scroll.WithScroll(time.Minute))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer response.Body.Close()

	// 	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
	// 		return fmt.Errorf("Error parsing the response body: %s", err)
	// 	}
	// 	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 		fmt.Println(hit.(map[string]interface{})["_source"].(map[string]interface{})["message"])

	// 	}
	// 	time.Sleep(1 * time.Second)
	// }

}
