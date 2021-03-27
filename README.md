# oauth-debugger

An app for debugging oauth code, a mock Service Api

## Build

This app can be built, and run, in many different ways.

### Docker

There is a Dockerfile with this project that can be used to run the application.

```bash
docker build -t oauthdebugger .
docker run --name oauthdebugger --rm -p 8090:8090 oauthdebugger:latest
```

This Docker build will build the Go app, the Webpack front end, and finally runs the app.

### Just the UI

The UI is built, or dev served, by Webpack.

```bash
npm install
npm run build # for production build
npm run dev # for dev server
```

### Just the Go Functions

There is a Go program that can host all the functions in one server for local development.

```bash
go build ./cmd/service/main.go
```

## Deploy

The app is deployed in two pieces, the UI and Go functions.

### UI

The UI is hosted using GCP Firebase. To deploy it using the CLI,

```bash
firebase deploy --only hosting:outh-debugger
```

### Go Functions

The Go functions are hosted as GCP cloud functions. These can be deployed via CLI in the following manner,

```bash
cd functions
gcloud functions deploy client --allow-unauthenticated --region=us-central1 \
  --runtime go113 --trigger-http --memory=128MB --entry-point=Client
```
