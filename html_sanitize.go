package main

import (
	"sync"

	"github.com/microcosm-cc/bluemonday"
)

var sanitizeHtmlOnce sync.Once

func sanitizeHtml(inHtml string) (string, error) {
	var p *bluemonday.Policy

	sanitizeHtmlOnce.Do(func() {
		p = bluemonday.UGCPolicy()
	})

	outHtml := p.Sanitize(inHtml)

	return outHtml, nil
}
