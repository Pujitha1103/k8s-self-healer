package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
    ReconcileDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
        Name: "reconcile_duration_seconds",
        Help: "Reconcile loop duration in seconds.",
    })
)

// RegisterMetrics registers prometheus metrics used by the controller.
func RegisterMetrics() {
    prometheus.MustRegister(ReconcileDuration)
}
