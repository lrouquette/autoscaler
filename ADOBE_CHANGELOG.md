# About

We maintain a set of patches on top of upstream K8s autoscaler. They need to be
semantically versioned as well so we understand, for example, that we're
upgrading a fix version of autoscaler vs. a fix version of our patches.

The scheme is quite simple. The product is first, followed by the upstream
version, followed by our version:

```
Product is cluster-autoscaler or vertical-pod-autoscaler or addon-resizer
K stands for K8s upstream version
A stands for Adobe's changes

{Product}-{K major}.{K minor}.{K fix}-{A major}.{A minor}.{A fix}-adobe
```

# Log

## cluster-autoscaler-1.20.0-1.0.0-adobe

- on scale-down, evict DaemonSet pods after all the other pods
