// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package auth

import (
	"context"

	"github.com/trinet2005/oss-console/pkg/auth/idp/oauth2"
	"github.com/trinet2005/oss-go-sdk/pkg/credentials"
	xoauth2 "golang.org/x/oauth2"
)

// IdentityProviderI interface with all functions to be implemented
// by mock when testing, it should include all IdentityProvider respective api calls
// that are used within this project.
type IdentityProviderI interface {
	VerifyIdentity(ctx context.Context, code, state string) (*credentials.Credentials, error)
	VerifyIdentityForOperator(ctx context.Context, code, state string) (*xoauth2.Token, error)
	GenerateLoginURL() string
}

// Interface implementation
//
// Define the structure of a IdentityProvider with Client inside and define the functions that are used
// during the authentication flow.
type IdentityProvider struct {
	KeyFunc oauth2.StateKeyFunc
	Client  *oauth2.Provider
	RoleARN string
}

// VerifyIdentity will verify the user identity against the idp using the authorization code flow
func (c IdentityProvider) VerifyIdentity(ctx context.Context, code, state string) (*credentials.Credentials, error) {
	return c.Client.VerifyIdentity(ctx, code, state, c.RoleARN, c.KeyFunc)
}

// VerifyIdentityForOperator will verify the user identity against the idp using the authorization code flow
func (c IdentityProvider) VerifyIdentityForOperator(ctx context.Context, code, state string) (*xoauth2.Token, error) {
	return c.Client.VerifyIdentityForOperator(ctx, code, state, c.KeyFunc)
}

// GenerateLoginURL returns a new URL used by the user to login against the idp
func (c IdentityProvider) GenerateLoginURL() string {
	return c.Client.GenerateLoginURL(c.KeyFunc, c.Client.IDPName)
}
