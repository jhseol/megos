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
  pushAddr       = flag.String("exporter.push-gateway", "", "Address to push metrics to the push-gateway")
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

  failedTasksDesc = prometheus.NewDesc(
    "failed_tasks",
    "Tasks failed to finish successfully.",
    variableLabels, nil,
  )
  finishedTasksDesc = prometheus.NewDesc(
    "finished_tasks",
    "Tasks finished successfully.",
    variableLabels, nil,
  )
  invalidStatusUpdatesDesc = prometheus.NewDesc(
    "invalid_status_updates",
    "Failed to update task status.",
    variableLabels, nil,
  )
  killedTasksDesc = prometheus.NewDesc(
    "killed_tasks",
    "Tasks were killed by the executor.",
    variableLabels, nil,
  )
  launchedTasksGaugeDesc = prometheus.NewDesc(
    "launched_tasks_gauge",
    "Sent to executor (TASK_STAGING, TASK_STARTING, TASK_RUNNING).",
    variableLabels, nil,
  )
  lostTasksDesc = prometheus.NewDesc(
    "lost_tasks",
    "Tasks failed but can be rescheduled.",
    variableLabels, nil,
  )
  queuedTasksGaugeDesc = prometheus.NewDesc(
    "queued_tasks_gauge",
    "Queued waiting for executor to register.",
    variableLabels, nil,
  )
  recoveryErrorsDesc = prometheus.NewDesc(
    "recovery_errors",
    "Indicates the number of errors ignored in '--no-strict' recovery mode",
    variableLabels, nil,
  )
  registeredDesc = prometheus.NewDesc(
    "registered",
    "Indicated the slave registered to master.",
    variableLabels, nil,
  )
  slaveCpusPercentDesc = prometheus.NewDesc(
    "slave_cpus_percent",
    "Relative CPU usage(slave) since the last query.",
    variableLabels, nil,
  )
  slaveCpusTotalDesc = prometheus.NewDesc(
    "slave_cpus_total",
    "Absolute CPUs total count.",
    variableLabels, nil,
  )
  slaveCpusUsedDesc = prometheus.NewDesc(
    "slave_cpus_used",
    "Slave/CPUs usage count.",
    variableLabels, nil,
  )
  slaveDiskPercentDesc = prometheus.NewDesc(
    "slave_disk_percent",
    "Slave/disk usage percentile.",
    variableLabels, nil,
  )
  slaveDiskTotalDesc = prometheus.NewDesc(
    "slave_disk_total_megabytes",
    "Slave/disk total size in MB.",
    variableLabels, nil,
  )
  slaveDiskUsedDesc = prometheus.NewDesc(
    "slave_disk_used_megabytes",
    "Slave/disk used in MB.",
    variableLabels, nil,
  )
  slaveExecutorsRegisteringDesc = prometheus.NewDesc(
    "slave_executors_registering",
    "Slave/executors registering.",
    variableLabels, nil,
  )
  slaveExecutorsRunningDesc = prometheus.NewDesc(
    "slave_executors_running",
    "Slave/executors running.",
    variableLabels, nil,
  )
  slaveExecutorsTerminatedDesc = prometheus.NewDesc(
    "slave_executors_terminated",
    "Slave/executors terminated.",
    variableLabels, nil,
  )
  slaveExecutorsTerminatingDesc = prometheus.NewDesc(
    "slave_executors_terminating",
    "Slave/executors terminating.",
    variableLabels, nil,
  )
  slaveFrameworksActiveDesc = prometheus.NewDesc(
    "slave_frameworks_active",
    "Slave/frameworks active.",
    variableLabels, nil,
  )
  slaveInvalidFrameworkMessagesDesc = prometheus.NewDesc(
    "slave_invalid_framework_messages",
    "Slave/invalid framework messages.",
    variableLabels, nil,
  )
  slaveInvalidStatusUpdatesDesc = prometheus.NewDesc(
    "slave_invalid_status_updates",
    "Slave/invalid status updates.",
    variableLabels, nil,
  )
  slaveMemPercentDesc = prometheus.NewDesc(
    "slave_mem_percent",
    "Slave/memory percentile.",
    variableLabels, nil,
  )
  slaveMemTotalDesc = prometheus.NewDesc(
    "slave_mem_total_megabytes",
    "Slave/memory total in MB.",
    variableLabels, nil,
  )
  slaveMemUsedDesc = prometheus.NewDesc(
    "slave_mem_used_megabytes",
    "Slave/memory used in MB.",
    variableLabels, nil,
  )
  slaveRecoveryErrorsDesc = prometheus.NewDesc(
    "slave_recovery_errors",
    "Slave/recovery errors.",
    variableLabels, nil,
  )
  slaveRegisteredDesc = prometheus.NewDesc(
    "slave_registered",
    "Slave/registered.",
    variableLabels, nil,
  )
  slaveTasksFailedDesc = prometheus.NewDesc(
    "slave_tasks_failed",
    "Slave/tasks failed.",
    variableLabels, nil,
  )
  slaveTasksFinishedDesc = prometheus.NewDesc(
    "slave_tasks_finished",
    "Slave/tasks finished.",
    variableLabels, nil,
  )
  slaveTasksKilledDesc = prometheus.NewDesc(
    "slave_tasks_killed",
    "Slave/tasks killed.",
    variableLabels, nil,
  )
  slaveTasksLostDesc = prometheus.NewDesc(
    "slave_tasks_lost",
    "Slave/tasks lost.",
    variableLabels, nil,
  )
  slaveTasksRunningDesc = prometheus.NewDesc(
    "slave_tasks_running",
    "Slave/tasks running.",
    variableLabels, nil,
  )
  slaveTasksStagingDesc = prometheus.NewDesc(
    "slave_tasks_staging",
    "Slave/tasks staging.",
    variableLabels, nil,
  )
  slaveTasksStartingDesc = prometheus.NewDesc(
    "slave_tasks_starting",
    "Slave/tasks starting.",
    variableLabels, nil,
  )
  slaveUptimeSecsDesc = prometheus.NewDesc(
    "slave_uptime_secs",
    "Slave/uptime in seconds.",
    variableLabels, nil,
  )
  slaveValidFrameworkMessagesDesc = prometheus.NewDesc(
    "slave_valid_framework_messages",
    "Slave/valid framework messages.",
    variableLabels, nil,
  )
  slaveValidStatusUpdatesDesc = prometheus.NewDesc(
    "slave_valid_status_updates",
    "Slave/vaild status updates.",
    variableLabels, nil,
  )
  stagedTasksDesc = prometheus.NewDesc(
    "staged_tasks",
    "Staged tasks.",
    variableLabels, nil,
  )
  startedTasksDesc = prometheus.NewDesc(
    "started_tasks",
    "Started tasks.",
    variableLabels, nil,
  )
  systemCpusTotalDesc = prometheus.NewDesc(
    "system_cpus_total",
    "System/CPUs total.",
    variableLabels, nil,
  )
  systemLoad15minDesc = prometheus.NewDesc(
    "system_load_15min",
    "System/loadavg in 15 mins.",
    variableLabels, nil,
  )
  systemLoad1minDesc = prometheus.NewDesc(
    "system_load_1min",
    "System/loadavg in 1 min.",
    variableLabels, nil,
  )
  systemLoad5minDesc = prometheus.NewDesc(
    "system_load_5min",
    "System/loadavg in 5 mins.",
    variableLabels, nil,
  )
  systemMemFreeBytesDesc = prometheus.NewDesc(
    "system_mem_free_bytes",
    "System/memory free in bytes.",
    variableLabels, nil,
  )
  systemMemTotalBytesDesc = prometheus.NewDesc(
    "system_mem_total_bytes",
    "System/memory total in bytes.",
    variableLabels, nil,
  )
  totalFrameworksDesc = prometheus.NewDesc(
    "total_frameworks",
    "Total frameworks.",
    variableLabels, nil,
  )
  uptimeDesc = prometheus.NewDesc(
    "uptime_nanosecs",
    "Uptime in nanoseconds.",
    variableLabels, nil,
  )
  validStatusUpdatesDesc = prometheus.NewDesc(
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
    failedTasksDesc,
    prometheus.CounterValue,
    float64(stats.FailedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    finishedTasksDesc,
    prometheus.CounterValue,
    float64(stats.FinishedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    invalidStatusUpdatesDesc,
    prometheus.CounterValue,
    float64(stats.InvalidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    killedTasksDesc,
    prometheus.CounterValue,
    float64(stats.KilledTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    launchedTasksGaugeDesc,
    prometheus.GaugeValue,
    float64(stats.LaunchedTasksGauge),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    lostTasksDesc,
    prometheus.CounterValue,
    float64(stats.LostTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    queuedTasksGaugeDesc,
    prometheus.GaugeValue,
    float64(stats.QueuedTasksGauge),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    recoveryErrorsDesc,
    prometheus.CounterValue,
    float64(stats.RecoveryErrors),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registeredDesc,
    prometheus.GaugeValue,
    float64(stats.Registered),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveCpusPercentDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveCpusPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveCpusTotalDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveCpusTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveCpusUsedDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveCpusUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveDiskPercentDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveDiskPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveDiskTotalDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveDiskTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveDiskUsedDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveDiskUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsRegisteringDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveExecutorsRegistering),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsRunningDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveExecutorsRunning),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsTerminatedDesc,
    prometheus.CounterValue,
    float64(stats.SlaveExecutorsTerminated),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveExecutorsTerminatingDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveExecutorsTerminating),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveFrameworksActiveDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveFrameworksActive),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveInvalidFrameworkMessagesDesc,
    prometheus.CounterValue,
    float64(stats.SlaveInvalidFrameworkMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveInvalidStatusUpdatesDesc,
    prometheus.CounterValue,
    float64(stats.SlaveInvalidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveMemPercentDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveMemPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveMemTotalDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveMemTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveMemUsedDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveMemUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveRecoveryErrorsDesc,
    prometheus.CounterValue,
    float64(stats.SlaveRecoveryErrors),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveRegisteredDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveRegistered),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksFailedDesc,
    prometheus.CounterValue,
    float64(stats.SlaveTasksFailed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksFinishedDesc,
    prometheus.CounterValue,
    float64(stats.SlaveTasksFinished),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksKilledDesc,
    prometheus.CounterValue,
    float64(stats.SlaveTasksKilled),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksLostDesc,
    prometheus.CounterValue,
    float64(stats.SlaveTasksLost),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksRunningDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveTasksRunning),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksStagingDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveTasksStaging),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveTasksStartingDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveTasksStarting),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveUptimeSecsDesc,
    prometheus.CounterValue,
    float64(stats.SlaveUptimeSecs),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveValidFrameworkMessagesDesc,
    prometheus.GaugeValue,
    float64(stats.SlaveValidFrameworkMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    slaveValidStatusUpdatesDesc,
    prometheus.CounterValue,
    float64(stats.SlaveValidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    stagedTasksDesc,
    prometheus.GaugeValue,
    float64(stats.StagedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    startedTasksDesc,
    prometheus.GaugeValue,
    float64(stats.StartedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemCpusTotalDesc,
    prometheus.GaugeValue,
    float64(stats.SystemCpusTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemLoad15minDesc,
    prometheus.GaugeValue,
    float64(stats.SystemLoad15min),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemLoad1minDesc,
    prometheus.GaugeValue,
    float64(stats.SystemLoad1min),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemLoad5minDesc,
    prometheus.GaugeValue,
    float64(stats.SystemLoad5min),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemMemFreeBytesDesc,
    prometheus.GaugeValue,
    float64(stats.SystemMemFreeBytes),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    systemMemTotalBytesDesc,
    prometheus.GaugeValue,
    float64(stats.SystemMemTotalBytes),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    totalFrameworksDesc,
    prometheus.GaugeValue,
    float64(stats.TotalFrameworks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    uptimeDesc,
    prometheus.CounterValue,
    float64(stats.Uptime),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    validStatusUpdatesDesc,
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

  if e.opts.pushGatewayURL != "" {
    if err := prometheus.PushCollectors("slave(1)_stats_json", hostname, e.opts.pushGatewayURL, e); err != nil {
      log.Printf("Could not push completion time to Pushgateway: %v\n", err)
    }
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
