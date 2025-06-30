package main

import (
	"context"
	"net/http"
)

func exposeHeader(url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Server", "Apache/2.4.1 (Unix)")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// GetSomethingWithContext is a dummy function to illustrate the anti-pattern.
func GetSomethingWithContext(name string, ctx context.Context) string {
	// do something with context
	log.Printf("Context value for key 'somekey': %v", ctx.Value("somekey"))
	return "something for " + name
}
