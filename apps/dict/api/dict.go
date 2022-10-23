package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/tqtcloud/cloud-manage/apps/dict"
)

func (h *handler) CrendentialType(r *restful.Request, w *restful.Response) {
	response.Success(w, dict.CrendentialTypes)
}

func (h *handler) Vendor(r *restful.Request, w *restful.Response) {
	response.Success(w, dict.Vendors)
}

func (h *handler) ResourceType(r *restful.Request, w *restful.Response) {
	response.Success(w, dict.ResourceTypes)
}

func (h *handler) VendorRegion(r *restful.Request, w *restful.Response) {
	response.Success(w, dict.Regions)
}
