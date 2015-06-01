package megos

import (
  "encoding/json"
  "fmt"
)

type MesosMasterTasks struct {
  Tasks []*MesosTasks `json:"tasks,omitempty"`
}

type MesosMasterStats struct {
  ActivatedSlaves                                        int64   `json:"activated_slaves,omitempty"`
  ActiveSchedulers                                       int64   `json:"active_schedulers,omitempty"`
  ActiveTasksGauge                                       int64   `json:"active_tasks_gauge,omitempty"`
  CpusPercent                                            float64 `json:"cpus_percent,omitempty"`
  CpusTotal                                              float64 `json:"cpus_total,omitempty"`
  CpusUsed                                               float64 `json:"cpus_used,omitempty"`
  DeactivatedSlaves                                      int64   `json:"deactivated_slaves,omitempty"`
  DiskPercent                                            float64 `json:"disk_percent,omitempty"`
  DiskTotal                                              float64 `json:"disk_total,omitempty"`
  DiskUsed                                               float64 `json:"disk_used,omitempty"`
  Elected                                                int64   `json:"elected,omitempty"`
  FailedTasks                                            int64   `json:"failed_tasks,omitempty"`
  FinishedTasks                                          int64   `json:"finished_tasks,omitempty"`
  InvalidStatusUpdates                                   int64   `json:"invalid_status_updates,omitempty"`
  KilledTasks                                            int64   `json:"killed_tasks,omitempty"`
  LostTasks                                              int64   `json:"lost_tasks,omitempty"`
  MasterCpusPercent                                      float64 `json:"master/cpus_percent,omitempty"`
  MasterCpusTotal                                        float64 `json:"master/cpus_total,omitempty"`
  MasterCpusUsed                                         float64 `json:"master/cpus_used,omitempty"`
  MasterDiskPercent                                      float64 `json:"master/disk_percent,omitempty"`
  MasterDiskTotal                                        float64 `json:"master/disk_total,omitempty"`
  MasterDiskUsed                                         float64 `json:"master/disk_used,omitempty"`
  MasterDroppedMessages                                  int64   `json:"master/dropped_messages,omitempty"`
  MasterElected                                          int64   `json:"master/elected,omitempty"`
  MasterEventQueueDispatches                             int64   `json:"master/event_queue_dispatches,omitempty"`
  MasterEventQueueHttpRequests                           int64   `json:"master/event_queue_http_requests,omitempty"`
  MasterEventQueueMessages                               int64   `json:"master/event_queue_messages,omitempty"`
  MasterFrameworksActive                                 int64   `json:"master/frameworks_active,omitempty"`
  MasterFrameworksConnected                              int64   `json:"master/frameworks_connected,omitempty"`
  MasterFrameworksDisconnected                           int64   `json:"master/frameworks_disconnected,omitempty"`
  MasterFrameworksInactive                               int64   `json:"master/frameworks_inactive,omitempty"`
  MasterInvalidFrameworkToExecutorMessages               int64   `json:"master/invalid_framework_to_executor_messages,omitempty"`
  MasterInvalidStatusUpdateAcknowledgements              int64   `json:"master/invalid_status_update_acknowledgements,omitempty"`
  MasterInvalidStatusUpdates                             int64   `json:"master/invalid_status_updates,omitempty"`
  MasterMemPercent                                       float64 `json:"master/mem_percent,omitempty"`
  MasterMemTotal                                         float64 `json:"master/mem_total,omitempty"`
  MasterMemUsed                                          float64 `json:"master/mem_used,omitempty"`
  MasterMessagesAuthenticate                             int64   `json:"master/messages_authenticate,omitempty"`
  MasterMessagesDeactivateFramework                      int64   `json:"master/messages_deactivate_framework,omitempty"`
  MasterMessagesDeclineOffers                            int64   `json:"master/messages_decline_offers,omitempty"`
  MasterMessagesExitedExecutor                           int64   `json:"master/messages_exited_executor,omitempty"`
  MasterMessagesFrameworkToExecutor                      int64   `json:"master/messages_framework_to_executor,omitempty"`
  MasterMessagesKillTask                                 int64   `json:"master/messages_kill_task,omitempty"`
  MasterMessagesLaunchTasks                              int64   `json:"master/messages_launch_tasks,omitempty"`
  MasterMessagesReconcileTasks                           int64   `json:"master/messages_reconcile_tasks,omitempty"`
  MasterMessagesRegisterFramework                        int64   `json:"master/messages_register_framework,omitempty"`
  MasterMessagesRegisterSlave                            int64   `json:"master/messages_register_slave,omitempty"`
  MasterMessagesReregisterFramework                      int64   `json:"master/messages_reregister_framework,omitempty"`
  MasterMessagesReregisterSlave                          int64   `json:"master/messages_reregister_slave,omitempty"`
  MasterMessagesResourceRequest                          int64   `json:"master/messages_resource_request,omitempty"`
  MasterMessagesReviveOffers                             int64   `json:"master/messages_revive_offers,omitempty"`
  MasterMessagesStatusUpdate                             int64   `json:"master/messages_status_update,omitempty"`
  MasterMessagesStatusUpdateAcknowledgement              int64   `json:"master/messages_status_update_acknowledgement,omitempty"`
  MasterMessagesUnregisterFramework                      int64   `json:"master/messages_unregister_framework,omitempty"`
  MasterMessagesUnregisterSlave                          int64   `json:"master/messages_unregister_slave,omitempty"`
  MasterOutstandingOffers                                int64   `json:"master/outstanding_offers,omitempty"`
  MasterRecoverySlaveRemovals                            int64   `json:"master/recovery_slave_removals,omitempty"`
  MasterSlaveRegistrations                               int64   `json:"master/slave_registrations,omitempty"`
  MasterSlaveRemovals                                    int64   `json:"master/slave_removals,omitempty"`
  MasterSlaveReregistrations                             int64   `json:"master/slave_reregistrations,omitempty"`
  MasterSlaveShutdownsCanceled                           int64   `json:"master/slave_shutdowns_canceled,omitempty"`
  MasterSlaveShutdownsScheduled                          int64   `json:"master/slave_shutdowns_scheduled,omitempty"`
  MasterSlavesActive                                     int64   `json:"master/slaves_active,omitempty"`
  MasterSlavesConnected                                  int64   `json:"master/slaves_connected,omitempty"`
  MasterSlavesDisconnected                               int64   `json:"master/slaves_disconnected,omitempty"`
  MasterSlavesInactive                                   int64   `json:"master/slaves_inactive,omitempty"`
  MasterTaskErrorsourceMasterReasonTaskInvalid           int64   `json:"master/task_error/source_master/reason_task_invalid,omitempty"`
  MasterTaskFailedsourceSlavereasonCommandExecutorFailed int64   `json:"master/task_failed/source_slave/reason_command_executor_failed,omitempty"`
  MasterTaskKilledsourceMasterReasonFrameworkRemoved     int64   `json:"master/task_killed/source_master/reason_framework_removed,omitempty"`
  MasterTaskKilledsourceSlavereasonExecutorUnregistered  int64   `json:"master/task_killed/source_slave/reason_executor_unregistered,omitempty"`
  MasterTaskLostsourceMasterReasonSlaveRemoved           int64   `json:"master/task_lost/source_master/reason_slave_removed,omitempty"`
  MasterTaskLostsourceSlavereasonExecutorTerminated      int64   `json:"master/task_lost/source_slave/reason_executor_terminated,omitempty"`
  MasterTasksError                                       int64   `json:"master/tasks_error,omitempty"`
  MasterTasksFailed                                      int64   `json:"master/tasks_failed,omitempty"`
  MasterTasksFinished                                    int64   `json:"master/tasks_finished,omitempty"`
  MasterTasksKilled                                      int64   `json:"master/tasks_killed,omitempty"`
  MasterTasksLost                                        int64   `json:"master/tasks_lost,omitempty"`
  MasterTasksRunning                                     int64   `json:"master/tasks_running,omitempty"`
  MasterTasksStaging                                     int64   `json:"master/tasks_staging,omitempty"`
  MasterTasksStarting                                    int64   `json:"master/tasks_starting,omitempty"`
  MasterUptimeSecs                                       float64 `json:"master/uptime_secs,omitempty"`
  MasterValidFrameworkToExecutorMessages                 int64   `json:"master/valid_framework_to_executor_messages,omitempty"`
  MasterValidStatusUpdateAcknowledgements                int64   `json:"master/valid_status_update_acknowledgements,omitempty"`
  MasterValidStatusUpdates                               int64   `json:"master/valid_status_updates,omitempty"`
  MemPercent                                             float64 `json:"mem_percent,omitempty"`
  MemTotal                                               float64 `json:"mem_total,omitempty"`
  MemUsed                                                float64 `json:"mem_used,omitempty"`
  OutstandingOffers                                      int64   `json:"outstanding_offers,omitempty"`
  RegistrarQueuedOperations                              int64   `json:"registrar/queued_operations,omitempty"`
  RegistrarRegistrySizeBytes                             int64   `json:"registrar/registry_size_bytes,omitempty"`
  RegistrarStateFetchMs                                  float64 `json:"registrar/state_fetch_ms,omitempty"`
  RegistrarStateStoreMs                                  float64 `json:"registrar/state_store_ms,omitempty"`
  RegistrarStateStoreMscount                             int64   `json:"registrar/state_store_ms/count,omitempty"`
  RegistrarStateStoreMsmax                               float64 `json:"registrar/state_store_ms/max,omitempty"`
  RegistrarStateStoreMsmin                               float64 `json:"registrar/state_store_ms/min,omitempty"`
  RegistrarStateStoreMsp50                               float64 `json:"registrar/state_store_ms/p50,omitempty"`
  RegistrarStateStoreMsp90                               float64 `json:"registrar/state_store_ms/p90,omitempty"`
  RegistrarStateStoreMsp95                               float64 `json:"registrar/state_store_ms/p95,omitempty"`
  RegistrarStateStoreMsp99                               float64 `json:"registrar/state_store_ms/p99,omitempty"`
  RegistrarStateStoreMsp999                              float64 `json:"registrar/state_store_ms/p999,omitempty"`
  RegistrarStateStoreMsp9999                             float64 `json:"registrar/state_store_ms/p9999,omitempty"`
  StagedTasks                                            int64   `json:"staged_tasks,omitempty"`
  StartedTasks                                           int64   `json:"started_tasks,omitempty"`
  SystemCpusTotal                                        int64   `json:"system/cpus_total,omitempty"`
  SystemLoad15min                                        float64 `json:"system/load_15min,omitempty"`
  SystemLoad1min                                         float64 `json:"system/load_1min,omitempty"`
  SystemLoad5min                                         float64 `json:"system/load_5min,omitempty"`
  SystemMemFreeBytes                                     int64   `json:"system/mem_free_bytes,omitempty"`
  SystemMemTotalBytes                                    int64   `json:"system/mem_total_bytes,omitempty"`
  TotalSchedulers                                        int64   `json:"total_schedulers,omitempty"`
  Uptime                                                 float64 `json:"uptime,omitempty"`
  ValidStatusUpdates                                     int64   `json:"valid_status_updates,omitempty"`
}

type MesosStatus struct {
  State     string  `json:"state,omitempty"`
  Timestamp float64 `json:"timestamp,omitempty"`
}

type MesosOffers struct {
  FrameworkId string      `json:"framework_id,omitempty"`
  Id          string      `json:"id,omitempty"`
  SlaveId     string      `json:"slave_id,omitempty"`
  Resources   interface{} `json:"resources,omitempty"`
}

type MesosTasks struct {
  ExecutorId  string         `json:"executor_id,omitempty"`
  FrameworkId string         `json:"framework_id,omitempty"`
  Id          string         `json:"id,omitempty"`
  Name        string         `json:"name,omitempty"`
  SlaveId     string         `json:"slave_id,omitempty"`
  State       string         `json:"state,omitempty"`
  Labels      []interface{}  `json:"labels,omitempty"`
  Resources   interface{}    `json:"resources,omitempty"`
  Statuses    []*MesosStatus `json:"statuses,omitempty"`
}

type MesosMasterFrameworks struct {
  Active           bool           `json:"active,omitempty"`
  Checkpoint       bool           `json:"checkpoint,omitempty"`
  FailoverTimeout  int64          `json:"failover_timeout,omitempty"`
  Hostname         string         `json:"hostname,omitempty"`
  Id               string         `json:"id,omitempty"`
  Name             string         `json:"name,omitempty"`
  RegisteredTime   float64        `json:"registered_time,omitempty"`
  ReregisteredTime float64        `json:"reregistered_time,omitempty"`
  Role             string         `json:"role,omitempty"`
  UnregisteredTime float64        `json:"unregistered_time,omitempty"`
  User             string         `json:"user,omitempty"`
  WebuiUrl         string         `json:"webui_url,omitempty"`
  CompletedTasks   []*MesosTasks  `json:"completed_tasks,omitempty"`
  OfferedResources interface{}    `json:"offered_resources,omitempty"`
  Offers           []*MesosOffers `json:"offers,omitempty"`
  Resources        interface{}    `json:"resources,omitempty"`
  Tasks            []*MesosTasks  `json:"tasks,omitempty"`
  UsedResources    interface{}    `json:"used_resources,omitempty"`
}

type MesosMasterFlags struct {
  AllocationInterval        string `json:"allocation_interval,omitempty"`
  Authenticate              string `json:"authenticate,omitempty"`
  AuthenticateSlaves        string `json:"authenticate_slaves,omitempty"`
  Authenticators            string `json:"authenticators,omitempty"`
  ExternalLogFile           string `json:"external_log_file,omitempty"`
  FrameworkSorter           string `json:"framework_sorter,omitempty"`
  Help                      string `json:"help,omitempty"`
  InitializeDriverLogging   string `json:"initialize_driver_logging,omitempty"`
  LogAutoInitialize         string `json:"log_auto_initialize,omitempty"`
  Logbufsecs                string `json:"logbufsecs,omitempty"`
  LoggingLevel              string `json:"logging_level,omitempty"`
  Port                      string `json:"port,omitempty"`
  Quiet                     string `json:"quiet,omitempty"`
  Quorum                    string `json:"quorum,omitempty"`
  RecoverySlaveRemovalLimit string `json:"recovery_slave_removal_limit,omitempty"`
  Registry                  string `json:"registry,omitempty"`
  RegistryFetchTimeout      string `json:"registry_fetch_timeout,omitempty"`
  RegistryStoreTimeout      string `json:"registry_store_timeout,omitempty"`
  RegistryStrict            string `json:"registry_strict,omitempty"`
  Roles                     string `json:"roles,omitempty"`
  RootSubmissions           string `json:"root_submissions,omitempty"`
  SlaveReregisterTimeout    string `json:"slave_reregister_timeout,omitempty"`
  UserSorter                string `json:"user_sorter,omitempty"`
  Version                   string `json:"version,omitempty"`
  WebuiDir                  string `json:"webui_dir,omitempty"`
  WorkDir                   string `json:"work_dir,omitempty"`
  Zk                        string `json:"zk,omitempty"`
  ZkSessionTimeout          string `json:"zk_session_timeout,omitempty"`
}

type MesosMasterState struct {
  ActivatedSlaves        int64                    `json:"activated_slaves,omitempty"`
  BuildDate              string                   `json:"build_date,omitempty"`
  BuildTime              int64                    `json:"build_time,omitempty"`
  BuildUser              string                   `json:"build_user,omitempty"`
  DeactivatedSlaves      int64                    `json:"deactivated_slaves,omitempty"`
  ElectedTime            float64                  `json:"elected_time,omitempty"`
  ExternalLogFile        string                   `json:"external_log_file,omitempty"`
  FailedTasks            int64                    `json:"failed_tasks,omitempty"`
  FinishedTasks          int64                    `json:"finished_tasks,omitempty"`
  Hostname               string                   `json:"hostname,omitempty"`
  Id                     string                   `json:"id,omitempty"`
  KilledTasks            int64                    `json:"killed_tasks,omitempty"`
  Leader                 string                   `json:"leader,omitempty"`
  LostTasks              int64                    `json:"lost_tasks,omitempty"`
  Pid                    string                   `json:"pid,omitempty"`
  StagedTasks            int64                    `json:"staged_tasks,omitempty"`
  StartTime              float64                  `json:"start_time,omitempty"`
  StartedTasks           int64                    `json:"started_tasks,omitempty"`
  Version                string                   `json:"version,omitempty"`
  Flags                  *MesosMasterFlags        `json:"flags,omitempty"`
  Frameworks             []*MesosMasterFrameworks `json:"frameworks,omitempty"`
  Slaves                 []*MesosSlaves           `json:"slaves,omitempty"`
  CompletedFrameworks    []*MesosMasterFrameworks `json:"completed_frameworks,omitempty"`
  UnregisteredFrameworks []interface{}            `json:"unregistered_frameworks,omitempty"`
  OrphanTasks            []interface{}            `json:"orphan_tasks,omitempty"`
}

type MesosMasterSlaves struct {
  Slaves []*MesosSlaves `json:"slaves,omitempty"`
}

type MesosSlaves struct {
  Active         bool        `json:"active,omitempty"`
  Hostname       string      `json:"hostname,omitempty"`
  Id             string      `json:"id,omitempty"`
  Pid            string      `json:"pid,omitempty"`
  RegisteredTime float64     `json:"registered_time,omitempty"`
  Resources      interface{} `json:"resources,omitempty"`
  Attributes     interface{} `json:"attributes,omitempty"`
}

type MesosRoles struct {
  Name       string      `json:"name,omitempty"`
  Weight     int64       `json:"weight,omitempty"`
  Frameworks []string    `json:"frameworks,omitempty"`
  Resources  interface{} `json:"resources,omitempty"`
}

type MesosMasterRoles struct {
  Roles []*MesosRoles `json:"roles,omitempty"`
}

type MesosMasterMetricsSnapshot struct {
  MasterCpusPercent                                      float64 `json:"master/cpus_percent,omitempty"`
  MasterCpusTotal                                        float64 `json:"master/cpus_total,omitempty"`
  MasterCpusUsed                                         float64 `json:"master/cpus_used,omitempty"`
  MasterDiskPercent                                      float64 `json:"master/disk_percent,omitempty"`
  MasterDiskTotal                                        float64 `json:"master/disk_total,omitempty"`
  MasterDiskUsed                                         float64 `json:"master/disk_used,omitempty"`
  MasterDroppedMessages                                  int64   `json:"master/dropped_messages,omitempty"`
  MasterElected                                          int64   `json:"master/elected,omitempty"`
  MasterEventQueueDispatches                             int64   `json:"master/event_queue_dispatches,omitempty"`
  MasterEventQueueHttpRequests                           int64   `json:"master/event_queue_http_requests,omitempty"`
  MasterEventQueueMessages                               int64   `json:"master/event_queue_messages,omitempty"`
  MasterFrameworksActive                                 int64   `json:"master/frameworks_active,omitempty"`
  MasterFrameworksConnected                              int64   `json:"master/frameworks_connected,omitempty"`
  MasterFrameworksDisconnected                           int64   `json:"master/frameworks_disconnected,omitempty"`
  MasterFrameworksInactive                               int64   `json:"master/frameworks_inactive,omitempty"`
  MasterInvalidFrameworkToExecutorMessages               int64   `json:"master/invalid_framework_to_executor_messages,omitempty"`
  MasterInvalidStatusUpdateAcknowledgements              int64   `json:"master/invalid_status_update_acknowledgements,omitempty"`
  MasterInvalidStatusUpdates                             int64   `json:"master/invalid_status_updates,omitempty"`
  MasterMemPercent                                       float64 `json:"master/mem_percent,omitempty"`
  MasterMemTotal                                         float64 `json:"master/mem_total,omitempty"`
  MasterMemUsed                                          float64 `json:"master/mem_used,omitempty"`
  MasterMessagesAuthenticate                             int64   `json:"master/messages_authenticate,omitempty"`
  MasterMessagesDeactivateFramework                      int64   `json:"master/messages_deactivate_framework,omitempty"`
  MasterMessagesDeclineOffers                            int64   `json:"master/messages_decline_offers,omitempty"`
  MasterMessagesExitedExecutor                           int64   `json:"master/messages_exited_executor,omitempty"`
  MasterMessagesFrameworkToExecutor                      int64   `json:"master/messages_framework_to_executor,omitempty"`
  MasterMessagesKillTask                                 int64   `json:"master/messages_kill_task,omitempty"`
  MasterMessagesLaunchTasks                              int64   `json:"master/messages_launch_tasks,omitempty"`
  MasterMessagesReconcileTasks                           int64   `json:"master/messages_reconcile_tasks,omitempty"`
  MasterMessagesRegisterFramework                        int64   `json:"master/messages_register_framework,omitempty"`
  MasterMessagesRegisterSlave                            int64   `json:"master/messages_register_slave,omitempty"`
  MasterMessagesReregisterFramework                      int64   `json:"master/messages_reregister_framework,omitempty"`
  MasterMessagesReregisterSlave                          int64   `json:"master/messages_reregister_slave,omitempty"`
  MasterMessagesResourceRequest                          int64   `json:"master/messages_resource_request,omitempty"`
  MasterMessagesReviveOffers                             int64   `json:"master/messages_revive_offers,omitempty"`
  MasterMessagesStatusUpdate                             int64   `json:"master/messages_status_update,omitempty"`
  MasterMessagesStatusUpdateAcknowledgement              int64   `json:"master/messages_status_update_acknowledgement,omitempty"`
  MasterMessagesUnregisterFramework                      int64   `json:"master/messages_unregister_framework,omitempty"`
  MasterMessagesUnregisterSlave                          int64   `json:"master/messages_unregister_slave,omitempty"`
  MasterOutstandingOffers                                int64   `json:"master/outstanding_offers,omitempty"`
  MasterRecoverySlaveRemovals                            int64   `json:"master/recovery_slave_removals,omitempty"`
  MasterSlaveRegistrations                               int64   `json:"master/slave_registrations,omitempty"`
  MasterSlaveRemovals                                    int64   `json:"master/slave_removals,omitempty"`
  MasterSlaveReregistrations                             int64   `json:"master/slave_reregistrations,omitempty"`
  MasterSlaveShutdownsCanceled                           int64   `json:"master/slave_shutdowns_canceled,omitempty"`
  MasterSlaveShutdownsScheduled                          int64   `json:"master/slave_shutdowns_scheduled,omitempty"`
  MasterSlavesActive                                     int64   `json:"master/slaves_active,omitempty"`
  MasterSlavesConnected                                  int64   `json:"master/slaves_connected,omitempty"`
  MasterSlavesDisconnected                               int64   `json:"master/slaves_disconnected,omitempty"`
  MasterSlavesInactive                                   int64   `json:"master/slaves_inactive,omitempty"`
  MasterTaskErrorsourceMasterReasonTaskInvalid           int64   `json:"master/task_error/source_master/reason_task_invalid,omitempty"`
  MasterTaskFailedsourceSlavereasonCommandExecutorFailed int64   `json:"master/task_failed/source_slave/reason_command_executor_failed,omitempty"`
  MasterTaskKilledsourceMasterReasonFrameworkRemoved     int64   `json:"master/task_killed/source_master/reason_framework_removed,omitempty"`
  MasterTaskKilledsourceSlavereasonExecutorUnregistered  int64   `json:"master/task_killed/source_slave/reason_executor_unregistered,omitempty"`
  MasterTaskLostsourceMasterReasonSlaveRemoved           int64   `json:"master/task_lost/source_master/reason_slave_removed,omitempty"`
  MasterTaskLostsourceSlavereasonExecutorTerminated      int64   `json:"master/task_lost/source_slave/reason_executor_terminated,omitempty"`
  MasterTasksError                                       int64   `json:"master/tasks_error,omitempty"`
  MasterTasksFailed                                      int64   `json:"master/tasks_failed,omitempty"`
  MasterTasksFinished                                    int64   `json:"master/tasks_finished,omitempty"`
  MasterTasksKilled                                      int64   `json:"master/tasks_killed,omitempty"`
  MasterTasksLost                                        int64   `json:"master/tasks_lost,omitempty"`
  MasterTasksRunning                                     int64   `json:"master/tasks_running,omitempty"`
  MasterTasksStaging                                     int64   `json:"master/tasks_staging,omitempty"`
  MasterTasksStarting                                    int64   `json:"master/tasks_starting,omitempty"`
  MasterUptimeSecs                                       float64 `json:"master/uptime_secs,omitempty"`
  MasterValidFrameworkToExecutorMessages                 int64   `json:"master/valid_framework_to_executor_messages,omitempty"`
  MasterValidStatusUpdateAcknowledgements                int64   `json:"master/valid_status_update_acknowledgements,omitempty"`
  MasterValidStatusUpdates                               int64   `json:"master/valid_status_updates,omitempty"`
  RegistrarQueuedOperations                              int64   `json:"registrar/queued_operations,omitempty"`
  RegistrarRegistrySizeBytes                             int64   `json:"registrar/registry_size_bytes,omitempty"`
  RegistrarStateFetchMs                                  float64 `json:"registrar/state_fetch_ms,omitempty"`
  RegistrarStateStoreMs                                  float64 `json:"registrar/state_store_ms,omitempty"`
  RegistrarStateStoreMscount                             int64   `json:"registrar/state_store_ms/count,omitempty"`
  RegistrarStateStoreMsmax                               float64 `json:"registrar/state_store_ms/max,omitempty"`
  RegistrarStateStoreMsmin                               float64 `json:"registrar/state_store_ms/min,omitempty"`
  RegistrarStateStoreMsp50                               float64 `json:"registrar/state_store_ms/p50,omitempty"`
  RegistrarStateStoreMsp90                               float64 `json:"registrar/state_store_ms/p90,omitempty"`
  RegistrarStateStoreMsp95                               float64 `json:"registrar/state_store_ms/p95,omitempty"`
  RegistrarStateStoreMsp99                               float64 `json:"registrar/state_store_ms/p99,omitempty"`
  RegistrarStateStoreMsp999                              float64 `json:"registrar/state_store_ms/p999,omitempty"`
  RegistrarStateStoreMsp9999                             float64 `json:"registrar/state_store_ms/p9999,omitempty"`
  SystemCpusTotal                                        int64   `json:"system/cpus_total,omitempty"`
  SystemLoad15min                                        float64 `json:"system/load_15min,omitempty"`
  SystemLoad1min                                         float64 `json:"system/load_1min,omitempty"`
  SystemLoad5min                                         float64 `json:"system/load_5min,omitempty"`
  SystemMemFreeBytes                                     int64   `json:"system/mem_free_bytes,omitempty"`
  SystemMemTotalBytes                                    int64   `json:"system/mem_total_bytes,omitempty"`
}

type MesosMasterRegistrar struct {
  Master *MesosRegistrarMaster `json:"master,omitempty"`
  Slaves *MesosRegistrarSlave  `json:"slaves,omitempty"`
}

type MesosRegistrarMaster struct {
  Info *MesosMasterInfo `json:"info,omitempty"`
}

type MesosRegistrarSlave struct {
  Slaves []*MesosSlave `json:"slaves,omitempty"`
}

type MesosMasterInfo struct {
  Hostname string `json:"hostname,omitempty"`
  Id       string `json:"id,omitempty"`
  Ip       int64  `json:"ip,omitempty"`
  Pid      string `json:"pid,omitempty"`
  Port     int64  `json:"port,omitempty"`
}

type MesosSlave struct {
  Info MesosSlaveInfo `json:"info,omitempty"`
}

type MesosSlaveInfo struct {
  Checkpoint bool          `json:"checkpoint,omitempty"`
  Hostname   string        `json:"hostname,omitempty"`
  Port       int64         `json:"port,omitempty"`
  ID         *MesosSlaveID `json:"id,omitempty"`
  // Attributes []*MesosSlaveAttribute `json:"attributes,omitempty"`
  // Resources  []*MesosSlaveResources `json:"resources,omitempty"`
  Attributes []interface{} `json:"attributes,omitempty"`
  Resources  []interface{} `json:"resources,omitempty"`
}

type MesosSlaveResources struct {
  Name   string      `json:"name,omitempty"`
  Role   string      `json:"role,omitempty"`
  Type   string      `json:"type,omitempty"`
  Ranges interface{} `json:"ranges,omitempty"`
  Scalar interface{} `json:"scalar,omitempty"`
}

type MesosSlaveResourceScalar struct {
  Value int64 `json:"value,omitempty"`
}

type MesosSlaveResourceRanges struct {
  Range []*MesosSlaveResourceRange `json:"range,omitempty"`
}

type MesosSlaveResourceRange struct {
  Begin int64 `json:"begin,omitempty"`
  End   int64 `json:"end,omitempty"`
}

type MesosSlaveID struct {
  Value string `json:"value,omitempty"`
}

type MesosSlaveAttributeText struct {
  Value string `json:"value,omitempty"`
}

type MesosSlaveAttribute struct {
  Name string                   `json:"name,omitempty"`
  Type string                   `json:"type,omitempty"`
  Text *MesosSlaveAttributeText `json:"text,omitempty"`
}

type MesosMasterOptions struct {
  Host string `json:"mesos_master_url,omitempty"`
}

type MesosMasterClient struct {
  opts *MesosMasterOptions
}

func NewMasterClient(opts *MesosMasterOptions) *MesosMasterClient {
  if opts.Host == "" {
    opts = &MesosMasterOptions{
      Host: MesosMasterHostDefault,
    }
  }

  s := &MesosMasterClient{
    opts: opts,
  }

  return s
}

func MesosAttributes(r interface{}) map[string]interface{} {
  attrs := make(map[string]interface{})

  switch t := r.(type) {
  case []interface{}:
    for _, at := range t {
      switch tt := at.(type) {
      case map[string]interface{}:
        switch tt["type"] {
        case "TEXT":
          switch s := tt["text"].(type) {
          case map[string]interface{}:
            attrs[tt["name"].(string)] = s["value"]
          }
        case "SCALAR":
          switch s := tt["scalar"].(type) {
          case map[string]interface{}:
            attrs[tt["name"].(string)] = s["value"]
          }
        }
      }
    }
  }

  return attrs
}

func MesosResources(r interface{}) map[string]interface{} {
  res := make(map[string]interface{})

  switch t := r.(type) {
  case []interface{}:
    for _, rs := range t {
      switch tt := rs.(type) {
      case map[string]interface{}:
        switch tt["type"] {
        case "SCALAR":
          switch s := tt["scalar"].(type) {
          case map[string]interface{}:
            res[tt["name"].(string)] = s["value"]
          }
        }
      }
    }
  }

  return res
}

func (mm *MesosMasterClient) MesosMasterRegistrar() (*MesosMasterRegistrar, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosMasterRegistrarURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosMasterRegistrar{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosMasterMetricsSnapshot() (*MesosMasterMetricsSnapshot, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosMasterMetricsSnapshotURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosMasterMetricsSnapshot{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosMasterRoles() (*MesosMasterRoles, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosMasterRolesURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosMasterRoles{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosMasterSlaves() (*MesosMasterSlaves, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosMasterSlavesURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosMasterSlaves{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosMasterState() (*MesosMasterState, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosMasterStateURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosMasterState{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosMasterTasks(query Query) (*MesosMasterTasks, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosMasterTasksURI, query)
  if err != nil {
    return nil, err
  }

  s := &MesosMasterTasks{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosMasterStats() (*MesosMasterStats, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosMasterStatsURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosMasterStats{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosSystemStats() (*MesosSystemStats, error) {
  body, err := httpGetRequest(mm.opts.Host+MesosSystemStatsURI, Query{})
  if err != nil {
    return nil, err
  }

  s := &MesosSystemStats{}
  if err := json.Unmarshal(body, s); err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  return s, nil
}

func (mm *MesosMasterClient) MesosMasterHealth() error {
  statusCode, err := httpHeadRequest(mm.opts.Host+MesosMasterHealthURI, Query{})
  if err != nil && statusCode != 200 {
    fmt.Errorf("%v: \n", err, statusCode)
  }
  return err
}
