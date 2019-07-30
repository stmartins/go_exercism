package twofer

func ShareWith(name string) string {
	nameToWrite := name
	if name == "" {
		nameToWrite = "you"
	}
	return "One for " + nameToWrite + ", one for me."
}
