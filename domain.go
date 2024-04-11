package serpstat

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

type GetDomainsInfoParams struct {
	Domains []string               `json:"domains"`
	Se      string                 `json:"se"`
	Filters map[string]interface{} `json:"filters,omitempty"`
}
type GetDomainsInfoRequest struct {
	ID     int                  `json:"id"`
	Method string               `json:"method"`
	Params GetDomainsInfoParams `json:"params"`
}

type SerpstatDomain struct {
	Domain string `json:"domain"`
	Traff  int    `json:"traff"`
}
type GetDomainsInfoResponse struct {
	ID     string        `json:"id"`
	Error  SerpstatError `json:"error"`
	Result struct {
		Data        []SerpstatDomain `json:"data"`
		SummaryInfo SerpstatSummary  `json:"summary_info"`
	} `json:"result"`
}

func (c *Client) GetDomainsInfo(req *GetDomainsInfoRequest) (*GetDomainsInfoResponse, error) {
	var resp GetDomainsInfoResponse
	requestURL := fmt.Sprintf("%s?token=%s", c.config.BaseURL, c.config.authToken)
	if err := requests.URL(requestURL).Client(c.HTTPClient).BodyJSON(&req).ToJSON(&resp).Fetch(context.TODO()); err != nil {
		return nil, err
	}

	return &resp, nil
}

type GetDomainKeywordsParams struct {
	Domain  string                 `json:"domain"`
	Se      string                 `json:"se"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Sort    map[string]string      `json:"sort"`
	Page    int                    `json:"page"`
	Size    int                    `json:"size"`
}
type GetDomainKeywordsRequest struct {
	ID     int                     `json:"id"`
	Method string                  `json:"method"`
	Params GetDomainKeywordsParams `json:"params"`
}

type GetDomainKeywordsResponse GetRelatedKeywordsResponse

func (c *Client) GetDomainKeywords(req *GetDomainKeywordsRequest) (*GetDomainKeywordsResponse, error) {
	var resp GetDomainKeywordsResponse
	requestURL := fmt.Sprintf("%s?token=%s", c.config.BaseURL, c.config.authToken)
	if err := requests.URL(requestURL).Client(c.HTTPClient).BodyJSON(&req).ToJSON(&resp).Fetch(context.TODO()); err != nil {
		return nil, err
	}

	return &resp, nil
}
