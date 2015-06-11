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

const (
  concurrentFetch = 100
  taskStaledTime  = 300
)

var (
  pushAddr       = flag.String("exporter.push-gateway", "", "Address to push metrics to the push-gateway")
  addr           = flag.String("web.listen-address", ":9105", "Address to listen on for web interface and telemetry")
  slaveURL       = flag.String("exporter.slave-url", "http://127.0.0.1:5051", "URL to the local Mesos slave")
  scrapeInterval = flag.Duration("exporter.interval", (10 * time.Second), "Scrape interval duration")
)

var (
  hostname, _ = os.Hostname()
)

var (
  variableLabels = []string{"source", "slave", "framework_id", "executor_id"}

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

  cpusUsrUsageDesc = prometheus.NewDesc(
    "mesos_task_cpus_user_usage",
    "Relative user CPU usage since the last query.",
    variableLabels, nil,
  )
  cpusSysUsageDesc = prometheus.NewDesc(
    "mesos_task_cpus_system_usage",
    "Relative CPU system usage since the last query.",
    variableLabels, nil,
  )
  cpusTotalUsageDesc = prometheus.NewDesc(
    "mesos_task_cpus_total_usage",
    "Relative combined CPU usage since the last query.",
    variableLabels, nil,
  )
)

type taskMetric struct {
  cpusUserTimeSecs   float64
  cpusSystemTimeSecs float64
  timestamp          float64
}

type taskMetrics map[string]taskMetric

type slaveExporterOpts struct {
  interval       time.Duration
  slaveURL       string
  pushGatewayURL string
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

func (e *periodicStatsExporter) fetch(metricsChan chan<- prometheus.Metric, tm taskMetrics, wg *sync.WaitGroup) {
  var taskID string
  var cpusLimit float64
  var cpusSystemTimeSecs float64
  var cpusUserTimeSecs float64
  var memLimitBytes int64
  var memRssBytes int64
  var timestamp float64

  var cpusSystemUsage float64
  var cpusUserUsage float64
  var cpusTotalUsage float64

  defer wg.Done()
  stats, err := e.slave.MesosSlaveMonitorStatistics()
  if err != nil {
    log.Printf("%v\n", err)
    return
  }

  for _, stat := range *stats {
    taskID = stat.Source
    cpusLimit = stat.Statistics.CpusLimit
    cpusSystemTimeSecs = stat.Statistics.CpusSystemTimeSecs
    cpusUserTimeSecs = stat.Statistics.CpusUserTimeSecs
    memLimitBytes = stat.Statistics.MemLimitBytes
    memRssBytes = stat.Statistics.MemRssBytes
    timestamp = stat.Statistics.Timestamp

    m, ok := tm[taskID]
    if ok {
      cpusSystemUsage = (cpusSystemTimeSecs - m.cpusSystemTimeSecs) / (timestamp - m.timestamp)
      cpusUserUsage = (cpusUserTimeSecs - m.cpusUserTimeSecs) / (timestamp - m.timestamp)
      cpusTotalUsage = cpusSystemUsage + cpusUserUsage

      metricsChan <- prometheus.MustNewConstMetric(
        cpusUsrUsageDesc,
        prometheus.GaugeValue,
        cpusUserUsage,
        stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
      )
      metricsChan <- prometheus.MustNewConstMetric(
        cpusSysUsageDesc,
        prometheus.GaugeValue,
        cpusSystemUsage,
        stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
      )
      metricsChan <- prometheus.MustNewConstMetric(
        cpusTotalUsageDesc,
        prometheus.GaugeValue,
        cpusTotalUsage,
        stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
      )
    }

    tm[taskID] = taskMetric{
      cpusSystemTimeSecs: cpusSystemTimeSecs,
      cpusUserTimeSecs:   cpusUserTimeSecs,
      timestamp:          timestamp,
    }

    metricsChan <- prometheus.MustNewConstMetric(
      cpusLimitDesc,
      prometheus.GaugeValue,
      cpusLimit,
      stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      cpusSysDesc,
      prometheus.CounterValue,
      cpusSystemTimeSecs,
      stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      cpusUsrDesc,
      prometheus.CounterValue,
      cpusUserTimeSecs,
      stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      memLimitDesc,
      prometheus.GaugeValue,
      float64(memLimitBytes),
      stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
    )
    metricsChan <- prometheus.MustNewConstMetric(
      memRssDesc,
      prometheus.GaugeValue,
      float64(memRssBytes),
      stat.Source, hostname, stat.FrameworkID, stat.ExecutorID,
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

  if e.opts.pushGatewayURL != "" {
    if err := prometheus.PushCollectors("monitor_statistics_json", hostname, e.opts.pushGatewayURL, e); err != nil {
      log.Printf("Could not push completion time to Pushgateway: %v\n", err)
    }
  }
}

func (e *periodicStatsExporter) scrapeSlaves(tm taskMetrics) {
  metricsChan := make(chan prometheus.Metric)
  go e.setMetrics(metricsChan)

  var wg sync.WaitGroup
  wg.Add(1)
  go e.fetch(metricsChan, tm, &wg)

  wg.Wait()
  close(metricsChan)

}

func deleteStaleTaskMetrics(tm taskMetrics) {
  now := time.Now().Unix()
  for tid, m := range tm {
    timeSinceLastUpdate := now - int64(m.timestamp)
    if timeSinceLastUpdate > taskStaledTime {
      delete(tm, tid)
    }
  }
}

func runEvery(f func(taskMetrics), interval time.Duration) {
  tm := make(taskMetrics)
  for _ = range time.NewTicker(interval).C {
    f(tm)
    deleteStaleTaskMetrics(tm)
  }
}

func main() {
  flag.Parse()

  opts := &slaveExporterOpts{
    interval:       *scrapeInterval,
    slaveURL:       *slaveURL,
    pushGatewayURL: *pushAddr,
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
