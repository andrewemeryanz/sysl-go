// Code generated by sysl DO NOT EDIT.
package simple

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/convert"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// *BusinessLogicError error
var BusinessLogicError = map[string]string{"name": "BusinessLogicError", "http_code": "1001", "http_message": "foo", "http_status": "500"}

// *BusinessLogicError2 error
var BusinessLogicError2 = map[string]string{"name": "BusinessLogicError2", "http_code": "1002", "http_message": "foo2", "http_status": "501"}

// Handler interface for Simple
type Handler interface {
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

// GeneratedMapError for Simple
func GeneratedMapError(ctx context.Context, err error) *common.HTTPError {
	if errors.As(err, BusinessLogicError) {
		return BusinessLogicError.HandleError()
	}

	if errors.As(err, BusinessLogicError2) {
		return BusinessLogicError2.HandleError()
	}

	return nil
}

// Handler Error
func (s *ServiceHandler) handleError(ctx context.Context, w http.ResponseWriter, kind common.Kind, message string, cause error) {
	httpError := GeneratedMapError(ctx, err)
	if httpError == nil {
		httpError := common.HandleError(ctx, err)
	}

	httpError.WriteError(ctx, w)
}

// GetOopsListHandler ...
func (s *ServiceHandler) GetOopsListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetOopsList == nil {
		s.handleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetOopsListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetOopsListClient{}

	response, err := s.serviceInterface.GetOopsList(ctx, &req, client)
	if err != nil {
		s.handleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, response, err)
}

// GetRawListHandler ...
func (s *ServiceHandler) GetRawListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetRawList == nil {
		s.handleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetRawListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetRawListClient{}

	str, err := s.serviceInterface.GetRawList(ctx, &req, client)
	if err != nil {
		s.handleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, str, err)
}

// GetRawIntListHandler ...
func (s *ServiceHandler) GetRawIntListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetRawIntList == nil {
		s.handleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetRawIntListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetRawIntListClient{}

	integer, err := s.serviceInterface.GetRawIntList(ctx, &req, client)
	if err != nil {
		s.handleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, integer, err)
}

// GetStuffListHandler ...
func (s *ServiceHandler) GetStuffListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetStuffList == nil {
		s.handleError(r.Context(), w, common.InternalError, "not implemented", nil)
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
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	req.St, convErr = convert.StringToStringPtr(ctx, StParam)
	if convErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	req.Bt, convErr = convert.StringToBoolPtr(ctx, BtParam)
	if convErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	req.It, convErr = convert.StringToIntPtr(ctx, ItParam)
	if convErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", convErr)
		return
	}

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := GetStuffListClient{}

	stuff, err := s.serviceInterface.GetStuffList(ctx, &req, client)
	if err != nil {
		s.handleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, stuff, err)
}

// PostStuffHandler ...
func (s *ServiceHandler) PostStuffHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.PostStuff == nil {
		s.handleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req PostStuffRequest

	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&req.Request)
	if decodeErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Error reading request body", decodeErr)
		return
	}

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		s.handleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	client := PostStuffClient{}

	str, err := s.serviceInterface.PostStuff(ctx, &req, client)
	if err != nil {
		s.handleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, str, err)
}
