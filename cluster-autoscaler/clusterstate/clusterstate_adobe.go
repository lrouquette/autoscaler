package clusterstate

func (csr *ClusterStateRegistry) ClearUnregisteredNodes() {
	csr.Lock()
	defer csr.Unlock()
	csr.unregisteredNodes = make(map[string]UnregisteredNode)
}
