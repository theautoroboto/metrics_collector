package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Step 1 add new Gage here

var dockerSLIGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "docker_metric_success",
	Help: "Docker SLI gauge use to collect Docker SLIs",
})

var ecsSLIGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "ecs_metric_success",
	Help: "ECS SLI gauge use to collect ECS SLIs",
})
var rabbitSLIGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "rabbitmq_shared_svc_creation_availability",
	Help: "Rabbit SLI gauge use to collect Rabbit Shared SLIs",
})

var rabbitondmdSLIGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "rabbitmq_ondmd_svc_creation_availability",
	Help: "Rabbit On-Demand SLI gauge use to collect Rabbit On-Demand SLIs",
})

var credhubSLIGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "credhub_svc_creation_availability",
	Help: "Credhub SLI gauge use to collect Credhub SLIs",
})
var scs3SLIGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "scs3_creation_availability",
	Help: "Spring Cloud Services v3 SLI gauge use to collect Spring Cloud v3 SLIs",
})
var redisSLIGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "redis_creation_availability",
	Help: "p-redis SLI gauge use to collect p-redis SLIs",
})

var requestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "request_duration_seconds",
		Help: "A histogram of latencies for requests.",
	},
	[]string{"code", "method", "handler"},
)

func main() {
	router := mux.NewRouter()
	staticPath := "./static"

	router.Handle("/", http.FileServer(http.Dir(staticPath)))
	router.Handle("/metrics", promhttp.Handler())

	//Step 2 register new endpoint
	router.HandleFunc("/docker_metric", instrument(docker_metric, "docker_metric"))
	router.HandleFunc("/ecs_metric", instrument(ecs_metric, "ecs_metric"))
	router.HandleFunc("/rabbitmq_shared_svc_creation_availability", instrument(rabbitmq_shared_svc_creation_availability, "rabbitmq_shared_svc_creation_availability"))
	router.HandleFunc("/rabbitmq_ondmd_svc_creation_availability", instrument(rabbitmq_ondmd_svc_creation_availability, "rabbitmq_ondmd_svc_creation_availability"))
	router.HandleFunc("/credhub_svc_creation_availability", instrument(credhub_svc_creation_availability, "credhub_svc_creation_availability"))
	router.HandleFunc("/scs3_creation_availability", instrument(scs3_creation_availability, "scs3_creation_availability"))
	router.HandleFunc("/redis_creation_availability", instrument(redis_creation_availability, "redis_creation_availability"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func instrument(handlerFunc http.HandlerFunc, name string) http.HandlerFunc {
	handlerDuration, err := requestDuration.CurryWith(prometheus.Labels{
		"handler": name,
	})

	if err != nil {
		panic(err)
	}

	return promhttp.InstrumentHandlerDuration(handlerDuration, handlerFunc)
}

//Step 3 create new endpoint handler function
func docker_metric(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("success") != "0" {
		dockerSLIGauge.Set(1.0)
	} else {
		dockerSLIGauge.Set(0.0)
	}

	w.Write([]byte("{}"))
}

func ecs_metric(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("success") != "0" {
		ecsSLIGauge.Set(1.0)
	} else {
		ecsSLIGauge.Set(0.0)
	}

	w.Write([]byte("{}"))
}
func rabbitmq_shared_svc_creation_availability(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("success") != "0" {
		rabbitSLIGauge.Set(1.0)
	} else {
		rabbitSLIGauge.Set(0.0)
	}

	w.Write([]byte("{}"))
}
func rabbitmq_ondmd_svc_creation_availability(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("success") != "0" {
		rabbitondmdSLIGauge.Set(1.0)
	} else {
		rabbitondmdSLIGauge.Set(0.0)
	}

	w.Write([]byte("{}"))
}
func credhub_svc_creation_availability(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("success") != "0" {
		credhubSLIGauge.Set(1.0)
	} else {
		credhubSLIGauge.Set(0.0)
	}

	w.Write([]byte("{}"))
}
func scs3_creation_availability(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("success") != "0" {
		scs3SLIGauge.Set(1.0)
	} else {
		scs3SLIGauge.Set(0.0)
	}

	w.Write([]byte("{}"))
}
func redis_creation_availability(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("success") != "0" {
		redisSLIGauge.Set(1.0)
	} else {
		redisSLIGauge.Set(0.0)
	}

	w.Write([]byte("{}"))
}