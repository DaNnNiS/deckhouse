package kube_events_manager

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/romana/rlog"
	"gopkg.in/satori/go.uuid.v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"

	"github.com/deckhouse/deckhouse/antiopa/kube"
)

var (
	KubeEventCh chan string
)

type InformerType int

const (
	OnAdd InformerType = iota
	OnUpdate
	OnDelete
)

type KubeEventsManager interface {
	Run(informerType InformerType, kind, namespace string, labelSelector *metav1.LabelSelector, jqFilter string) (string, error)
	Stop(configId string) error
}

type MainKubeEventsManager struct {
	KubeEventsInformersByConfigId map[string]*KubeEventsInformer
}

func NewMainKubeEventsManager() *MainKubeEventsManager {
	em := &MainKubeEventsManager{}
	em.KubeEventsInformersByConfigId = make(map[string]*KubeEventsInformer)
	return em
}

func Init() (KubeEventsManager, error) {
	em := NewMainKubeEventsManager()
	KubeEventCh = make(chan string, 1)
	return em, nil
}

func (em *MainKubeEventsManager) Run(informerType InformerType, kind, namespace string, labelSelector *metav1.LabelSelector, jqFilter string) (string, error) {
	kubeEventsInformer, err := em.addKubeEventsInformer(kind, namespace, labelSelector, jqFilter, func(kubeEventsInformer *KubeEventsInformer) cache.ResourceEventHandlerFuncs {
		return cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				err := kubeEventsInformer.HandleKubeEvent(obj, runtimeObjectMd5(obj, jqFilter), informerType == OnAdd)
				if err != nil {
					rlog.Error("Kube events manager: %s", err)
				}
			},
			UpdateFunc: func(_ interface{}, obj interface{}) {
				err := kubeEventsInformer.HandleKubeEvent(obj, runtimeObjectMd5(obj, jqFilter), informerType == OnUpdate)
				if err != nil {
					rlog.Error("Kube events manager: %s", err)
				}
			},
			DeleteFunc: func(obj interface{}) {
				err := kubeEventsInformer.HandleKubeEvent(obj, "", informerType == OnDelete)
				if err != nil {
					rlog.Error("Kube events manager: %s", err)
				}
			},
		}
	})

	if err != nil {
		return "", err
	}

	go kubeEventsInformer.Run()

	return kubeEventsInformer.ConfigId, nil
}

func (em *MainKubeEventsManager) addKubeEventsInformer(kind, namespace string, labelSelector *metav1.LabelSelector, jqFilter string, resourceEventHandlerFuncs func(kubeEventsInformer *KubeEventsInformer) cache.ResourceEventHandlerFuncs) (*KubeEventsInformer, error) {
	kubeEventsInformer := NewKubeEventsInformer()

	listOptions := &metav1.ListOptions{}

	if labelSelector != nil {
		listOptions.LabelSelector = labelSelector.String()
	}

	configMaps, _ := kube.KubernetesClient.CoreV1().ConfigMaps(namespace).List(*listOptions)
	for _, configMap := range configMaps.Items {
		configMapId := fmt.Sprintf("%s-%s", configMap.Name, configMap.Namespace)
		kubeEventsInformer.Checksum[configMapId] = runtimeObjectMd5(configMap, jqFilter)
	}

	optionsModifier := func(options *metav1.ListOptions) {
		if labelSelector != nil {
			*options = *listOptions
		}
	}

	restKubeClient := kube.KubernetesClient.CoreV1().RESTClient()
	lw := cache.NewFilteredListWatchFromClient(restKubeClient, kind, namespace, optionsModifier)

	kubeEventsInformer.SharedInformer = cache.NewSharedInformer(lw, &v1.ConfigMap{}, time.Duration(15)*time.Second)
	kubeEventsInformer.SharedInformer.AddEventHandler(resourceEventHandlerFuncs(kubeEventsInformer))
	kubeEventsInformer.ConfigId = uuid.NewV4().String()

	em.KubeEventsInformersByConfigId[kubeEventsInformer.ConfigId] = kubeEventsInformer

	return kubeEventsInformer, nil
}

func runtimeObjectMd5(obj interface{}, _ string) string {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	h := md5.New()
	io.WriteString(h, string(data))

	return string(h.Sum(nil))
}

func (em *MainKubeEventsManager) Stop(configId string) error {
	kubeEventsInformer, ok := em.KubeEventsInformersByConfigId[configId]
	if ok {
		kubeEventsInformer.Stop()
	} else {
		rlog.Errorf("Kube events informer '%s' not found!", configId)
	}
	return nil
}

type KubeEventsInformer struct {
	ConfigId           string
	Checksum           map[string]string
	SharedInformer     cache.SharedInformer
	SharedInformerStop chan struct{}
}

func NewKubeEventsInformer() *KubeEventsInformer {
	kubeEventsInformer := &KubeEventsInformer{}
	kubeEventsInformer.Checksum = make(map[string]string)
	kubeEventsInformer.SharedInformerStop = make(chan struct{}, 1)
	return kubeEventsInformer
}

func (ei *KubeEventsInformer) HandleKubeEvent(obj interface{}, newChecksum string, sendSignal bool) error {
	objectId, err := runtimeObjectId(obj.(runtime.Object))
	if err != nil {
		return fmt.Errorf("failed to get object id: %s", err)
	}

	if ei.Checksum[objectId] != newChecksum {
		ei.Checksum[objectId] = newChecksum

		if sendSignal {
			KubeEventCh <- ei.ConfigId
		}
	}

	return nil
}

func runtimeObjectId(obj runtime.Object) (string, error) {
	accessor := meta.NewAccessor()

	name, err := accessor.Name(obj)
	if err != nil {
		return "", err
	}

	namespace, err := accessor.Namespace(obj)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", name, namespace), nil
}

func (ei *KubeEventsInformer) Run() {
	ei.SharedInformer.Run(ei.SharedInformerStop)
}

func (ei *KubeEventsInformer) Stop() {
	ei.SharedInformerStop <- struct{}{}
}
