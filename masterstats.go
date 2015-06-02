// +build ignore

// TODO: Write full descriptions for metrics

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
  masterURL      = flag.String("exporter.master-url", "http://127.0.0.1:5051", "URL to the local Mesos master")
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
  variableLabels = []string{"master"}

  activatedSlaves = prometheus.NewDesc(
    "activated_slaves",
    "activated_slaves",
    variableLabels, nil)
  activeSchedulers = prometheus.NewDesc(
    "active_schedulers",
    "active_schedulers",
    variableLabels, nil)
  activeTasksGauge = prometheus.NewDesc(
    "active_tasks_gauge",
    "active_tasks_gauge",
    variableLabels, nil)
  cpusPercent = prometheus.NewDesc(
    "cpus_percent",
    "cpus_percent",
    variableLabels, nil)
  cpusTotal = prometheus.NewDesc(
    "cpus_total",
    "cpus_total",
    variableLabels, nil)
  cpusUsed = prometheus.NewDesc(
    "cpus_used",
    "cpus_used",
    variableLabels, nil)
  deactivatedSlaves = prometheus.NewDesc(
    "deactivated_slaves",
    "deactivated_slaves",
    variableLabels, nil)
  diskPercent = prometheus.NewDesc(
    "disk_percent",
    "disk_percent",
    variableLabels, nil)
  diskTotal = prometheus.NewDesc(
    "disk_total",
    "disk_total",
    variableLabels, nil)
  diskUsed = prometheus.NewDesc(
    "disk_used",
    "disk_used",
    variableLabels, nil)
  elected = prometheus.NewDesc(
    "elected",
    "elected",
    variableLabels, nil)
  failedTasks = prometheus.NewDesc(
    "failed_tasks",
    "failed_tasks",
    variableLabels, nil)
  finishedTasks = prometheus.NewDesc(
    "finished_tasks",
    "finished_tasks",
    variableLabels, nil)
  invalidStatusUpdates = prometheus.NewDesc(
    "invalid_status_updates",
    "invalid_status_updates",
    variableLabels, nil)
  killedTasks = prometheus.NewDesc(
    "killed_tasks",
    "killed_tasks",
    variableLabels, nil)
  lostTasks = prometheus.NewDesc(
    "lost_tasks",
    "lost_tasks",
    variableLabels, nil)
  masterCpusPercent = prometheus.NewDesc(
    "master_cpus_percent",
    "master_cpus_percent",
    variableLabels, nil)
  masterCpusTotal = prometheus.NewDesc(
    "master_cpus_total",
    "master_cpus_total",
    variableLabels, nil)
  masterCpusUsed = prometheus.NewDesc(
    "master_cpus_used",
    "master_cpus_used",
    variableLabels, nil)
  masterDiskPercent = prometheus.NewDesc(
    "master_disk_percent",
    "master_disk_percent",
    variableLabels, nil)
  masterDiskTotal = prometheus.NewDesc(
    "master_disk_total",
    "master_disk_total",
    variableLabels, nil)
  masterDiskUsed = prometheus.NewDesc(
    "master_disk_used",
    "master_disk_used",
    variableLabels, nil)
  masterDroppedMessages = prometheus.NewDesc(
    "master_dropped_messages",
    "master_dropped_messages",
    variableLabels, nil)
  masterElected = prometheus.NewDesc(
    "master_elected",
    "master_elected",
    variableLabels, nil)
  masterEventQueueDispatches = prometheus.NewDesc(
    "master_event_queue_dispatches",
    "master_event_queue_dispatches",
    variableLabels, nil)
  masterEventQueueHttpRequests = prometheus.NewDesc(
    "master_event_queue_http_requests",
    "master_event_queue_http_requests",
    variableLabels, nil)
  masterEventQueueMessages = prometheus.NewDesc(
    "master_event_queue_messages",
    "master_event_queue_messages",
    variableLabels, nil)
  masterFrameworksActive = prometheus.NewDesc(
    "master_frameworks_active",
    "master_frameworks_active",
    variableLabels, nil)
  masterFrameworksConnected = prometheus.NewDesc(
    "master_frameworks_connected",
    "master_frameworks_connected",
    variableLabels, nil)
  masterFrameworksDisconnected = prometheus.NewDesc(
    "master_frameworks_disconnected",
    "master_frameworks_disconnected",
    variableLabels, nil)
  masterFrameworksInactive = prometheus.NewDesc(
    "master_frameworks_inactive",
    "master_frameworks_inactive",
    variableLabels, nil)
  masterInvalidFrameworkToExecutorMessages = prometheus.NewDesc(
    "master_invalid_framework_to_executor_messages",
    "master_invalid_framework_to_executor_messages",
    variableLabels, nil)
  masterInvalidStatusUpdateAcknowledgements = prometheus.NewDesc(
    "master_invalid_status_update_acknowledgements",
    "master_invalid_status_update_acknowledgements",
    variableLabels, nil)
  masterInvalidStatusUpdates = prometheus.NewDesc(
    "master_invalid_status_updates",
    "master_invalid_status_updates",
    variableLabels, nil)
  masterMemPercent = prometheus.NewDesc(
    "master_mem_percent",
    "master_mem_percent",
    variableLabels, nil)
  masterMemTotal = prometheus.NewDesc(
    "master_mem_total",
    "master_mem_total",
    variableLabels, nil)
  masterMemUsed = prometheus.NewDesc(
    "master_mem_used",
    "master_mem_used",
    variableLabels, nil)
  masterMessagesAuthenticate = prometheus.NewDesc(
    "master_messages_authenticate",
    "master_messages_authenticate",
    variableLabels, nil)
  masterMessagesDeactivateFramework = prometheus.NewDesc(
    "master_messages_deactivate_framework",
    "master_messages_deactivate_framework",
    variableLabels, nil)
  masterMessagesDeclineOffers = prometheus.NewDesc(
    "master_messages_decline_offers",
    "master_messages_decline_offers",
    variableLabels, nil)
  masterMessagesExitedExecutor = prometheus.NewDesc(
    "master_messages_exited_executor",
    "master_messages_exited_executor",
    variableLabels, nil)
  masterMessagesFrameworkToExecutor = prometheus.NewDesc(
    "master_messages_framework_to_executor",
    "master_messages_framework_to_executor",
    variableLabels, nil)
  masterMessagesKillTask = prometheus.NewDesc(
    "master_messages_kill_task",
    "master_messages_kill_task",
    variableLabels, nil)
  masterMessagesLaunchTasks = prometheus.NewDesc(
    "master_messages_launch_tasks",
    "master_messages_launch_tasks",
    variableLabels, nil)
  masterMessagesReconcileTasks = prometheus.NewDesc(
    "master_messages_reconcile_tasks",
    "master_messages_reconcile_tasks",
    variableLabels, nil)
  masterMessagesRegisterFramework = prometheus.NewDesc(
    "master_messages_register_framework",
    "master_messages_register_framework",
    variableLabels, nil)
  masterMessagesRegisterSlave = prometheus.NewDesc(
    "master_messages_register_slave",
    "master_messages_register_slave",
    variableLabels, nil)
  masterMessagesReregisterFramework = prometheus.NewDesc(
    "master_messages_reregister_framework",
    "master_messages_reregister_framework",
    variableLabels, nil)
  masterMessagesReregisterSlave = prometheus.NewDesc(
    "master_messages_reregister_slave",
    "master_messages_reregister_slave",
    variableLabels, nil)
  masterMessagesResourceRequest = prometheus.NewDesc(
    "master_messages_resource_request",
    "master_messages_resource_request",
    variableLabels, nil)
  masterMessagesReviveOffers = prometheus.NewDesc(
    "master_messages_revive_offers",
    "master_messages_revive_offers",
    variableLabels, nil)
  masterMessagesStatusUpdate = prometheus.NewDesc(
    "master_messages_status_update",
    "master_messages_status_update",
    variableLabels, nil)
  masterMessagesStatusUpdateAcknowledgement = prometheus.NewDesc(
    "master_messages_status_update_acknowledgement",
    "master_messages_status_update_acknowledgement",
    variableLabels, nil)
  masterMessagesUnregisterFramework = prometheus.NewDesc(
    "master_messages_unregister_framework",
    "master_messages_unregister_framework",
    variableLabels, nil)
  masterMessagesUnregisterSlave = prometheus.NewDesc(
    "master_messages_unregister_slave",
    "master_messages_unregister_slave",
    variableLabels, nil)
  masterOutstandingOffers = prometheus.NewDesc(
    "master_outstanding_offers",
    "master_outstanding_offers",
    variableLabels, nil)
  masterRecoverySlaveRemovals = prometheus.NewDesc(
    "master_recovery_slave_removals",
    "master_recovery_slave_removals",
    variableLabels, nil)
  masterSlaveRegistrations = prometheus.NewDesc(
    "master_slave_registrations",
    "master_slave_registrations",
    variableLabels, nil)
  masterSlaveRemovals = prometheus.NewDesc(
    "master_slave_removals",
    "master_slave_removals",
    variableLabels, nil)
  masterSlaveReregistrations = prometheus.NewDesc(
    "master_slave_reregistrations",
    "master_slave_reregistrations",
    variableLabels, nil)
  masterSlaveShutdownsCanceled = prometheus.NewDesc(
    "master_slave_shutdowns_canceled",
    "master_slave_shutdowns_canceled",
    variableLabels, nil)
  masterSlaveShutdownsScheduled = prometheus.NewDesc(
    "master_slave_shutdowns_scheduled",
    "master_slave_shutdowns_scheduled",
    variableLabels, nil)
  masterSlavesActive = prometheus.NewDesc(
    "master_slaves_active",
    "master_slaves_active",
    variableLabels, nil)
  masterSlavesConnected = prometheus.NewDesc(
    "master_slaves_connected",
    "master_slaves_connected",
    variableLabels, nil)
  masterSlavesDisconnected = prometheus.NewDesc(
    "master_slaves_disconnected",
    "master_slaves_disconnected",
    variableLabels, nil)
  masterSlavesInactive = prometheus.NewDesc(
    "master_slaves_inactive",
    "master_slaves_inactive",
    variableLabels, nil)
  masterTaskErrorsourceMasterReasonTaskInvalid = prometheus.NewDesc(
    "master_task_error_source_master_reason_task_invalid",
    "master_task_error_source_master_reason_task_invalid",
    variableLabels, nil)
  masterTaskFailedsourceSlavereasonCommandExecutorFailed = prometheus.NewDesc(
    "master_task_failed_source_slave_reason_command_executor_failed",
    "master_task_failed_source_slave_reason_command_executor_failed",
    variableLabels, nil)
  masterTaskKilledsourceMasterReasonFrameworkRemoved = prometheus.NewDesc(
    "master_task_killed_source_master_reason_framework_removed",
    "master_task_killed_source_master_reason_framework_removed",
    variableLabels, nil)
  masterTaskKilledsourceSlavereasonExecutorUnregistered = prometheus.NewDesc(
    "master_task_killed_source_slave_reason_executor_unregistered",
    "master_task_killed_source_slave_reason_executor_unregistered",
    variableLabels, nil)
  masterTaskLostsourceMasterReasonSlaveRemoved = prometheus.NewDesc(
    "master_task_lost_source_master_reason_slave_removed",
    "master_task_lost_source_master_reason_slave_removed",
    variableLabels, nil)
  masterTaskLostsourceSlavereasonExecutorTerminated = prometheus.NewDesc(
    "master_task_lost_source_slave_reason_executor_terminated",
    "master_task_lost_source_slave_reason_executor_terminated",
    variableLabels, nil)
  masterTasksError = prometheus.NewDesc(
    "master_tasks_error",
    "master_tasks_error",
    variableLabels, nil)
  masterTasksFailed = prometheus.NewDesc(
    "master_tasks_failed",
    "master_tasks_failed",
    variableLabels, nil)
  masterTasksFinished = prometheus.NewDesc(
    "master_tasks_finished",
    "master_tasks_finished",
    variableLabels, nil)
  masterTasksKilled = prometheus.NewDesc(
    "master_tasks_killed",
    "master_tasks_killed",
    variableLabels, nil)
  masterTasksLost = prometheus.NewDesc(
    "master_tasks_lost",
    "master_tasks_lost",
    variableLabels, nil)
  masterTasksRunning = prometheus.NewDesc(
    "master_tasks_running",
    "master_tasks_running",
    variableLabels, nil)
  masterTasksStaging = prometheus.NewDesc(
    "master_tasks_staging",
    "master_tasks_staging",
    variableLabels, nil)
  masterTasksStarting = prometheus.NewDesc(
    "master_tasks_starting",
    "master_tasks_starting",
    variableLabels, nil)
  masterUptimeSecs = prometheus.NewDesc(
    "master_uptime_secs",
    "master_uptime_secs",
    variableLabels, nil)
  masterValidFrameworkToExecutorMessages = prometheus.NewDesc(
    "master_valid_framework_to_executor_messages",
    "master_valid_framework_to_executor_messages",
    variableLabels, nil)
  masterValidStatusUpdateAcknowledgements = prometheus.NewDesc(
    "master_valid_status_update_acknowledgements",
    "master_valid_status_update_acknowledgements",
    variableLabels, nil)
  masterValidStatusUpdates = prometheus.NewDesc(
    "master_valid_status_updates",
    "master_valid_status_updates",
    variableLabels, nil)
  memPercent = prometheus.NewDesc(
    "mem_percent",
    "mem_percent",
    variableLabels, nil)
  memTotal = prometheus.NewDesc(
    "mem_total",
    "mem_total",
    variableLabels, nil)
  memUsed = prometheus.NewDesc(
    "mem_used",
    "mem_used",
    variableLabels, nil)
  outstandingOffers = prometheus.NewDesc(
    "outstanding_offers",
    "outstanding_offers",
    variableLabels, nil)
  registrarQueuedOperations = prometheus.NewDesc(
    "registrar_queued_operations",
    "registrar_queued_operations",
    variableLabels, nil)
  registrarRegistrySizeBytes = prometheus.NewDesc(
    "registrar_registry_size_bytes",
    "registrar_registry_size_bytes",
    variableLabels, nil)
  registrarStateFetchMs = prometheus.NewDesc(
    "registrar_state_fetch_ms",
    "registrar_state_fetch_ms",
    variableLabels, nil)
  registrarStateStoreMs = prometheus.NewDesc(
    "registrar_state_store_ms",
    "registrar_state_store_ms",
    variableLabels, nil)
  registrarStateStoreMscount = prometheus.NewDesc(
    "registrar_state_store_ms_count",
    "registrar_state_store_ms_count",
    variableLabels, nil)
  registrarStateStoreMsmax = prometheus.NewDesc(
    "registrar_state_store_ms_max",
    "registrar_state_store_ms_max",
    variableLabels, nil)
  registrarStateStoreMsmin = prometheus.NewDesc(
    "registrar_state_store_ms_min",
    "registrar_state_store_ms_min",
    variableLabels, nil)
  registrarStateStoreMsp50 = prometheus.NewDesc(
    "registrar_state_store_ms_p50",
    "registrar_state_store_ms_p50",
    variableLabels, nil)
  registrarStateStoreMsp90 = prometheus.NewDesc(
    "registrar_state_store_ms_p90",
    "registrar_state_store_ms_p90",
    variableLabels, nil)
  registrarStateStoreMsp95 = prometheus.NewDesc(
    "registrar_state_store_ms_p95",
    "registrar_state_store_ms_p95",
    variableLabels, nil)
  registrarStateStoreMsp99 = prometheus.NewDesc(
    "registrar_state_store_ms_p99",
    "registrar_state_store_ms_p99",
    variableLabels, nil)
  registrarStateStoreMsp999 = prometheus.NewDesc(
    "registrar_state_store_ms_p999",
    "registrar_state_store_ms_p999",
    variableLabels, nil)
  registrarStateStoreMsp9999 = prometheus.NewDesc(
    "registrar_state_store_ms_p9999",
    "registrar_state_store_ms_p9999",
    variableLabels, nil)
  stagedTasks = prometheus.NewDesc(
    "staged_tasks",
    "staged_tasks",
    variableLabels, nil)
  startedTasks = prometheus.NewDesc(
    "started_tasks",
    "started_tasks",
    variableLabels, nil)
  systemCpusTotal = prometheus.NewDesc(
    "system_cpus_total",
    "system_cpus_total",
    variableLabels, nil)
  systemLoad15min = prometheus.NewDesc(
    "system_load_15min",
    "system_load_15min",
    variableLabels, nil)
  systemLoad1min = prometheus.NewDesc(
    "system_load_1min",
    "system_load_1min",
    variableLabels, nil)
  systemLoad5min = prometheus.NewDesc(
    "system_load_5min",
    "system_load_5min",
    variableLabels, nil)
  systemMemFreeBytes = prometheus.NewDesc(
    "system_mem_free_bytes",
    "system_mem_free_bytes",
    variableLabels, nil)
  systemMemTotalBytes = prometheus.NewDesc(
    "system_mem_total_bytes",
    "system_mem_total_bytes",
    variableLabels, nil)
  totalSchedulers = prometheus.NewDesc(
    "total_schedulers",
    "total_schedulers",
    variableLabels, nil)
  uptime = prometheus.NewDesc(
    "uptime",
    "uptime",
    variableLabels, nil)
  validStatusUpdates = prometheus.NewDesc(
    "valid_status_updates",
    "valid_status_updates",
    variableLabels, nil)
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

  go runEvery(e.scrapeMasters, e.opts.interval)

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
  stats, err := e.master.MesosMasterStats()
  if err != nil {
    log.Printf("%v\n", err)
    return
  }

  metricsChan <- prometheus.MustNewConstMetric(
    activatedSlaves,
    prometheus.GaugeValue,
    float64(stats.ActivatedSlaves),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    activeSchedulers,
    prometheus.GaugeValue,
    float64(stats.ActiveSchedulers),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    activeTasksGauge,
    prometheus.GaugeValue,
    float64(stats.ActiveTasksGauge),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    cpusPercent,
    prometheus.GaugeValue,
    float64(stats.CpusPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    cpusTotal,
    prometheus.GaugeValue,
    float64(stats.CpusTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    cpusUsed,
    prometheus.GaugeValue,
    float64(stats.CpusUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    deactivatedSlaves,
    prometheus.GaugeValue,
    float64(stats.DeactivatedSlaves),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    diskPercent,
    prometheus.GaugeValue,
    float64(stats.DiskPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    diskTotal,
    prometheus.GaugeValue,
    float64(stats.DiskTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    diskUsed,
    prometheus.GaugeValue,
    float64(stats.DiskUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    elected,
    prometheus.GaugeValue,
    float64(stats.Elected),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    failedTasks,
    prometheus.GaugeValue,
    float64(stats.FailedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    finishedTasks,
    prometheus.GaugeValue,
    float64(stats.FinishedTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    invalidStatusUpdates,
    prometheus.GaugeValue,
    float64(stats.InvalidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    killedTasks,
    prometheus.GaugeValue,
    float64(stats.KilledTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    lostTasks,
    prometheus.GaugeValue,
    float64(stats.LostTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterCpusPercent,
    prometheus.GaugeValue,
    float64(stats.MasterCpusPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterCpusTotal,
    prometheus.GaugeValue,
    float64(stats.MasterCpusTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterCpusUsed,
    prometheus.GaugeValue,
    float64(stats.MasterCpusUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterDiskPercent,
    prometheus.GaugeValue,
    float64(stats.MasterDiskPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterDiskTotal,
    prometheus.GaugeValue,
    float64(stats.MasterDiskTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterDiskUsed,
    prometheus.GaugeValue,
    float64(stats.MasterDiskUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterDroppedMessages,
    prometheus.GaugeValue,
    float64(stats.MasterDroppedMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterElected,
    prometheus.GaugeValue,
    float64(stats.MasterElected),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterEventQueueDispatches,
    prometheus.GaugeValue,
    float64(stats.MasterEventQueueDispatches),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterEventQueueHttpRequests,
    prometheus.GaugeValue,
    float64(stats.MasterEventQueueHttpRequests),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterEventQueueMessages,
    prometheus.GaugeValue,
    float64(stats.MasterEventQueueMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterFrameworksActive,
    prometheus.GaugeValue,
    float64(stats.MasterFrameworksActive),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterFrameworksConnected,
    prometheus.GaugeValue,
    float64(stats.MasterFrameworksConnected),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterFrameworksDisconnected,
    prometheus.GaugeValue,
    float64(stats.MasterFrameworksDisconnected),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterFrameworksInactive,
    prometheus.GaugeValue,
    float64(stats.MasterFrameworksInactive),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterInvalidFrameworkToExecutorMessages,
    prometheus.GaugeValue,
    float64(stats.MasterInvalidFrameworkToExecutorMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterInvalidStatusUpdateAcknowledgements,
    prometheus.GaugeValue,
    float64(stats.MasterInvalidStatusUpdateAcknowledgements),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterInvalidStatusUpdates,
    prometheus.GaugeValue,
    float64(stats.MasterInvalidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMemPercent,
    prometheus.GaugeValue,
    float64(stats.MasterMemPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMemTotal,
    prometheus.GaugeValue,
    float64(stats.MasterMemTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMemUsed,
    prometheus.GaugeValue,
    float64(stats.MasterMemUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesAuthenticate,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesAuthenticate),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesDeactivateFramework,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesDeactivateFramework),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesDeclineOffers,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesDeclineOffers),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesExitedExecutor,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesExitedExecutor),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesFrameworkToExecutor,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesFrameworkToExecutor),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesKillTask,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesKillTask),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesLaunchTasks,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesLaunchTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesReconcileTasks,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesReconcileTasks),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesRegisterFramework,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesRegisterFramework),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesRegisterSlave,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesRegisterSlave),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesReregisterFramework,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesReregisterFramework),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesReregisterSlave,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesReregisterSlave),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesResourceRequest,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesResourceRequest),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesReviveOffers,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesReviveOffers),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesStatusUpdate,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesStatusUpdate),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesStatusUpdateAcknowledgement,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesStatusUpdateAcknowledgement),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesUnregisterFramework,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesUnregisterFramework),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterMessagesUnregisterSlave,
    prometheus.GaugeValue,
    float64(stats.MasterMessagesUnregisterSlave),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterOutstandingOffers,
    prometheus.GaugeValue,
    float64(stats.MasterOutstandingOffers),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterRecoverySlaveRemovals,
    prometheus.GaugeValue,
    float64(stats.MasterRecoverySlaveRemovals),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlaveRegistrations,
    prometheus.GaugeValue,
    float64(stats.MasterSlaveRegistrations),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlaveRemovals,
    prometheus.GaugeValue,
    float64(stats.MasterSlaveRemovals),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlaveReregistrations,
    prometheus.GaugeValue,
    float64(stats.MasterSlaveReregistrations),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlaveShutdownsCanceled,
    prometheus.GaugeValue,
    float64(stats.MasterSlaveShutdownsCanceled),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlaveShutdownsScheduled,
    prometheus.GaugeValue,
    float64(stats.MasterSlaveShutdownsScheduled),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlavesActive,
    prometheus.GaugeValue,
    float64(stats.MasterSlavesActive),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlavesConnected,
    prometheus.GaugeValue,
    float64(stats.MasterSlavesConnected),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlavesDisconnected,
    prometheus.GaugeValue,
    float64(stats.MasterSlavesDisconnected),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterSlavesInactive,
    prometheus.GaugeValue,
    float64(stats.MasterSlavesInactive),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTaskErrorsourceMasterReasonTaskInvalid,
    prometheus.GaugeValue,
    float64(stats.MasterTaskErrorsourceMasterReasonTaskInvalid),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTaskFailedsourceSlavereasonCommandExecutorFailed,
    prometheus.GaugeValue,
    float64(stats.MasterTaskFailedsourceSlavereasonCommandExecutorFailed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTaskKilledsourceMasterReasonFrameworkRemoved,
    prometheus.GaugeValue,
    float64(stats.MasterTaskKilledsourceMasterReasonFrameworkRemoved),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTaskKilledsourceSlavereasonExecutorUnregistered,
    prometheus.GaugeValue,
    float64(stats.MasterTaskKilledsourceSlavereasonExecutorUnregistered),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTaskLostsourceMasterReasonSlaveRemoved,
    prometheus.GaugeValue,
    float64(stats.MasterTaskLostsourceMasterReasonSlaveRemoved),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTaskLostsourceSlavereasonExecutorTerminated,
    prometheus.GaugeValue,
    float64(stats.MasterTaskLostsourceSlavereasonExecutorTerminated),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksError,
    prometheus.GaugeValue,
    float64(stats.MasterTasksError),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksFailed,
    prometheus.GaugeValue,
    float64(stats.MasterTasksFailed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksFinished,
    prometheus.GaugeValue,
    float64(stats.MasterTasksFinished),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksKilled,
    prometheus.GaugeValue,
    float64(stats.MasterTasksKilled),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksLost,
    prometheus.GaugeValue,
    float64(stats.MasterTasksLost),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksRunning,
    prometheus.GaugeValue,
    float64(stats.MasterTasksRunning),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksStaging,
    prometheus.GaugeValue,
    float64(stats.MasterTasksStaging),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterTasksStarting,
    prometheus.GaugeValue,
    float64(stats.MasterTasksStarting),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterUptimeSecs,
    prometheus.CounterValue,
    float64(stats.MasterUptimeSecs),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterValidFrameworkToExecutorMessages,
    prometheus.GaugeValue,
    float64(stats.MasterValidFrameworkToExecutorMessages),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterValidStatusUpdateAcknowledgements,
    prometheus.GaugeValue,
    float64(stats.MasterValidStatusUpdateAcknowledgements),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    masterValidStatusUpdates,
    prometheus.GaugeValue,
    float64(stats.MasterValidStatusUpdates),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    memPercent,
    prometheus.GaugeValue,
    float64(stats.MemPercent),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    memTotal,
    prometheus.GaugeValue,
    float64(stats.MemTotal),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    memUsed,
    prometheus.GaugeValue,
    float64(stats.MemUsed),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    outstandingOffers,
    prometheus.GaugeValue,
    float64(stats.OutstandingOffers),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarQueuedOperations,
    prometheus.GaugeValue,
    float64(stats.RegistrarQueuedOperations),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarRegistrySizeBytes,
    prometheus.GaugeValue,
    float64(stats.RegistrarRegistrySizeBytes),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateFetchMs,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateFetchMs),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMs,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMs),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMscount,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMscount),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsmax,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsmax),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsmin,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsmin),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsp50,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsp50),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsp90,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsp90),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsp95,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsp95),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsp99,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsp99),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsp999,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsp999),
    hostname,
  )
  metricsChan <- prometheus.MustNewConstMetric(
    registrarStateStoreMsp9999,
    prometheus.GaugeValue,
    float64(stats.RegistrarStateStoreMsp9999),
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
    totalSchedulers,
    prometheus.GaugeValue,
    float64(stats.TotalSchedulers),
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
    prometheus.GaugeValue,
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

  if err := prometheus.PushCollectors("master_stats_json", hostname, e.opts.pushGatewayURL, e); err != nil {
    log.Printf("Could not push completion time to Pushgateway: %v\n", err)
  }
}

func (e *periodicStatsExporter) scrapeMasters() {
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
