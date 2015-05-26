// +build ignore

package main

import (
  "encoding/json"
  "fmt"
  "github.com/jhseol/megos"
)

func ExportMasterMetrics(m *megos.MesosMasterClient) error {

}

func ExportSlaveMetrics(m *megos.MesosSlaveClient) error {

}

func main() {
  var host string
  var master bool
  var slave bool

  flag.StringVar(&host, "host", "", "Mesos host address(e.g 127.0.0.1:5050)")
  flag.StringVar(&host, "h", "", "Mesos host address(e.g 127.0.0.1:5050)")
  flag.StringVar(&mesosType, "type", "", "Mesos type(e.g '(m)aster' or (s)lave")
  flag.StringVar(&mesosType, "t", "", "Mesos type(e.g '(m)aster' or (s)lave")
  flag.Parse()

  command := flag.Arg(0)

  switch mesosType {
  case "master", "m":
    opts := &megos.MesosMasterOptions{Host: host}
    m := megos.NewMasterClient(opts)
    _ = ExportMasterMetrics(m)
  case "slave", "s":
    opts := &megos.MesosSlaveOptions{Host: host}
    s := megos.NewSlaveClient(opts)
    _ = ExportSlaveMetrics(s)
  default:
    Println(flag.Usage())
  }
}
