package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/bygui86/go-graphql/server/graph/generated"
	"github.com/bygui86/go-graphql/server/graph/model"
)

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (m *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	order := model.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}

	err := m.DB.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (m *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {
	updatedOrder := model.Order{
		ID:           orderID,
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}

	err := m.DB.Save(&updatedOrder).Error
	if err != nil {
		return nil, err
	}

	return &updatedOrder, nil
}

func (m *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	queryWhereId := "order_id = ?"

	itemErr := m.DB.Where(queryWhereId, orderID).Delete(&model.Item{}).Error
	if itemErr != nil {
		return false, itemErr
	}

	orderErr := m.DB.Where(queryWhereId, orderID).Delete(&model.Order{}).Error
	if orderErr != nil {
		return false, orderErr
	}

	return true, nil
}

func (q *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order

	// Preload() method ensures that associations are preloaded while using the Find()
	//err := q.DB.Preload("Items").Find(&orders).Error

	err := q.DB.Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func mapItemsFromInput(itemsInput []*model.ItemInput) []*model.Item {
	var items []*model.Item
	for _, itemInput := range itemsInput {
		items = append(items, &model.Item{
			ProductCode: itemInput.ProductCode,
			ProductName: itemInput.ProductName,
			Quantity:    itemInput.Quantity,
		})
	}
	return items
}
