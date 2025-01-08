// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type CreateDestinationDestinationsPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Destination *components.Destination
}

func (o *CreateDestinationDestinationsPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateDestinationDestinationsPostResponse) GetDestination() *components.Destination {
	if o == nil {
		return nil
	}
	return o.Destination
}
