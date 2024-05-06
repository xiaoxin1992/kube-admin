package k8s

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

type Containers struct {
	Name            string        `json:"name"`
	Image           string        `json:"image"`
	ImagePullPolicy v1.PullPolicy `json:"imagePullPolicy"`
}

type Status struct {
	Phase        v1.PodPhase    `json:"phase"`
	StartTime    string         `json:"startTime"`
	Ready        string         `json:"ready"`
	QosClass     v1.PodQOSClass `json:"qosClass"`
	RestartCount int            `json:"restartCount"`
}
type Pod struct {
	NodeName      string            `json:"nodeName"`
	Namespace     string            `json:"namespace"`
	Name          string            `json:"name"`
	Labels        map[string]string `json:"labels"`
	RestartPolicy v1.RestartPolicy  `json:"restartPolicy"`
	Containers    []Containers      `json:"Containers"`
	Status        Status            `json:"status"`
}

func PodAnalysis(pod *v1.Pod) *Pod {
	p := Pod{
		NodeName:      pod.Spec.NodeName,
		Namespace:     pod.Namespace,
		Name:          pod.Name,
		Labels:        pod.GetLabels(),
		RestartPolicy: pod.Spec.RestartPolicy,
		Containers:    make([]Containers, 0),
	}
	for _, container := range pod.Spec.Containers {
		c := Containers{
			Name:            container.Name,
			Image:           container.Image,
			ImagePullPolicy: container.ImagePullPolicy,
		}
		p.Containers = append(p.Containers, c)
	}
	phase := pod.Status.Phase
	if pod.GetDeletionTimestamp() != nil {
		phase = "Terminating"
	}
	s := Status{
		Phase:     phase,
		StartTime: pod.GetCreationTimestamp().Format("2006-01-02 15:04:05"),
		QosClass:  pod.Status.QOSClass,
	}
	ready := 0
	total := 0
	restartCount := 0
	for _, state := range pod.Status.ContainerStatuses {
		total += 1
		if state.Ready {
			ready += 1
		}
		restartCount = int(state.RestartCount)
	}
	s.RestartCount = restartCount
	s.Ready = fmt.Sprintf("%d/%d", ready, total)
	p.Status = s
	return &p
}
