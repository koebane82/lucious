#!/bin/bash
os="centos6"
repo_dir="/usr/local/repo/${REPO}"
REPO="${repo_dir}/go/src/${MODULE}"
GOPATH="${repo_dir}/go"

export GOPATH

go_bin='/usr/local/go/bin/go'
if [ $PLUGIN ];then
  OUTPUT_DIR="${repo_dir}/go/bin/${os}/plugins"
else
  OUTPUT_DIR="${repo_dir}/go/bin/${os}"
fi

if [ ! -d $OUTPUT_DIR ];then
  mkdir -p $OUTPUT_DIR
fi

cd $REPO

if [ $PLUGIN ];then
  echo "BUILDING PLUGIN - ${MODULE}"
  $go_bin build -buildmode=plugin -o ${OUTPUT_DIR}/${MODULE}.so ${MODULE}
else
  echo "BUILDING APP - ${MODULE}"
  $go_bin build -o ${OUTPUT_DIR}/${MODULE} ${MODULE}
  chmod 755 ${OUTPUT_DIR}/${MODULE}
fi

echo "Checking for UnitTest"
if [ -f "unit_test.sh" ];then
  echo "Unit Test Found running"
  ./unit_test.sh ${os}
fi