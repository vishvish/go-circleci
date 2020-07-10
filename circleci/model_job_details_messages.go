/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2. API v2 is currently in Preview. Additional documentation for this API can be found in the [API Preview Docs](https://github.com/CircleCI-Public/api-preview-docs/tree/master/docs). Breaking changes to the API will be announced in the [Breaking Changes log](https://github.com/CircleCI-Public/api-preview-docs/blob/master/docs/breaking.md).
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package circleci

// Message from CircleCI execution platform.
type JobDetailsMessages struct {
	// Message type.
	Type_ string `json:"type"`
	// Information describing message.
	Message string `json:"message"`
	// Value describing the reason for message to be added to the job.
	Reason string `json:"reason,omitempty"`
}