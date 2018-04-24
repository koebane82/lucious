package main
import (
  "time"
  "syscall"
  "os"
)

func file_info(fi os.FileInfo)(int,int,uint64,bool,int64,int64,int64){
  stat := fi.Sys().(*syscall.Stat_t)
  atime := time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec)).UnixNano()
  ctime := time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec)).UnixNano()
  mtime := time.Unix(int64(stat.Mtimespec.Sec), int64(stat.Mtimespec.Nsec)).UnixNano()
  inode := uint64(stat.Ino)
  nlink := uint32(stat.Nlink)
  is_hardlink := false
  if nlink >= 2 {
    is_hardlink = true
  }

  return int(stat.Uid), int(stat.Gid), inode, is_hardlink, ctime, atime, mtime
}
