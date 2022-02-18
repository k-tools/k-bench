package internal

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1")

func (c Client) CreatePo(ctx context.Context, p StructPodsOptions) error {

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
			Labels:    map[string]string{"name": p.Name},
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
	listOpts.LabelSelector = "name=" + p.Name

	w, err := c.Clientset.CoreV1().Pods(p.Namespace).Watch(ctx, listOpts)
	if err != nil {
		return err
	}

	ch := w.ResultChan()

	for {
		event := <-ch
		po := event.Object.(*v1.Pod)
		if po.Status.Phase == "Pending" {
		} else {
			fmt.Printf("Pod %s created \n", po.Name)
			break
		}
	}

	return nil
}
