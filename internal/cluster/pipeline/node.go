package pipeline

import (
	"log"

	"github.com/layer5io/meshsync/internal/model"
	broker "github.com/layer5io/meshsync/pkg/broker"
	discovery "github.com/layer5io/meshsync/pkg/discovery"
	"github.com/myntra/pipeline"
)

// Node will implement step interface for Nodes
type Node struct {
	pipeline.StepContext
	client *discovery.Client
	broker broker.Handler
}

// NewNode - constructor
func NewNode(client *discovery.Client, broker broker.Handler) *Node {
	return &Node{
		client: client,
		broker: broker,
	}
}

// Exec - step interface
func (n *Node) Exec(request *pipeline.Request) *pipeline.Result {
	// it will contain a pipeline to run
	log.Println("Node Discovery Started")

	// get nodes
	nodes, err := n.client.ListNodes()
	if err != nil {
		return &pipeline.Result{
			Error: err,
		}
	}

	// processing
	for _, node := range nodes {
		// publishing discovered node
		err := n.broker.Publish(Subject, model.ConvModelObject(
			node.TypeMeta,
			node.ObjectMeta,
			node.Spec,
			node.Status,
		))
		if err != nil {
			log.Printf("Error publishing node named %s", node.Name)
		} else {
			log.Printf("Published node named %s", node.Name)
		}
	}

	// no data is feeded to future steps or stages
	return &pipeline.Result{
		Error: nil,
	}
}

// Cancel - step interface
func (n *Node) Cancel() error {
	n.Status("cancel step")
	return nil
}
