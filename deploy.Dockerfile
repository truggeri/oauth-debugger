FROM gcr.io/google.com/cloudsdktool/cloud-sdk:latest as node-build

# GCP-PROJECT - Name of the Google Cloud/Firebase project for deployment
# GCP-ACCOUNT - Name of the GCP Service Account used for credentials

RUN gcloud version

ENV REGION us-central1
ENV RUNTIME go113
ENV MEM 128MB

VOLUME [ "/secrets" ]

RUN mkdir /deploy
COPY Makefile /deploy
COPY go.* /deploy/
COPY templates /deploy
COPY functions /deploy

WORKDIR /deploy
CMD [ "make", "gcp-deploy-all" ]