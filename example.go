package main

import (
	"fmt"
	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"log"
	"strconv"
)

func main() {
	client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "test")
	//searchQuery(client)
	//searchQueryPaged(client)
	itemQuery(client)
}

// curl -H 'X-Api-Key: test' 'http://content.guardianapis.com/football/live/2017/jun/29/germany-v-mexico-confederations-cup-semi-final-live?show-fields=bodyText&format=json&api-key=test'
// http://content.guardianapis.com/search?order-by=newest&q=football&api-key=test
// http://content.guardianapis.com/football/live/2017/jun/29/germany-v-mexico-confederations-cup-semi-final-live?show-fields=all&format=json&api-key=test
// http://content.guardianapis.com/football/live/2017/jun/29/germany-v-mexico-confederations-cup-semi-final-live?show-fields=bodyText&format=json&api-key=test
// http://content.guardianapis.com/football/live/2017/jun/29/germany-v-mexico-confederations-cup-semi-final-live?show-fields=bodyText&format=json&api-key=test

func searchQueryPaged(client *gocapiclient.GuardianContentClient) {
	searchQuery := queries.NewSearchQuery()
	searchQuery.PageOffset = int64(10)

	showParam := queries.StringParam{"q", "sausages"}
	params := []queries.Param{&showParam}

	searchQuery.Params = params

	iterator := client.SearchQueryIterator(searchQuery)

	for page := range iterator {
		fmt.Println("Page: " + strconv.FormatInt(int64(page.SearchResponse.CurrentPage), 10))
		for _, v := range page.SearchResponse.Results {
			fmt.Println(v.ID)
		}
	}
}

func searchQuery(client *gocapiclient.GuardianContentClient) {
	searchQuery := queries.NewSearchQuery()

	showParam := queries.StringParam{"q", "sausages"}
	params := []queries.Param{&showParam}

	searchQuery.Params = params

	err := client.GetResponse(searchQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(searchQuery.Response.Status)
	fmt.Println(searchQuery.Response.Total)

	for _, v := range searchQuery.Response.Results {
		fmt.Println(v.ID)
	}
}

func itemQuery(client *gocapiclient.GuardianContentClient) {
	itemQuery := queries.NewItemQuery("football/live/2017/jun/29/germany-v-mexico-confederations-cup-semi-final-live")

	showParam := queries.StringParam{"show-fields", "all"}
	params := []queries.Param{&showParam}

/*	showParam := queries.StringParam{"show-blocks", "body:latest:10"}
	formatParam := queries.StringParam{"format", "json"}
	params := []queries.Param{&showParam, &formatParam}*/

	itemQuery.Params = params

	err := client.GetResponse(itemQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(itemQuery) // prints response
	fmt.Println(itemQuery.Response.Status)
	fmt.Println(itemQuery.Response.Content.WebTitle)

	//fmt.Println(len(itemQuery.Response.Content.Elements))
	//fmt.Println(len(itemQuery.Response.Content.Tags))
	//fmt.Println(itemQuery.Response.Content.References)
}