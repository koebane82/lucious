#!/bin/bash
os=$1
user=$(whoami)
repo_name='lucious'
module_name='michael'
repos_base="/Users/${user}/REPOS"
if [ $os == "darwin" ];then
  if [ $(uname) != "Darwin" ];then
    echo "You can't build Darwin in Docker and this OS isn't Darwin"
  else
    out_dir="${repos_base}/${repo_name}/go/bin/darwin/plugins"
    build_file="${repos_base}/${repo_name}/go/${module_name}/"
    if [ ! -d $out_dir ];then
      mkdir -p $out_dir
    fi
    go build -o "${out_dir}/${module_name}.so" -buildmode=plugin ${module_name}
  fi
else
  docker run --mount type=bind,source=${repos_base},target=/usr/local/repo -e REPO=$repo_name -e MODULE=${module_name} -e PLUGIN=true "${os}-go:latest"
fi

