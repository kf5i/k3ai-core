package kctl

func nameSpaceExists(config Config, nameSpace string) bool {
	err := execute(config, k3sExec, kubectl, "get", "namespace", nameSpace, "--no-headers")
	if err != nil {
		return false
	}
	return true
}

func createNameSpace(config Config, nameSpace string) error {
	if !nameSpaceExists(config, nameSpace) {
		return execute(config, k3sExec, kubectl, create, "namespace", nameSpace)
	}
	return nil
}

func deleteNameSpace(config Config, nameSpace string) error {
	if nameSpaceExists(config, nameSpace) {
		return execute(config, k3sExec, kubectl, delete, "namespace", nameSpace)
	}
	return nil
}
