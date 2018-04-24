package main

import (
  "fmt"
  "os"
  "flag"
  "plugin"
)

type PluginReturn interface {
  RunPlugin(p string)(string)
}

func main() {
  pluginPtr := flag.String("plugin","none","A Plugin File to test")
  argPtr := flag.String("arg","no_arg","Argument to pass to plugin")
  flag.Parse()

  plug,err := plugin.Open(*pluginPtr)
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

  fmt.Println(pluginreturn.RunPlugin(*argPtr))
}