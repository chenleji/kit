package handling

import (
	"context"
	"time"

	"github.com/chenleji/kit/endpoint"

	"github.com/chenleji/kit/examples/shipping/cargo"
	"github.com/chenleji/kit/examples/shipping/location"
	"github.com/chenleji/kit/examples/shipping/voyage"
)

type registerIncidentRequest struct {
	ID             cargo.TrackingID
	Location       location.UNLocode
	Voyage         voyage.Number
	EventType      cargo.HandlingEventType
	CompletionTime time.Time
}

type registerIncidentResponse struct {
	Err error `json:"error,omitempty"`
}

func (r registerIncidentResponse) error() error { return r.Err }

func makeRegisterIncidentEndpoint(hs Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerIncidentRequest)
		err := hs.RegisterHandlingEvent(req.CompletionTime, req.ID, req.Voyage, req.Location, req.EventType)
		return registerIncidentResponse{Err: err}, nil
	}
}
