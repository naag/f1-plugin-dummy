// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create user API
*/
func (a *CreateUserRequest) Do() (*CreateUserCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateUser",
		Method:             "POST",
		PathPattern:        "/security/users",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateUserReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserCreated), nil

}

func (a *CreateUserRequest) MustDo() *CreateUserCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create user credentials API
*/
func (a *CreateUserCredentialsRequest) Do() (*CreateUserCredentialsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateUserCredentials",
		Method:             "POST",
		PathPattern:        "/security/users/{user_id}/credentials",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateUserCredentialsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserCredentialsCreated), nil

}

func (a *CreateUserCredentialsRequest) MustDo() *CreateUserCredentialsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create user role API
*/
func (a *CreateUserRoleRequest) Do() (*CreateUserRoleCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateUserRole",
		Method:             "POST",
		PathPattern:        "/security/users/{user_id}/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateUserRoleReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserRoleCreated), nil

}

func (a *CreateUserRoleRequest) MustDo() *CreateUserRoleCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
delete user API
*/
func (a *DeleteUserRequest) Do() (*DeleteUserNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUser",
		Method:             "DELETE",
		PathPattern:        "/security/users/{user_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &DeleteUserReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserNoContent), nil

}

func (a *DeleteUserRequest) MustDo() *DeleteUserNoContent {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
delete user credential API
*/
func (a *DeleteUserCredentialRequest) Do() (*DeleteUserCredentialNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUserCredential",
		Method:             "DELETE",
		PathPattern:        "/security/users/{user_id}/credentials/{client_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &DeleteUserCredentialReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserCredentialNoContent), nil

}

func (a *DeleteUserCredentialRequest) MustDo() *DeleteUserCredentialNoContent {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
delete user role API
*/
func (a *DeleteUserRoleRequest) Do() (*DeleteUserRoleNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUserRole",
		Method:             "DELETE",
		PathPattern:        "/security/users/{user_id}/roles/{role_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &DeleteUserRoleReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserRoleNoContent), nil

}

func (a *DeleteUserRoleRequest) MustDo() *DeleteUserRoleNoContent {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get user API
*/
func (a *GetUserRequest) Do() (*GetUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUser",
		Method:             "GET",
		PathPattern:        "/security/users/{user_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetUserReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserOK), nil

}

func (a *GetUserRequest) MustDo() *GetUserOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get user aces API
*/
func (a *GetUserAcesRequest) Do() (*GetUserAcesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUserAces",
		Method:             "GET",
		PathPattern:        "/security/users/{user_id}/aces",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetUserAcesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserAcesOK), nil

}

func (a *GetUserAcesRequest) MustDo() *GetUserAcesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list user credentials API
*/
func (a *ListUserCredentialsRequest) Do() (*ListUserCredentialsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListUserCredentials",
		Method:             "GET",
		PathPattern:        "/security/users/{user_id}/credentials",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListUserCredentialsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUserCredentialsOK), nil

}

func (a *ListUserCredentialsRequest) MustDo() *ListUserCredentialsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list user roles API
*/
func (a *ListUserRolesRequest) Do() (*ListUserRolesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListUserRoles",
		Method:             "GET",
		PathPattern:        "/security/users/{user_id}/roles",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListUserRolesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUserRolesOK), nil

}

func (a *ListUserRolesRequest) MustDo() *ListUserRolesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list users API
*/
func (a *ListUsersRequest) Do() (*ListUsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListUsers",
		Method:             "GET",
		PathPattern:        "/security/users",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListUsersReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUsersOK), nil

}

func (a *ListUsersRequest) MustDo() *ListUsersOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
modify user API
*/
func (a *ModifyUserRequest) Do() (*ModifyUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyUser",
		Method:             "PATCH",
		PathPattern:        "/security/users/{user_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ModifyUserReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyUserOK), nil

}

func (a *ModifyUserRequest) MustDo() *ModifyUserOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
