package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ReportCollection represents the ReportCollection schema from the OpenAPI specification
type ReportCollection struct {
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []ReportRecordContract `json:"value,omitempty"` // Page values.
}

// ReportRecordContract represents the ReportRecordContract schema from the OpenAPI specification
type ReportRecordContract struct {
	Apiregion string `json:"apiRegion,omitempty"` // API region identifier.
	Apitimemin float64 `json:"apiTimeMin,omitempty"` // Minimum time it took to process request.
	Callcountother int `json:"callCountOther,omitempty"` // Number of other calls.
	Cachehitcount int `json:"cacheHitCount,omitempty"` // Number of times when content was served from cache policy.
	Productid string `json:"productId,omitempty"` // Product identifier path. /products/{productId}
	Bandwidth int64 `json:"bandwidth,omitempty"` // Bandwidth consumed.
	Name string `json:"name,omitempty"` // Name depending on report endpoint specifies product, API, operation or developer name.
	Operationid string `json:"operationId,omitempty"` // Operation identifier path. /apis/{apiId}/operations/{operationId}
	Callcountblocked int `json:"callCountBlocked,omitempty"` // Number of calls blocked due to invalid credentials. This includes calls returning HttpStatusCode.Unauthorized and HttpStatusCode.Forbidden and HttpStatusCode.TooManyRequests
	Callcountfailed int `json:"callCountFailed,omitempty"` // Number of calls failed due to proxy or backend errors. This includes calls returning HttpStatusCode.BadRequest(400) and any Code between HttpStatusCode.InternalServerError (500) and 600
	Region string `json:"region,omitempty"` // Country region to which this record data is related.
	Callcounttotal int `json:"callCountTotal,omitempty"` // Total number of calls.
	Timestamp string `json:"timestamp,omitempty"` // Start of aggregation period. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Userid string `json:"userId,omitempty"` // User identifier path. /users/{userId}
	Cachemisscount int `json:"cacheMissCount,omitempty"` // Number of times content was fetched from backend.
	Country string `json:"country,omitempty"` // Country to which this record data is related.
	Servicetimeavg float64 `json:"serviceTimeAvg,omitempty"` // Average time it took to process request on backend.
	Servicetimemin float64 `json:"serviceTimeMin,omitempty"` // Minimum time it took to process request on backend.
	Zip string `json:"zip,omitempty"` // Zip code to which this record data is related.
	Apitimeavg float64 `json:"apiTimeAvg,omitempty"` // Average time it took to process request.
	Subscriptionid string `json:"subscriptionId,omitempty"` // Subscription identifier path. /subscriptions/{subscriptionId}
	Apitimemax float64 `json:"apiTimeMax,omitempty"` // Maximum time it took to process request.
	Interval int64 `json:"interval,omitempty"` // Length of aggregation period.
	Callcountsuccess int `json:"callCountSuccess,omitempty"` // Number of successful calls. This includes calls returning HttpStatusCode <= 301 and HttpStatusCode.NotModified and HttpStatusCode.TemporaryRedirect
	Servicetimemax float64 `json:"serviceTimeMax,omitempty"` // Maximum time it took to process request on backend.
	Apiid string `json:"apiId,omitempty"` // API identifier path. /apis/{apiId}
}
