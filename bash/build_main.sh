#!/bin/bash
if [[ $BUILDTOP == "" ]]; then
  echo "BUILDTOP required but not defined"
  exit
fi

cd $BUILDTOP/cmd/stargazer;

go build -o ../../bin/stargazer;

if [ $? -eq 0 ]; then
  echo "Build successful!"
else
  echo "Build failed."
fi