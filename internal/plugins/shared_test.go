package plugins

func getRootTestData() string {
	return "../../local_repo/core/"
}

func joinWithRootData(fileURI string) string {
	return getRootTestData() + fileURI
}
