/*
 * Todo Manager API
 *
 * This is todo manager backend for coding tests.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UsersForRes struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	EmailAddress string `json:"emailAddress,omitempty"`

	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// AssertUsersForResRequired checks if the required fields are not zero-ed
func AssertUsersForResRequired(obj UsersForRes) error {
	return nil
}

// AssertRecurseUsersForResRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UsersForRes (e.g. [][]UsersForRes), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUsersForResRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUsersForRes, ok := obj.(UsersForRes)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUsersForResRequired(aUsersForRes)
	})
}
