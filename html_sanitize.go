package main

import (
	"sync"

	"github.com/microcosm-cc/bluemonday"
)

var sanitizeHtmlOnce sync.Once
var p *bluemonday.Policy

func sanitizeHtml(inHtml string) (string, error) {

	sanitizeHtmlOnce.Do(func() {
		p = bluemonday.UGCPolicy()
	})

	outHtml := p.Sanitize(inHtml)

	return outHtml, nil
}
