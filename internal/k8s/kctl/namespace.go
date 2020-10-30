package kctl

func nameSpaceExists(config Config, nameSpace string) bool {
	cmd, args := prepareCommand(config, "get", "namespace", nameSpace, "--no-headers")
	err := execute(config, cmd, args...)
	if err != nil {
		return false
	}
	return true
}

func createNameSpace(config Config, nameSpace string) error {
	if !nameSpaceExists(config, nameSpace) {
		cmd, args := prepareCommand(config, create, "namespace", nameSpace)
		return execute(config, cmd, args...)
	}
	return nil
}

func deleteNameSpace(config Config, nameSpace string) error {
	if nameSpaceExists(config, nameSpace) {
		cmd, args := prepareCommand(config, delete, "namespace", nameSpace)
		return execute(config, cmd, args...)
	}
	return nil
}
