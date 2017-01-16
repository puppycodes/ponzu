package api

import "net/http"

// Run adds Handlers to default http listener for API
func Run() {
	http.HandleFunc("/api/contents", Record(CORS(contentsHandler)))

	http.HandleFunc("/api/content", Record(CORS(contentHandler)))

	http.HandleFunc("/api/content/external", Record(externalContentHandler))
}
