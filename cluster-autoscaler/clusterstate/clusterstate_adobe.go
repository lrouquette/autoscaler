package clusterstate

// ClearUnregisteredNodes clears the unregistered node cache
func (csr *ClusterStateRegistry) ClearUnregisteredNodes() {
	csr.Lock()
	defer csr.Unlock()
	csr.unregisteredNodes = make(map[string]UnregisteredNode)
}
