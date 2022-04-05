/*
 * Todo Manager API
 *
 * This is todo manager backend for coding tests.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject struct {

	ActivationCode string `json:"activationCode,omitempty"`
}

// AssertInlineObjectRequired checks if the required fields are not zero-ed
func AssertInlineObjectRequired(obj InlineObject) error {
	return nil
}

// AssertRecurseInlineObjectRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of InlineObject (e.g. [][]InlineObject), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInlineObjectRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInlineObject, ok := obj.(InlineObject)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInlineObjectRequired(aInlineObject)
	})
}