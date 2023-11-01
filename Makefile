# note: call scripts from /scripts

.PHONY: app \
	clean \
	container \
	container-clean \
	container-run \
	container-stop

help:
	@echo "run 'make app args_appname={应用名称}' to build"
	@echo "run 'make clean args_appname={应用名称}' to clean"
	@echo "run 'make container args_appname={应用名称}' to build docker image"
	@echo "run 'make container-clean args_appname={应用名称}' to delete docker image"
	@echo "run 'make container-run args_appname={应用名称}' to run docker container"
	@echo "run 'make container-stop args_appname={应用名称}' to stop docker container"

app:
	@sh scripts/build.sh ${args_appname}

app-with-docker:
	@sh scripts/build-with-docker.sh ${args_appname}

clean:
	@sh scripts/clean.sh ${args_appname}

container:
	@sh scripts/container-build.sh ${args_appname} 

container-clean: container-stop
	@sh scripts/container-clean.sh ${args_appname}

container-run: container
	@sh scripts/container-run.sh ${args_appname}

container-stop:
	@sh scripts/container-stop.sh ${args_appname}