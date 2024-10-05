package metrics

// func MetricListen(address string) error {
// 	mux := http.NewServeMux()
// 	mux.Handle("/metrics", promhttp.Handler())

// 	return http.ListenAndServe(address, mux)
// }

// пихать в каждый роут
// start := time.Now()
// defer func() {
// 	metrics.ObserveRequest(time.Since(start), c.Response().StatusCode())
// }()
