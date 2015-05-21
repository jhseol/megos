package megos

import (
  "encoding/json"
  "fmt"
)

// Mesos version 0.22.0 compatible

type MesosCompletedTasks struct {
  ExecutorID  string        `json:"executor_id,omitempty"`
  FrameworkID string        `json:"framework_id,omitempty"`
  ID          string        `json:"id,omitempty"`
  Labels      []interface{} `json:"labels,omitempty"`
  Name        string        `json:"name,omitempty"`
  Resources   interface{}   `json:"resources,omitempty"`
  SlaveID     string        `json:"slave_id,omitempty"`
  State       string        `json:"state,omitempty"`
  Statuses    []interface{} `json:"statuses,omitempty"`
}

type MesosCompletedExecutors struct {
  CompletedTasks []*MesosCompletedTasks `json:"completed_tasks,omitempty"`
  Container      string                 `json:"container,omitempty"`
  Directory      string                 `json:"directory,omitempty"`
  ID             string                 `json:"id,omitempty"`
  Name           string                 `json:"name,omitempty"`
  QueuedTasks    []interface{}          `json:"queued_tasks,omitempty"`
  Resources      interface{}            `json:"resources,omitempty"`
  Source         string                 `json:"source,omitempty"`
  Tasks          []interface{}          `json:"tasks,omitempty"`
}

type MesosCompletedFramework struct {
  Checkpoint         bool                       `json:"checkpoint,omitempty"`
  CompletedExecutors []*MesosCompletedExecutors `json:"completed_executors,omitempty"`
  Executors          []interface{}              `json:"executors,omitempty"`
  FailoverTimeout    int64                      `json:"failover_timeout,omitempty"`
  Hostname           string                     `json:"hostname,omitempty"`
  ID                 string                     `json:"id,omitempty"`
  Name               string                     `json:"name,omitempty"`
  Role               string                     `json:"role,omitempty"`
  User               string                     `json:"user,omitempty"`
}

type MesosExecutors struct {
  CompletedTasks []interface{} `json:"completed_tasks,omitempty"`
  Container      string        `json:"container,omitempty"`
  Directory      string        `json:"directory,omitempty"`
  ID             string        `json:"id,omitempty"`
  Name           string        `json:"name,omitempty"`
  QueuedTasks    []interface{} `json:"queued_tasks,omitempty"`
  Resources      interface{}   `json:"resources,omitempty"`
  Source         string        `json:"source,omitempty"`
  Tasks          []interface{} `json:"tasks,omitempty"`
}

type MesosFramework struct {
  Checkpoint         bool              `json:"checkpoint,omitempty"`
  CompletedExecutors []interface{}     `json:"completed_executors,omitempty"`
  Executors          []*MesosExecutors `json:"executors,omitempty"`
  FailoverTimeout    int64             `json:"failover_timeout,omitempty"`
  Hostname           string            `json:"hostname,omitempty"`
  ID                 string            `json:"id,omitempty"`
  Name               string            `json:"name,omitempty"`
  Role               string            `json:"role,omitempty"`
  User               string            `json:"user,omitempty"`
}

type MesosSlaveFlags struct {
  Authenticatee               string `json:"authenticatee,omitempty"`
  CgroupsEnableCfs            string `json:"cgroups_enable_cfs,omitempty"`
  CgroupsHierarchy            string `json:"cgroups_hierarchy,omitempty"`
  CgroupsLimitSwap            string `json:"cgroups_limit_swap,omitempty"`
  CgroupsRoot                 string `json:"cgroups_root,omitempty"`
  ContainerDiskWatchInterval  string `json:"container_disk_watch_interval,omitempty"`
  Containerizers              string `json:"containerizers,omitempty"`
  DefaultRole                 string `json:"default_role,omitempty"`
  DiskWatchInterval           string `json:"disk_watch_interval,omitempty"`
  Docker                      string `json:"docker,omitempty"`
  DockerRemoveDelay           string `json:"docker_remove_delay,omitempty"`
  DockerSandboxDirectory      string `json:"docker_sandbox_directory,omitempty"`
  DockerStopTimeout           string `json:"docker_stop_timeout,omitempty"`
  EnforceContainerDiskQuota   string `json:"enforce_container_disk_quota,omitempty"`
  ExecutorRegistrationTimeout string `json:"executor_registration_timeout,omitempty"`
  ExecutorShutdownGracePeriod string `json:"executor_shutdown_grace_period,omitempty"`
  FrameworksHome              string `json:"frameworks_home,omitempty"`
  GcDelay                     string `json:"gc_delay,omitempty"`
  GcDiskHeadroom              string `json:"gc_disk_headroom,omitempty"`
  HadoopHome                  string `json:"hadoop_home,omitempty"`
  Help                        string `json:"help,omitempty"`
  InitializeDriverLogging     string `json:"initialize_driver_logging,omitempty"`
  Isolation                   string `json:"isolation,omitempty"`
  LauncherDir                 string `json:"launcher_dir,omitempty"`
  Logbufsecs                  string `json:"logbufsecs,omitempty"`
  LoggingLevel                string `json:"logging_level,omitempty"`
  Master                      string `json:"master,omitempty"`
  PerfDuration                string `json:"perf_duration,omitempty"`
  PerfInterval                string `json:"perf_interval,omitempty"`
  Port                        string `json:"port,omitempty"`
  Quiet                       string `json:"quiet,omitempty"`
  Recover                     string `json:"recover,omitempty"`
  RecoveryTimeout             string `json:"recovery_timeout,omitempty"`
  RegistrationBackoffFactor   string `json:"registration_backoff_factor,omitempty"`
  ResourceMonitoringInterval  string `json:"resource_monitoring_interval,omitempty"`
  Strict                      string `json:"strict,omitempty"`
  SwitchUser                  string `json:"switch_user,omitempty"`
  Version                     string `json:"version,omitempty"`
  WorkDir                     string `json:"work_dir,omitempty"`
}

type MesosSlaveState struct {
  Attributes          interface{}                `json:"attributes,omitempty"`
  BuildDate           string                     `json:"build_date,omitempty"`
  BuildTime           int64                      `json:"build_time,omitempty"`
  BuildUser           string                     `json:"build_user,omitempty"`
  CompletedFrameworks []*MesosCompletedFramework `json:"completed_frameworks,omitempty"`
  FailedTasks         int64                      `json:"failed_tasks,omitempty"`
  FinishedTasks       int64                      `json:"finished_tasks,omitempty"`
  Flags               *MesosSlaveFlags           `json:"flags,omitempty"`
  Frameworks          []*MesosFramework          `json:"frameworks,omitempty"`
  Hostname            string                     `json:"hostname,omitempty"`
  ID                  string                     `json:"id,omitempty"`
  KilledTasks         int64                      `json:"killed_tasks,omitempty"`
  LostTasks           int64                      `json:"lost_tasks,omitempty"`
  MasterHostname      string                     `json:"master_hostname,omitempty"`
  Pid                 string                     `json:"pid,omitempty"`
  Resources           interface{}                `json:"resources,omitempty"`
  StagedTasks         int64                      `json:"staged_tasks,omitempty"`
  StartTime           float64                    `json:"start_time,omitempty"`
  StartedTasks        int64                      `json:"started_tasks,omitempty"`
  Version             string                     `json:"version,omitempty"`
}

type MesosSlaveStats struct {
  FailedTasks                   int64   `json:"failed_tasks,omitempty"`
  FinishedTasks                 int64   `json:"finished_tasks,omitempty"`
  InvalidStatusUpdates          int64   `json:"invalid_status_updates,omitempty"`
  KilledTasks                   int64   `json:"killed_tasks,omitempty"`
  LaunchedTasksGauge            int64   `json:"launched_tasks_gauge,omitempty"`
  LostTasks                     int64   `json:"lost_tasks,omitempty"`
  QueuedTasksGauge              int64   `json:"queued_tasks_gauge,omitempty"`
  RecoveryErrors                int64   `json:"recovery_errors,omitempty"`
  Registered                    int64   `json:"registered,omitempty"`
  SlaveCpusPercent              float64 `json:"slave/cpus_percent,omitempty"`
  SlaveCpusTotal                float64 `json:"slave/cpus_total,omitempty"`
  SlaveCpusUsed                 float64 `json:"slave/cpus_used,omitempty"`
  SlaveDiskPercent              float64 `json:"slave/disk_percent,omitempty"`
  SlaveDiskTotal                float64 `json:"slave/disk_total,omitempty"`
  SlaveDiskUsed                 float64 `json:"slave/disk_used,omitempty"`
  SlaveExecutorsRegistering     int64   `json:"slave/executors_registering,omitempty"`
  SlaveExecutorsRunning         int64   `json:"slave/executors_running,omitempty"`
  SlaveExecutorsTerminated      int64   `json:"slave/executors_terminated,omitempty"`
  SlaveExecutorsTerminating     int64   `json:"slave/executors_terminating,omitempty"`
  SlaveFrameworksActive         int64   `json:"slave/frameworks_active,omitempty"`
  SlaveInvalidFrameworkMessages int64   `json:"slave/invalid_framework_messages,omitempty"`
  SlaveInvalidStatusUpdates     int64   `json:"slave/invalid_status_updates,omitempty"`
  SlaveMemPercent               float64 `json:"slave/mem_percent,omitempty"`
  SlaveMemTotal                 float64 `json:"slave/mem_total,omitempty"`
  SlaveMemUsed                  float64 `json:"slave/mem_used,omitempty"`
  SlaveRecoveryErrors           int64   `json:"slave/recovery_errors,omitempty"`
  SlaveRegistered               int64   `json:"slave/registered,omitempty"`
  SlaveTasksFailed              int64   `json:"slave/tasks_failed,omitempty"`
  SlaveTasksFinished            int64   `json:"slave/tasks_finished,omitempty"`
  SlaveTasksKilled              int64   `json:"slave/tasks_killed,omitempty"`
  SlaveTasksLost                int64   `json:"slave/tasks_lost,omitempty"`
  SlaveTasksRunning             int64   `json:"slave/tasks_running,omitempty"`
  SlaveTasksStaging             int64   `json:"slave/tasks_staging,omitempty"`
  SlaveTasksStarting            int64   `json:"slave/tasks_starting,omitempty"`
  SlaveUptimeSecs               float64 `json:"slave/uptime_secs,omitempty"`
  SlaveValidFrameworkMessages   int64   `json:"slave/valid_framework_messages,omitempty"`
  SlaveValidStatusUpdates       int64   `json:"slave/valid_status_updates,omitempty"`
  StagedTasks                   int64   `json:"staged_tasks,omitempty"`
  StartedTasks                  int64   `json:"started_tasks,omitempty"`
  SystemCpusTotal               int64   `json:"system/cpus_total,omitempty"`
  SystemLoad15min               float64 `json:"system/load_15min,omitempty"`
  SystemLoad1min                float64 `json:"system/load_1min,omitempty"`
  SystemLoad5min                float64 `json:"system/load_5min,omitempty"`
  SystemMemFreeBytes            float64 `json:"system/mem_free_bytes,omitempty"`
  SystemMemTotalBytes           float64 `json:"system/mem_total_bytes,omitempty"`
  TotalFrameworks               int64   `json:"total_frameworks,omitempty"`
  Uptime                        float64 `json:"uptime,omitempty"`
  ValidStatusUpdates            int64   `json:"valid_status_updates,omitempty"`
}

type MesosMonitorStatistics struct {
  CpusLimit          float64 `json:"cpus_limit,omitempty"`
  CpusSystemTimeSecs float64 `json:"cpus_system_time_secs,omitempty"`
  CpusUserTimeSecs   float64 `json:"cpus_user_time_secs,omitempty"`
  MemLimitBytes      int64   `json:"mem_limit_bytes,omitempty"`
  MemRssBytes        int64   `json:"mem_rss_bytes,omitempty"`
  Timestamp          float64 `json:"timestamp,omitempty"`
}

type MesosMonitor struct {
  ExecutorID   string                  `json:"executor_id,omitempty"`
  ExecutorName string                  `json:"executor_name,omitempty"`
  FrameworkID  string                  `json:"framework_id,omitempty"`
  Source       string                  `json:"source,omitempty"`
  Statistics   *MesosMonitorStatistics `json:"statistics,omitempty"`
}

type MesosSlaveMonitorStatistics []*MesosMonitor

type MesosSlaveOptions struct {
  Host string `json:"mesos_slave_url,omitempty"`
}

type MesosSlaveClient struct {
  opts *MesosSlaveOptions
}

func NewSlaveClient(opts *MesosSlaveOptions) *MesosSlaveClient {
  if opts.Host == "" {
    opts = &MesosSlaveOptions{
      Host: MesosSlaveHostDefault,
    }
  }

  s := &MesosSlaveClient{
    opts: opts,
  }

  return s
}

func (ms *MesosSlaveClient) MesosSystemStats() (*MesosSystemStats, error) {
  body, err := httpGetRequest(ms.opts.Host+MesosSystemStatsURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosSystemStats{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Printf("%v: %s\n", err, string(body))
    return nil, err
  }

  return s, nil
}

func (ms *MesosSlaveClient) MesosSlaveStats() (*MesosSlaveStats, error) {
  body, err := httpGetRequest(ms.opts.Host+MesosSlaveStatsURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosSlaveStats{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Printf("%v: %s\n", err, string(body))
    return nil, err
  }

  return s, nil
}

func (ms *MesosSlaveClient) MesosSlaveState() (*MesosSlaveState, error) {
  body, err := httpGetRequest(ms.opts.Host+MesosSlaveStateURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosSlaveState{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Printf("%v: %s\n", err, string(body))
    return nil, err
  }

  return s, nil
}

func (ms *MesosSlaveClient) MesosSlaveMonitorStatistics() (*MesosSlaveMonitorStatistics, error) {
  body, err := httpGetRequest(ms.opts.Host+MesosSlaveMonitorStatisticsURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosSlaveMonitorStatistics{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Printf("%v: %s\n", err, string(body))
    return nil, err
  }

  return s, nil
}

func (ms *MesosSlaveClient) MesosSlaveHealth() error {
  statusCode, err := httpHeadRequest(ms.opts.Host+MesosSlaveHealthURI, Query{})
  if err != nil && statusCode != 200 {
    fmt.Printf("%v: \n", err, statusCode)
  }
  return err
}
