/*
 * MERS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type TrialsPostRequest struct {

	Datetime time.Time `json:"datetime,omitempty"`

	Subject TrialsPostRequestSubject `json:"subject"`

	Ratings []TrialsPostRequestRatingsInner `json:"ratings"`
}

