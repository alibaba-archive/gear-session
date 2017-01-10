// https://github.com/gorilla/sessions

package session

import (
	"encoding/base32"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Store is an interface for custom session stores.
//
// See CookieStore and FilesystemStore for examples.
type Store interface {
	// Get should return a cached session.
	Get(sid string) (val string, error)

	// New should create and return a new session.
	//
	// Note that New should never return a nil session, even in the case of
	// an error if using the Registry infrastructure to cache the session.
	Set(sid, val string) error

	// Save should persist session to the underlying store implementation.
	Destroy(sid string) error
}
