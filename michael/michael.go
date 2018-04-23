package main

import (
  "fmt"
  "os"
  "io/ioutil"
)

type pluginexec string

func list_file(p string,id int,pid int) (string,bool,int) {
  file, err := os.Stat(p)
  if err != nil {
    return "", false,  1
  }

  filemode := file.Mode()
  ret_val := fmt.Sprintf("%d^%d^%s^%o^%t^%t",id,pid,file.Name(),filemode.Perm(),filemode.IsDir(),filemode.IsRegular())
  return ret_val, filemode.IsDir(), 0
}

//func list_dir(p string) ([]string,int) {
func list_dir(p string,pid int)(string) {
  files, err := ioutil.ReadDir(p)
  return_string := ""

  id := pid + 1
  if err == nil {
    for _, file := range files {
      info, dir, err := list_file(fmt.Sprintf("%s/%s",p,file.Name()),id,pid)
      if dir == true {
        //do nothing
      }
      if err == 0 {
        return_string = fmt.Sprintf("%s~#~%s",return_string,info)
      }
      id = id + 1
    }
  }
  return return_string
}

func (pl pluginexec)RunPlugin(p string)(string){
  file,is_dir,err := list_file(p,0,0)
  return_string := ""
  if err == 0 {
    return_string = file
    if is_dir == true {

      return_string = fmt.Sprintf("%s~#~%s",return_string,list_dir(p,0))
    }
  }
  return return_string
}

var PluginReturn pluginexec
