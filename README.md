# oauth-debugger

An app for debugging oauth code, a mock Service Api

## Documentation

See our [api documentation](https://oauth-debugger.truggeri.com/docs) for full details of each endpoint that's available.

## Build

This app can be built, and run, in many different ways.

### Docker

There is a Dockerfile with this project that can be used to run the application.
This Docker build will build the Go app, the Webpack front end, and finally runs the app.

```bash
docker build -t oauthdebugger .
docker run --name oauthdebugger --rm -p 8090:8090 oauthdebugger:latest
```

or

```bash
make docker-run
```

### Just the UI

The UI is built, or dev served, by Webpack.

```bash
npm install
npm run build # for production build
npm run dev # for dev build
npm run serve # for dev server
```

### Just the Go Functions

There is a Go program that can host all the functions in one server for local development.

```bash
go build ./cmd/service/main.go
```

## Deploy

The app is deployed in two pieces, the UI and Go functions.

### Docker Deploy

There is a Dockerfile made specifically for deploying the app to Google Cloud. This app consumes two environment variables for configuration, and mounts secrets in a volume.

```bash
docker build -t oad-deploy --file deploy.Dockerfile .
docker run -it --name dply --rm -e "GCP_PROJECT=xxx" -e "GCP_ACCOUNT=yyy" -v secrets:/secrets oad-deploy 
```

or

```bash
GCP_PROJECT=xxx GCP_ACCOUNT=yyy make deploy-run
```

| ENV | Value |
| ------------- | ------------- |
| GCP_PROJECT | Project ID |
| GCP_ACCOUNT | Email of GCP Service Account used for deployment |

### UI

The UI is hosted using GCP Firebase. To deploy it using the CLI,

```bash
firebase deploy --only hosting:outh-debugger
```

or

```bash
make firebase-deploy
```

### Go Functions

The Go functions are hosted as GCP cloud functions. These can be deployed via CLI in the following manner,
but note that you must manually copy other files and folders in. I would strongly advise using the Docker container
to deploy.

```bash
cd functions
gcloud functions deploy client --allow-unauthenticated --region=us-central1 \
  --runtime go113 --trigger-http --memory=128MB --entry-point=Client
```

## Licensing

### Graphics

Graphics are provided by [Twemoji](https://twemoji.twitter.com/). Thank you for providing free SVG graphics.

```text
Copyright 2020 Twitter, Inc and other contributors
Code licensed under the MIT License: http://opensource.org/licenses/MIT
Graphics licensed under CC-BY 4.0: https://creativecommons.org/licenses/by/4.0/
```
