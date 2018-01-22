#!/usr/bin/env bash
cd /go/src/stash.ges.symantec.com/scm/oasis-log-ingestor
glide install
cd main && go build -o oasis-log-ingestor
