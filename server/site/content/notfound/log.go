package notfound

import (
	_log "log"
	"net/url"
)

// needs work

func log(url *url.URL) {
	// todo
	// - only log if option verbose is set
	// - only if not on on ignore list
	// - have a list of urls that will get your IP banned if you visit them, e.g. those urls that hackers call in an attempt to get access
	_log.Printf("404 for url %q", url)
}
