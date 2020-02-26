#  Copyright 2020 The Docker Applications Manager Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
GO_VERSION := 1.13.4
PWD=$(shell pwd)
PROJECT=$(shell (basename ${PWD}))

#if we have tty. then pseudo_tty="-t".
pseudo_tty = $(shell (tty 2>&1 > /dev/null) && (echo "-t"))

all: test
	build

build:
	docker run 	-i \
				$(pseudo_tty) \
				--rm \
				-v "${PWD}":/go/src/${PROJECT} \
				-v "${PWD}/_build":/go/src \
				-w /go/src/${PROJECT} \
				golang:${GO_VERSION} /bin/bash -c "go get -d ./ && \
					go build -o ${PROJECT} main.go"

test:
	docker run 	-i \
				$(pseudo_tty) \
				--rm \
				-v "${PWD}":/go/src/${PROJECT} \
				-v "${PWD}/_build":/go/src \
				-w /go/src/${PROJECT} \
				golang:${GO_VERSION} /bin/bash -c "go get -d ./ && \
					go test -v ${PROJECT}/run"

clean:
	docker run 	-i \
    			$(pseudo_tty) \
    			--rm \
    			-v "${PWD}":/pwd \
    			-w /pwd \
    			golang:${GO_VERSION} /bin/bash -c "rm -rf ./_build ./${PROJECT}"
