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
	rm -rf vendor || true

clean-binary:
	rm -rf _build/linux/${PROJECT_NAME} || true
	rm -rf _build/windows/${PROJECT_NAME} || true
	rm -rf src/build/dam-linux/${PROJECT_NAME} || true

clean-docker:
	$(call clear_func,windows)
	$(call clear_func,linux)

app-linux: clean-binary	build-linux
	cp -f _build/linux/${PROJECT_NAME} src/build/dam-linux/${PROJECT_NAME}
	src/build/dam-linux/${PROJECT_NAME} -x create src/build/dam-linux/ --name ${PROJECT_NAME} --version ${PROJECT_VERSION}
