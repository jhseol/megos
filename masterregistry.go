// +build ignore

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
  pushAddr       = flag.String("exporter.push-gateway", "localhost:9091", "Address to push metrics to the push-gateway")
  addr           = flag.String("web.listen-address", ":9105", "Address to listen on for web interface and telemetry")
  masterURL      = flag.String("exporter.master-url", "http://127.0.0.1:5050", "URL to the local Mesos master")
  scrapeInterval = flag.Duration("exporter.interval", (10 * time.Second), "Scrape interval duration")
)

var (
  hostname, _ = os.Hostname()
)

var (
  variableLabels = []string{"rackid", "slaveid", "hostname"}

  resourcesCPUsDesc = prometheus.NewDesc(
    "mesos_task_cpus_limit",
    "Fractional CPU limit.",
    variableLabels, nil,
  )
  resourcesMemDesc = prometheus.NewDesc(
    "mesos_task_cpus_system_time_secs",
    "Cumulative system CPU time in seconds.",
    variableLabels, nil,
  )
  resourcesDiskDesc = prometheus.NewDesc(
    "mesos_task_cpus_user_time_secs",
    "Cumulative user CPU time in seconds.",
    variableLabels, nil,
  )
)

type masterExporterOpts struct {
  interval       time.Duration
  masterURL      string
  pushGatewayURL string
}

type periodicStatsExporter struct {
  sync.RWMutex
  errors  *prometheus.CounterVec
  metrics []prometheus.Metric
  opts    *masterExporterOpts
  master  *megos.MesosMasterClient
}

func newMesosStatsExporter(master *megos.MesosMasterClient, opts *masterExporterOpts) *periodicStatsExporter {
  e := &periodicStatsExporter{
    errors: prometheus.NewCounterVec(
      prometheus.CounterOpts{
        Namespace: "mesos_stats_exporter",
        Name:      "master_scrape_errors_total",
        Help:      "Current total scrape erros",
      },
      []string{"master"},
    ),

    opts:   opts,
    master: master,
  }

  go runEvery(e.scrapeMaster, e.opts.interval)

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

func (e *periodicStatsExporter) fetch(metricsChan chan<- prometheus.Metric, wg *sync.WaitGroup) {
  defer wg.Done()
  stats, err := e.master.MesosMasterMonitorStatistics()
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

  if err := prometheus.PushCollectors("registrar(1)_registry", hostname, e.opts.pushGatewayURL, e); err != nil {
    log.Printf("Could not push completion time to Pushgateway: %v\n", err)
  }
}

func (e *periodicStatsExporter) scrapeMaster() {
  metricsChan := make(chan prometheus.Metric)
  go e.setMetrics(metricsChan)

  var wg sync.WaitGroup
  wg.Add(1)
  go e.fetch(metricsChan, &wg)

  wg.Wait()
  close(metricsChan)
}

func runEvery(f func(), interval time.Duration) {
  for _ = range time.NewTicker(interval).C {
    f()
  }
}

func main() {
  flag.Parse()

  opts := &masterExporterOpts{
    interval:       *scrapeInterval,
    masterURL:      *masterURL,
    pushGatewayURL: *pushAddr,
  }
  sopts := &megos.MesosMasterOptions{Host: opts.masterURL}
  master := megos.NewMasterClient(sopts)

  exporter := newMesosStatsExporter(master, opts)
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
