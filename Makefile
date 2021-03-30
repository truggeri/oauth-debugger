server:
	go run cmd/server/main.go

# ==============================================================================
# Modules support

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
  gcloud functions deploy $(1) --entry-point=$(2) --project=$(GCP-PROJECT) --allow-unauthenticated --region=$(REGION) --runtime $(RUNTIME) --trigger-http --memory=$(MEM) --env-vars-file=/secrets/.env.yaml
endef

gcp-authorize:
	gcloud auth activate-service-account $(GCP-ACCOUNT) --key-file=/secrets/key.json

gcp-deploy-all: gcp-deploy-authorize gcp-deploy-create-client gcp-deploy-code-grant gcp-deploy-token

gcp-deploy-authorize: gcp-authorize
	$(call gcp_deploy,authorize,Authorize)

gcp-deploy-create-client: gcp-authorize
	$(call gcp_deploy,create_client,CreateClient)

gcp-deploy-code-grant: gcp-authorize
	$(call gcp_deploy,code_grant,CodeGrant)
	
gcp-deploy-token: gcp-authorize
	$(call gcp_deploy,token,Token)
	