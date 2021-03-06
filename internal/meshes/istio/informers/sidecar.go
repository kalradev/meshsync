package informers

import (
	"log"

	"github.com/layer5io/meshsync/internal/model"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
)

func (i *Istio) SidecarInformer() cache.SharedIndexInformer {
	// get informer
	SidecarInformer := i.client.GetSidecarInformer().Informer()

	// register event handlers
	SidecarInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				Sidecar := obj.(*v1beta1.Sidecar)
				log.Printf("Sidecar Named: %s - added", Sidecar.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					Sidecar.TypeMeta,
					Sidecar.ObjectMeta,
					Sidecar.Spec,
					Sidecar.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing Sidecar")
				}
			},
			UpdateFunc: func(new interface{}, old interface{}) {
				Sidecar := new.(*v1beta1.Sidecar)
				log.Printf("Sidecar Named: %s - updated", Sidecar.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					Sidecar.TypeMeta,
					Sidecar.ObjectMeta,
					Sidecar.Spec,
					Sidecar.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing Sidecar")
				}
			},
			DeleteFunc: func(obj interface{}) {
				Sidecar := obj.(*v1beta1.Sidecar)
				log.Printf("Sidecar Named: %s - deleted", Sidecar.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					Sidecar.TypeMeta,
					Sidecar.ObjectMeta,
					Sidecar.Spec,
					Sidecar.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing Sidecar")
				}
			},
		},
	)

	return SidecarInformer
}
