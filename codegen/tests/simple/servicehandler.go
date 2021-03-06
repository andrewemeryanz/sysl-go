// Code generated by sysl DO NOT EDIT.
package simple

import (
	"encoding/json"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/convert"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Handler interface for Simple
type Handler interface {
	GetJustOkAndJustErrorListHandler(w http.ResponseWriter, r *http.Request)
	GetJustReturnErrorListHandler(w http.ResponseWriter, r *http.Request)
	GetJustReturnOkListHandler(w http.ResponseWriter, r *http.Request)
	GetOkTypeAndJustErrorListHandler(w http.ResponseWriter, r *http.Request)
	GetOopsListHandler(w http.ResponseWriter, r *http.Request)
	GetRawListHandler(w http.ResponseWriter, r *http.Request)
	GetRawIntListHandler(w http.ResponseWriter, r *http.Request)
	GetStuffListHandler(w http.ResponseWriter, r *http.Request)
	PostStuffHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for Simple API
type ServiceHandler struct {
	genCallback      GenCallback
	serviceInterface *ServiceInterface
}

// NewServiceHandler for Simple
func NewServiceHandler(genCallback GenCallback, serviceInterface *ServiceInterface) *ServiceHandler {
	return &ServiceHandler{genCallback, serviceInterface}
}

// GetJustOkAndJustErrorListHandler ...
func (s *ServiceHandler) GetJustOkAndJustErrorListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetJustOkAndJustErrorList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetJustOkAndJustErrorListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetJustOkAndJustErrorListClient{}

	err := s.serviceInterface.GetJustOkAndJustErrorList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, nil)
}

// GetJustReturnErrorListHandler ...
func (s *ServiceHandler) GetJustReturnErrorListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetJustReturnErrorList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetJustReturnErrorListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetJustReturnErrorListClient{}

	err := s.serviceInterface.GetJustReturnErrorList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, nil)
}

// GetJustReturnOkListHandler ...
func (s *ServiceHandler) GetJustReturnOkListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetJustReturnOkList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetJustReturnOkListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetJustReturnOkListClient{}

	err := s.serviceInterface.GetJustReturnOkList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, nil)
}

// GetOkTypeAndJustErrorListHandler ...
func (s *ServiceHandler) GetOkTypeAndJustErrorListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetOkTypeAndJustErrorList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetOkTypeAndJustErrorListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetOkTypeAndJustErrorListClient{}

	response, err := s.serviceInterface.GetOkTypeAndJustErrorList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, response)
}

// GetOopsListHandler ...
func (s *ServiceHandler) GetOopsListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetOopsList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetOopsListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetOopsListClient{}

	response, err := s.serviceInterface.GetOopsList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, response)
}

// GetRawListHandler ...
func (s *ServiceHandler) GetRawListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetRawList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetRawListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetRawListClient{}

	str, err := s.serviceInterface.GetRawList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, str)
}

// GetRawIntListHandler ...
func (s *ServiceHandler) GetRawIntListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetRawIntList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetRawIntListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetRawIntListClient{}

	integer, err := s.serviceInterface.GetRawIntList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, integer)
}

// GetStuffListHandler ...
func (s *ServiceHandler) GetStuffListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetStuffList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetStuffListRequest

	var DtParam string

	var StParam string

	var BtParam string

	var ItParam string

	var convErr error

	DtParam = restlib.GetQueryParam(r, "dt")
	StParam = restlib.GetQueryParam(r, "st")
	BtParam = restlib.GetQueryParam(r, "bt")
	ItParam = restlib.GetQueryParam(r, "it")
	req.Dt, convErr = convert.StringToTimePtr(ctx, DtParam)
	if convErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	req.St, convErr = convert.StringToStringPtr(ctx, StParam)
	if convErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	req.Bt, convErr = convert.StringToBoolPtr(ctx, BtParam)
	if convErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	req.It, convErr = convert.StringToIntPtr(ctx, ItParam)
	if convErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetStuffListClient{}

	stuff, err := s.serviceInterface.GetStuffList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, stuff)
}

// PostStuffHandler ...
func (s *ServiceHandler) PostStuffHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.PostStuff == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req PostStuffRequest

	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&req.Request)
	if decodeErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Error reading request body", decodeErr)
		return
	}

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := PostStuffClient{}

	str, err := s.serviceInterface.PostStuff(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, str)
}
