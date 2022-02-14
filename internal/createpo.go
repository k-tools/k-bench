package internal

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Client) CreatePo(ctx context.Context, p StructPodsOptions) error {

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []v1.Container{
				{
					Name:  p.Name,
					Image: p.Image,
				},
			},
		},
	}

	createOpts := metav1.CreateOptions{}
	_, err := c.Clientset.CoreV1().Pods(p.Namespace).Create(ctx, pod, createOpts)
	if err != nil {
		return err
	}

	listOpts := metav1.ListOptions{}
	//listOpts.LabelSelector = "name=" + p.Name

	/*w, err := c.Clientset.CoreV1().Pods(p.Namespace).Watch(ctx, listOpts)
	if err != nil {
		return err
	}

	ch := w.ResultChan()

	for {
		event := <-ch
		po := event.Object.(*v1.Pod)
		fmt.Printf("debug %s \n", po.Status.Phase)
		if po.Status.Phase == "Pending" {
			fmt.Printf("Pod Phase: %v \n", po.Status.Phase)
			fmt.Printf("Pod Creation: %v \n", po.CreationTimestamp)
			fmt.Printf("Pod Start: %v \n", po.Status.StartTime)
		} else {
			fmt.Printf("Pod Phase: %v \n", po.Status.Phase)
			fmt.Printf("Pod Creation: %v \n", po.CreationTimestamp)
			fmt.Printf("Pod Start: %v \n", po.Status.StartTime)
			break
		}
	}*/

	w, err := c.Clientset.CoreV1().Events(p.Namespace).Watch(ctx, listOpts)
	if err != nil {
		return err
	}

	ch := w.ResultChan()

	for {
		event := <-ch
		e := event.Object.(*v1.Event)
		if e.Reason == "Pulled" {
			fmt.Printf("%s: %v \n", e.Reason, e.Message)
			//break
		}

	}

	return nil
}
