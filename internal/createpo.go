package internal

import (
	"context"

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

	opts := metav1.CreateOptions{}
	_, err := c.Clientset.CoreV1().Pods(p.Namespace).Create(ctx, pod, opts)

	return err
}
