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

## cluster-autoscaler-1.24.2-1.5.0-adobe
- built from CA v1.24.2

## cluster-autoscaler-1.24.1-1.5.0-adobe
- built from CA v1.24.1

## cluster-autoscaler-1.23.1-1.5.0-adobe
- skip deletion of unregistered nodes if there are more than 50% of discovered nodes

## cluster-autoscaler-1.23.1-1.4.0-adobe
- built from CA v1.23.1
- backport https://github.com/kubernetes/autoscaler/pull/4970
- remove code from 1.0.0-adobe and 1.2.0-adobe

## cluster-autoscaler-1.22.3-1.3.0-adobe
- built from CA v1.22.3
- backport https://github.com/kubernetes/autoscaler/pull/4970

## cluster-autoscaler-1.22.2-1.2.1-adobe
- built from CA v1.22.2

## cluster-autoscaler-1.21.2-1.2.1-adobe
- built from CA v1.21.2

## cluster-autoscaler-1.20.1-1.2.1-adobe
- also use `topology.kubernetes.io/zone` when evaluating topology spread constraints

## cluster-autoscaler-1.20.1-1.2.0-adobe
- improve scale out time when using topology spread constraints

## cluster-autoscaler-1.20.1-1.1.0-adobe
- upgraded ca

## cluster-autoscaler-1.20.0-1.1.0-adobe
- fix uneven scaling with topology contraints
- backport https://github.com/kubernetes/autoscaler/pull/3883

## cluster-autoscaler-1.20.0-1.0.0-adobe
- on scale-down, evict DaemonSet pods after all the other pods

# Procedure

```bash
# Assuming cluster-autoscaler-1.21-adobe is the latest "Adobe" release, built
# "on top" of of 1.21.0, and we want to upgrade to 1.21.2
git checkout cluster-autoscaler-1.21-adobe
# Reset to the upstream version we originally built on
git reset cluster-autoscaler-1.21.0
# Stash "Adobe" changes
git stash
# Get the latest stuff from upstream
git fetch upstream
# Update the branch to upstream 1.21.2
git merge -s recursive -X theirs cluster-autoscaler-1.21.2
# Apply "Adobe" changes, on top of 1.21.2
git stash pop
# Resolve conflicts if needed, then git add/git commit all the files (including this one)
```
