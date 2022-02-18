package internal

import (
	"context"
	"fmt"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Client) GetEvent(ctx context.Context, p StructPodsOptions) error {

	listOpts := metav1.ListOptions{}
	w, err := c.Clientset.CoreV1().Events(p.Namespace).Watch(ctx, listOpts)
	if err != nil {
		return err
	}
	ch := w.ResultChan()
	go func(){
		for {
			event := <-ch
			e := event.Object.(*v1.Event)
			if e.Reason == "Pulled" {
				if strings.Contains(e.InvolvedObject.Name, p.Name ) {
					fmt.Printf("%s: %v \n", e.Reason, e.Message)
				}
			}
		}
	}() 
    time.Sleep(1 * time.Second)
	return nil
}
