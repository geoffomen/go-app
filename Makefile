# note: call scripts from /scripts

help:
	@echo "run 'make myapp' to build"
	@echo "run 'make clean' to clean"

myapp:
	@sh scripts/build.sh myapp

myapp-clean:
	@sh scripts/clean.sh myapp

myapp-pod:
	@sh scripts/pod_create.sh pod_myapp

myapp-pod-clean:
	@sh scripts/pod_clean.sh pod_myapp
	
myapp-container:
	@sh scripts/container-build.sh myapp build/package/Dockerfile_myapp

myapp-container-run: myapp-container myapp-pod myapp-db-container-run
	@sh scripts/container-run.sh pod_myapp myapp

myapp-container-clean: 
	@sh scripts/container-clean.sh myapp

myapp-db-container-run: 
	@sh scripts/mysql-container-run.sh pod_myapp password /tmp/mysql_data

myapp-db-container-clean: 
	@sh scripts/container-clean.sh mysql

myapp-db-container-data-import: 
	@sh scripts/mysql-import.sh password ./docs/myapp.sql

myapp-db-container-data-export: 
	@sh scripts/mysql-export.sh password myapp /tmp/myapp_database_dump.sql