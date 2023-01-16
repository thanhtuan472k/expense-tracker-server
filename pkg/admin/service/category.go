package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/external/util/ptime"
	"expense-tracker-server/pkg/admin/dao"
	"expense-tracker-server/pkg/admin/errorcode"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

// CategoryInterface ...
type CategoryInterface interface {
	// Create new category ...
	Create(ctx context.Context, payload requestmodel.CategoryBodyCreate) (categoryID string, err error)

	// All ...
	All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseCategoryAll)

	// Detail ...
	Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseCategoryAdmin, err error)

	// Update ...
	Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryBodyUpdate) (categoryID string, err error)

	// ChangeStatus ...
	ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryChangeStatus) (result responsemodel.ResponseChangeStatus, err error)
}

// Category ...
func Category() CategoryInterface {
	return categoryImplement{}
}

// categoryImplement ...
type categoryImplement struct{}

//
// PUBLIC METHODS
//

// Create ...
func (s categoryImplement) Create(ctx context.Context, payload requestmodel.CategoryBodyCreate) (categoryID string, err error) {
	var (
		d   = dao.Category()
		doc = payload.ConvertToBSON()
	)

	// Create
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	categoryID = doc.ID.Hex()
	return
}

// All ...
func (s categoryImplement) All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseCategoryAll) {
	var (
		d  = dao.Category()
		wg = sync.WaitGroup{}
	)

	// Assign cond
	cond := bson.D{}
	q.ExpenseTracker.AssignKeyword(&cond)
	q.ExpenseTracker.AssignStatus(&cond)
	q.ExpenseTracker.AssignCategoryType(&cond)

	// Find
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Init data
		result.List = make([]responsemodel.ResponseCategoryAdmin, 0)

		// Find options
		findOpts := q.GetFindOptionsWithPage()

		// Find
		docs := d.FindByCondition(ctx, cond, findOpts)
		for _, doc := range docs {
			result.List = append(result.List, s.detail(ctx, doc))
		}
	}()

	// Assign total
	wg.Add(1)
	go func() {
		defer wg.Done()
		result.Total = d.CountByCondition(ctx, cond)
	}()

	wg.Wait()

	// Assign limit
	result.Limit = q.Limit

	return

}

// Detail ...
func (s categoryImplement) Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseCategoryAdmin, err error) {
	var (
		d    = dao.Category()
		cond = bson.M{"_id": id}
	)

	category := d.FindOneByCondition(ctx, cond)
	if category.ID.IsZero() {
		err = errors.New(errorcode.CategoryNotFound)
		return
	}

	result = s.detail(ctx, category)

	return
}

// Update ...
func (s categoryImplement) Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryBodyUpdate) (categoryID string, err error) {
	// Find category
	category, err := s.FindByID(ctx, id)
	if err != nil {
		return
	}

	var (
		d             = dao.Category()
		data          = payload.ConvertToBSON()
		payloadUpdate = bson.M{
			"name":      data.Name,
			"updatedAt": ptime.Now(),
		}
		cond = bson.D{{"_id", category.ID}}
	)

	// Update
	if err = d.UpdateOneByCondition(ctx, cond, bson.M{"$set": payloadUpdate}); err != nil {
		return
	}

	// Response
	categoryID = category.ID.Hex()
	return
}

// ChangeStatus ...
func (s categoryImplement) ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryChangeStatus) (result responsemodel.ResponseChangeStatus, err error) {
	//TODO implement me
	panic("implement me")
}

// FindByID ...
func (s categoryImplement) FindByID(ctx context.Context, id primitive.ObjectID) (result mgexpense.Category, err error) {
	var (
		d    = dao.Category()
		cond = bson.D{{"_id", id}}
	)

	category := d.FindOneByCondition(ctx, cond)
	if category.ID.IsZero() {
		err = errors.New(errorcode.CategoryNotFound)
		return
	}
	result = category
	return
}

//
// PRIVATE ...
//

// detail ...
func (s categoryImplement) detail(ctx context.Context, doc mgexpense.Category) (result responsemodel.ResponseCategoryAdmin) {
	return responsemodel.ResponseCategoryAdmin{
		ID:        doc.ID.Hex(),
		Name:      doc.Name,
		Type:      doc.Type,
		Status:    doc.Status,
		CreatedAt: ptime.TimeResponseInit(doc.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(doc.UpdatedAt),
	}
}
