#!/bin/bash
if [[ $BUILDTOP == "" ]]; then
  echo "BUILDTOP required but not defined"
  exit
fi

cd $BUILDTOP/cmd/strekker;

go build -o ../../bin/strekker;

if [ $? -eq 0 ]; then
  echo "Build successful!"
else
  echo "Build failed."
fi