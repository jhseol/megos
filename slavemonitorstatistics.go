package main

import (
  "flag"
  "log"
  "net/http"
  "os"
  "sync"
  "time"

  "github.com/jhseol/megos"
  "github.com/prometheus/client_golang/prometheus"
)

const concurrentFetch = 100

var (
  addr           = flag.String("web.listen-address", ":9105", "Address to listen on for web interface and telemetry")
  slaveURL       = flag.String("exporter.slave-url", "http://127.0.0.1:5051", "URL to the local Mesos slave")
  scrapeInterval = flag.Duration("exporter.interval", (60 * time.Second), "Scrape interval duration")
)

var (
  hostname, _ = os.Hostname()
)

var (
  variableLabels = []string{"source", "slave", "framework_id"}

  cpusLimitDesc = prometheus.NewDesc(
    "mesos_task_cpus_limit",
    "Fractional CPU limit.",
    variableLabels, nil,
  )
  cpusSysDesc = prometheus.NewDesc(
    "mesos_task_cpus_system_time_secs",
    "Cumulative system CPU time in seconds.",
    variableLabels, nil,
  )
  cpusUsrDesc = prometheus.NewDesc(
    "mesos_task_cpus_user_time_secs",
    "Cumulative user CPU time in seconds.",
    variableLabels, nil,
  )
  memLimitDesc = prometheus.NewDesc(
    "mesos_task_mem_limit_bytes",
    "Task memory limit in bytes",
    variableLabels, nil,
  )
  memRssDesc = prometheus.NewDesc(
    "mesos_task_mem_rss_bytes",
    "Task memory RSS usage in bytes",
    variableLabels, nil,
  )
)

type slaveExporterOpts struct {
  interval time.Duration
  slaveURL string
}

type periodicStatsExporter struct {
  sync.RWMutex
  errors  *prometheus.CounterVec
  metrics []prometheus.Metric
  opts    *slaveExporterOpts
  slave   *megos.MesosSlaveClient
}

func newMesosStatsExporter(slave *megos.MesosSlaveClient, opts *slaveExporterOpts) *periodicStatsExporter {
  e := &periodicStatsExporter{
    errors: prometheus.NewCounterVec(
      prometheus.CounterOpts{
        Namespace: "mesos_stats_exporter",
        Name:      "slave_scrape_errors_total",
        Help:      "Current total scrape erros",
      },
      []string{"slave"},
    ),

    opts:  opts,
    slave: slave,
  }

  go runEvery(e.scrapeSlaves, e.opts.interval)

  return e
}

func (e *periodicStatsExporter) Describe(ch chan<- *prometheus.Desc) {
  e.rLockMetrics(func() {
    for _, m := range e.metrics {
      ch <- m.Desc()
    }
  })
  e.errors.MetricVec.Describe(ch)
}

func (e *periodicStatsExporter) Collect(ch chan<- prometheus.Metric) {
  e.rLockMetrics(func() {
    for _, m := range e.metrics {
      ch <- m
    }
  })
  e.errors.MetricVec.Collect(ch)
}

func (e *periodicStatsExporter) fetch(metricsChan chan<- prometheus.Metric) {
  stats, err := e.slave.MesosSlaveMonitorStatistics()
  if err != nil {
    log.Printf("%v\n", err)
    return
  }

  for _, stat := range *stats {
    metricsChan <- prometheus.MustNewConstMetric(
      cpusLimitDesc,
      prometheus.GaugeValue,
      float64(stat.Statistics.CpusLimit),
      stat.Source, hostname, stat.FrameworkID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      cpusSysDesc,
      prometheus.CounterValue,
      float64(stat.Statistics.CpusSystemTimeSecs),
      stat.Source, hostname, stat.FrameworkID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      cpusUsrDesc,
      prometheus.CounterValue,
      float64(stat.Statistics.CpusUserTimeSecs),
      stat.Source, hostname, stat.FrameworkID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      memLimitDesc,
      prometheus.GaugeValue,
      float64(stat.Statistics.MemLimitBytes),
      stat.Source, hostname, stat.FrameworkID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      memRssDesc,
      prometheus.GaugeValue,
      float64(stat.Statistics.MemRssBytes),
      stat.Source, hostname, stat.FrameworkID,
    )
  }
}

func (e *periodicStatsExporter) rLockMetrics(f func()) {
  e.RLock()
  defer e.RUnlock()
  f()
}

func (e *periodicStatsExporter) setMetrics(ch chan prometheus.Metric) {
  metrics := make([]prometheus.Metric, 0)
  for metric := range ch {
    metrics = append(metrics, metric)
  }

  e.Lock()
  e.metrics = metrics
  e.Unlock()
}

func (e *periodicStatsExporter) scrapeSlaves() {
  metricsChan := make(chan prometheus.Metric)
  go e.setMetrics(metricsChan)

  go e.fetch(metricsChan)

  close(metricsChan)
}

func runEvery(f func(), interval time.Duration) {
  for _ = range time.NewTicker(interval).C {
    f()
  }
}

func main() {
  flag.Parse()

  opts := &slaveExporterOpts{
    interval: *scrapeInterval,
    slaveURL: *slaveURL,
  }
  sopts := &megos.MesosSlaveOptions{Host: opts.slaveURL}
  slave := megos.NewSlaveClient(sopts)

  exporter := newMesosStatsExporter(slave, opts)
  prometheus.MustRegister(exporter)

  http.Handle("/metrics", prometheus.Handler())
  http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%v, OK", w)
  })
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
  })

  log.Printf("starting mesos_exporter on %v\n", *addr)

  log.Fatal(http.ListenAndServe(*addr, nil))
}
