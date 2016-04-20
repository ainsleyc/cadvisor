#!/bin/bash

# Copyright 2015 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GIT_ROOT=$(dirname "${BASH_SOURCE}")/..

STATUS=0

go get -u github.com/jteeuwen/go-bindata/...

for f in $GIT_ROOT/pages/assets/html/*; do
  if [ "$f" -nt $GIT_ROOT/pages/templates.go ]; then
    go-bindata -o $GIT_ROOT/pages/templates.go -pkg pages $GIT_ROOT/pages/assets/html/... || STATUS=$?
    break
  fi
done

for f in $GIT_ROOT/pages/assets/js/* $GIT_ROOT/pages/assets/styles/*; do
  if [ "$f" -nt $GIT_ROOT/pages/static/assets.go ]; then
    go-bindata -o $GIT_ROOT/pages/static/assets.go -pkg static $GIT_ROOT/pages/assets/js/... $GIT_ROOT/pages/assets/styles/... || STATUS=$?
    break
  fi
done

exit $STATUS

