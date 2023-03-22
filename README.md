# gerror

gerror provides a set of wrappers around the Go `error` type to allow
carrying clearly distinguished private / public messages and metadata
that can be used to determine severity and status codes to use for
errors throughout a codebase. Refer to the package's
[godocs](https://pkg.go.dev/github.com/grafana/gerror) for
documentation and examples of how to use the package to construct
errors.

> **Warning**
> _gerror_ does not currently have a stable API.

_This is not an officially supported Grafana Labs product._

## Licensing

Copyright 2023 Grafana Labs

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.