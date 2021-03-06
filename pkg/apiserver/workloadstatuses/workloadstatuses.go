/*

 Copyright 2019 The KubeSphere Authors.

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

package workloadstatuses

import (
	"github.com/emicklei/go-restful"
	"net/http"

	"kubesphere.io/kubesphere/pkg/models/status"
	"kubesphere.io/kubesphere/pkg/server/errors"
)

func GetClusterAbnormalWorkloads(req *restful.Request, resp *restful.Response) {
	res, err := status.GetClusterResourceStatus()
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, errors.Wrap(err))
		return
	}
	resp.WriteAsJson(res)
}

func GetNamespacedAbnormalWorkloads(req *restful.Request, resp *restful.Response) {
	res, err := status.GetNamespacesResourceStatus(req.PathParameter("namespace"))
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, errors.Wrap(err))
		return
	}
	resp.WriteAsJson(res)
}
