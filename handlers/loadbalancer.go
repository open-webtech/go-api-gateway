package handlers

import (
	"net/http"

	"github.com/open-webtech/go-api-gateway/services"
	loadbalancer "github.com/open-webtech/go-load-balancer"
	"github.com/open-webtech/go-load-balancer/iterator"
)

func LoadBalancer(iter iterator.Iterator, sc services.Context) *loadbalancer.LoadBalancer {
	lb := loadbalancer.NewLoadBalancer(iter)
	lb.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		sc.Responder.SendServerError(w, err.Error(), http.StatusInternalServerError)
	}
	return lb
}
