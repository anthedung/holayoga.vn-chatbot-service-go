#!/usr/bin/env bash

docker-compose run --rm test ./build.sh

docker run --rm -v `pwd`:/build -e BUILD_OS=ye7 -e JENKINS_BUILD_NUMBER=1 -e BUILD_NUMBER=5  -w /build artifactory.ges.symantec.com/ase-docker/fpm:yel7 bash -c "rm -rf *.rpm; chmod +x create-rpm.sh; BUILD_OS=ye7 ./create-rpm.sh"
