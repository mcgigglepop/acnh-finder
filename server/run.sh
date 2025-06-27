#!/bin/bash
go build -o acnh-finder-build cmd/web/*.go && ./acnh-finder-build  -cache=false -production=false -cognito-user-pool-id=us-east-1_X2nr3u6aE -cognito-client-id=5ah5not8osfrbmootsercpe0ns