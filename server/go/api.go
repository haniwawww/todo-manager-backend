/*
 * Todo Manager API
 *
 * This is todo manager backend for coding tests.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)

// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface {
	UsersPost(http.ResponseWriter, *http.Request)
	UsersUserIdDelete(http.ResponseWriter, *http.Request)
	UsersUserIdGet(http.ResponseWriter, *http.Request)
	UsersUserIdPatch(http.ResponseWriter, *http.Request)
	UsersUserIdPost(http.ResponseWriter, *http.Request)
	UsersUserIdTasksGet(http.ResponseWriter, *http.Request)
	UsersUserIdTasksPost(http.ResponseWriter, *http.Request)
	UsersUserIdTasksTaskIdDelete(http.ResponseWriter, *http.Request)
	UsersUserIdTasksTaskIdGet(http.ResponseWriter, *http.Request)
	UsersUserIdTasksTaskIdPatch(http.ResponseWriter, *http.Request)
}

func UsersPost() {

}

// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface {
	UsersPost(context.Context, UsersForReq) (ImplResponse, error)
	UsersUserIdDelete(context.Context, string) (ImplResponse, error)
	UsersUserIdGet(context.Context, string) (ImplResponse, error)
	UsersUserIdPatch(context.Context, string, UsersForReq) (ImplResponse, error)
	UsersUserIdPost(context.Context, string, InlineObject) (ImplResponse, error)
	UsersUserIdTasksGet(context.Context, string) (ImplResponse, error)
	UsersUserIdTasksPost(context.Context, string, TasksForReq) (ImplResponse, error)
	UsersUserIdTasksTaskIdDelete(context.Context, string, string) (ImplResponse, error)
	UsersUserIdTasksTaskIdGet(context.Context, string, string) (ImplResponse, error)
	UsersUserIdTasksTaskIdPatch(context.Context, string, string) (ImplResponse, error)
}
