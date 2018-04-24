package main

import (
  "fmt"
  "os"
  "io/ioutil"
)

type pluginexec string
// this fucntion should be parsed with ^ as field seperator
// 0 - file id
// 1 - parent file id
// 2 - filename
// 3 - link_parent
// 4 - inode
// 5 - file permissions
// 6 - is Dir
// 7 - is hardlink
// 8 - user id
// 9 - group id
// 10 - filesize - bytes
// 11 - ctime - epoch
// 12 - atime - epoch
// 13 - mtime - epoch
func list_file(p string,id int,pid int) (string,bool,int) {
  file, err := os.Stat(p)
  if err != nil {
    return "", false,  1
  }

  filemode := file.Mode()

  link_parent := "-"

  if filemode&os.ModeSymlink != 0 {
    link, err := os.Readlink(file.Name())
    if err == nil {
        link_parent = link
    }
  }

  uid, gid, inode, is_hardlink,ctime, atime, mtime := file_info(file)

  ret_val := fmt.Sprintf("%d^%d^%s^%s^%d^%o^%t^%t^%d^%d^%d^%d^%d^%d",id,pid,file.Name(),link_parent,inode,filemode.Perm(),filemode.IsDir(),is_hardlink,uid,gid,file.Size(),ctime,atime,mtime)
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
