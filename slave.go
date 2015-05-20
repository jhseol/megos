package megos

import (
  "encoding/json"
  "fmt"
)

// Mesos version 0.22.0 compatible

type MesosCompletedTasks struct {
  ExecutorID  string        `json:"executor_id"`
  FrameworkID string        `json:"framework_id"`
  ID          string        `json:"id"`
  Labels      []interface{} `json:"labels"`
  Name        string        `json:"name"`
  Resources   interface{}   `json:"resources"`
  SlaveID     string        `json:"slave_id"`
  State       string        `json:"state"`
  Statuses    []interface{} `json:"statuses"`
}

type MesosCompletedExecutors struct {
  CompletedTasks []MesosCompletedTasks `json:"completed_tasks"`
  Container      string                `json:"container"`
  Directory      string                `json:"directory"`
  ID             string                `json:"id"`
  Name           string                `json:"name"`
  QueuedTasks    []interface{}         `json:"queued_tasks"`
  Resources      interface{}           `json:"resources"`
  Source         string                `json:"source"`
  Tasks          []interface{}         `json:"tasks"`
}

type MesosCompletedFramework struct {
  Checkpoint         bool                      `json:"checkpoint"`
  CompletedExecutors []MesosCompletedExecutors `json:"completed_executors"`
  Executors          []interface{}             `json:"executors"`
  FailoverTimeout    int64                     `json:"failover_timeout"`
  Hostname           string                    `json:"hostname"`
  ID                 string                    `json:"id"`
  Name               string                    `json:"name"`
  Role               string                    `json:"role"`
  User               string                    `json:"user"`
}

type MesosExecutors struct {
  CompletedTasks []interface{} `json:"completed_tasks"`
  Container      string        `json:"container"`
  Directory      string        `json:"directory"`
  ID             string        `json:"id"`
  Name           string        `json:"name"`
  QueuedTasks    []interface{} `json:"queued_tasks"`
  Resources      interface{}   `json:"resources"`
  Source         string        `json:"source"`
  Tasks          []interface{} `json:"tasks"`
}

type MesosFramework struct {
  Checkpoint         bool             `json:"checkpoint"`
  CompletedExecutors []interface{}    `json:"completed_executors"`
  Executors          []MesosExecutors `json:"executors"`
  FailoverTimeout    int64            `json:"failover_timeout"`
  Hostname           string           `json:"hostname"`
  ID                 string           `json:"id"`
  Name               string           `json:"name"`
  Role               string           `json:"role"`
  User               string           `json:"user"`
}

type MesosSlaveFlags struct {
  Authenticatee               string `json:"authenticatee"`
  CgroupsEnableCfs            string `json:"cgroups_enable_cfs"`
  CgroupsHierarchy            string `json:"cgroups_hierarchy"`
  CgroupsLimitSwap            string `json:"cgroups_limit_swap"`
  CgroupsRoot                 string `json:"cgroups_root"`
  ContainerDiskWatchInterval  string `json:"container_disk_watch_interval"`
  Containerizers              string `json:"containerizers"`
  DefaultRole                 string `json:"default_role"`
  DiskWatchInterval           string `json:"disk_watch_interval"`
  Docker                      string `json:"docker"`
  DockerRemoveDelay           string `json:"docker_remove_delay"`
  DockerSandboxDirectory      string `json:"docker_sandbox_directory"`
  DockerStopTimeout           string `json:"docker_stop_timeout"`
  EnforceContainerDiskQuota   string `json:"enforce_container_disk_quota"`
  ExecutorRegistrationTimeout string `json:"executor_registration_timeout"`
  ExecutorShutdownGracePeriod string `json:"executor_shutdown_grace_period"`
  FrameworksHome              string `json:"frameworks_home"`
  GcDelay                     string `json:"gc_delay"`
  GcDiskHeadroom              string `json:"gc_disk_headroom"`
  HadoopHome                  string `json:"hadoop_home"`
  Help                        string `json:"help"`
  InitializeDriverLogging     string `json:"initialize_driver_logging"`
  Isolation                   string `json:"isolation"`
  LauncherDir                 string `json:"launcher_dir"`
  Logbufsecs                  string `json:"logbufsecs"`
  LoggingLevel                string `json:"logging_level"`
  Master                      string `json:"master"`
  PerfDuration                string `json:"perf_duration"`
  PerfInterval                string `json:"perf_interval"`
  Port                        string `json:"port"`
  Quiet                       string `json:"quiet"`
  Recover                     string `json:"recover"`
  RecoveryTimeout             string `json:"recovery_timeout"`
  RegistrationBackoffFactor   string `json:"registration_backoff_factor"`
  ResourceMonitoringInterval  string `json:"resource_monitoring_interval"`
  Strict                      string `json:"strict"`
  SwitchUser                  string `json:"switch_user"`
  Version                     string `json:"version"`
  WorkDir                     string `json:"work_dir"`
}

type MesosSlaveState struct {
  Attributes          interface{}               `json:"attributes"`
  BuildDate           string                    `json:"build_date"`
  BuildTime           int64                     `json:"build_time"`
  BuildUser           string                    `json:"build_user"`
  CompletedFrameworks []MesosCompletedFramework `json:"completed_frameworks"`
  FailedTasks         int64                     `json:"failed_tasks"`
  FinishedTasks       int64                     `json:"finished_tasks"`
  Flags               MesosSlaveFlags           `json:"flags"`
  Frameworks          []MesosFramework          `json:"frameworks"`
  Hostname            string                    `json:"hostname"`
  ID                  string                    `json:"id"`
  KilledTasks         int64                     `json:"killed_tasks"`
  LostTasks           int64                     `json:"lost_tasks"`
  MasterHostname      string                    `json:"master_hostname"`
  Pid                 string                    `json:"pid"`
  Resources           interface{}               `json:"resources"`
  StagedTasks         int64                     `json:"staged_tasks"`
  StartTime           float64                   `json:"start_time"`
  StartedTasks        int64                     `json:"started_tasks"`
  Version             string                    `json:"version"`
}

type MesosSlaveStats struct {
  FailedTasks                   int64   `json:"failed_tasks"`
  FinishedTasks                 int64   `json:"finished_tasks"`
  InvalidStatusUpdates          int64   `json:"invalid_status_updates"`
  KilledTasks                   int64   `json:"killed_tasks"`
  LaunchedTasksGauge            int64   `json:"launched_tasks_gauge"`
  LostTasks                     int64   `json:"lost_tasks"`
  QueuedTasksGauge              int64   `json:"queued_tasks_gauge"`
  RecoveryErrors                int64   `json:"recovery_errors"`
  Registered                    int64   `json:"registered"`
  SlaveCpusPercent              float64 `json:"slave/cpus_percent"`
  SlaveCpusTotal                float64 `json:"slave/cpus_total"`
  SlaveCpusUsed                 float64 `json:"slave/cpus_used"`
  SlaveDiskPercent              float64 `json:"slave/disk_percent"`
  SlaveDiskTotal                float64 `json:"slave/disk_total"`
  SlaveDiskUsed                 float64 `json:"slave/disk_used"`
  SlaveExecutorsRegistering     int64   `json:"slave/executors_registering"`
  SlaveExecutorsRunning         int64   `json:"slave/executors_running"`
  SlaveExecutorsTerminated      int64   `json:"slave/executors_terminated"`
  SlaveExecutorsTerminating     int64   `json:"slave/executors_terminating"`
  SlaveFrameworksActive         int64   `json:"slave/frameworks_active"`
  SlaveInvalidFrameworkMessages int64   `json:"slave/invalid_framework_messages"`
  SlaveInvalidStatusUpdates     int64   `json:"slave/invalid_status_updates"`
  SlaveMemPercent               float64 `json:"slave/mem_percent"`
  SlaveMemTotal                 float64 `json:"slave/mem_total"`
  SlaveMemUsed                  float64 `json:"slave/mem_used"`
  SlaveRecoveryErrors           int64   `json:"slave/recovery_errors"`
  SlaveRegistered               int64   `json:"slave/registered"`
  SlaveTasksFailed              int64   `json:"slave/tasks_failed"`
  SlaveTasksFinished            int64   `json:"slave/tasks_finished"`
  SlaveTasksKilled              int64   `json:"slave/tasks_killed"`
  SlaveTasksLost                int64   `json:"slave/tasks_lost"`
  SlaveTasksRunning             int64   `json:"slave/tasks_running"`
  SlaveTasksStaging             int64   `json:"slave/tasks_staging"`
  SlaveTasksStarting            int64   `json:"slave/tasks_starting"`
  SlaveUptimeSecs               float64 `json:"slave/uptime_secs"`
  SlaveValidFrameworkMessages   int64   `json:"slave/valid_framework_messages"`
  SlaveValidStatusUpdates       int64   `json:"slave/valid_status_updates"`
  StagedTasks                   int64   `json:"staged_tasks"`
  StartedTasks                  int64   `json:"started_tasks"`
  SystemCpusTotal               int64   `json:"system/cpus_total"`
  SystemLoad15min               float64 `json:"system/load_15min"`
  SystemLoad1min                float64 `json:"system/load_1min"`
  SystemLoad5min                float64 `json:"system/load_5min"`
  SystemMemFreeBytes            float64 `json:"system/mem_free_bytes"`
  SystemMemTotalBytes           float64 `json:"system/mem_total_bytes"`
  TotalFrameworks               int64   `json:"total_frameworks"`
  Uptime                        float64 `json:"uptime"`
  ValidStatusUpdates            int64   `json:"valid_status_updates"`
}

type MesosMonitorStatistics struct {
  CpusLimit          float64 `json:"cpus_limit"`
  CpusSystemTimeSecs float64 `json:"cpus_system_time_secs"`
  CpusUserTimeSecs   float64 `json:"cpus_user_time_secs"`
  MemLimitBytes      int64   `json:"mem_limit_bytes"`
  MemRssBytes        int64   `json:"mem_rss_bytes"`
  Timestamp          float64 `json:"timestamp"`
}

type MesosMonitor struct {
  ExecutorID   string                 `json:"executor_id"`
  ExecutorName string                 `json:"executor_name"`
  FrameworkID  string                 `json:"framework_id"`
  Source       string                 `json:"source"`
  Statistics   MesosMonitorStatistics `json:"statistics"`
}

type MesosSlaveMonitorStatistics []MesosMonitor

type MesosSlaveOptions struct {
  Host string `json:"mesos_slave_url"`
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
