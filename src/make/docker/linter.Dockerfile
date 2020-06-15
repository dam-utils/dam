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

FROM golangci/golangci-lint:v1.27

ARG PROJECT_NAME

WORKDIR /go/src/${PROJECT_NAME}
COPY . .

RUN mv /go/src/${PROJECT_NAME}/_build/goget_cache/github.com /go/src/github.com
RUN go get -d ./
RUN golangci-lint run