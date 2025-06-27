#!/bin/bash
go build -o acnh-finder-build cmd/web/*.go && ./acnh-finder-build  -cache=false -production=false -cognito-user-pool-id=1234 -cognito-client-id=1234