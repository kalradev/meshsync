package pipeline

import (
	"log"

	"github.com/layer5io/meshsync/internal/model"
	broker "github.com/layer5io/meshsync/pkg/broker"
	discovery "github.com/layer5io/meshsync/pkg/discovery"
	"github.com/myntra/pipeline"
)

// Sidecar will implement step interface for Sidecar
type Sidecar struct {
	pipeline.StepContext
	// clients
	client *discovery.Client
	broker broker.Handler
}

// NewSidecar - constructor
func NewSidecar(client *discovery.Client, broker broker.Handler) *Sidecar {
	return &Sidecar{
		client: client,
		broker: broker,
	}
}

// Exec - step interface
func (s *Sidecar) Exec(request *pipeline.Request) *pipeline.Result {
	// it will contain a pipeline to run
	log.Println("Sidecar Discovery Started")

	for _, namespace := range Namespaces {
		sidecars, err := s.client.ListSidecars(namespace)
		if err != nil {
			return &pipeline.Result{
				Error: err,
			}
		}

		// processing
		for _, sidecar := range sidecars {
			// publishing discovered sidecar
			err := s.broker.Publish(Subject, model.ConvModelObject(
				sidecar.TypeMeta,
				sidecar.ObjectMeta,
				sidecar.Spec,
				sidecar.Status,
			))
			if err != nil {
				log.Printf("Error publishing sidecar named %s", sidecar.Name)
			} else {
				log.Printf("Published sidecar named %s", sidecar.Name)
			}
		}
	}

	// no data is feeded to future steps or stages
	return &pipeline.Result{
		Error: nil,
	}
}

// Cancel - step interface
func (s *Sidecar) Cancel() error {
	s.Status("cancel step")
	return nil
}
