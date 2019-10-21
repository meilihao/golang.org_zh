// https://github.com/davecheney/godoc2md
// base: https://github.com/golang/tools/blob/master/godoc/static

package main

import (
	"bytes"
	"strings"
)

var (
	funcs = map[string]interface{}{
		"comment_md":   commentMdFunc,
		"md":           mdFunc,
		"pre":          preFunc,
		"kebab":        kebabFunc,
		"bitscape":     bitscapeFunc, //Escape [] for bitbucket confusion
		"trim_prefix":  strings.TrimPrefix,
		"appendSrcUrl": appendSrcUrl,
	}
)

func commentMdFunc(comment string) string {
	var buf bytes.Buffer
	ToMD(&buf, comment)
	return buf.String()
}

func mdFunc(text string) string {
	text = strings.Replace(text, "*", "\\*", -1)
	text = strings.Replace(text, "_", "\\_", -1)
	return text
}

func preFunc(text string) string {
	return "``` go\n" + text + "\n```"
}

func kebabFunc(text string) string {
	s := strings.Replace(strings.ToLower(text), " ", "-", -1)
	s = strings.Replace(s, ".", "-", -1)
	s = strings.Replace(s, "\\*", "42", -1)
	return s
}

func bitscapeFunc(text string) string {
	s := strings.Replace(text, "[", "\\[", -1)
	s = strings.Replace(s, "]", "\\]", -1)
	return s
}

func appendSrcUrl(s string) string {
	return strings.TrimSuffix(pkgUrl, "/pkg/") + s
}
