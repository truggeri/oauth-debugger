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

RUNTIME = go113
ZONE = us-central1

gcp-deploy-all: gcp-deploy-authorize gcp-deploy-client gcp-deploy-token

gcp-deploy-authorize:
	cd functions &&	gcloud functions deploy authorize --allow-unauthenticated --region=$(ZONE) --runtime $(RUNTIME) --trigger-http --memory=128MB --entry-point=Authorize

gcp-deploy-client:
	cd functions
	gcloud functions deploy client --allow-unauthenticated --region=$(ZONE) --runtime $(RUNTIME) --trigger-http --memory=128MB --entry-point=Client
	
gcp-deploy-token:
	cd functions
	gcloud functions deploy token --allow-unauthenticated --region=$(ZONE) --runtime $(RUNTIME) --trigger-http --memory=128MB --entry-point=Token
	