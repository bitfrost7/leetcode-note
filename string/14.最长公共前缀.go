package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			return ""
		}
	}
	return prefix
}
func lcp(s1, s2 string) string {
	var i int
	for i < len(s1) && i < len(s2) && s1[i] == s2[i] {
		i++
	}
	return s1[:i]
}
