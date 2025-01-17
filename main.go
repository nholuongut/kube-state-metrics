/*
Copyright 2015 Nho Luong DevOps All rights reserved.

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

package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"k8s.io/kube-state-metrics/v2/internal"
	"k8s.io/kube-state-metrics/v2/pkg/options"
)

func main() {
	opts := options.NewOptions()
	cmd := options.InitCommand
	cmd.Run = func(_ *cobra.Command, _ []string) {
		internal.RunKubeStateMetricsWrapper(opts)
	}
	opts.AddFlags(cmd)
	if err := opts.Parse(); err != nil {
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	if err := opts.Validate(); err != nil {
		klog.ErrorS(err, "Validating options error")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
}
