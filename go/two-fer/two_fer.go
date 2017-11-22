package twofer

func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	out := "One for " + name + ", one for me."
	return out
}
