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

// SubCategoryInterface ...
type SubCategoryInterface interface {
	// Create new subCategory ...
	Create(ctx context.Context, categoryID primitive.ObjectID, payload requestmodel.SubCategoryBodyCreate) (subCategoryID string, err error)

	// All ...
	All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseSubCategoryAll)

	// Detail ...
	Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseSubCategoryAdmin, err error)

	// Update ...
	Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryBodyUpdate) (subCategoryID string, err error)

	// ChangeStatus ...
	ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryChangeStatus) (subCategoryID string, err error)
}

// SubCategory ...
func SubCategory() SubCategoryInterface {
	return subCategoryImplement{}
}

// subCategoryImplement ...
type subCategoryImplement struct{}

//
// PUBLIC METHODS
//

// Create ...
func (s subCategoryImplement) Create(ctx context.Context, categoryID primitive.ObjectID, payload requestmodel.SubCategoryBodyCreate) (subCategoryID string, err error) {
	// Find category by id
	var categorySvc = categoryImplement{}
	category, _ := categorySvc.FindByID(ctx, categoryID)

	var (
		d   = dao.SubCategory()
		doc = payload.ConvertToBSON(category)
	)

	// Check category is active
	if !category.IsActiveCategory() {
		err = errors.New(errorcode.CategoryStatusInactive)
		return
	}

	// Create
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	subCategoryID = doc.ID.Hex()
	return

}

// All ...
func (s subCategoryImplement) All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseSubCategoryAll) {
	var (
		d  = dao.SubCategory()
		wg = sync.WaitGroup{}
	)

	// Assign condition
	cond := bson.D{}
	q.ExpenseTracker.AssignCategoryID(&cond)
	q.ExpenseTracker.AssignStatus(&cond)
	q.ExpenseTracker.AssignKeyword(&cond)

	// Assign total
	wg.Add(1)
	go func() {
		defer wg.Done()
		result.Total = d.CountByCondition(ctx, cond)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Init data
		result.List = make([]responsemodel.ResponseSubCategoryAdmin, 0)
		// Find options
		findOpts := q.GetFindOptionsWithPage()
		docs := d.FindByCondition(ctx, cond, findOpts)
		for _, doc := range docs {
			result.List = append(result.List, s.brief(ctx, doc))
		}
	}()

	wg.Wait()

	// Assign limit
	result.Limit = q.Limit

	// Response
	return

}

// Detail ...
func (s subCategoryImplement) Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseSubCategoryAdmin, err error) {
	// Find subCategory
	subCategory, err := s.FindByID(ctx, id)
	if err != nil {
		return
	}

	// Response
	result = s.detail(ctx, subCategory)
	return
}

// Update ...
func (s subCategoryImplement) Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryBodyUpdate) (categoryID string, err error) {
	//TODO implement me
	panic("implement me")
}

// ChangeStatus ...
func (s subCategoryImplement) ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryChangeStatus) (categoryID string, err error) {
	//TODO implement me
	panic("implement me")
}

// FindByID ...
func (s subCategoryImplement) FindByID(ctx context.Context, id primitive.ObjectID) (result mgexpense.SubCategory, err error) {
	var (
		d    = dao.SubCategory()
		cond = bson.D{{"_id", id}}
	)

	subCategory := d.FindOneByCondition(ctx, cond)
	if subCategory.ID.IsZero() {
		err = errors.New(errorcode.SubCategoryNotFound)
		return
	}
	result = subCategory
	return
}

//
// PRIVATE ...
//

// brief ...
func (s subCategoryImplement) brief(ctx context.Context, doc mgexpense.SubCategory) (result responsemodel.ResponseSubCategoryAdmin) {
	return responsemodel.ResponseSubCategoryAdmin{
		ID:        doc.ID.Hex(),
		Name:      doc.Name,
		Type:      doc.Type,
		Status:    doc.Status,
		CreatedAt: ptime.TimeResponseInit(doc.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(doc.UpdatedAt),
	}
}

// detail ...
func (s subCategoryImplement) detail(ctx context.Context, doc mgexpense.SubCategory) (result responsemodel.ResponseSubCategoryAdmin) {
	return responsemodel.ResponseSubCategoryAdmin{
		ID:        doc.ID.Hex(),
		Name:      doc.Name,
		Type:      doc.Type,
		Status:    doc.Status,
		CreatedAt: ptime.TimeResponseInit(doc.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(doc.UpdatedAt),
	}
}
