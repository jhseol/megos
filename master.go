package megos

import (
  "encoding/json"
  "errors"
  "fmt"
)

type MesosMasterTasks struct {
  Tasks []MesosTasks `json:"tasks"`
}

type MesosMasterStats struct {
  ActivatedSlaves                                        int64   `json:"activated_slaves"`
  ActiveSchedulers                                       int64   `json:"active_schedulers"`
  ActiveTasksGauge                                       int64   `json:"active_tasks_gauge"`
  CpusPercent                                            float64 `json:"cpus_percent"`
  CpusTotal                                              int64   `json:"cpus_total"`
  CpusUsed                                               float64 `json:"cpus_used"`
  DeactivatedSlaves                                      int64   `json:"deactivated_slaves"`
  DiskPercent                                            float64 `json:"disk_percent"`
  DiskTotal                                              float64 `json:"disk_total"`
  DiskUsed                                               float64 `json:"disk_used"`
  Elected                                                int64   `json:"elected"`
  FailedTasks                                            int64   `json:"failed_tasks"`
  FinishedTasks                                          int64   `json:"finished_tasks"`
  InvalidStatusUpdates                                   int64   `json:"invalid_status_updates"`
  KilledTasks                                            int64   `json:"killed_tasks"`
  LostTasks                                              int64   `json:"lost_tasks"`
  MasterCpusPercent                                      float64 `json:"master/cpus_percent"`
  MasterCpusTotal                                        float64 `json:"master/cpus_total"`
  MasterCpusUsed                                         float64 `json:"master/cpus_used"`
  MasterDiskPercent                                      float64 `json:"master/disk_percent"`
  MasterDiskTotal                                        float64 `json:"master/disk_total"`
  MasterDiskUsed                                         float64 `json:"master/disk_used"`
  MasterDroppedMessages                                  int64   `json:"master/dropped_messages"`
  MasterElected                                          int64   `json:"master/elected"`
  MasterEventQueueDispatches                             int64   `json:"master/event_queue_dispatches"`
  MasterEventQueueHttpRequests                           int64   `json:"master/event_queue_http_requests"`
  MasterEventQueueMessages                               int64   `json:"master/event_queue_messages"`
  MasterFrameworksActive                                 int64   `json:"master/frameworks_active"`
  MasterFrameworksConnected                              int64   `json:"master/frameworks_connected"`
  MasterFrameworksDisconnected                           int64   `json:"master/frameworks_disconnected"`
  MasterFrameworksInactive                               int64   `json:"master/frameworks_inactive"`
  MasterInvalidFrameworkToExecutorMessages               int64   `json:"master/invalid_framework_to_executor_messages"`
  MasterInvalidStatusUpdateAcknowledgements              int64   `json:"master/invalid_status_update_acknowledgements"`
  MasterInvalidStatusUpdates                             int64   `json:"master/invalid_status_updates"`
  MasterMemPercent                                       float64 `json:"master/mem_percent"`
  MasterMemTotal                                         float64 `json:"master/mem_total"`
  MasterMemUsed                                          float64 `json:"master/mem_used"`
  MasterMessagesAuthenticate                             int64   `json:"master/messages_authenticate"`
  MasterMessagesDeactivateFramework                      int64   `json:"master/messages_deactivate_framework"`
  MasterMessagesDeclineOffers                            int64   `json:"master/messages_decline_offers"`
  MasterMessagesExitedExecutor                           int64   `json:"master/messages_exited_executor"`
  MasterMessagesFrameworkToExecutor                      int64   `json:"master/messages_framework_to_executor"`
  MasterMessagesKillTask                                 int64   `json:"master/messages_kill_task"`
  MasterMessagesLaunchTasks                              int64   `json:"master/messages_launch_tasks"`
  MasterMessagesReconcileTasks                           int64   `json:"master/messages_reconcile_tasks"`
  MasterMessagesRegisterFramework                        int64   `json:"master/messages_register_framework"`
  MasterMessagesRegisterSlave                            int64   `json:"master/messages_register_slave"`
  MasterMessagesReregisterFramework                      int64   `json:"master/messages_reregister_framework"`
  MasterMessagesReregisterSlave                          int64   `json:"master/messages_reregister_slave"`
  MasterMessagesResourceRequest                          int64   `json:"master/messages_resource_request"`
  MasterMessagesReviveOffers                             int64   `json:"master/messages_revive_offers"`
  MasterMessagesStatusUpdate                             int64   `json:"master/messages_status_update"`
  MasterMessagesStatusUpdateAcknowledgement              int64   `json:"master/messages_status_update_acknowledgement"`
  MasterMessagesUnregisterFramework                      int64   `json:"master/messages_unregister_framework"`
  MasterMessagesUnregisterSlave                          int64   `json:"master/messages_unregister_slave"`
  MasterOutstandingOffers                                int64   `json:"master/outstanding_offers"`
  MasterRecoverySlaveRemovals                            int64   `json:"master/recovery_slave_removals"`
  MasterSlaveRegistrations                               int64   `json:"master/slave_registrations"`
  MasterSlaveRemovals                                    int64   `json:"master/slave_removals"`
  MasterSlaveReregistrations                             int64   `json:"master/slave_reregistrations"`
  MasterSlaveShutdownsCanceled                           int64   `json:"master/slave_shutdowns_canceled"`
  MasterSlaveShutdownsScheduled                          int64   `json:"master/slave_shutdowns_scheduled"`
  MasterSlavesActive                                     int64   `json:"master/slaves_active"`
  MasterSlavesConnected                                  int64   `json:"master/slaves_connected"`
  MasterSlavesDisconnected                               int64   `json:"master/slaves_disconnected"`
  MasterSlavesInactive                                   int64   `json:"master/slaves_inactive"`
  MasterTaskErrorsourceMasterReasonTaskInvalid           int64   `json:"master/task_error/source_master/reason_task_invalid"`
  MasterTaskFailedsourceSlavereasonCommandExecutorFailed int64   `json:"master/task_failed/source_slave/reason_command_executor_failed"`
  MasterTaskKilledsourceMasterReasonFrameworkRemoved     int64   `json:"master/task_killed/source_master/reason_framework_removed"`
  MasterTaskKilledsourceSlavereasonExecutorUnregistered  int64   `json:"master/task_killed/source_slave/reason_executor_unregistered"`
  MasterTaskLostsourceMasterReasonSlaveRemoved           int64   `json:"master/task_lost/source_master/reason_slave_removed"`
  MasterTaskLostsourceSlavereasonExecutorTerminated      int64   `json:"master/task_lost/source_slave/reason_executor_terminated"`
  MasterTasksError                                       int64   `json:"master/tasks_error"`
  MasterTasksFailed                                      int64   `json:"master/tasks_failed"`
  MasterTasksFinished                                    int64   `json:"master/tasks_finished"`
  MasterTasksKilled                                      int64   `json:"master/tasks_killed"`
  MasterTasksLost                                        int64   `json:"master/tasks_lost"`
  MasterTasksRunning                                     int64   `json:"master/tasks_running"`
  MasterTasksStaging                                     int64   `json:"master/tasks_staging"`
  MasterTasksStarting                                    int64   `json:"master/tasks_starting"`
  MasterUptimeSecs                                       float64 `json:"master/uptime_secs"`
  MasterValidFrameworkToExecutorMessages                 int64   `json:"master/valid_framework_to_executor_messages"`
  MasterValidStatusUpdateAcknowledgements                int64   `json:"master/valid_status_update_acknowledgements"`
  MasterValidStatusUpdates                               int64   `json:"master/valid_status_updates"`
  MemPercent                                             float64 `json:"mem_percent"`
  MemTotal                                               int64   `json:"mem_total"`
  MemUsed                                                float64 `json:"mem_used"`
  OutstandingOffers                                      int64   `json:"outstanding_offers"`
  RegistrarQueuedOperations                              int64   `json:"registrar/queued_operations"`
  RegistrarRegistrySizeBytes                             int64   `json:"registrar/registry_size_bytes"`
  RegistrarStateFetchMs                                  float64 `json:"registrar/state_fetch_ms"`
  RegistrarStateStoreMs                                  float64 `json:"registrar/state_store_ms"`
  RegistrarStateStoreMscount                             int64   `json:"registrar/state_store_ms/count"`
  RegistrarStateStoreMsmax                               float64 `json:"registrar/state_store_ms/max"`
  RegistrarStateStoreMsmin                               float64 `json:"registrar/state_store_ms/min"`
  RegistrarStateStoreMsp50                               float64 `json:"registrar/state_store_ms/p50"`
  RegistrarStateStoreMsp90                               float64 `json:"registrar/state_store_ms/p90"`
  RegistrarStateStoreMsp95                               float64 `json:"registrar/state_store_ms/p95"`
  RegistrarStateStoreMsp99                               float64 `json:"registrar/state_store_ms/p99"`
  RegistrarStateStoreMsp999                              float64 `json:"registrar/state_store_ms/p999"`
  RegistrarStateStoreMsp9999                             float64 `json:"registrar/state_store_ms/p9999"`
  StagedTasks                                            int64   `json:"staged_tasks"`
  StartedTasks                                           int64   `json:"started_tasks"`
  SystemCpusTotal                                        int64   `json:"system/cpus_total"`
  SystemLoad15min                                        float64 `json:"system/load_15min"`
  SystemLoad1min                                         float64 `json:"system/load_1min"`
  SystemLoad5min                                         float64 `json:"system/load_5min"`
  SystemMemFreeBytes                                     int64   `json:"system/mem_free_bytes"`
  SystemMemTotalBytes                                    int64   `json:"system/mem_total_bytes"`
  TotalSchedulers                                        int64   `json:"total_schedulers"`
  Uptime                                                 float64 `json:"uptime"`
  ValidStatusUpdates                                     int64   `json:"valid_status_updates"`
}

type MesosStatus struct {
  State     string  `json:"state"`
  Timestamp float64 `json:"timestamp"`
}

type MesosOffers struct {
  FrameworkId string      `json:"framework_id"`
  Id          string      `json:"id"`
  SlaveId     string      `json:"slave_id"`
  Resources   interface{} `json:"resources"`
}

type MesosTasks struct {
  ExecutorId  string        `json:"executor_id"`
  FrameworkId string        `json:"framework_id"`
  Id          string        `json:"id"`
  Name        string        `json:"name"`
  SlaveId     string        `json:"slave_id"`
  State       string        `json:"state"`
  Labels      []interface{} `json:"labels"`
  Resources   interface{}   `json:"resources"`
  Statuses    []MesosStatus `json:"statuses"`
}

type MesosMasterFrameworks struct {
  Active           bool          `json:"active"`
  Checkpoint       bool          `json:"checkpoint"`
  FailoverTimeout  int64         `json:"failover_timeout"`
  Hostname         string        `json:"hostname"`
  Id               string        `json:"id"`
  Name             string        `json:"name"`
  RegisteredTime   float64       `json:"registered_time"`
  ReregisteredTime float64       `json:"reregistered_time"`
  Role             string        `json:"role"`
  UnregisteredTime float64       `json:"unregistered_time"`
  User             string        `json:"user"`
  WebuiUrl         string        `json:"webui_url"`
  CompletedTasks   []MesosTasks  `json:"completed_tasks"`
  OfferedResources interface{}   `json:"offered_resources"`
  Offers           []MesosOffers `json:"offers"`
  Resources        interface{}   `json:"resources"`
  Tasks            []MesosTasks  `json:"tasks"`
  UsedResources    interface{}   `json:"used_resources"`
}

type MesosMasterFlags struct {
  AllocationInterval        string `json:"allocation_interval"`
  Authenticate              string `json:"authenticate"`
  AuthenticateSlaves        string `json:"authenticate_slaves"`
  Authenticators            string `json:"authenticators"`
  ExternalLogFile           string `json:"external_log_file"`
  FrameworkSorter           string `json:"framework_sorter"`
  Help                      string `json:"help"`
  InitializeDriverLogging   string `json:"initialize_driver_logging"`
  LogAutoInitialize         string `json:"log_auto_initialize"`
  Logbufsecs                string `json:"logbufsecs"`
  LoggingLevel              string `json:"logging_level"`
  Port                      string `json:"port"`
  Quiet                     string `json:"quiet"`
  Quorum                    string `json:"quorum"`
  RecoverySlaveRemovalLimit string `json:"recovery_slave_removal_limit"`
  Registry                  string `json:"registry"`
  RegistryFetchTimeout      string `json:"registry_fetch_timeout"`
  RegistryStoreTimeout      string `json:"registry_store_timeout"`
  RegistryStrict            string `json:"registry_strict"`
  Roles                     string `json:"roles"`
  RootSubmissions           string `json:"root_submissions"`
  SlaveReregisterTimeout    string `json:"slave_reregister_timeout"`
  UserSorter                string `json:"user_sorter"`
  Version                   string `json:"version"`
  WebuiDir                  string `json:"webui_dir"`
  WorkDir                   string `json:"work_dir"`
  Zk                        string `json:"zk"`
  ZkSessionTimeout          string `json:"zk_session_timeout"`
}

type MesosMasterState struct {
  ActivatedSlaves        int64                   `json:"activated_slaves"`
  BuildDate              string                  `json:"build_date"`
  BuildTime              int64                   `json:"build_time"`
  BuildUser              string                  `json:"build_user"`
  DeactivatedSlaves      int64                   `json:"deactivated_slaves"`
  ElectedTime            float64                 `json:"elected_time"`
  ExternalLogFile        string                  `json:"external_log_file"`
  FailedTasks            int64                   `json:"failed_tasks"`
  FinishedTasks          int64                   `json:"finished_tasks"`
  Hostname               string                  `json:"hostname"`
  Id                     string                  `json:"id"`
  KilledTasks            int64                   `json:"killed_tasks"`
  Leader                 string                  `json:"leader"`
  LostTasks              int64                   `json:"lost_tasks"`
  Pid                    string                  `json:"pid"`
  StagedTasks            int64                   `json:"staged_tasks"`
  StartTime              float64                 `json:"start_time"`
  StartedTasks           int64                   `json:"started_tasks"`
  Version                string                  `json:"version"`
  Flags                  MesosMasterFlags        `json:"flags"`
  Frameworks             []MesosMasterFrameworks `json:"frameworks"`
  Slaves                 []MesosSlaves           `json:"slaves"`
  CompletedFrameworks    []MesosMasterFrameworks `json:"completed_frameworks"`
  UnregisteredFrameworks []interface{}           `json:"unregistered_frameworks"`
  OrphanTasks            []interface{}           `json:"orphan_tasks"`
}

type MesosMasterSlaves struct {
  Slaves []MesosSlaves `json:"slaves"`
}

type MesosSlaves struct {
  Active         bool        `json:"active"`
  Hostname       string      `json:"hostname"`
  Id             string      `json:"id"`
  Pid            string      `json:"pid"`
  RegisteredTime float64     `json:"registered_time"`
  Resources      interface{} `json:"resources"`
  Attributes     interface{} `json:"attributes"`
}

type MesosRoles struct {
  Name       string      `json:"name"`
  Weight     int64       `json:"weight"`
  Frameworks []string    `json:"frameworks"`
  Resources  interface{} `json:"resources"`
}

type MesosMasterRoles struct {
  Roles []MesosRoles `json:"roles"`
}

type MesosMasterMetricsSnapshot struct {
  MasterCpusPercent                                      float64 `json:"master/cpus_percent"`
  MasterCpusTotal                                        float64 `json:"master/cpus_total"`
  MasterCpusUsed                                         float64 `json:"master/cpus_used"`
  MasterDiskPercent                                      float64 `json:"master/disk_percent"`
  MasterDiskTotal                                        float64 `json:"master/disk_total"`
  MasterDiskUsed                                         float64 `json:"master/disk_used"`
  MasterDroppedMessages                                  int64   `json:"master/dropped_messages"`
  MasterElected                                          int64   `json:"master/elected"`
  MasterEventQueueDispatches                             int64   `json:"master/event_queue_dispatches"`
  MasterEventQueueHttpRequests                           int64   `json:"master/event_queue_http_requests"`
  MasterEventQueueMessages                               int64   `json:"master/event_queue_messages"`
  MasterFrameworksActive                                 int64   `json:"master/frameworks_active"`
  MasterFrameworksConnected                              int64   `json:"master/frameworks_connected"`
  MasterFrameworksDisconnected                           int64   `json:"master/frameworks_disconnected"`
  MasterFrameworksInactive                               int64   `json:"master/frameworks_inactive"`
  MasterInvalidFrameworkToExecutorMessages               int64   `json:"master/invalid_framework_to_executor_messages"`
  MasterInvalidStatusUpdateAcknowledgements              int64   `json:"master/invalid_status_update_acknowledgements"`
  MasterInvalidStatusUpdates                             int64   `json:"master/invalid_status_updates"`
  MasterMemPercent                                       float64 `json:"master/mem_percent"`
  MasterMemTotal                                         float64 `json:"master/mem_total"`
  MasterMemUsed                                          float64 `json:"master/mem_used"`
  MasterMessagesAuthenticate                             int64   `json:"master/messages_authenticate"`
  MasterMessagesDeactivateFramework                      int64   `json:"master/messages_deactivate_framework"`
  MasterMessagesDeclineOffers                            int64   `json:"master/messages_decline_offers"`
  MasterMessagesExitedExecutor                           int64   `json:"master/messages_exited_executor"`
  MasterMessagesFrameworkToExecutor                      int64   `json:"master/messages_framework_to_executor"`
  MasterMessagesKillTask                                 int64   `json:"master/messages_kill_task"`
  MasterMessagesLaunchTasks                              int64   `json:"master/messages_launch_tasks"`
  MasterMessagesReconcileTasks                           int64   `json:"master/messages_reconcile_tasks"`
  MasterMessagesRegisterFramework                        int64   `json:"master/messages_register_framework"`
  MasterMessagesRegisterSlave                            int64   `json:"master/messages_register_slave"`
  MasterMessagesReregisterFramework                      int64   `json:"master/messages_reregister_framework"`
  MasterMessagesReregisterSlave                          int64   `json:"master/messages_reregister_slave"`
  MasterMessagesResourceRequest                          int64   `json:"master/messages_resource_request"`
  MasterMessagesReviveOffers                             int64   `json:"master/messages_revive_offers"`
  MasterMessagesStatusUpdate                             int64   `json:"master/messages_status_update"`
  MasterMessagesStatusUpdateAcknowledgement              int64   `json:"master/messages_status_update_acknowledgement"`
  MasterMessagesUnregisterFramework                      int64   `json:"master/messages_unregister_framework"`
  MasterMessagesUnregisterSlave                          int64   `json:"master/messages_unregister_slave"`
  MasterOutstandingOffers                                int64   `json:"master/outstanding_offers"`
  MasterRecoverySlaveRemovals                            int64   `json:"master/recovery_slave_removals"`
  MasterSlaveRegistrations                               int64   `json:"master/slave_registrations"`
  MasterSlaveRemovals                                    int64   `json:"master/slave_removals"`
  MasterSlaveReregistrations                             int64   `json:"master/slave_reregistrations"`
  MasterSlaveShutdownsCanceled                           int64   `json:"master/slave_shutdowns_canceled"`
  MasterSlaveShutdownsScheduled                          int64   `json:"master/slave_shutdowns_scheduled"`
  MasterSlavesActive                                     int64   `json:"master/slaves_active"`
  MasterSlavesConnected                                  int64   `json:"master/slaves_connected"`
  MasterSlavesDisconnected                               int64   `json:"master/slaves_disconnected"`
  MasterSlavesInactive                                   int64   `json:"master/slaves_inactive"`
  MasterTaskErrorsourceMasterReasonTaskInvalid           int64   `json:"master/task_error/source_master/reason_task_invalid"`
  MasterTaskFailedsourceSlavereasonCommandExecutorFailed int64   `json:"master/task_failed/source_slave/reason_command_executor_failed"`
  MasterTaskKilledsourceMasterReasonFrameworkRemoved     int64   `json:"master/task_killed/source_master/reason_framework_removed"`
  MasterTaskKilledsourceSlavereasonExecutorUnregistered  int64   `json:"master/task_killed/source_slave/reason_executor_unregistered"`
  MasterTaskLostsourceMasterReasonSlaveRemoved           int64   `json:"master/task_lost/source_master/reason_slave_removed"`
  MasterTaskLostsourceSlavereasonExecutorTerminated      int64   `json:"master/task_lost/source_slave/reason_executor_terminated"`
  MasterTasksError                                       int64   `json:"master/tasks_error"`
  MasterTasksFailed                                      int64   `json:"master/tasks_failed"`
  MasterTasksFinished                                    int64   `json:"master/tasks_finished"`
  MasterTasksKilled                                      int64   `json:"master/tasks_killed"`
  MasterTasksLost                                        int64   `json:"master/tasks_lost"`
  MasterTasksRunning                                     int64   `json:"master/tasks_running"`
  MasterTasksStaging                                     int64   `json:"master/tasks_staging"`
  MasterTasksStarting                                    int64   `json:"master/tasks_starting"`
  MasterUptimeSecs                                       float64 `json:"master/uptime_secs"`
  MasterValidFrameworkToExecutorMessages                 int64   `json:"master/valid_framework_to_executor_messages"`
  MasterValidStatusUpdateAcknowledgements                int64   `json:"master/valid_status_update_acknowledgements"`
  MasterValidStatusUpdates                               int64   `json:"master/valid_status_updates"`
  RegistrarQueuedOperations                              int64   `json:"registrar/queued_operations"`
  RegistrarRegistrySizeBytes                             int64   `json:"registrar/registry_size_bytes"`
  RegistrarStateFetchMs                                  float64 `json:"registrar/state_fetch_ms"`
  RegistrarStateStoreMs                                  float64 `json:"registrar/state_store_ms"`
  RegistrarStateStoreMscount                             int64   `json:"registrar/state_store_ms/count"`
  RegistrarStateStoreMsmax                               float64 `json:"registrar/state_store_ms/max"`
  RegistrarStateStoreMsmin                               float64 `json:"registrar/state_store_ms/min"`
  RegistrarStateStoreMsp50                               float64 `json:"registrar/state_store_ms/p50"`
  RegistrarStateStoreMsp90                               float64 `json:"registrar/state_store_ms/p90"`
  RegistrarStateStoreMsp95                               float64 `json:"registrar/state_store_ms/p95"`
  RegistrarStateStoreMsp99                               float64 `json:"registrar/state_store_ms/p99"`
  RegistrarStateStoreMsp999                              float64 `json:"registrar/state_store_ms/p999"`
  RegistrarStateStoreMsp9999                             float64 `json:"registrar/state_store_ms/p9999"`
  SystemCpusTotal                                        int64   `json:"system/cpus_total"`
  SystemLoad15min                                        float64 `json:"system/load_15min"`
  SystemLoad1min                                         float64 `json:"system/load_1min"`
  SystemLoad5min                                         float64 `json:"system/load_5min"`
  SystemMemFreeBytes                                     int64   `json:"system/mem_free_bytes"`
  SystemMemTotalBytes                                    int64   `json:"system/mem_total_bytes"`
}

type MesosMasterRegistrar struct {
  Master MesosRegistrarMaster `json:"master"`
  Slaves MesosRegistrarSlaves `json:"slaves"`
}

type MesosRegistrarMaster struct {
  Info MesosMasterInfo `json:"info"`
}

type MesosRegistrarSlaves struct {
  Slaves []MesosSlave `json:"slaves"`
}

type MesosMasterInfo struct {
  Hostname string `json:"hostname"`
  Id       string `json:"id"`
  Ip       int64  `json:"ip"`
  Pid      string `json:"pid"`
  Port     int64  `json:"port"`
}

type MesosSlave struct {
  Info MesosSlaveInfo `json:"info"`
}

type MesosSlaveInfo struct {
  Checkpoint bool                  `json:"checkpoint"`
  Hostname   string                `json:"hostname"`
  Port       int64                 `json:"port"`
  ID         MesosSlaveID          `json:"id"`
  Attributes []MesosSlaveAttribute `json:"attributes"`
  Resources  []MesosSlaveResources `json:"resources"`
}

type MesosSlaveResources struct {
  Name   string      `json:"name"`
  Role   string      `json:"role"`
  Type   string      `json:"type"`
  Ranges interface{} `json:"ranges,omitempty"`
  Scalar interface{} `json:"scalar,omitempty"`
}

type MesosSlaveResourceScalar struct {
  Value int64 `json:"value,omitempty"`
}

type MesosSlaveResourceRanges struct {
  Range []MesosSlaveResourceRange `json:"range,omitempty"`
}

type MesosSlaveResourceRange struct {
  Begin int64 `json:"begin"`
  End   int64 `json:"end"`
}

type MesosSlaveID struct {
  Value string `json:"value"`
}

type MesosSlaveAttributeText struct {
  Value string `json:"value"`
}

type MesosSlaveAttribute struct {
  Name string                  `json:"name"`
  Type string                  `json:"type"`
  Text MesosSlaveAttributeText `json:"text"`
}

type MesosMasterOptions struct {
  Host string `json:"mesos_master_url"`
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

  encoded, err := json.Marshal(s)
  if err != nil {
    fmt.Errorf("%v\n", err)
    return nil, err
  }

  if string(encoded) != string(body) {
    fmt.Printf("Encoded %v and body %v is not equal\n", string(encoded), string(body))
    return nil, errors.New("Encode error")
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
