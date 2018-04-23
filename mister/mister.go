package main

import (
  "plugin"
  "fmt"
  "os"
)

type PluginReturn interface {
  RunPlugin(p string)(string)
}

func main() {
  plug_file := "plugins/michael.so"

  plug,err := plugin.Open(plug_file)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  symPlugin, err := plug.Lookup("PluginReturn")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  var pluginreturn PluginReturn
  pluginreturn, ok := symPlugin.(PluginReturn)
  if !ok {
    fmt.Println("unexpected type from module symbol")
    os.Exit(1)
  }

  fmt.Println(pluginreturn.RunPlugin("/"))
}