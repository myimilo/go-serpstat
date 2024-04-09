package serpstat

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

type GetRelatedKeywordsParams struct {
	Keyword string                 `json:"keyword"`
	Se      string                 `json:"se"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Sort    map[string]string      `json:"sort"`
	Page    int                    `json:"page"`
	Size    int                    `json:"size"`
}
type GetRelatedKeywordsRequest struct {
	ID     int                      `json:"id"`
	Method string                   `json:"method"`
	Params GetRelatedKeywordsParams `json:"params"`
}

type SerpstatKeyword struct {
	Keyword            string  `json:"keyword"`
	RegionQueriesCount int     `json:"region_queries_count"`
	Difficulty         float64 `json:"difficulty"`
	Cost               float64 `json:"cost"`
}
type SerpstatSummary struct {
	Page      int `json:"page"`
	Total     int `json:"total"`
	LeftLines int `json:"left_lines"`
}
type GetRelatedKeywordsResponse struct {
	ID     string        `json:"id"`
	Error  SerpstatError `json:"error"`
	Result struct {
		Data        []SerpstatKeyword `json:"data"`
		SummaryInfo SerpstatSummary   `json:"summary_info"`
	} `json:"result"`
}

func (c *Client) GetRelatedKeywords(req *GetRelatedKeywordsRequest) (*GetRelatedKeywordsResponse, error) {
	var resp GetRelatedKeywordsResponse
	requestURL := fmt.Sprintf("%s?token=%s", c.config.BaseURL, c.config.authToken)
	if err := requests.URL(requestURL).Client(c.HTTPClient).BodyJSON(&req).ToJSON(&resp).Fetch(context.TODO()); err != nil {
		return nil, err
	}

	return &resp, nil
}

type GetKeywordsInfoParams struct {
	Keywords []string          `json:"keywords"`
	Se       string            `json:"se"`
	Sort     map[string]string `json:"sort"`
}
type GetKeywordsInfoRequest struct {
	ID     int                   `json:"id"`
	Method string                `json:"method"`
	Params GetKeywordsInfoParams `json:"params"`
}
type GetKeywordsInfoResponse GetRelatedKeywordsResponse

func (c *Client) GetKeywordsInfo(req *GetKeywordsInfoRequest) (*GetKeywordsInfoResponse, error) {
	var resp GetKeywordsInfoResponse
	requestURL := fmt.Sprintf("%s?token=%s", c.config.BaseURL, c.config.authToken)
	if err := requests.URL(requestURL).Client(c.HTTPClient).BodyJSON(&req).ToJSON(&resp).Fetch(context.TODO()); err != nil {
		return nil, err
	}

	return &resp, nil
}

type GetKeywordsRequest GetRelatedKeywordsRequest
type GetKeywordsResponse GetRelatedKeywordsResponse

func (c *Client) GetKeywords(req *GetKeywordsRequest) (*GetKeywordsResponse, error) {
	var resp GetKeywordsResponse
	requestURL := fmt.Sprintf("%s?token=%s", c.config.BaseURL, c.config.authToken)
	if err := requests.URL(requestURL).Client(c.HTTPClient).BodyJSON(&req).ToJSON(&resp).Fetch(context.TODO()); err != nil {
		return nil, err
	}

	return &resp, nil
}
