package informers

import (
	"log"

	"github.com/layer5io/meshsync/internal/model"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	"k8s.io/client-go/tools/cache"
)

func (i *Istio) RequestAuthenticationInformer() cache.SharedIndexInformer {
	// get informer
	RequestAuthenticationInformer := i.client.GetRequestAuthenticationInformer().Informer()

	// register event handlers
	RequestAuthenticationInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				RequestAuthentication := obj.(*v1beta1.RequestAuthentication)
				log.Printf("RequestAuthentication Named: %s - added", RequestAuthentication.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					RequestAuthentication.TypeMeta,
					RequestAuthentication.ObjectMeta,
					RequestAuthentication.Spec,
					RequestAuthentication.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing RequestAuthentication")
				}
			},
			UpdateFunc: func(new interface{}, old interface{}) {
				RequestAuthentication := new.(*v1beta1.RequestAuthentication)
				log.Printf("RequestAuthentication Named: %s - updated", RequestAuthentication.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					RequestAuthentication.TypeMeta,
					RequestAuthentication.ObjectMeta,
					RequestAuthentication.Spec,
					RequestAuthentication.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing RequestAuthentication")
				}
			},
			DeleteFunc: func(obj interface{}) {
				RequestAuthentication := obj.(*v1beta1.RequestAuthentication)
				log.Printf("RequestAuthentication Named: %s - deleted", RequestAuthentication.Name)
				err := i.broker.Publish(Subject, model.ConvModelObject(
					RequestAuthentication.TypeMeta,
					RequestAuthentication.ObjectMeta,
					RequestAuthentication.Spec,
					RequestAuthentication.Status,
				))
				if err != nil {
					log.Println("NATS: Error publishing RequestAuthentication")
				}
			},
		},
	)

	return RequestAuthenticationInformer
}
