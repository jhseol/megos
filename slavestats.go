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

  hostname, _ = os.Hostname()
)

type promDescs struct {
  fqName         string
  help           string
  variableLabels string
  constLabels    prometheus.Labels
}

var (
  variableLabels = []string{"slave"}

  failedTasks = prometheus.NewDesc(
    "failed_tasks",
    "Tasks failed to finish successfully.",
    variableLabels, nil,
  )
  finishedTasks = prometheus.NewDesc(
    "finished_tasks",
    "Tasks finished successfully.",
    variableLabels, nil,
  )
  invalidStatusUpdates = prometheus.NewDesc(
    "invalid_status_updates",
    "Failed to update task status.",
    variableLabels, nil,
  )
  killedTasks = prometheus.NewDesc(
    "killed_tasks",
    "Tasks were killed by the executor.",
    variableLabels, nil,
  )
  launchedTasksGauge = prometheus.NewDesc(
    "launched_tasks_gauge",
    "Sent to executor (TASK_STAGING, TASK_STARTING, TASK_RUNNING).",
    variableLabels, nil,
  )
  lostTasks = prometheus.NewDesc(
    "lost_tasks",
    "Tasks failed but can be rescheduled.",
    variableLabels, nil,
  )
  queuedTasksGauge = prometheus.NewDesc(
    "queued_tasks_gauge",
    "Queued waiting for executor to register.",
    variableLabels, nil,
  )
  recoveryErrors = prometheus.NewDesc(
    "recovery_errors",
    "Indicates the number of errors ignored in '--no-strict' recovery mode",
    variableLabels, nil,
  )
  registered = prometheus.NewDesc(
    "registered",
    "Indicated the slave registered to master.",
    variableLabels, nil,
  )
  slaveCpusPercent = prometheus.NewDesc(
    "slave_cpus_percent",
    "Relative CPU usage(slave) since the last query.",
    variableLabels, nil,
  )
  slaveCpusTotal = prometheus.NewDesc(
    "slave_cpus_total",
    "Absolute CPUs total count.",
    variableLabels, nil,
  )
  slaveCpusUsed = prometheus.NewDesc(
    "slave_cpus_used",
    "Slave/CPUs usage count.",
    variableLabels, nil,
  )
  slaveDiskPercent = prometheus.NewDesc(
    "slave_disk_percent",
    "Slave/disk usage percentile.",
    variableLabels, nil,
  )
  slaveDiskTotal = prometheus.NewDesc(
    "slave_disk_total",
    "Slave/disk total size.",
    variableLabels, nil,
  )
  slaveDiskUsed = prometheus.NewDesc(
    "slave_disk_used",
    "Slave/disk used.",
    variableLabels, nil,
  )
  slaveExecutorsRegistering = prometheus.NewDesc(
    "slave_executors_registering",
    "Slave/executors registering.",
    variableLabels, nil,
  )
  slaveExecutorsRunning = prometheus.NewDesc(
    "slave_executors_running",
    "Slave/executors running.",
    variableLabels, nil,
  )
  slaveExecutorsTerminated = prometheus.NewDesc(
    "slave_executors_terminated",
    "Slave/executors terminated.",
    variableLabels, nil,
  )
  slaveExecutorsTerminating = prometheus.NewDesc(
    "slave_executors_terminating",
    "Slave/executors terminating.",
    variableLabels, nil,
  )
  slaveFrameworksActive = prometheus.NewDesc(
    "slave_frameworks_active",
    "Slave/frameworks active.",
    variableLabels, nil,
  )
  slaveInvalidFrameworkMessages = prometheus.NewDesc(
    "slave_invalid_framework_messages",
    "Slave/invalid framework messages.",
    variableLabels, nil,
  )
  slaveInvalidStatusUpdates = prometheus.NewDesc(
    "slave_invalid_status_updates",
    "Slave/invalid status updates.",
    variableLabels, nil,
  )
  slaveMemPercent = prometheus.NewDesc(
    "slave_mem_percent",
    "Slave/memory percentile.",
    variableLabels, nil,
  )
  slaveMemTotal = prometheus.NewDesc(
    "slave_mem_total",
    "Slave/memory total.",
    variableLabels, nil,
  )
  slaveMemUsed = prometheus.NewDesc(
    "slave_mem_used",
    "Slave/memory used.",
    variableLabels, nil,
  )
  slaveRecoveryErrors = prometheus.NewDesc(
    "slave_recovery_errors",
    "Slave/recovery errors.",
    variableLabels, nil,
  )
  slaveRegistered = prometheus.NewDesc(
    "slave_registered",
    "Slave/registered.",
    variableLabels, nil,
  )
  slaveTasksFailed = prometheus.NewDesc(
    "slave_tasks_failed",
    "Slave/tasks failed.",
    variableLabels, nil,
  )
  slaveTasksFinished = prometheus.NewDesc(
    "slave_tasks_finished",
    "Slave/tasks finished.",
    variableLabels, nil,
  )
  slaveTasksKilled = prometheus.NewDesc(
    "slave_tasks_killed",
    "Slave/tasks killed.",
    variableLabels, nil,
  )
  slaveTasksLost = prometheus.NewDesc(
    "slave_tasks_lost",
    "Slave/tasks lost.",
    variableLabels, nil,
  )
  slaveTasksRunning = prometheus.NewDesc(
    "slave_tasks_running",
    "Slave/tasks running.",
    variableLabels, nil,
  )
  slaveTasksStaging = prometheus.NewDesc(
    "slave_tasks_staging",
    "Slave/tasks staging.",
    variableLabels, nil,
  )
  slaveTasksStarting = prometheus.NewDesc(
    "slave_tasks_starting",
    "Slave/tasks starting.",
    variableLabels, nil,
  )
  slaveUptimeSecs = prometheus.NewDesc(
    "slave_uptime_secs",
    "Slave/uptime in seconds.",
    variableLabels, nil,
  )
  slaveValidFrameworkMessages = prometheus.NewDesc(
    "slave_valid_framework_messages",
    "Slave/valid framework messages.",
    variableLabels, nil,
  )
  slaveValidStatusUpdates = prometheus.NewDesc(
    "slave_valid_status_updates",
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
    "System/memory free in bytes.",
    variableLabels, nil,
  )
  systemMemTotalBytes = prometheus.NewDesc(
    "system/mem_total_bytes",
    "System/memory total in bytes.",
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
  stats, err := e.slave.MesosSlaveStats()
  if err != nil {
    log.Printf("%v\n", err)
    return
  }

  metricsChan <- prometheus.MustNewConstMetric(
    failedTasks,
    prometheus.CounterValue,
    float64(stats.FailedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    finishedTasks,
    prometheus.CounterValue,
    float64(stats.FinishedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    invalidStatusUpdates,
    prometheus.CounterValue,
    float64(stats.InvalidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    killedTasks,
    prometheus.CounterValue,
    float64(stats.KilledTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    launchedTasksGauge,
    prometheus.GaugeValue,
    float64(stats.LaunchedTasksGauge),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    lostTasks,
    prometheus.CounterValue,
    float64(stats.LostTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    queuedTasksGauge,
    prometheus.GaugeValue,
    float64(stats.QueuedTasksGauge),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    recoveryErrors,
    prometheus.CounterValue,
    float64(stats.RecoveryErrors),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registered,
    prometheus.GaugeValue,
    float64(stats.Registered),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveCpusPercent,
    prometheus.GaugeValue,
    float64(stats.SlaveCpusPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveCpusTotal,
    prometheus.GaugeValue,
    float64(stats.SlaveCpusTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveCpusUsed,
    prometheus.GaugeValue,
    float64(stats.SlaveCpusUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveDiskPercent,
    prometheus.GaugeValue,
    float64(stats.SlaveDiskPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveDiskTotal,
    prometheus.GaugeValue,
    float64(stats.SlaveDiskTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveDiskUsed,
    prometheus.GaugeValue,
    float64(stats.SlaveDiskUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsRegistering,
    prometheus.GaugeValue,
    float64(stats.SlaveExecutorsRegistering),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsRunning,
    prometheus.GaugeValue,
    float64(stats.SlaveExecutorsRunning),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsTerminated,
    prometheus.CounterValue,
    float64(stats.SlaveExecutorsTerminated),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsTerminating,
    prometheus.GaugeValue,
    float64(stats.SlaveExecutorsTerminating),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveFrameworksActive,
    prometheus.GaugeValue,
    float64(stats.SlaveFrameworksActive),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveInvalidFrameworkMessages,
    prometheus.CounterValue,
    float64(stats.SlaveInvalidFrameworkMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveInvalidStatusUpdates,
    prometheus.CounterValue,
    float64(stats.SlaveInvalidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveMemPercent,
    prometheus.GaugeValue,
    float64(stats.SlaveMemPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveMemTotal,
    prometheus.GaugeValue,
    float64(stats.SlaveMemTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveMemUsed,
    prometheus.GaugeValue,
    float64(stats.SlaveMemUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveRecoveryErrors,
    prometheus.CounterValue,
    float64(stats.SlaveRecoveryErrors),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveRegistered,
    prometheus.GaugeValue,
    float64(stats.SlaveRegistered),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksFailed,
    prometheus.CounterValue,
    float64(stats.SlaveTasksFailed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksFinished,
    prometheus.CounterValue,
    float64(stats.SlaveTasksFinished),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksKilled,
    prometheus.CounterValue,
    float64(stats.SlaveTasksKilled),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksLost,
    prometheus.CounterValue,
    float64(stats.SlaveTasksLost),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksRunning,
    prometheus.GaugeValue,
    float64(stats.SlaveTasksRunning),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksStaging,
    prometheus.GaugeValue,
    float64(stats.SlaveTasksStaging),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksStarting,
    prometheus.GaugeValue,
    float64(stats.SlaveTasksStarting),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveUptimeSecs,
    prometheus.CounterValue,
    float64(stats.SlaveUptimeSecs),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveValidFrameworkMessages,
    prometheus.GaugeValue,
    float64(stats.SlaveValidFrameworkMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveValidStatusUpdates,
    prometheus.CounterValue,
    float64(stats.SlaveValidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    stagedTasks,
    prometheus.GaugeValue,
    float64(stats.StagedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    startedTasks,
    prometheus.GaugeValue,
    float64(stats.StartedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemCpusTotal,
    prometheus.GaugeValue,
    float64(stats.SystemCpusTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemLoad15min,
    prometheus.GaugeValue,
    float64(stats.SystemLoad15min),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemLoad1min,
    prometheus.GaugeValue,
    float64(stats.SystemLoad1min),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemLoad5min,
    prometheus.GaugeValue,
    float64(stats.SystemLoad5min),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemMemFreeBytes,
    prometheus.GaugeValue,
    float64(stats.SystemMemFreeBytes),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemMemTotalBytes,
    prometheus.GaugeValue,
    float64(stats.SystemMemTotalBytes),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    totalFrameworks,
    prometheus.GaugeValue,
    float64(stats.TotalFrameworks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    uptime,
    prometheus.CounterValue,
    float64(stats.Uptime),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    validStatusUpdates,
    prometheus.CounterValue,
    float64(stats.ValidStatusUpdates),
    hostname,
  )

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

  if err := prometheus.PushCollectors("slave(1)_stats_json", hostname, e.opts.pushGatewayURL, e); err != nil {
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
