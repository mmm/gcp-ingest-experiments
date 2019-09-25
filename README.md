# Ingestion

Here's a basic serverless ingestion pipeline

    Run -> PubSub -> DataFlow -> GCS -> BQ

do we actually _need_ k8s for anything serverless?
maybe to back up Run services?


## Layers

### Core

- network
- IAM / service accts
- bastion access (nec for serverless?)
- prometheus or stackdriver
- cert config
- config mgmt of some sort?

### Pipeline

- Pub/Sub
- DataFlow
- Buckets
- BQ tables
- BT tables

### Ingest

- Run
  - how do these work elastically?
  - how cold are the starts?
  - are we better off with a simple k8s deployment+svc?
  - do we need to muck about with knative explicitly?
- ESP
- LB?

### Debugging


## Auth

- IAM for service accounts
- Cert stuff?
- ESP config and ACLs of any sort?
- Cloud IOT?


## Performance

- stackdriver
- seige engine of some sort?


## Example Scenario

- generate characteristic test data
- generate characteristic queries
- assemble a pipeline to glue those together
