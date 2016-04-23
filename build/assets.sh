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
YEAR=$(date +%Y)

TMP_ASSETS_OUTPUT="/tmp/cadvisor_assets.go"
ASSETS_OUTPUT=$GIT_ROOT/pages/static/assets.go

TMP_TEMPLATE_OUTPUT="/tmp/cadvisor_templates.go"
TEMPLATE_OUTPUT=$GIT_ROOT/pages/templates.go

go get -u github.com/jteeuwen/go-bindata/...

for f in $GIT_ROOT/pages/assets/js/* $GIT_ROOT/pages/assets/styles/*; do
  if [ "$f" -nt $ASSETS_OUTPUT -o ! -e $ASSETS_OUTPUT ]; then
    go-bindata -o $ASSETS_OUTPUT -pkg static $GIT_ROOT/pages/assets/js/... $GIT_ROOT/pages/assets/styles/...
    cat build/boilerplate/boilerplate.go.txt | sed "s/YEAR/$YEAR/" > "${TMP_ASSETS_OUTPUT}"
    echo -e "// generated by build/assets.sh; DO NOT EDIT\n" >> "${TMP_ASSETS_OUTPUT}"
    cat "${ASSETS_OUTPUT}" >> "${TMP_ASSETS_OUTPUT}"
    gofmt -w -s "${TMP_ASSETS_OUTPUT}"
    mv "${TMP_ASSETS_OUTPUT}" "${ASSETS_OUTPUT}"
    break
  fi
done

for f in $GIT_ROOT/pages/assets/html/*; do
  if [ "$f" -nt $TEMPLATE_OUTPUT -o ! -e $TEMPLATE_OUTPUT ]; then
    go-bindata -o $TEMPLATE_OUTPUT -pkg pages $GIT_ROOT/pages/assets/html/...
    cat build/boilerplate/boilerplate.go.txt | sed "s/YEAR/$YEAR/" > "${TMP_TEMPLATE_OUTPUT}"
    echo -e "// generated by build/assets.sh; DO NOT EDIT\n" >> "${TMP_TEMPLATE_OUTPUT}"
    cat "${TEMPLATE_OUTPUT}" >> "${TMP_TEMPLATE_OUTPUT}"
    gofmt -w -s "${TMP_TEMPLATE_OUTPUT}"
    mv "${TMP_TEMPLATE_OUTPUT}" "${TEMPLATE_OUTPUT}"
    break
  fi
done

exit 0 
