package main

import (
	"net/http"
	"fmt"
	"runtime"
	"log"
	"time"
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	request_latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_latency_seconds",
			Help: "latency seconds duration in seconds",
		},
		[]string{"endpoint"})
	login_users = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "login_users",
			Help: "login count",
		})
	buy_count = prometheus.NewCounter(prometheus.CounterOpts{
			Name: "buy_count",
			Help: "buy counter",
		})

	registry = prometheus.NewRegistry()
)

func init() {
	// Metrics have to be registered to be exposed:
	registry.MustRegister(request_latency)
	registry.MustRegister(login_users)
	registry.MustRegister(buy_count)
}

func main() {
	runtime.GOMAXPROCS(4)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/buy", buyHandler)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Fatal(http.ListenAndServe(":9876", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	fmt.Fprintf(w, "Login\n")
	fmt.Printf("Login\n")

	login_users.Inc()
	elapsed := time.Since(start).Seconds()
	request_latency.With(prometheus.Labels{"endpoint":"/login"}).Observe(elapsed)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	fmt.Fprintf(w, "Logout\n")
	fmt.Printf("Logout\n")

	login_users.Dec()
	elapsed := time.Since(start).Seconds()
	request_latency.With(prometheus.Labels{"endpoint":"/logout"}).Observe(elapsed)
}

func buyHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	delay := time.Duration( rand.Int() % 10000 ) * time.Millisecond // random delay
	time.Sleep(delay)
	fmt.Fprintf(w, "Buy with delay %fs\n", delay.Seconds())
	fmt.Printf("Buy with delay %fs\n", delay.Seconds())

	buy_count.Inc()
	elapsed := time.Since(start).Seconds()
	request_latency.With(prometheus.Labels{"endpoint":"/buy"}).Observe(elapsed)
}