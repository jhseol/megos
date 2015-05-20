// +build ignore

package main

import (
  "flag"
  "fmt"
  "github.com/jhseol/megos"
)

func TestMasterAPIs(m *megos.MesosMasterClient) error {
  r, err := m.MesosMasterRegistrar()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", r)

  rs, err := m.MesosMasterRoles()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", rs)

  s, err := m.MesosMasterMetricsSnapshot()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", s)

  sv, err := m.MesosMasterSlaves()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", sv)

  ss, err := m.MesosSystemStats()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", ss)

  ms, err := m.MesosMasterState()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", ms)

  mss, err := m.MesosMasterStats()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", mss)

  mt, err := m.MesosMasterTasks(megos.Query{"limit": "100", "offset": "0", "order": "asc"})
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", mt)

  if err := m.MesosMasterHealth(); err != nil {
    return err
  }
  fmt.Println("Mesos master health is OK.")

  return nil
}

func TestSlaveAPIs(s *megos.MesosSlaveClient) error {
  state, err := s.MesosSlaveState()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", state)

  stats, err := s.MesosSlaveStats()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", stats)

  sstats, err := s.MesosSystemStats()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", sstats)

  mstats, err := s.MesosSlaveMonitorStatistics()
  if err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", mstats)

  if err := s.MesosSlaveHealth(); err != nil {
    fmt.Printf("%v\n", err)
    return err
  }
  fmt.Printf("%+v\n", "Mesos slave health is OK.")

  return nil
}

func main() {
  var host string

  flag.StringVar(&host, "host", "", "Mesos host address(e.g 127.0.0.1:5050)")
  flag.Parse()

  mesosType := flag.Arg(0)

  if mesosType == "master" || mesosType == "m" {
    opts := &megos.MesosMasterOptions{Host: host}
    m := megos.NewMasterClient(opts)
    _ = TestMasterAPIs(m)
  } else if mesosType == "slave" || mesosType == "s" {
    opts := &megos.MesosSlaveOptions{Host: host}
    s := megos.NewSlaveClient(opts)
    _ = TestSlaveAPIs(s)
  } else {
    fmt.Println("Argument is missed: 'master' or 'slave'")
    return
  }
}
