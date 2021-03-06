package minimux

import (
	"net/http"
	"path"
	"strings"
)

//MiniMux Just another router
type MiniMux struct {
	pathNotFound   http.Handler
	sortedRegPaths []string
	routes         map[string]http.Handler
}

//New creates empty routers
func New() *MiniMux {
	return &MiniMux{
		pathNotFound:   http.NotFoundHandler(),
		sortedRegPaths: []string{},
		routes:         make(map[string]http.Handler),
	}
}

func (m *MiniMux) matchPath(methodwpath string) (string, bool) {

	// Get URL without parameters
	if i := strings.Index(methodwpath, "?"); i != -1 {
		methodwpath = methodwpath[:i]
	}

	// Clean URL
	methodwpath = path.Clean(methodwpath)

	// Direct match first
	_, ok := m.routes[methodwpath]
	if ok {
		return methodwpath, true
	}

	// ... same old... same old ...
	for _, path := range m.sortedRegPaths {
		if strings.HasPrefix(methodwpath+"/", path) && path[len(path)-1] == '/' {
			return path, true
		}
	}

	return "", false
}

func (m *MiniMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodWPath := r.Method + r.URL.Path

	path, ok := m.matchPath(methodWPath)

	if !ok {
		m.pathNotFound.ServeHTTP(w, r)
		return
	}

	m.routes[path].ServeHTTP(w, r)

}

//Head requests handler registrator
func (m *MiniMux) Head(pat string, h http.Handler) {
	m.add("HEAD", pat, h)
}

//Get requests handler registrator
func (m *MiniMux) Get(path string, handler http.Handler) {
	m.add("GET", path, handler)
}

//Post requests handler registrator
func (m *MiniMux) Post(path string, handler http.Handler) {
	m.add("POST", path, handler)
}

//Put requests handler registrator
func (m *MiniMux) Put(pat string, h http.Handler) {
	m.add("PUT", pat, h)
}

//Delete requests handler registrator
func (m *MiniMux) Delete(pat string, h http.Handler) {
	m.add("DELETE", pat, h)
}

//Options requests handler registrator
func (m *MiniMux) Options(pat string, h http.Handler) {
	m.add("OPTIONS", pat, h)
}

//Patch requests handler registrator.
func (m *MiniMux) Patch(pat string, h http.Handler) {
	m.add("PATCH", pat, h)
}

func (m *MiniMux) add(method string, path string, handler http.Handler) {
	if path == "" {
		panic("error: invalid path")
	}

	if handler == nil {
		panic("error: empty handler")
	}

	methodWPath := method + path
	_, ok := m.routes[methodWPath]
	if ok {
		panic("error: already registered path " + path)
	}

	m.routes[methodWPath] = handler

	m.sortedRegPaths = addToSortedList(m.sortedRegPaths, methodWPath)
}

//addToSortedList adds a string in the sorted by lentgh array
func addToSortedList(strStore []string, strItem string) []string {

	strLen := len(strItem)
	strPlace := 0
	for id, item := range strStore {
		strPlace = id
		if strLen > len(item) {
			break
		}
		// If 'strItem' is shortest then it has to go at very end
		strPlace++
	}

	// Expand slice
	strStore = append(strStore, "")
	// Shift
	copy(strStore[strPlace+1:], strStore[strPlace:])
	strStore[strPlace] = strItem

	return strStore
}
