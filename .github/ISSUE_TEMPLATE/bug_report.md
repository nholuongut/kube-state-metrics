---
name: Bug report
about: Report a bug encountered while running kube-state-metrics
title: ''
labels: kind/bug
assignees: ''

---

<!-- Please use this template while reporting a bug and provide as much info as possible. Not doing so may result in your bug not being addressed in a timely manner. Thanks!

If the matter is security related, please disclose it privately see https://github.com/nholuongut/kube-state-metrics/blob/main/SECURITY.md
-->

**What happened**:

**What you expected to happen**:

**How to reproduce it (as minimally and precisely as possible)**:

```bash
# An example: https://github.com/nholuongut/kube-state-metrics/issues/2223#issuecomment-1792850276
minikube start
...
go run main.go --custom-resource-state-only --custom-resource-state-config-file ksm-2223/custom-resource-config-file.yaml --kubeconfig ~/.kube/config
```

**Anything else we need to know?**:

**Environment**:

* kube-state-metrics version:
* Kubernetes version (use `kubectl version`):
* Cloud provider or hardware configuration:
* Other info:
