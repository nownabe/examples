bundle exec coaster help
bundle exec coaster schema schema_definitions.yaml
bundle exec coaster java_class \
  --package com.nownabe.examples.coaster \
  --class-name UserActivityLog \
  --dataflow \
  schema_definitions.yaml
bundle exec coaster dataflow_converter \
  --package com.nownabe.examples.coaster \
  --class-name BigQueryConverter \
  --source-class UserActivityLog \
  schema_definitions.yaml
