package serpstat

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

type GetBacklinkSummaryParams struct {
	Query      string `json:"query"`
	SearchType string `json:"searchType,omitempty"`
}
type GetBacklinkSummaryRequest struct {
	ID     int                      `json:"id"`
	Method string                   `json:"method"`
	Params GetBacklinkSummaryParams `json:"params"`
}

type SerpstatBacklinkSummary struct {
	SersptatDomainRank int `json:"sersptat_domain_rank"`
}
type GetBacklinkSummaryResponse struct {
	ID     string        `json:"id"`
	Error  SerpstatError `json:"error"`
	Result struct {
		Data        SerpstatBacklinkSummary `json:"data"`
		SummaryInfo SerpstatSummary         `json:"summary_info"`
	} `json:"result"`
}

func (c *Client) GetBacklinkSummary(req *GetBacklinkSummaryRequest) (*GetBacklinkSummaryResponse, error) {
	var resp GetBacklinkSummaryResponse
	requestURL := fmt.Sprintf("%s?token=%s", c.config.BaseURL, c.config.authToken)
	if err := requests.URL(requestURL).Client(c.HTTPClient).BodyJSON(&req).ToJSON(&resp).Fetch(context.TODO()); err != nil {
		return nil, err
	}

	return &resp, nil
}
