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
PROJECT_NAME := $(shell grep 'PROJECT_NAME' config/sys.config.go | awk -F '=' '{print $$2}' | awk -F '\"' '{print $$2}')
PROJECT_VERSION := $(shell grep 'PROJECT_VERSION' config/sys.config.go | awk -F '=' '{print $$2}' | awk -F '\"' '{print $$2}')

# Удалить за собой временные контейнеры docker
CLEAR_BUILD_CONTAINER	:= true
# Удалить за собой временные образы docker
CLEAR_BUILD_IMAGE		:= false
# true - кэширует промежуточные слои для образов docker
NO_DOCKER_IMAGE_CACHE	:= false
# Сохранять результат `go get` в кэше для ускорения сборки
# (!) Предполагаю, что кэш go пакетов для разных архитектур одинаковый
USE_GO_GET_CACHE		:= true

#preparing
include src/make/Makefile.funcs

build:	build-windows build-linux

build-windows:
	$(call build_func,windows,amd64)

build-linux:
	$(call build_func,linux,amd64)

test:
	$(call test_func)

lint:
	$(call lint_func)

clean: clean-docker clean-binary
	rm -rf _build || true

clean-binary:
	rm -rf _build/linux/${PROJECT_NAME} || true
	rm -rf _build/windows/${PROJECT_NAME} || true
	rm -rf src/build/dam-linux/${PROJECT_NAME} || true

clean-docker:
	$(call clear_func,windows)
	$(call clear_func,linux)

app-linux: clean-binary	build-linux
	cp -f _build/linux/${PROJECT_NAME} src/build/dam-linux/${PROJECT_NAME}
	src/build/dam-linux/${PROJECT_NAME} create src/build/dam-linux/ --name ${PROJECT_NAME} --version ${PROJECT_VERSION}
