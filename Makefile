# note: call scripts from /scripts

.PHONY: app \
	clean \
	container \
	container_clean \
	container_run \
	container_stop

help:
	@echo "run 'make app args_appname={应用名称}' to build"
	@echo "run 'make app-with-docker args_appname={应用名称}' to cross build with docker image"
	@echo "run 'make clean args_appname={应用名称}' to clean"
	@echo "run 'make container args_appname={应用名称}' to build docker image"
	@echo "run 'make container-clean args_appname={应用名称}' to delete docker image"
	@echo "run 'make container-run args_appname={应用名称}' to run docker container"
	@echo "run 'make container-stop args_appname={应用名称}' to stop docker container"

app:
	@sh scripts/build.sh ${args_appname}

app-with-docker:
	@sh scripts/build_with_docker.sh ${args_appname}

clean:
	@sh scripts/clean.sh ${args_appname}

container:
	@sh scripts/container_build.sh ${args_appname} 

container-clean: container_stop
	@sh scripts/container_clean.sh ${args_appname}

container-run: container
	@sh scripts/container_run.sh ${args_appname}

container-stop:
	@sh scripts/container_stop.sh ${args_appname}