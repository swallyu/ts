package ts

type AntMatch struct {
	pathSeparator  string
	wildcard_chars *[]string
	caseSensitive  bool
}

func (a *AntMatch) IsPattern(path string) bool {
	uriVar := false
	i := 0
	for i < len(path) {
		ch := path[i]
		if ch == '*' || ch == '?' {
			return true
		}
		if ch == '{' {
			uriVar = true
			continue
		}
		if ch == '}' && uriVar {
			return true
		}
	}
	return false
}

func (a *AntMatch) Match(pattern string, path string) bool {
	return a.doMatch(path, path, true, nil)
}

func (a *AntMatch) MatchStart(pattern string, path string) bool {
	return a.doMatch(path, path, true, nil)
}

func (a *AntMatch) doMatch(pattern string, path string, fullMatch bool, uriTplVari map[string]string) bool {
	if (path[0] == '/') != (pattern[0] == '/') {
		return false
	}

	pattDirs := tokenizePattern(pattern)
	if fullMatch && a.caseSensitive && !isPotentialMatch(path, pattDirs) {
		return false
	}
	return false
}

func isPotentialMatch(path string, pattDirs []string) bool {

	return false
}

func tokenizePattern(pattern string) []string {
	return []string{}
}
