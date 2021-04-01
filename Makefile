server:
	go run cmd/server/main.go

# ==============================================================================
# Docker

docker-build:
	docker build -t oauthdebugger .

docker-run:
	docker run --name oauthdebugger --rm -p 8090:8090 oauthdebugger:latest

deploy-build:
	docker build -t oad-deploy --file deploy.Dockerfile .

deploy-run: deploy-build
	export SECRETS_DIR=`pwd`/secrets
	docker run -it --entrypoint="/bin/bash" --name dply --rm -e "GCP_PROJECT=$(GCP_PROJECT)" -e "GCP_ACCOUNT=$(GCP_ACCOUNT)" -v $(SECRETS_DIR):/secrets oad-deploy

deploy-all: deploy-build
	export SECRETS_DIR=`pwd`/secrets
	docker run --name dply --rm -e "GCP_PROJECT=$(GCP_PROJECT)" -e "GCP_ACCOUNT=$(GCP_ACCOUNT)" -v $(SECRETS_DIR):/secrets oad-deploy

# ==============================================================================
# Go Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

# ==============================================================================
# GCP

define gcp_deploy
  gcloud functions deploy $(1) --entry-point=$(2) --project=$(GCP_PROJECT) --allow-unauthenticated --region=$(REGION) --runtime $(RUNTIME) --trigger-http --memory=$(MEM) --env-vars-file=/secrets/.env.yaml
endef

gcp-authorize:
	gcloud auth activate-service-account $(GCP_ACCOUNT) --key-file=/secrets/key.json

gcp-deploy-all: gcp-deploy-authorize gcp-deploy-create-client gcp-deploy-code-grant gcp-deploy-token

gcp-deploy-authorize: gcp-authorize
	$(call gcp_deploy,authorize,Authorize)

gcp-deploy-create-client: gcp-authorize
	$(call gcp_deploy,create_client,CreateClient)

gcp-deploy-code-grant: gcp-authorize
	$(call gcp_deploy,code_grant,CodeGrant)
	
gcp-deploy-token: gcp-authorize
	$(call gcp_deploy,token,Token)

# ==============================================================================
# Front end

firebase-deploy:
	firebase deploy --only hosting:outh-debugger