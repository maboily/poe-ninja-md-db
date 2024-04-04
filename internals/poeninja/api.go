package poeninja

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/rotisserie/eris"
	"strings"
)

const HttpBaseUrl = "https://poe.ninja/api/data"

type HttpClient struct {
	League string
}

type HttpErrorResponse struct {
	Title   string `json:"title"`
	TraceId string `json:"traceId"`
	Status  int    `json:"status"`

	Errors struct {
		Type []string `json:"type"`
	} `json:"errors"`
}

func (her *HttpErrorResponse) ToError() error {
	errLines := make([]string, 0)
	for _, errType := range her.Errors.Type {
		errLines = append(errLines, fmt.Sprintf(" - %s", errType))
	}
	errText := fmt.Sprintf(
		"poe.ninja API error occurred (Status: %d): %s\n%s",
		her.Status,
		her.Title,
		strings.Join(errLines, "\n"),
	)
	return eris.New(errText)
}

func NewHttpClient(league string) *HttpClient {
	return &HttpClient{League: league}
}

func (hc *HttpClient) GetItemsOverview(itemType ItemType) (response *ItemOverviewResponse, err error) {
	var errResponse *HttpErrorResponse
	err = req.C().SetBaseURL(HttpBaseUrl).
		Get("itemoverview").
		SetQueryParam("league", hc.League).
		SetQueryParam("itemtype", string(itemType)).
		SetErrorResult(&errResponse).
		Do().
		Into(&response)
	if err != nil {
		return nil, eris.Wrap(err, "Failed to get item overview from poe.ninja")
	}

	if errResponse.Status != 0 {
		return nil, errResponse.ToError()
	}

	return response, nil
}
