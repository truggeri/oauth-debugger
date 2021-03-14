# oauth-debugger

An app for debugging oauth code, a mock Service Api

## Deploys

### UI

```bash
firebase deploy --only hosting:outh-debugger
```

### Function

```bash
cd functions
gcloud functions deploy client --allow-unauthenticated --region=us-central1 \
  --runtime go113 --trigger-http --memory=128MB --entry-point=Client
```
