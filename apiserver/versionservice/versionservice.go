package versionservice

/*
Copyright 2018 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/crunchydata/postgres-operator/apiserver"
	"net/http"
)

// VersionHandler ...
// pgo version
func VersionHandler(w http.ResponseWriter, r *http.Request) {

	log.Debug("versionservice.VersionHandler called")

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	err := apiserver.Authn(apiserver.VERSION_PERM, w, r)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := Version()

	json.NewEncoder(w).Encode(resp)
}
