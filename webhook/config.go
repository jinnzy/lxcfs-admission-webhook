/*
Copyright 2018 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"crypto/tls"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/klog/v2"
)

const (
	admissionWebhookAnnotationMutateKey = "lxcfs-admission-webhook.aliyun.com/mutate"
	admissionWebhookAnnotationStatusKey = "lxcfs-admission-webhook.aliyun.com/status"
)

var (
	ignoredNamespaces = []string{
		metav1.NamespaceSystem,
		metav1.NamespacePublic,
	}
)

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

var volumeMountsTemplate = []corev1.VolumeMount{

	{
		Name:      "lxcfs-proc-cpuinfo",
		MountPath: "/proc/cpuinfo",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-meminfo",
		MountPath: "/proc/meminfo",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-diskstats",
		MountPath: "/proc/diskstats",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-stat",
		MountPath: "/proc/stat",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-swaps",
		MountPath: "/proc/swaps",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-uptime",
		MountPath: "/proc/uptime",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-loadavg",
		MountPath: "/proc/loadavg",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-sys-devices-system-cpu-online",
		MountPath: "/sys/devices/system/cpu/online",
		ReadOnly:  true,
	},
}

var volumesTemplate = []corev1.Volume{
	{
		Name: "lxcfs-proc-cpuinfo",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/cpuinfo",
			},
		},
	},
	{
		Name: "lxcfs-proc-diskstats",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/diskstats",
			},
		},
	},
	{
		Name: "lxcfs-proc-meminfo",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/meminfo",
			},
		},
	},
	{
		Name: "lxcfs-proc-stat",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/stat",
			},
		},
	},
	{
		Name: "lxcfs-proc-swaps",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/swaps",
			},
		},
	},
	{
		Name: "lxcfs-proc-uptime",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/uptime",
			},
		},
	},
	{
		Name: "lxcfs-proc-loadavg",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/loadavg",
			},
		},
	},
	{
		Name: "lxcfs-sys-devices-system-cpu-online",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/sys/devices/system/cpu/online",
			},
		},
	},
}

// Config contains the server (the webhook) cert and key.
type Config struct {
	CertFile string
	KeyFile  string
}

func configTLS(config Config) *tls.Config {
	sCert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	if err != nil {
		klog.Fatal(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{sCert},
		// TODO: uses mutual tls after we agree on what cert the apiserver should use.
		// ClientAuth:   tls.RequireAndVerifyClientCert,
	}
}
