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
  slaveURL       = flag.String("exporter.slave-url", "http://127.0.0.1:5051", "URL to the local Mesos slave")
  scrapeInterval = flag.Duration("exporter.interval", (10 * time.Second), "Scrape interval duration")
)

var (
  hostname, _ = os.Hostname()
)

var (
  variableLabels = []string{"source", "slave", "framework_id"}

  failedTasks = prometheus.NewDesc(
    "failed_tasks",
    "Failed tasks.",
    variableLabels, nil,
  )
  finishedTasks = prometheus.NewDesc(
    "finished_tasks",
    "Finished tasks.",
    variableLabels, nil,
  )
  invalidStatusUpdates = prometheus.NewDesc(
    "invalid_status_updates",
    "Invalid status updates.",
    variableLabels, nil,
  )
  killedTasks = prometheus.NewDesc(
    "killed_tasks",
    "Killed tasks.",
    variableLabels, nil,
  )
  launchedTasksGauge = prometheus.NewDesc(
    "launched_tasks_gauge",
    "Launched tasks gauge.",
    variableLabels, nil,
  )
  lostTasks = prometheus.NewDesc(
    "lost_tasks",
    "Lost tasks.",
    variableLabels, nil,
  )
  queuedTasksGauge = prometheus.NewDesc(
    "queued_tasks_gauge",
    "Queued tasks gauge",
    variableLabels, nil,
  )
  recoveryErrors = prometheus.NewDesc(
    "recovery_errors",
    "Recovery errors",
    variableLabels, nil,
  )
  registered = prometheus.NewDesc(
    "registered",
    "Registered.",
    variableLabels, nil,
  )
  slaveCpusPercent = prometheus.NewDesc(
    "slave/cpus_percent",
    "Slave/CPUs percentile.",
    variableLabels, nil,
  )
  slaveCpusTotal = prometheus.NewDesc(
    "slave/cpus_total",
    "Slave/CPUs total count.",
    variableLabels, nil,
  )
  slaveCpusUsed = prometheus.NewDesc(
    "slave/cpus_used",
    "CPUs used.",
    variableLabels, nil,
  )
  slaveDiskPercent = prometheus.NewDesc(
    "slave/disk_percent",
    "Slave/disk percentile.",
    variableLabels, nil,
  )
  slaveDiskTotal = prometheus.NewDesc(
    "slave/disk_total",
    "Slave/disk total.",
    variableLabels, nil,
  )
  slaveDiskUsed = prometheus.NewDesc(
    "slave/disk_used",
    "Slave/disk used.",
    variableLabels, nil,
  )
  slaveExecutorsRegistering = prometheus.NewDesc(
    "slave/executors_registering",
    "Slave/executors registering.",
    variableLabels, nil,
  )
  slaveExecutorsRunning = prometheus.NewDesc(
    "slave/executors_running",
    "Slave/executors running.",
    variableLabels, nil,
  )
  slaveExecutorsTerminated = prometheus.NewDesc(
    "slave/executors_terminated",
    "Slave/executors terminated.",
    variableLabels, nil,
  )
  slaveExecutorsTerminating = prometheus.NewDesc(
    "slave/executors_terminating",
    "Slave/executors terminating.",
    variableLabels, nil,
  )
  slaveFrameworksActive = prometheus.NewDesc(
    "slave/frameworks_active",
    "Slave/frameworks active.",
    variableLabels, nil,
  )
  slaveInvalidFrameworkMessages = prometheus.NewDesc(
    "slave/invalid_framework_messages",
    "Slave/invalid framework messages.",
    variableLabels, nil,
  )
  slaveInvalidStatusUpdates = prometheus.NewDesc(
    "slave/invalid_status_updates",
    "Slave/invalid status updates.",
    variableLabels, nil,
  )
  slaveMemPercent = prometheus.NewDesc(
    "slave/mem_percent",
    "Slave/mem percentile.",
    variableLabels, nil,
  )
  slaveMemTotal = prometheus.NewDesc(
    "slave/mem_total",
    "Slave/mem total.",
    variableLabels, nil,
  )
  slaveMemUsed = prometheus.NewDesc(
    "slave/mem_used",
    "Slave/mem used.",
    variableLabels, nil,
  )
  slaveRecoveryErrors = prometheus.NewDesc(
    "slave/recovery_errors",
    "Slave/recovery errors.",
    variableLabels, nil,
  )
  slaveRegistered = prometheus.NewDesc(
    "slave/registered",
    "Slave/registered.",
    variableLabels, nil,
  )
  slaveTasksFailed = prometheus.NewDesc(
    "slave/tasks_failed",
    "Slave/tasks failed.",
    variableLabels, nil,
  )
  slaveTasksFinished = prometheus.NewDesc(
    "slave/tasks_finished",
    "Slave/tasks finished.",
    variableLabels, nil,
  )
  slaveTasksKilled = prometheus.NewDesc(
    "slave/tasks_killed",
    "Slave/tasks killed.",
    variableLabels, nil,
  )
  slaveTasksLost = prometheus.NewDesc(
    "slave/tasks_lost",
    "Slave/tasks lost.",
    variableLabels, nil,
  )
  slaveTasksRunning = prometheus.NewDesc(
    "slave/tasks_running",
    "Slave/tasks running.",
    variableLabels, nil,
  )
  slaveTasksStaging = prometheus.NewDesc(
    "slave/tasks_staging",
    "Slave/tasks staging.",
    variableLabels, nil,
  )
  slaveTasksStarting = prometheus.NewDesc(
    "slave/tasks_starting",
    "Slave/tasks starting.",
    variableLabels, nil,
  )
  slaveUptimeSecs = prometheus.NewDesc(
    "slave/uptime_secs",
    "Slave/uptime in seconds.",
    variableLabels, nil,
  )
  slaveValidFrameworkMessages = prometheus.NewDesc(
    "slave/valid_framework_messages",
    "Slave/valid framework messages.",
    variableLabels, nil,
  )
  slaveValidStatusUpdates = prometheus.NewDesc(
    "slave/valid_status_updates",
    "Slave/vaild status updates.",
    variableLabels, nil,
  )
  stagedTasks = prometheus.NewDesc(
    "staged_tasks",
    "Staged tasks.",
    variableLabels, nil,
  )
  startedTasks = prometheus.NewDesc(
    "started_tasks",
    "Started tasks.",
    variableLabels, nil,
  )
  systemCpusTotal = prometheus.NewDesc(
    "system/cpus_total",
    "System/CPUs total.",
    variableLabels, nil,
  )
  systemLoad15min = prometheus.NewDesc(
    "system/load_15min",
    "System/loadavg in 15 mins.",
    variableLabels, nil,
  )
  systemLoad1min = prometheus.NewDesc(
    "system/load_1min",
    "System/loadavg in 1 min.",
    variableLabels, nil,
  )
  systemLoad5min = prometheus.NewDesc(
    "system/load_5min",
    "System/loadavg in 5 mins.",
    variableLabels, nil,
  )
  systemMemFreeBytes = prometheus.NewDesc(
    "system/mem_free_bytes",
    "System/mem free in bytes.",
    variableLabels, nil,
  )
  systemMemTotalBytes = prometheus.NewDesc(
    "system/mem_total_bytes",
    "System/mem total in bytes.",
    variableLabels, nil,
  )
  totalFrameworks = prometheus.NewDesc(
    "total_frameworks",
    "Total frameworks.",
    variableLabels, nil,
  )
  uptime = prometheus.NewDesc(
    "uptime",
    "Uptime in nanoseconds.",
    variableLabels, nil,
  )
  validStatusUpdates = prometheus.NewDesc(
    "valid_status_updates",
    "Vaild status updates.",
    variableLabels, nil,
  )
)

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

func (e *periodicStatsExporter) fetch(metricsChan chan<- prometheus.Metric, wg *sync.WaitGroup) {
  defer wg.Done()
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

  if err := prometheus.PushCollectors("monitor_statistics_json", hostname, e.opts.pushGatewayURL, e); err != nil {
    log.Printf("Could not push completion time to Pushgateway: %v\n", err)
  }
}

func (e *periodicStatsExporter) scrapeSlaves() {
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
