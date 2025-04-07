# Loki Recording Rules

This Lab contains Loki recording rules. We have setup Loki ruler in two
different storage backends: S3 and local (filesystem).

- **S3 bucket**: In this approach Grafana has write permission to the rules
defined in S3 bucket, and we can define new rules via Grafana.

- **Local**: In this approach Grafana cannot edit the rules defined in local
filesystem. We can setup GitOps to manage the rules defined in local.