/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2. API v2 is currently in Preview. Additional documentation for this API can be found in the [API Preview Docs](https://github.com/CircleCI-Public/api-preview-docs/tree/master/docs). Breaking changes to the API will be announced in the [Breaking Changes log](https://github.com/CircleCI-Public/api-preview-docs/blob/master/docs/breaking.md).
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Metrics relating to the duration of runs for a workflow.
type InlineResponse200MetricsDurationMetrics struct {
	// The minimum duration, in seconds, among a group of runs.
	Min int64 `json:"min"`
	// The mean duration, in seconds, among a group of runs.
	Mean int64 `json:"mean"`
	// The median duration, in seconds, among a group of runs.
	Median int64 `json:"median"`
	// The 95th percentile duration, in seconds, among a group of runs.
	P95 int64 `json:"p95"`
	// The max duration, in seconds, among a group of runs.
	Max int64 `json:"max"`
	// The standard deviation, in seconds, among a group of runs.
	StandardDeviation float32 `json:"standard_deviation"`
}
