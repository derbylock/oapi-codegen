// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v0.0.0-00010101000000-000000000000 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"github.com/oapi-codegen/runtime"
	strictiris "github.com/oapi-codegen/runtime/strictmiddleware/iris"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /json)
	JSONExample(ctx iris.Context)

	// (POST /multipart)
	MultipartExample(ctx iris.Context)

	// (POST /multiple)
	MultipleRequestAndResponseTypes(ctx iris.Context)

	// (GET /reserved-go-keyword-parameters/{type})
	ReservedGoKeywordParameters(ctx iris.Context, pType string)

	// (POST /reusable-responses)
	ReusableResponses(ctx iris.Context)

	// (POST /text)
	TextExample(ctx iris.Context)

	// (POST /unknown)
	UnknownExample(ctx iris.Context)

	// (POST /unspecified-content-type)
	UnspecifiedContentType(ctx iris.Context)

	// (POST /urlencoded)
	URLEncodedExample(ctx iris.Context)

	// (POST /with-headers)
	HeadersExample(ctx iris.Context, params HeadersExampleParams)

	// (POST /with-union)
	UnionExample(ctx iris.Context)
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc iris.Handler

// JSONExample converts iris context to params.
func (w *ServerInterfaceWrapper) JSONExample(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.JSONExample(ctx)
}

// MultipartExample converts iris context to params.
func (w *ServerInterfaceWrapper) MultipartExample(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.MultipartExample(ctx)
}

// MultipleRequestAndResponseTypes converts iris context to params.
func (w *ServerInterfaceWrapper) MultipleRequestAndResponseTypes(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.MultipleRequestAndResponseTypes(ctx)
}

// ReservedGoKeywordParameters converts iris context to params.
func (w *ServerInterfaceWrapper) ReservedGoKeywordParameters(ctx iris.Context) {

	var err error

	// ------------- Path parameter "type" -------------
	var pType string

	err = runtime.BindStyledParameterWithLocation("simple", false, "type", runtime.ParamLocationPath, ctx.Params().Get("type"), &pType)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter type: %s", err)
		return
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.ReservedGoKeywordParameters(ctx, pType)
}

// ReusableResponses converts iris context to params.
func (w *ServerInterfaceWrapper) ReusableResponses(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.ReusableResponses(ctx)
}

// TextExample converts iris context to params.
func (w *ServerInterfaceWrapper) TextExample(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.TextExample(ctx)
}

// UnknownExample converts iris context to params.
func (w *ServerInterfaceWrapper) UnknownExample(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.UnknownExample(ctx)
}

// UnspecifiedContentType converts iris context to params.
func (w *ServerInterfaceWrapper) UnspecifiedContentType(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.UnspecifiedContentType(ctx)
}

// URLEncodedExample converts iris context to params.
func (w *ServerInterfaceWrapper) URLEncodedExample(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.URLEncodedExample(ctx)
}

// HeadersExample converts iris context to params.
func (w *ServerInterfaceWrapper) HeadersExample(ctx iris.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params HeadersExampleParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "header1" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header1")]; found {
		var Header1 string
		n := len(valueList)
		if n != 1 {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.Writef("Expected one value for header1, got %d", n)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "header1", runtime.ParamLocationHeader, valueList[0], &Header1)
		if err != nil {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.Writef("Invalid format for parameter header1: %s", err)
			return
		}

		params.Header1 = Header1
	} else {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("Header header1 is required, but not found")
		return
	}
	// ------------- Optional header parameter "header2" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header2")]; found {
		var Header2 int
		n := len(valueList)
		if n != 1 {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.Writef("Expected one value for header2, got %d", n)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "header2", runtime.ParamLocationHeader, valueList[0], &Header2)
		if err != nil {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.Writef("Invalid format for parameter header2: %s", err)
			return
		}

		params.Header2 = &Header2
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.HeadersExample(ctx, params)
}

// UnionExample converts iris context to params.
func (w *ServerInterfaceWrapper) UnionExample(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.UnionExample(ctx)
}

// IrisServerOption is the option for iris server
type IrisServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *iris.Application, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, IrisServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *iris.Application, si ServerInterface, options IrisServerOptions) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.Post(options.BaseURL+"/json", wrapper.JSONExample)
	router.Post(options.BaseURL+"/multipart", wrapper.MultipartExample)
	router.Post(options.BaseURL+"/multiple", wrapper.MultipleRequestAndResponseTypes)
	router.Get(options.BaseURL+"/reserved-go-keyword-parameters/:type", wrapper.ReservedGoKeywordParameters)
	router.Post(options.BaseURL+"/reusable-responses", wrapper.ReusableResponses)
	router.Post(options.BaseURL+"/text", wrapper.TextExample)
	router.Post(options.BaseURL+"/unknown", wrapper.UnknownExample)
	router.Post(options.BaseURL+"/unspecified-content-type", wrapper.UnspecifiedContentType)
	router.Post(options.BaseURL+"/urlencoded", wrapper.URLEncodedExample)
	router.Post(options.BaseURL+"/with-headers", wrapper.HeadersExample)
	router.Post(options.BaseURL+"/with-union", wrapper.UnionExample)

	router.Build()
}

type BadrequestResponse struct {
}

type ReusableresponseResponseHeaders struct {
	Header1 string
	Header2 int
}
type ReusableresponseJSONResponse struct {
	Body Example

	Headers ReusableresponseResponseHeaders
}

type JSONExampleRequestObject struct {
	Body *JSONExampleJSONRequestBody
}

type JSONExampleResponseObject interface {
	VisitJSONExampleResponse(ctx iris.Context) error
}

type JSONExample200JSONResponse Example

func (response JSONExample200JSONResponse) VisitJSONExampleResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	ctx.StatusCode(200)

	return ctx.JSON(&response)
}

type JSONExample400Response = BadrequestResponse

func (response JSONExample400Response) VisitJSONExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type JSONExampledefaultResponse struct {
	StatusCode int
}

func (response JSONExampledefaultResponse) VisitJSONExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type MultipartExampleRequestObject struct {
	Body *multipart.Reader
}

type MultipartExampleResponseObject interface {
	VisitMultipartExampleResponse(ctx iris.Context) error
}

type MultipartExample200MultipartResponse func(writer *multipart.Writer) error

func (response MultipartExample200MultipartResponse) VisitMultipartExampleResponse(ctx iris.Context) error {
	writer := multipart.NewWriter(ctx.ResponseWriter())
	ctx.ResponseWriter().Header().Set("Content-Type", writer.FormDataContentType())
	ctx.StatusCode(200)

	defer writer.Close()
	return response(writer)
}

type MultipartExample400Response = BadrequestResponse

func (response MultipartExample400Response) VisitMultipartExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type MultipartExampledefaultResponse struct {
	StatusCode int
}

func (response MultipartExampledefaultResponse) VisitMultipartExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type MultipleRequestAndResponseTypesRequestObject struct {
	JSONBody      *MultipleRequestAndResponseTypesJSONRequestBody
	FormdataBody  *MultipleRequestAndResponseTypesFormdataRequestBody
	Body          io.Reader
	MultipartBody *multipart.Reader
	TextBody      *MultipleRequestAndResponseTypesTextRequestBody
}

type MultipleRequestAndResponseTypesResponseObject interface {
	VisitMultipleRequestAndResponseTypesResponse(ctx iris.Context) error
}

type MultipleRequestAndResponseTypes200JSONResponse Example

func (response MultipleRequestAndResponseTypes200JSONResponse) VisitMultipleRequestAndResponseTypesResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	ctx.StatusCode(200)

	return ctx.JSON(&response)
}

type MultipleRequestAndResponseTypes200FormdataResponse Example

func (response MultipleRequestAndResponseTypes200FormdataResponse) VisitMultipleRequestAndResponseTypesResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "application/x-www-form-urlencoded")
	ctx.StatusCode(200)

	if form, err := runtime.MarshalForm(response, nil); err != nil {
		return err
	} else {
		_, err := ctx.WriteString(form.Encode())
		return err
	}
}

type MultipleRequestAndResponseTypes200ImagePngResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response MultipleRequestAndResponseTypes200ImagePngResponse) VisitMultipleRequestAndResponseTypesResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "image/png")
	if response.ContentLength != 0 {
		ctx.ResponseWriter().Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	ctx.StatusCode(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(ctx.ResponseWriter(), response.Body)
	return err
}

type MultipleRequestAndResponseTypes200MultipartResponse func(writer *multipart.Writer) error

func (response MultipleRequestAndResponseTypes200MultipartResponse) VisitMultipleRequestAndResponseTypesResponse(ctx iris.Context) error {
	writer := multipart.NewWriter(ctx.ResponseWriter())
	ctx.ResponseWriter().Header().Set("Content-Type", writer.FormDataContentType())
	ctx.StatusCode(200)

	defer writer.Close()
	return response(writer)
}

type MultipleRequestAndResponseTypes200TextResponse string

func (response MultipleRequestAndResponseTypes200TextResponse) VisitMultipleRequestAndResponseTypesResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "text/plain")
	ctx.StatusCode(200)

	_, err := ctx.WriteString(string(response))
	return err
}

type MultipleRequestAndResponseTypes400Response = BadrequestResponse

func (response MultipleRequestAndResponseTypes400Response) VisitMultipleRequestAndResponseTypesResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type ReservedGoKeywordParametersRequestObject struct {
	Type string `json:"type"`
}

type ReservedGoKeywordParametersResponseObject interface {
	VisitReservedGoKeywordParametersResponse(ctx iris.Context) error
}

type ReservedGoKeywordParameters200TextResponse string

func (response ReservedGoKeywordParameters200TextResponse) VisitReservedGoKeywordParametersResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "text/plain")
	ctx.StatusCode(200)

	_, err := ctx.WriteString(string(response))
	return err
}

type ReusableResponsesRequestObject struct {
	Body *ReusableResponsesJSONRequestBody
}

type ReusableResponsesResponseObject interface {
	VisitReusableResponsesResponse(ctx iris.Context) error
}

type ReusableResponses200JSONResponse struct{ ReusableresponseJSONResponse }

func (response ReusableResponses200JSONResponse) VisitReusableResponsesResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.ResponseWriter().Header().Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	ctx.StatusCode(200)

	return ctx.JSON(&response.Body)
}

type ReusableResponses400Response = BadrequestResponse

func (response ReusableResponses400Response) VisitReusableResponsesResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type ReusableResponsesdefaultResponse struct {
	StatusCode int
}

func (response ReusableResponsesdefaultResponse) VisitReusableResponsesResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type TextExampleRequestObject struct {
	Body *TextExampleTextRequestBody
}

type TextExampleResponseObject interface {
	VisitTextExampleResponse(ctx iris.Context) error
}

type TextExample200TextResponse string

func (response TextExample200TextResponse) VisitTextExampleResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "text/plain")
	ctx.StatusCode(200)

	_, err := ctx.WriteString(string(response))
	return err
}

type TextExample400Response = BadrequestResponse

func (response TextExample400Response) VisitTextExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type TextExampledefaultResponse struct {
	StatusCode int
}

func (response TextExampledefaultResponse) VisitTextExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type UnknownExampleRequestObject struct {
	Body io.Reader
}

type UnknownExampleResponseObject interface {
	VisitUnknownExampleResponse(ctx iris.Context) error
}

type UnknownExample200VideoMp4Response struct {
	Body          io.Reader
	ContentLength int64
}

func (response UnknownExample200VideoMp4Response) VisitUnknownExampleResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "video/mp4")
	if response.ContentLength != 0 {
		ctx.ResponseWriter().Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	ctx.StatusCode(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(ctx.ResponseWriter(), response.Body)
	return err
}

type UnknownExample400Response = BadrequestResponse

func (response UnknownExample400Response) VisitUnknownExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type UnknownExampledefaultResponse struct {
	StatusCode int
}

func (response UnknownExampledefaultResponse) VisitUnknownExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type UnspecifiedContentTypeRequestObject struct {
	ContentType string
	Body        io.Reader
}

type UnspecifiedContentTypeResponseObject interface {
	VisitUnspecifiedContentTypeResponse(ctx iris.Context) error
}

type UnspecifiedContentType200VideoResponse struct {
	Body          io.Reader
	ContentType   string
	ContentLength int64
}

func (response UnspecifiedContentType200VideoResponse) VisitUnspecifiedContentTypeResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", response.ContentType)
	if response.ContentLength != 0 {
		ctx.ResponseWriter().Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	ctx.StatusCode(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(ctx.ResponseWriter(), response.Body)
	return err
}

type UnspecifiedContentType400Response = BadrequestResponse

func (response UnspecifiedContentType400Response) VisitUnspecifiedContentTypeResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type UnspecifiedContentType401Response struct {
}

func (response UnspecifiedContentType401Response) VisitUnspecifiedContentTypeResponse(ctx iris.Context) error {
	ctx.StatusCode(401)
	return nil
}

type UnspecifiedContentType403Response struct {
}

func (response UnspecifiedContentType403Response) VisitUnspecifiedContentTypeResponse(ctx iris.Context) error {
	ctx.StatusCode(403)
	return nil
}

type UnspecifiedContentTypedefaultResponse struct {
	StatusCode int
}

func (response UnspecifiedContentTypedefaultResponse) VisitUnspecifiedContentTypeResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type URLEncodedExampleRequestObject struct {
	Body *URLEncodedExampleFormdataRequestBody
}

type URLEncodedExampleResponseObject interface {
	VisitURLEncodedExampleResponse(ctx iris.Context) error
}

type URLEncodedExample200FormdataResponse Example

func (response URLEncodedExample200FormdataResponse) VisitURLEncodedExampleResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("Content-Type", "application/x-www-form-urlencoded")
	ctx.StatusCode(200)

	if form, err := runtime.MarshalForm(response, nil); err != nil {
		return err
	} else {
		_, err := ctx.WriteString(form.Encode())
		return err
	}
}

type URLEncodedExample400Response = BadrequestResponse

func (response URLEncodedExample400Response) VisitURLEncodedExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type URLEncodedExampledefaultResponse struct {
	StatusCode int
}

func (response URLEncodedExampledefaultResponse) VisitURLEncodedExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type HeadersExampleRequestObject struct {
	Params HeadersExampleParams
	Body   *HeadersExampleJSONRequestBody
}

type HeadersExampleResponseObject interface {
	VisitHeadersExampleResponse(ctx iris.Context) error
}

type HeadersExample200ResponseHeaders struct {
	Header1 string
	Header2 int
}

type HeadersExample200JSONResponse struct {
	Body    Example
	Headers HeadersExample200ResponseHeaders
}

func (response HeadersExample200JSONResponse) VisitHeadersExampleResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.ResponseWriter().Header().Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	ctx.StatusCode(200)

	return ctx.JSON(&response.Body)
}

type HeadersExample400Response = BadrequestResponse

func (response HeadersExample400Response) VisitHeadersExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type HeadersExampledefaultResponse struct {
	StatusCode int
}

func (response HeadersExampledefaultResponse) VisitHeadersExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

type UnionExampleRequestObject struct {
	Body *UnionExampleJSONRequestBody
}

type UnionExampleResponseObject interface {
	VisitUnionExampleResponse(ctx iris.Context) error
}

type UnionExample200ResponseHeaders struct {
	Header1 string
	Header2 int
}

type UnionExample200ApplicationAlternativePlusJSONResponse struct {
	Body    Example
	Headers UnionExample200ResponseHeaders
}

func (response UnionExample200ApplicationAlternativePlusJSONResponse) VisitUnionExampleResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.ResponseWriter().Header().Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.ResponseWriter().Header().Set("Content-Type", "application/alternative+json")
	ctx.StatusCode(200)

	return ctx.JSON(&response.Body)
}

type UnionExample200JSONResponse struct {
	Body struct {
		union json.RawMessage
	}
	Headers UnionExample200ResponseHeaders
}

func (response UnionExample200JSONResponse) VisitUnionExampleResponse(ctx iris.Context) error {
	ctx.ResponseWriter().Header().Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.ResponseWriter().Header().Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	ctx.StatusCode(200)

	return ctx.JSON(&response.Body.union)
}

type UnionExample400Response = BadrequestResponse

func (response UnionExample400Response) VisitUnionExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(400)
	return nil
}

type UnionExampledefaultResponse struct {
	StatusCode int
}

func (response UnionExampledefaultResponse) VisitUnionExampleResponse(ctx iris.Context) error {
	ctx.StatusCode(response.StatusCode)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /json)
	JSONExample(ctx context.Context, request JSONExampleRequestObject) (JSONExampleResponseObject, error)

	// (POST /multipart)
	MultipartExample(ctx context.Context, request MultipartExampleRequestObject) (MultipartExampleResponseObject, error)

	// (POST /multiple)
	MultipleRequestAndResponseTypes(ctx context.Context, request MultipleRequestAndResponseTypesRequestObject) (MultipleRequestAndResponseTypesResponseObject, error)

	// (GET /reserved-go-keyword-parameters/{type})
	ReservedGoKeywordParameters(ctx context.Context, request ReservedGoKeywordParametersRequestObject) (ReservedGoKeywordParametersResponseObject, error)

	// (POST /reusable-responses)
	ReusableResponses(ctx context.Context, request ReusableResponsesRequestObject) (ReusableResponsesResponseObject, error)

	// (POST /text)
	TextExample(ctx context.Context, request TextExampleRequestObject) (TextExampleResponseObject, error)

	// (POST /unknown)
	UnknownExample(ctx context.Context, request UnknownExampleRequestObject) (UnknownExampleResponseObject, error)

	// (POST /unspecified-content-type)
	UnspecifiedContentType(ctx context.Context, request UnspecifiedContentTypeRequestObject) (UnspecifiedContentTypeResponseObject, error)

	// (POST /urlencoded)
	URLEncodedExample(ctx context.Context, request URLEncodedExampleRequestObject) (URLEncodedExampleResponseObject, error)

	// (POST /with-headers)
	HeadersExample(ctx context.Context, request HeadersExampleRequestObject) (HeadersExampleResponseObject, error)

	// (POST /with-union)
	UnionExample(ctx context.Context, request UnionExampleRequestObject) (UnionExampleResponseObject, error)
}

type StrictHandlerFunc = strictiris.StrictIrisHandlerFunc
type StrictMiddlewareFunc = strictiris.StrictIrisMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// JSONExample operation middleware
func (sh *strictHandler) JSONExample(ctx iris.Context) {
	var request JSONExampleRequestObject

	var body JSONExampleJSONRequestBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}
	request.Body = &body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.JSONExample(ctx, request.(JSONExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "JSONExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(JSONExampleResponseObject); ok {
		if err := validResponse.VisitJSONExampleResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// MultipartExample operation middleware
func (sh *strictHandler) MultipartExample(ctx iris.Context) {
	var request MultipartExampleRequestObject

	if reader, err := ctx.Request().MultipartReader(); err == nil {
		request.Body = reader
	} else {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.MultipartExample(ctx, request.(MultipartExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipartExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(MultipartExampleResponseObject); ok {
		if err := validResponse.VisitMultipartExampleResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// MultipleRequestAndResponseTypes operation middleware
func (sh *strictHandler) MultipleRequestAndResponseTypes(ctx iris.Context) {
	var request MultipleRequestAndResponseTypesRequestObject

	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/json") {

		var body MultipleRequestAndResponseTypesJSONRequestBody
		if err := ctx.ReadJSON(&body); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
		request.JSONBody = &body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/x-www-form-urlencoded") {
		if err := ctx.Request().ParseForm(); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
		var body MultipleRequestAndResponseTypesFormdataRequestBody
		if err := runtime.BindForm(&body, ctx.Request().Form, nil, nil); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
		request.FormdataBody = &body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "image/png") {
		request.Body = ctx.Request().Body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "multipart/form-data") {
		if reader, err := ctx.Request().MultipartReader(); err == nil {
			request.MultipartBody = reader
		} else {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "text/plain") {
		data, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
		body := MultipleRequestAndResponseTypesTextRequestBody(data)
		request.TextBody = &body
	}

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.MultipleRequestAndResponseTypes(ctx, request.(MultipleRequestAndResponseTypesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipleRequestAndResponseTypes")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(MultipleRequestAndResponseTypesResponseObject); ok {
		if err := validResponse.VisitMultipleRequestAndResponseTypesResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// ReservedGoKeywordParameters operation middleware
func (sh *strictHandler) ReservedGoKeywordParameters(ctx iris.Context, pType string) {
	var request ReservedGoKeywordParametersRequestObject

	request.Type = pType

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReservedGoKeywordParameters(ctx, request.(ReservedGoKeywordParametersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReservedGoKeywordParameters")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(ReservedGoKeywordParametersResponseObject); ok {
		if err := validResponse.VisitReservedGoKeywordParametersResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// ReusableResponses operation middleware
func (sh *strictHandler) ReusableResponses(ctx iris.Context) {
	var request ReusableResponsesRequestObject

	var body ReusableResponsesJSONRequestBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}
	request.Body = &body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReusableResponses(ctx, request.(ReusableResponsesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReusableResponses")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(ReusableResponsesResponseObject); ok {
		if err := validResponse.VisitReusableResponsesResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// TextExample operation middleware
func (sh *strictHandler) TextExample(ctx iris.Context) {
	var request TextExampleRequestObject

	data, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}
	body := TextExampleTextRequestBody(data)
	request.Body = &body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TextExample(ctx, request.(TextExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TextExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(TextExampleResponseObject); ok {
		if err := validResponse.VisitTextExampleResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// UnknownExample operation middleware
func (sh *strictHandler) UnknownExample(ctx iris.Context) {
	var request UnknownExampleRequestObject

	request.Body = ctx.Request().Body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UnknownExample(ctx, request.(UnknownExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnknownExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(UnknownExampleResponseObject); ok {
		if err := validResponse.VisitUnknownExampleResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// UnspecifiedContentType operation middleware
func (sh *strictHandler) UnspecifiedContentType(ctx iris.Context) {
	var request UnspecifiedContentTypeRequestObject

	request.ContentType = ctx.GetContentTypeRequested()

	request.Body = ctx.Request().Body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UnspecifiedContentType(ctx, request.(UnspecifiedContentTypeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnspecifiedContentType")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(UnspecifiedContentTypeResponseObject); ok {
		if err := validResponse.VisitUnspecifiedContentTypeResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// URLEncodedExample operation middleware
func (sh *strictHandler) URLEncodedExample(ctx iris.Context) {
	var request URLEncodedExampleRequestObject

	if err := ctx.Request().ParseForm(); err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}
	var body URLEncodedExampleFormdataRequestBody
	if err := runtime.BindForm(&body, ctx.Request().Form, nil, nil); err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}
	request.Body = &body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.URLEncodedExample(ctx, request.(URLEncodedExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "URLEncodedExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(URLEncodedExampleResponseObject); ok {
		if err := validResponse.VisitURLEncodedExampleResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// HeadersExample operation middleware
func (sh *strictHandler) HeadersExample(ctx iris.Context, params HeadersExampleParams) {
	var request HeadersExampleRequestObject

	request.Params = params

	var body HeadersExampleJSONRequestBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}
	request.Body = &body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.HeadersExample(ctx, request.(HeadersExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HeadersExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(HeadersExampleResponseObject); ok {
		if err := validResponse.VisitHeadersExampleResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// UnionExample operation middleware
func (sh *strictHandler) UnionExample(ctx iris.Context) {
	var request UnionExampleRequestObject

	var body UnionExampleJSONRequestBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	}
	request.Body = &body

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UnionExample(ctx, request.(UnionExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnionExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(UnionExampleResponseObject); ok {
		if err := validResponse.VisitUnionExampleResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYS2/jNhD+KwTb01aynGxOunWDxbbdtimc5FTkQIsjm7sSyQ5HVgzD/72gKL9ixbW3",
	"fhRBb3oMvxl+8+BwZjwzpTUaNDmezjiCs0Y7aF6GQiL8VYEj/ybBZagsKaN5yj8IOWj/zSOOUDkxLGCx",
	"3MtnRhPoZqmwtlCZ8EuTL86vn3GXjaEU/ul7hJyn/LtkZUoS/roEnkVpC+Dz+Tx6YcHdZx7xMQgJ2Fgb",
	"Hq82sWlqgafcESo94h4kiF13iilNMAL02rxoa4QXWNiRzrhFYwFJBY4moqigW1P7xQy/QEZhB0rnZpvL",
	"W6NJKO2YVHkOCJpYSx7zGI65ylqDBJINp8xryIg5wAkgjzgp8obx+/XvrDXY8YhPAF1QdNXr9/reX8aC",
	"FlbxlL9vPkXcCho3G1o6yJouv/9yf/c7U46JikwpSGWiKKasFOjGogDJlCbjTawycj3eaMLG8T/LdvXH",
	"lkofNU0AfTByeoqAaeJyLZyv+/0zxeU84jdBWRfG0qhkLcEamFxURQfnj/qrNrVmgGiw3VlSVgUpK5DW",
	"fbXJ9m8LkX0oX+IlucEyloLEiVg/lqaLEt/Wgs4cuR+b2rGxqRkZJkEUrFY0ZouFL5JbaSaYU3pUAFsY",
	"FXV6soC25P6o5aDdy4PHOHkuRRsoz3Fd13HjvAoL0JmRIL8NVpViBInVo83lHlsQT/lwSj5st4vrkYIo",
	"4gTPlNhCKL375DhTOfmf6aMldkhXhOZElPHIxF9hWhuUsRUoSiBAl8y89rkHHkFHKv+xlGSZ0GwITIsS",
	"JBM5AbJPhrWQbitlB63eT+ZzEFlBNcft8iX9c8Y9Jc0RzCPuFfA0sBLyWqF3OmEF0Q7anv4xPv+VAxZs",
	"hkYv3lDVXQYXJWpJHULufEns8lwHf0HTYE3iMg3D7ojban3PcQZ5T75+7j/A815H/hFL37lz+1DCqvDx",
	"dc7aVfvQ9o2VdA8WJ0qCSUp7cyDyxUh1FjKVK5Bxu4s42PZaSbg1OkOgzRbIXye0IbYE87ccGgMLDETM",
	"GVYDKytHzArnmKKmihQq3JQkbBWPx5Vlt0HTw6qc7vLquxP59N2lPHrTvzp8yfsTx81GK/NKPg5+/Rhk",
	"Dr0vHq1nOrDjO57eC6Wzv6TEawOV7hT+KQiszvQM1MR3RFoyBKpQg2QTJRZDgK3cbAFWbu3qhYIZq25o",
	"Mdw5pCGKdmJd82jXAOjpDY8nTjk2O1ecVlrtGlM9+t+s7aFfng3K6P/oEEoUBKgFqQn8cJwb5DaK0XCX",
	"N5n2wsvRnhqe3l5UzSMe5qahBFVY+DpBZNMkCfPWnqvFaATYUyYRVnkW/g4AAP//Pk3lbjwXAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
