#!/bin/bash
os=$1

if [ ! -f ../../bin/${os}/listlus ];then
  ../../app/listlus/build.sh $os
fi

cd ../../bin/${os}

args=("/" "/bin" "/usr")

for arg in "${args[@]}";do
  echo "testing with '${arg}'"
  ls -l $arg
  ./listlus -plugin plugins/michael.so -arg $arg
done