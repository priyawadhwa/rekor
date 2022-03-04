#!/bin/bash
#
# Copyright 2021 The Sigstore Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -ex
# Spin up services as usual

echo "starting services"
docker-compose up -d

echo "building CLI and server"
go build -o rekor-cli ./cmd/rekor-cli
go build -o rekor-server ./cmd/rekor-server

count=0

echo -n "waiting up to 60 sec for system to start"
until [ $(docker-compose ps | grep -c "(healthy)") == 3 ];
do
    if [ $count -eq 6 ]; then
       echo "! timeout reached"
       exit 1
    else
       echo -n "."
       sleep 10
       let 'count+=1'
    fi
done

echo


# Add some things to the tlog :)
rekor-cli upload --artifact test_file.txt --signature test_file.sig --public-key test_public_key.key --rekor_server http://localhost:3000
cd sharding-testdata
rekor-cli upload --artifact file1 --signature file1.sig --pki-format=x509 --public-key=ec_public.pem --rekor_server http://localhost:3000
rekor-cli upload --artifact file2 --signature file2.sig --pki-format=x509 --public-key=ec_public.pem --rekor_server http://localhost:3000
cd ..

# Now, we want to shard the log.
# First, spin down the rekor server

