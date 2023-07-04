// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/sourcenetwork/defradb/client"

	datastore "github.com/sourcenetwork/defradb/datastore"

	mock "github.com/stretchr/testify/mock"
)

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

type Collection_Expecter struct {
	mock *mock.Mock
}

func (_m *Collection) EXPECT() *Collection_Expecter {
	return &Collection_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Collection) Create(_a0 context.Context, _a1 *client.Document) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.Document) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Collection_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Collection_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *client.Document
func (_e *Collection_Expecter) Create(_a0 interface{}, _a1 interface{}) *Collection_Create_Call {
	return &Collection_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *Collection_Create_Call) Run(run func(_a0 context.Context, _a1 *client.Document)) *Collection_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*client.Document))
	})
	return _c
}

func (_c *Collection_Create_Call) Return(_a0 error) *Collection_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_Create_Call) RunAndReturn(run func(context.Context, *client.Document) error) *Collection_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateIndex provides a mock function with given fields: _a0, _a1
func (_m *Collection) CreateIndex(_a0 context.Context, _a1 client.IndexDescription) (client.IndexDescription, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.IndexDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.IndexDescription) (client.IndexDescription, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.IndexDescription) client.IndexDescription); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.IndexDescription)
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.IndexDescription) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_CreateIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateIndex'
type Collection_CreateIndex_Call struct {
	*mock.Call
}

// CreateIndex is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.IndexDescription
func (_e *Collection_Expecter) CreateIndex(_a0 interface{}, _a1 interface{}) *Collection_CreateIndex_Call {
	return &Collection_CreateIndex_Call{Call: _e.mock.On("CreateIndex", _a0, _a1)}
}

func (_c *Collection_CreateIndex_Call) Run(run func(_a0 context.Context, _a1 client.IndexDescription)) *Collection_CreateIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.IndexDescription))
	})
	return _c
}

func (_c *Collection_CreateIndex_Call) Return(_a0 client.IndexDescription, _a1 error) *Collection_CreateIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_CreateIndex_Call) RunAndReturn(run func(context.Context, client.IndexDescription) (client.IndexDescription, error)) *Collection_CreateIndex_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMany provides a mock function with given fields: _a0, _a1
func (_m *Collection) CreateMany(_a0 context.Context, _a1 []*client.Document) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*client.Document) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Collection_CreateMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMany'
type Collection_CreateMany_Call struct {
	*mock.Call
}

// CreateMany is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []*client.Document
func (_e *Collection_Expecter) CreateMany(_a0 interface{}, _a1 interface{}) *Collection_CreateMany_Call {
	return &Collection_CreateMany_Call{Call: _e.mock.On("CreateMany", _a0, _a1)}
}

func (_c *Collection_CreateMany_Call) Run(run func(_a0 context.Context, _a1 []*client.Document)) *Collection_CreateMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*client.Document))
	})
	return _c
}

func (_c *Collection_CreateMany_Call) Return(_a0 error) *Collection_CreateMany_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_CreateMany_Call) RunAndReturn(run func(context.Context, []*client.Document) error) *Collection_CreateMany_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *Collection) Delete(_a0 context.Context, _a1 client.DocKey) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.DocKey) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Collection_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.DocKey
func (_e *Collection_Expecter) Delete(_a0 interface{}, _a1 interface{}) *Collection_Delete_Call {
	return &Collection_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *Collection_Delete_Call) Run(run func(_a0 context.Context, _a1 client.DocKey)) *Collection_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.DocKey))
	})
	return _c
}

func (_c *Collection_Delete_Call) Return(_a0 bool, _a1 error) *Collection_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_Delete_Call) RunAndReturn(run func(context.Context, client.DocKey) (bool, error)) *Collection_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteWith provides a mock function with given fields: ctx, target
func (_m *Collection) DeleteWith(ctx context.Context, target interface{}) (*client.DeleteResult, error) {
	ret := _m.Called(ctx, target)

	var r0 *client.DeleteResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (*client.DeleteResult, error)); ok {
		return rf(ctx, target)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *client.DeleteResult); ok {
		r0 = rf(ctx, target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DeleteResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_DeleteWith_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWith'
type Collection_DeleteWith_Call struct {
	*mock.Call
}

// DeleteWith is a helper method to define mock.On call
//   - ctx context.Context
//   - target interface{}
func (_e *Collection_Expecter) DeleteWith(ctx interface{}, target interface{}) *Collection_DeleteWith_Call {
	return &Collection_DeleteWith_Call{Call: _e.mock.On("DeleteWith", ctx, target)}
}

func (_c *Collection_DeleteWith_Call) Run(run func(ctx context.Context, target interface{})) *Collection_DeleteWith_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *Collection_DeleteWith_Call) Return(_a0 *client.DeleteResult, _a1 error) *Collection_DeleteWith_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_DeleteWith_Call) RunAndReturn(run func(context.Context, interface{}) (*client.DeleteResult, error)) *Collection_DeleteWith_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteWithFilter provides a mock function with given fields: ctx, filter
func (_m *Collection) DeleteWithFilter(ctx context.Context, filter interface{}) (*client.DeleteResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 *client.DeleteResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (*client.DeleteResult, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *client.DeleteResult); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DeleteResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_DeleteWithFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWithFilter'
type Collection_DeleteWithFilter_Call struct {
	*mock.Call
}

// DeleteWithFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - filter interface{}
func (_e *Collection_Expecter) DeleteWithFilter(ctx interface{}, filter interface{}) *Collection_DeleteWithFilter_Call {
	return &Collection_DeleteWithFilter_Call{Call: _e.mock.On("DeleteWithFilter", ctx, filter)}
}

func (_c *Collection_DeleteWithFilter_Call) Run(run func(ctx context.Context, filter interface{})) *Collection_DeleteWithFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *Collection_DeleteWithFilter_Call) Return(_a0 *client.DeleteResult, _a1 error) *Collection_DeleteWithFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_DeleteWithFilter_Call) RunAndReturn(run func(context.Context, interface{}) (*client.DeleteResult, error)) *Collection_DeleteWithFilter_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteWithKey provides a mock function with given fields: _a0, _a1
func (_m *Collection) DeleteWithKey(_a0 context.Context, _a1 client.DocKey) (*client.DeleteResult, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *client.DeleteResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey) (*client.DeleteResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey) *client.DeleteResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DeleteResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.DocKey) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_DeleteWithKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWithKey'
type Collection_DeleteWithKey_Call struct {
	*mock.Call
}

// DeleteWithKey is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.DocKey
func (_e *Collection_Expecter) DeleteWithKey(_a0 interface{}, _a1 interface{}) *Collection_DeleteWithKey_Call {
	return &Collection_DeleteWithKey_Call{Call: _e.mock.On("DeleteWithKey", _a0, _a1)}
}

func (_c *Collection_DeleteWithKey_Call) Run(run func(_a0 context.Context, _a1 client.DocKey)) *Collection_DeleteWithKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.DocKey))
	})
	return _c
}

func (_c *Collection_DeleteWithKey_Call) Return(_a0 *client.DeleteResult, _a1 error) *Collection_DeleteWithKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_DeleteWithKey_Call) RunAndReturn(run func(context.Context, client.DocKey) (*client.DeleteResult, error)) *Collection_DeleteWithKey_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteWithKeys provides a mock function with given fields: _a0, _a1
func (_m *Collection) DeleteWithKeys(_a0 context.Context, _a1 []client.DocKey) (*client.DeleteResult, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *client.DeleteResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []client.DocKey) (*client.DeleteResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []client.DocKey) *client.DeleteResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DeleteResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []client.DocKey) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_DeleteWithKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWithKeys'
type Collection_DeleteWithKeys_Call struct {
	*mock.Call
}

// DeleteWithKeys is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []client.DocKey
func (_e *Collection_Expecter) DeleteWithKeys(_a0 interface{}, _a1 interface{}) *Collection_DeleteWithKeys_Call {
	return &Collection_DeleteWithKeys_Call{Call: _e.mock.On("DeleteWithKeys", _a0, _a1)}
}

func (_c *Collection_DeleteWithKeys_Call) Run(run func(_a0 context.Context, _a1 []client.DocKey)) *Collection_DeleteWithKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]client.DocKey))
	})
	return _c
}

func (_c *Collection_DeleteWithKeys_Call) Return(_a0 *client.DeleteResult, _a1 error) *Collection_DeleteWithKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_DeleteWithKeys_Call) RunAndReturn(run func(context.Context, []client.DocKey) (*client.DeleteResult, error)) *Collection_DeleteWithKeys_Call {
	_c.Call.Return(run)
	return _c
}

// Description provides a mock function with given fields:
func (_m *Collection) Description() client.CollectionDescription {
	ret := _m.Called()

	var r0 client.CollectionDescription
	if rf, ok := ret.Get(0).(func() client.CollectionDescription); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(client.CollectionDescription)
	}

	return r0
}

// Collection_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type Collection_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *Collection_Expecter) Description() *Collection_Description_Call {
	return &Collection_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *Collection_Description_Call) Run(run func()) *Collection_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Collection_Description_Call) Return(_a0 client.CollectionDescription) *Collection_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_Description_Call) RunAndReturn(run func() client.CollectionDescription) *Collection_Description_Call {
	_c.Call.Return(run)
	return _c
}

// DropIndex provides a mock function with given fields: ctx, indexName
func (_m *Collection) DropIndex(ctx context.Context, indexName string) error {
	ret := _m.Called(ctx, indexName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, indexName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Collection_DropIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropIndex'
type Collection_DropIndex_Call struct {
	*mock.Call
}

// DropIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - indexName string
func (_e *Collection_Expecter) DropIndex(ctx interface{}, indexName interface{}) *Collection_DropIndex_Call {
	return &Collection_DropIndex_Call{Call: _e.mock.On("DropIndex", ctx, indexName)}
}

func (_c *Collection_DropIndex_Call) Run(run func(ctx context.Context, indexName string)) *Collection_DropIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Collection_DropIndex_Call) Return(_a0 error) *Collection_DropIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_DropIndex_Call) RunAndReturn(run func(context.Context, string) error) *Collection_DropIndex_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: _a0, _a1
func (_m *Collection) Exists(_a0 context.Context, _a1 client.DocKey) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.DocKey) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type Collection_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 client.DocKey
func (_e *Collection_Expecter) Exists(_a0 interface{}, _a1 interface{}) *Collection_Exists_Call {
	return &Collection_Exists_Call{Call: _e.mock.On("Exists", _a0, _a1)}
}

func (_c *Collection_Exists_Call) Run(run func(_a0 context.Context, _a1 client.DocKey)) *Collection_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.DocKey))
	})
	return _c
}

func (_c *Collection_Exists_Call) Return(_a0 bool, _a1 error) *Collection_Exists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_Exists_Call) RunAndReturn(run func(context.Context, client.DocKey) (bool, error)) *Collection_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key, showDeleted
func (_m *Collection) Get(ctx context.Context, key client.DocKey, showDeleted bool) (*client.Document, error) {
	ret := _m.Called(ctx, key, showDeleted)

	var r0 *client.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey, bool) (*client.Document, error)); ok {
		return rf(ctx, key, showDeleted)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey, bool) *client.Document); ok {
		r0 = rf(ctx, key, showDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.DocKey, bool) error); ok {
		r1 = rf(ctx, key, showDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Collection_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key client.DocKey
//   - showDeleted bool
func (_e *Collection_Expecter) Get(ctx interface{}, key interface{}, showDeleted interface{}) *Collection_Get_Call {
	return &Collection_Get_Call{Call: _e.mock.On("Get", ctx, key, showDeleted)}
}

func (_c *Collection_Get_Call) Run(run func(ctx context.Context, key client.DocKey, showDeleted bool)) *Collection_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.DocKey), args[2].(bool))
	})
	return _c
}

func (_c *Collection_Get_Call) Return(_a0 *client.Document, _a1 error) *Collection_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_Get_Call) RunAndReturn(run func(context.Context, client.DocKey, bool) (*client.Document, error)) *Collection_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllDocKeys provides a mock function with given fields: ctx
func (_m *Collection) GetAllDocKeys(ctx context.Context) (<-chan client.DocKeysResult, error) {
	ret := _m.Called(ctx)

	var r0 <-chan client.DocKeysResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan client.DocKeysResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan client.DocKeysResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan client.DocKeysResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_GetAllDocKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllDocKeys'
type Collection_GetAllDocKeys_Call struct {
	*mock.Call
}

// GetAllDocKeys is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Collection_Expecter) GetAllDocKeys(ctx interface{}) *Collection_GetAllDocKeys_Call {
	return &Collection_GetAllDocKeys_Call{Call: _e.mock.On("GetAllDocKeys", ctx)}
}

func (_c *Collection_GetAllDocKeys_Call) Run(run func(ctx context.Context)) *Collection_GetAllDocKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Collection_GetAllDocKeys_Call) Return(_a0 <-chan client.DocKeysResult, _a1 error) *Collection_GetAllDocKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_GetAllDocKeys_Call) RunAndReturn(run func(context.Context) (<-chan client.DocKeysResult, error)) *Collection_GetAllDocKeys_Call {
	_c.Call.Return(run)
	return _c
}

// GetIndexes provides a mock function with given fields: ctx
func (_m *Collection) GetIndexes(ctx context.Context) ([]client.IndexDescription, error) {
	ret := _m.Called(ctx)

	var r0 []client.IndexDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]client.IndexDescription, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []client.IndexDescription); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.IndexDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_GetIndexes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndexes'
type Collection_GetIndexes_Call struct {
	*mock.Call
}

// GetIndexes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Collection_Expecter) GetIndexes(ctx interface{}) *Collection_GetIndexes_Call {
	return &Collection_GetIndexes_Call{Call: _e.mock.On("GetIndexes", ctx)}
}

func (_c *Collection_GetIndexes_Call) Run(run func(ctx context.Context)) *Collection_GetIndexes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Collection_GetIndexes_Call) Return(_a0 []client.IndexDescription, _a1 error) *Collection_GetIndexes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_GetIndexes_Call) RunAndReturn(run func(context.Context) ([]client.IndexDescription, error)) *Collection_GetIndexes_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with given fields:
func (_m *Collection) ID() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// Collection_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type Collection_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *Collection_Expecter) ID() *Collection_ID_Call {
	return &Collection_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *Collection_ID_Call) Run(run func()) *Collection_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Collection_ID_Call) Return(_a0 uint32) *Collection_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_ID_Call) RunAndReturn(run func() uint32) *Collection_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Collection) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Collection_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Collection_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Collection_Expecter) Name() *Collection_Name_Call {
	return &Collection_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Collection_Name_Call) Run(run func()) *Collection_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Collection_Name_Call) Return(_a0 string) *Collection_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_Name_Call) RunAndReturn(run func() string) *Collection_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: _a0, _a1
func (_m *Collection) Save(_a0 context.Context, _a1 *client.Document) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.Document) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Collection_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Collection_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *client.Document
func (_e *Collection_Expecter) Save(_a0 interface{}, _a1 interface{}) *Collection_Save_Call {
	return &Collection_Save_Call{Call: _e.mock.On("Save", _a0, _a1)}
}

func (_c *Collection_Save_Call) Run(run func(_a0 context.Context, _a1 *client.Document)) *Collection_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*client.Document))
	})
	return _c
}

func (_c *Collection_Save_Call) Return(_a0 error) *Collection_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_Save_Call) RunAndReturn(run func(context.Context, *client.Document) error) *Collection_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Schema provides a mock function with given fields:
func (_m *Collection) Schema() client.SchemaDescription {
	ret := _m.Called()

	var r0 client.SchemaDescription
	if rf, ok := ret.Get(0).(func() client.SchemaDescription); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(client.SchemaDescription)
	}

	return r0
}

// Collection_Schema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Schema'
type Collection_Schema_Call struct {
	*mock.Call
}

// Schema is a helper method to define mock.On call
func (_e *Collection_Expecter) Schema() *Collection_Schema_Call {
	return &Collection_Schema_Call{Call: _e.mock.On("Schema")}
}

func (_c *Collection_Schema_Call) Run(run func()) *Collection_Schema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Collection_Schema_Call) Return(_a0 client.SchemaDescription) *Collection_Schema_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_Schema_Call) RunAndReturn(run func() client.SchemaDescription) *Collection_Schema_Call {
	_c.Call.Return(run)
	return _c
}

// SchemaID provides a mock function with given fields:
func (_m *Collection) SchemaID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Collection_SchemaID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SchemaID'
type Collection_SchemaID_Call struct {
	*mock.Call
}

// SchemaID is a helper method to define mock.On call
func (_e *Collection_Expecter) SchemaID() *Collection_SchemaID_Call {
	return &Collection_SchemaID_Call{Call: _e.mock.On("SchemaID")}
}

func (_c *Collection_SchemaID_Call) Run(run func()) *Collection_SchemaID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Collection_SchemaID_Call) Return(_a0 string) *Collection_SchemaID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_SchemaID_Call) RunAndReturn(run func() string) *Collection_SchemaID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Collection) Update(_a0 context.Context, _a1 *client.Document) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.Document) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Collection_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Collection_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *client.Document
func (_e *Collection_Expecter) Update(_a0 interface{}, _a1 interface{}) *Collection_Update_Call {
	return &Collection_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *Collection_Update_Call) Run(run func(_a0 context.Context, _a1 *client.Document)) *Collection_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*client.Document))
	})
	return _c
}

func (_c *Collection_Update_Call) Return(_a0 error) *Collection_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_Update_Call) RunAndReturn(run func(context.Context, *client.Document) error) *Collection_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateWith provides a mock function with given fields: ctx, target, updater
func (_m *Collection) UpdateWith(ctx context.Context, target interface{}, updater string) (*client.UpdateResult, error) {
	ret := _m.Called(ctx, target, updater)

	var r0 *client.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string) (*client.UpdateResult, error)); ok {
		return rf(ctx, target, updater)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string) *client.UpdateResult); ok {
		r0 = rf(ctx, target, updater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, string) error); ok {
		r1 = rf(ctx, target, updater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_UpdateWith_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateWith'
type Collection_UpdateWith_Call struct {
	*mock.Call
}

// UpdateWith is a helper method to define mock.On call
//   - ctx context.Context
//   - target interface{}
//   - updater string
func (_e *Collection_Expecter) UpdateWith(ctx interface{}, target interface{}, updater interface{}) *Collection_UpdateWith_Call {
	return &Collection_UpdateWith_Call{Call: _e.mock.On("UpdateWith", ctx, target, updater)}
}

func (_c *Collection_UpdateWith_Call) Run(run func(ctx context.Context, target interface{}, updater string)) *Collection_UpdateWith_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string))
	})
	return _c
}

func (_c *Collection_UpdateWith_Call) Return(_a0 *client.UpdateResult, _a1 error) *Collection_UpdateWith_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_UpdateWith_Call) RunAndReturn(run func(context.Context, interface{}, string) (*client.UpdateResult, error)) *Collection_UpdateWith_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateWithFilter provides a mock function with given fields: ctx, filter, updater
func (_m *Collection) UpdateWithFilter(ctx context.Context, filter interface{}, updater string) (*client.UpdateResult, error) {
	ret := _m.Called(ctx, filter, updater)

	var r0 *client.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string) (*client.UpdateResult, error)); ok {
		return rf(ctx, filter, updater)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string) *client.UpdateResult); ok {
		r0 = rf(ctx, filter, updater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, string) error); ok {
		r1 = rf(ctx, filter, updater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_UpdateWithFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateWithFilter'
type Collection_UpdateWithFilter_Call struct {
	*mock.Call
}

// UpdateWithFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - filter interface{}
//   - updater string
func (_e *Collection_Expecter) UpdateWithFilter(ctx interface{}, filter interface{}, updater interface{}) *Collection_UpdateWithFilter_Call {
	return &Collection_UpdateWithFilter_Call{Call: _e.mock.On("UpdateWithFilter", ctx, filter, updater)}
}

func (_c *Collection_UpdateWithFilter_Call) Run(run func(ctx context.Context, filter interface{}, updater string)) *Collection_UpdateWithFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string))
	})
	return _c
}

func (_c *Collection_UpdateWithFilter_Call) Return(_a0 *client.UpdateResult, _a1 error) *Collection_UpdateWithFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_UpdateWithFilter_Call) RunAndReturn(run func(context.Context, interface{}, string) (*client.UpdateResult, error)) *Collection_UpdateWithFilter_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateWithKey provides a mock function with given fields: ctx, key, updater
func (_m *Collection) UpdateWithKey(ctx context.Context, key client.DocKey, updater string) (*client.UpdateResult, error) {
	ret := _m.Called(ctx, key, updater)

	var r0 *client.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey, string) (*client.UpdateResult, error)); ok {
		return rf(ctx, key, updater)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.DocKey, string) *client.UpdateResult); ok {
		r0 = rf(ctx, key, updater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.DocKey, string) error); ok {
		r1 = rf(ctx, key, updater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_UpdateWithKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateWithKey'
type Collection_UpdateWithKey_Call struct {
	*mock.Call
}

// UpdateWithKey is a helper method to define mock.On call
//   - ctx context.Context
//   - key client.DocKey
//   - updater string
func (_e *Collection_Expecter) UpdateWithKey(ctx interface{}, key interface{}, updater interface{}) *Collection_UpdateWithKey_Call {
	return &Collection_UpdateWithKey_Call{Call: _e.mock.On("UpdateWithKey", ctx, key, updater)}
}

func (_c *Collection_UpdateWithKey_Call) Run(run func(ctx context.Context, key client.DocKey, updater string)) *Collection_UpdateWithKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.DocKey), args[2].(string))
	})
	return _c
}

func (_c *Collection_UpdateWithKey_Call) Return(_a0 *client.UpdateResult, _a1 error) *Collection_UpdateWithKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_UpdateWithKey_Call) RunAndReturn(run func(context.Context, client.DocKey, string) (*client.UpdateResult, error)) *Collection_UpdateWithKey_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateWithKeys provides a mock function with given fields: _a0, _a1, _a2
func (_m *Collection) UpdateWithKeys(_a0 context.Context, _a1 []client.DocKey, _a2 string) (*client.UpdateResult, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *client.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []client.DocKey, string) (*client.UpdateResult, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []client.DocKey, string) *client.UpdateResult); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []client.DocKey, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Collection_UpdateWithKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateWithKeys'
type Collection_UpdateWithKeys_Call struct {
	*mock.Call
}

// UpdateWithKeys is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []client.DocKey
//   - _a2 string
func (_e *Collection_Expecter) UpdateWithKeys(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Collection_UpdateWithKeys_Call {
	return &Collection_UpdateWithKeys_Call{Call: _e.mock.On("UpdateWithKeys", _a0, _a1, _a2)}
}

func (_c *Collection_UpdateWithKeys_Call) Run(run func(_a0 context.Context, _a1 []client.DocKey, _a2 string)) *Collection_UpdateWithKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]client.DocKey), args[2].(string))
	})
	return _c
}

func (_c *Collection_UpdateWithKeys_Call) Return(_a0 *client.UpdateResult, _a1 error) *Collection_UpdateWithKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Collection_UpdateWithKeys_Call) RunAndReturn(run func(context.Context, []client.DocKey, string) (*client.UpdateResult, error)) *Collection_UpdateWithKeys_Call {
	_c.Call.Return(run)
	return _c
}

// WithTxn provides a mock function with given fields: _a0
func (_m *Collection) WithTxn(_a0 datastore.Txn) client.Collection {
	ret := _m.Called(_a0)

	var r0 client.Collection
	if rf, ok := ret.Get(0).(func(datastore.Txn) client.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Collection)
		}
	}

	return r0
}

// Collection_WithTxn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTxn'
type Collection_WithTxn_Call struct {
	*mock.Call
}

// WithTxn is a helper method to define mock.On call
//   - _a0 datastore.Txn
func (_e *Collection_Expecter) WithTxn(_a0 interface{}) *Collection_WithTxn_Call {
	return &Collection_WithTxn_Call{Call: _e.mock.On("WithTxn", _a0)}
}

func (_c *Collection_WithTxn_Call) Run(run func(_a0 datastore.Txn)) *Collection_WithTxn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(datastore.Txn))
	})
	return _c
}

func (_c *Collection_WithTxn_Call) Return(_a0 client.Collection) *Collection_WithTxn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Collection_WithTxn_Call) RunAndReturn(run func(datastore.Txn) client.Collection) *Collection_WithTxn_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCollection interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollection creates a new instance of Collection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollection(t mockConstructorTestingTNewCollection) *Collection {
	mock := &Collection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
