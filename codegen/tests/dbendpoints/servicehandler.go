// Code generated by sysl DO NOT EDIT.
package dbendpoints

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/convert"
	"github.com/anz-bank/sysl-go/database"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Handler interface for DbEndpoints
type Handler interface {
	GetCompanyLocationListHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for DbEndpoints API
type ServiceHandler struct {
	genCallback      GenCallback
	serviceInterface *ServiceInterface
	DB               *sql.DB
}

// NewServiceHandler for DbEndpoints
func NewServiceHandler(genCallback GenCallback, serviceInterface *ServiceInterface) *ServiceHandler {
	db, err := database.GetDBHandle()
	if err != nil {
		return nil
	}

	return &ServiceHandler{genCallback, serviceInterface, db}
}

// Handler Error
func (s *ServiceHandler) handleError(ctx context.Context, w http.ResponseWriter, kind common.Kind, message string, cause error) {
	serverError := common.CreateError(ctx, kind, message, cause)
	httpError := s.genCallback.MapError(ctx, serverError)
	if httpError != nil {
		httpError.WriteError(ctx, w)
		return
	}

	t, ok := cause.(common.CustomError)
	if ok {
		e := t.HTTPError()
		httpError = &e
		httpError.WriteError(ctx, w)
		return
	}

	commonError := common.HandleError(ctx, serverError)
	httpError = &commonError
	httpError.WriteError(ctx, w)
}

// GetCompanyLocationListHandler ...
func (s *ServiceHandler) GetCompanyLocationListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetCompanyLocationList == nil {
		s.handleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetCompanyLocationListRequest

	req.DeptLoc = restlib.GetQueryParam(r, "deptLoc")
	var CompanyNameParam string

	var convErr error

	CompanyNameParam = restlib.GetQueryParam(r, "companyName")
	req.CompanyName, convErr = convert.StringToStringPtr(ctx, CompanyNameParam)
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

	conn, err := s.DB.Conn(ctx)
	if err != nil {
		s.handleError(ctx, w, common.InternalError, "Database connection could not be retrieved", err)
		return
	}

	defer conn.Close()
	retrievebycompanyandlocationStmt, err_retrievebycompanyandlocation := conn.PrepareContext(ctx, "select company.abnnumber, company.companyname, company.companycountry, department.deptid, department.deptname, department.deptloc from company JOIN department ON company.abnnumber=department.abn WHERE department.deptloc=? and company.companyname=? order by company.abnnumber;")
	if err_retrievebycompanyandlocation != nil {
		s.handleError(ctx, w, common.InternalError, "could not parse the sql query with the name sql_retrieveByCompanyAndLocation", err_retrievebycompanyandlocation)
		return
	}

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		s.handleError(ctx, w, common.DownstreamUnavailableError, "DB Transaction could not be created", err)
		return
	}

	client := GetCompanyLocationListClient{
		conn:                         conn,
		retrievebycompanyandlocation: retrievebycompanyandlocationStmt,
	}

	getcompanylocationresponse, err := s.serviceInterface.GetCompanyLocationList(ctx, &req, client)
	if err != nil {
		tx.Rollback()
		s.handleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		s.handleError(ctx, w, common.InternalError, "Failed to commit the transaction", commitErr)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, getcompanylocationresponse)
}
