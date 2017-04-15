#!/usr/bin/env bash

set -e

BUILD_TYPE="${BUILD_TYPE:-dev}"         # dev or prod

if [ ${BUILD_TYPE} == "dev" ];
then
    MAP_VOLUME="-v ${PWD}/..:/go/src/github3.cisco.com/robot/pulse "
    echo ${MAP_VOLUME}
fi

docker run -it ${MAP_VOLUME} \
    --name=gomonit --hostname=gomonit \
    --rm \
    narenm/golang