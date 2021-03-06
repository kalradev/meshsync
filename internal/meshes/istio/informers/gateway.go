package informers

import (
	"log"

	"github.com/layer5io/meshsync/internal/model"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
)

func (i *Istio) GatewayInformer() cache.SharedIndexInformer {
	// get informer
	GatewayInformer := i.client.GetGatewayInformer().Informer()

	// register event handlers
	GatewayInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				Gateway := obj.(*v1beta1.Gateway)
				log.Printf("Gateway Named: %s - added", Gateway.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					Gateway.TypeMeta,
					Gateway.ObjectMeta,
					Gateway.Spec,
					Gateway.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing Gateway")
				}
			},
			UpdateFunc: func(new interface{}, old interface{}) {
				Gateway := new.(*v1beta1.Gateway)
				log.Printf("Gateway Named: %s - updated", Gateway.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					Gateway.TypeMeta,
					Gateway.ObjectMeta,
					Gateway.Spec,
					Gateway.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing Gateway")
				}
			},
			DeleteFunc: func(obj interface{}) {
				Gateway := obj.(*v1beta1.Gateway)
				log.Printf("Gateway Named: %s - deleted", Gateway.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					Gateway.TypeMeta,
					Gateway.ObjectMeta,
					Gateway.Spec,
					Gateway.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing Gateway")
				}
			},
		},
	)

	return GatewayInformer
}
