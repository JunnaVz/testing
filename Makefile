test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing" && \
	go test -tags=unit /home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/unit_tests/unit_services/... \
	/home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/unit_tests/unit_repositories/... --race --parallel 11
	cp environment.properties allure-results

allure:
	cp -R allure-reports/history allure-results
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure
#report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
 	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
 	export DB_INIT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing/db/sql/init.sql" && \
 	go test -tags=unit /home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/unit_tests/unit_services/... \
	/home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/unit_tests/unit_repositories/... --race

local-unit:
	export ALLURE_OUTPUT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing" && \
 	export DB_INIT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing/db/sql/init.sql" && \
 	go test -tags=unit /home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/unit_tests/unit_services/... \
	/home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/unit_tests/unit_repositories/... --race

ci-integration:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
 	export DB_INIT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing/db/sql/init.sql" && \
	go test -tags=integration /home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/integration/category_test.go --race

local-integration:
	export ALLURE_OUTPUT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing" && \
 	export DB_INIT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing/db/sql/init.sql" && \
	go test -tags=integration /home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/integration/category_test.go --race

ci-e2e:
	docker compose up -d
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="e2e-allure" && \
	go test -tags=e2e /home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/integration/e2e_test.go --race
	docker compose down
	docker image rm testing-backend:latest bitnami/postgresql:16 alpine:latest

local-e2e:
	docker compose up -d
	export ALLURE_OUTPUT_PATH="/home/pikasoft/Documents/jovana/sem7/TEST/testing" && \
	go test -tags=e2e /home/pikasoft/Documents/jovana/sem7/TEST/testing/tests/integration/e2e_test.go --race

rm-compose:
	docker compose down
	docker image rm testing-backend:latest

ci-concat-reports:
	mkdir allure-results
	cp unit-allure/* allure-results/
	cp integration-allure/* allure-results/
	cp e2e-allure/* allure-results/
	cp environment.properties allure-results

.PHONY: test allure report ci-unit local-unit ci-integration local-integration ci-e2e local-e2e rm-compose ci-concat-reports
