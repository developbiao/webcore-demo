package contract

import "net/http"

const KernelKey = "web:kernel"

// Kernel interface provider framwork kernel structure
type Kernel interface {
	// Represet net/http, real is gin.Engine
	HttpEngine() http.Handler
}
