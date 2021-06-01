/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pipeline

// CommonError - Generic error object. Deprecated: use Error schema instead.
type CommonError struct {

	// HTTP status code
	Code int32 `json:"code,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`

	// Error message
	Error string `json:"error,omitempty"`
}