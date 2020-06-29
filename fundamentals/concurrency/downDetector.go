package concurrency

// WebsiteChecker : a function which returns whether a website is reachable
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites : Check wether a list of URLs are reachable or not
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
