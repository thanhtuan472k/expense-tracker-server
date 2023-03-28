package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/external/util/pagetoken"
	"expense-tracker-server/external/util/ptime"
	"expense-tracker-server/pkg/app/dao"
	"expense-tracker-server/pkg/app/errorcode"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	responsemodel "expense-tracker-server/pkg/app/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncomeMoneyInterface ...
type IncomeMoneyInterface interface {
	// Create ...
	Create(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeMoneyBodyCreate) (incomeID string, err error)

	// Update ...
	Update(ctx context.Context, userID, incomeID primitive.ObjectID, payload requestmodel.IncomeMoneyBodyUpdate) (result string, err error)

	// All ...
	All(ctx context.Context, q mgquerry.AppQuery, userID primitive.ObjectID) (result responsemodel.ResponseIncomeMoneyAll, err error)

	// Detail ...
	Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseIncomeMoneyInfo, err error)
}

// IncomeMoney ...
func IncomeMoney() IncomeMoneyInterface {
	return incomeMoneyImplement{}
}

// incomeMoneyImplement ...
type incomeMoneyImplement struct{}

//
// PUBLIC METHODS
//

// Create ...
func (s incomeMoneyImplement) Create(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeMoneyBodyCreate) (incomeID string, err error) {
	// Convert payload string to ObjectId
	categoryID, _ := mongodb.NewIDFromString(payload.Category)

	// Find doc category is available
	categorySvc := categoryImplement{}
	category, err := categorySvc.findDocCategoryTypeIncomeAvailableByID(ctx, categoryID)
	if err != nil {
		return
	}

	var (
		d   = dao.IncomeMoney()
		doc = payload.ConvertToBSON(userID, category)
	)

	// Insert doc income money
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	incomeID = doc.ID.Hex()
	return
}

// Update ...
func (s incomeMoneyImplement) Update(ctx context.Context, userID, incomeID primitive.ObjectID, payload requestmodel.IncomeMoneyBodyUpdate) (result string, err error) {
	// Find incomeMoney
	incomeMoney, err := s.FindByID(ctx, incomeID)
	if err != nil {
		return
	}

	// Convert payload string to ObjectId
	categoryID, _ := mongodb.NewIDFromString(payload.Category)

	// Find doc category is available
	categorySvc := categoryImplement{}
	category, err := categorySvc.findDocCategoryTypeIncomeAvailableByID(ctx, categoryID)
	if err != nil {
		return
	}

	var (
		d             = dao.IncomeMoney()
		data          = payload.ConvertToBSON(category)
		payloadUpdate = bson.M{
			"money":     data.Money,
			"note":      data.Note,
			"category":  data.Category,
			"updatedAt": data.UpdatedAt,
		}
		cond = bson.D{{"_id", incomeMoney.ID}}
	)

	// Update
	if err = d.UpdateOneByCondition(ctx, cond, bson.M{"$set": payloadUpdate}); err != nil {
		return
	}

	// Response
	result = incomeMoney.ID.Hex()
	return
}

// All ...
func (s incomeMoneyImplement) All(ctx context.Context, q mgquerry.AppQuery, userID primitive.ObjectID) (result responsemodel.ResponseIncomeMoneyAll, err error) {
	var (
		d    = dao.IncomeMoney()
		cond = bson.D{
			{"user", userID},
		}
	)

	// Assign query
	s.assignQueryIncomeMoney(&q, &cond)

	// Assign sort
	s.assignQuerySort(&q)

	// Init data
	var list = make([]responsemodel.ResponseIncomeMoneyInfo, 0)

	// Find docs
	docs := d.FindByCondition(ctx, cond, q.GetFindOptionsWithPage())
	for _, doc := range docs {
		list = append(list, s.brief(ctx, doc))
	}

	// Page token
	total := len(list)
	endData := total < int(q.Limit)
	var nextPageToken = ""
	if len(list) == int(q.Limit) {
		nextPageToken = pagetoken.PageTokenUsingPage(int(q.Page) + 1)
	}

	// Response
	result = responsemodel.ResponseIncomeMoneyAll{
		List:          list,
		EndData:       endData,
		NextPageToken: nextPageToken,
		Total:         int64(total),
	}

	return
}

// Detail ...
func (s incomeMoneyImplement) Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseIncomeMoneyInfo, err error) {
	return
}

// FindByID ...
func (s incomeMoneyImplement) FindByID(ctx context.Context, id primitive.ObjectID) (result mgexpense.IncomeMoney, err error) {
	var (
		d    = dao.IncomeMoney()
		cond = bson.D{{"_id", id}}
	)

	incomeMoney := d.FindOneByCondition(ctx, cond)
	if incomeMoney.ID.IsZero() {
		err = errors.New(errorcode.IncomeMoneyNotFound)
		return
	}
	result = incomeMoney
	return
}

//
// PRIVATE METHODS
//

// assignQueryIncomeMoney ...
func (s incomeMoneyImplement) assignQueryIncomeMoney(q *mgquerry.AppQuery, cond *bson.D) {
	q.ExpenseTracker.AssignFromToAt(cond)
}

// assignQuerySort ...
func (s incomeMoneyImplement) assignQuerySort(q *mgquerry.AppQuery) {
	switch q.SortString {
	case "created_at_first":
		q.SortInterface = bson.D{
			{"createdAt", 1},
			{"_id", 1},
		}
	default:
		q.SortInterface = bson.D{
			{"createdAt", -1},
			{"_id", -1},
		}
	}
}

// brief ...
func (s incomeMoneyImplement) brief(ctx context.Context, doc mgexpense.IncomeMoney) responsemodel.ResponseIncomeMoneyInfo {
	return responsemodel.ResponseIncomeMoneyInfo{
		ID: doc.ID.Hex(),
		Category: mgexpense.CategoryShort{
			ID:   doc.Category.ID,
			Name: doc.Category.Name,
		},
		Money:     doc.Money, // TODO: refactor bỏ format tiền
		Note:      doc.Note,
		CreatedAt: ptime.TimeResponseInit(doc.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(doc.UpdatedAt),
	}
}
