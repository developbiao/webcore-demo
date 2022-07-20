package contract

import "net/http"

const KernelKey = "web:kernel"

// Kernel interface provider framwork kernel structure
type Kernel interface {
	HttpEngine() http.Handler
}
