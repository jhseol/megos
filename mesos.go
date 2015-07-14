package megos

const (
  MesosSlaveHostDefault  = "http://127.0.0.1:5051"
  MesosMasterHostDefault = "http://127.0.0.1:5050"

  MesosSlaveStateURI             = "/slave(1)/state.json"
  MesosSlaveStatsURI             = "/slave(1)/stats.json"
  MesosSlaveHealthURI            = "/slave(1)/health"
  MesosSlaveMonitorStatisticsURI = "/monitor/statistics.json"
  MesosSlaveMetricsSnapshotURI   = "/metrics/snapshot"

  MesosMasterRegistrarURI       = "/registrar(1)/registry"
  MesosMasterMetricsSnapshotURI = "/metrics/snapshot"
  MesosSystemStatsURI           = "/system/stats.json"
  MesosMasterHealthURI          = "/master/health"
  MesosMasterObserveURI         = "/master/observe"
  MesosMasterRedirectURI        = "/master/redirect"
  MesosMasterRolesURI           = "/master/roles.json"
  MesosMasterShutdownURI        = "/master/shutdown"
  MesosMasterSlavesURI          = "/master/slaves"
  MesosMasterStateURI           = "/master/state.json"
  MesosMasterStatsURI           = "/master/stats.json"
  MesosMasterTasksURI           = "/master/tasks.json"
)

type MesosSystemStats struct {
  AvgLoad15min  float64 `json:"avg_load_15min"`
  AvgLoad1min   float64 `json:"avg_load_1min"`
  AvgLoad5min   float64 `json:"avg_load_5min"`
  CpusTotal     int64   `json:"cpus_total"`
  MemFreeBytes  int64   `json:"mem_free_bytes"`
  MemTotalBytes int64   `json:"mem_total_bytes"`
}
