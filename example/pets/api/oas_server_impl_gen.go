// Code generated by entc, DO NOT EDIT.

package api

import (
	"context"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct{}

// CreateCategory handles POST /categories requests.
func (OgentHandler) CreateCategory(ctx context.Context, req CreateCategoryReq) (CreateCategoryRes, error) {
	panic("unimplemented")
}

// ReadCategory handles GET /categories/{id} requests.
func (OgentHandler) ReadCategory(ctx context.Context, params ReadCategoryParams) (ReadCategoryRes, error) {
	panic("unimplemented")
}

// UpdateCategory handles PATCH /categories/{id} requests.
func (OgentHandler) UpdateCategory(ctx context.Context, req UpdateCategoryReq, params UpdateCategoryParams) (UpdateCategoryRes, error) {
	panic("unimplemented")
}

// DeleteCategory handles DELETE /categories/{id} requests.
func (OgentHandler) DeleteCategory(ctx context.Context, params DeleteCategoryParams) (DeleteCategoryRes, error) {
	panic("unimplemented")
}

// ListCategory handles GET /categories requests.
func (OgentHandler) ListCategory(ctx context.Context, params ListCategoryParams) (ListCategoryRes, error) {
	panic("unimplemented")
}

// CreateCategoryPets handles POST /categories/{id}/pets requests.
func (OgentHandler) CreateCategoryPets(ctx context.Context, req CreateCategoryPetsReq, params CreateCategoryPetsParams) (CreateCategoryPetsRes, error) {
	panic("unimplemented")
}

// ListCategoryPets handles GET /categories/{id}/pets requests.
func (OgentHandler) ListCategoryPets(ctx context.Context, params ListCategoryPetsParams) (ListCategoryPetsRes, error) {
	panic("unimplemented")
}

// CreatePet handles POST /pets requests.
func (OgentHandler) CreatePet(ctx context.Context, req CreatePetReq) (CreatePetRes, error) {
	panic("unimplemented")
}

// ReadPet handles GET /pets/{id} requests.
func (OgentHandler) ReadPet(ctx context.Context, params ReadPetParams) (ReadPetRes, error) {
	panic("unimplemented")
}

// UpdatePet handles PATCH /pets/{id} requests.
func (OgentHandler) UpdatePet(ctx context.Context, req UpdatePetReq, params UpdatePetParams) (UpdatePetRes, error) {
	panic("unimplemented")
}

// DeletePet handles DELETE /pets/{id} requests.
func (OgentHandler) DeletePet(ctx context.Context, params DeletePetParams) (DeletePetRes, error) {
	panic("unimplemented")
}

// ListPet handles GET /pets requests.
func (OgentHandler) ListPet(ctx context.Context, params ListPetParams) (ListPetRes, error) {
	panic("unimplemented")
}

// CreatePetCategories handles POST /pets/{id}/categories requests.
func (OgentHandler) CreatePetCategories(ctx context.Context, req CreatePetCategoriesReq, params CreatePetCategoriesParams) (CreatePetCategoriesRes, error) {
	panic("unimplemented")
}

// ListPetCategories handles GET /pets/{id}/categories requests.
func (OgentHandler) ListPetCategories(ctx context.Context, params ListPetCategoriesParams) (ListPetCategoriesRes, error) {
	panic("unimplemented")
}

// CreatePetOwner handles POST /pets/{id}/owner requests.
func (OgentHandler) CreatePetOwner(ctx context.Context, req CreatePetOwnerReq, params CreatePetOwnerParams) (CreatePetOwnerRes, error) {
	panic("unimplemented")
}

// ReadPetOwner handles GET /pets/{id}/owner requests.
func (OgentHandler) ReadPetOwner(ctx context.Context, params ReadPetOwnerParams) (ReadPetOwnerRes, error) {
	panic("unimplemented")
}

// DeletePetOwner handles DELETE /pets/{id}/owner requests.
func (OgentHandler) DeletePetOwner(ctx context.Context, params DeletePetOwnerParams) (DeletePetOwnerRes, error) {
	panic("unimplemented")
}

// CreatePetFriends handles POST /pets/{id}/friends requests.
func (OgentHandler) CreatePetFriends(ctx context.Context, req CreatePetFriendsReq, params CreatePetFriendsParams) (CreatePetFriendsRes, error) {
	panic("unimplemented")
}

// ListPetFriends handles GET /pets/{id}/friends requests.
func (OgentHandler) ListPetFriends(ctx context.Context, params ListPetFriendsParams) (ListPetFriendsRes, error) {
	panic("unimplemented")
}

// CreateUser handles POST /users requests.
func (OgentHandler) CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error) {
	panic("unimplemented")
}

// ReadUser handles GET /users/{id} requests.
func (OgentHandler) ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error) {
	panic("unimplemented")
}

// UpdateUser handles PATCH /users/{id} requests.
func (OgentHandler) UpdateUser(ctx context.Context, req UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error) {
	panic("unimplemented")
}

// DeleteUser handles DELETE /users/{id} requests.
func (OgentHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error) {
	panic("unimplemented")
}

// ListUser handles GET /users requests.
func (OgentHandler) ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error) {
	panic("unimplemented")
}

// CreateUserPets handles POST /users/{id}/pets requests.
func (OgentHandler) CreateUserPets(ctx context.Context, req CreateUserPetsReq, params CreateUserPetsParams) (CreateUserPetsRes, error) {
	panic("unimplemented")
}

// ListUserPets handles GET /users/{id}/pets requests.
func (OgentHandler) ListUserPets(ctx context.Context, params ListUserPetsParams) (ListUserPetsRes, error) {
	panic("unimplemented")
}

var _ Handler = (*OgentHandler)(nil)
