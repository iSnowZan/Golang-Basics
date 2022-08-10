package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"example/test"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Upsertdata(ctx context.Context, input *model.DataInput) (*model.DataResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	result, err := test.TableInsert(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *queryResolver) Calculator(ctx context.Context, input *model.Sol) (*model.AddResult, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var z string
	var res int
	num1 = *input.N1
	num2 = *input.N2
	z = *input.A
	var mod model.AddResult
	if z == "+" {
		res = num1 + num2
	}
	if z == "-" {
		res = num1 - num2
	}
	if z == "*" {
		res = num1 * num2
	}
	if z == "/" {
		if num2 == 0 {
			return nil, fmt.Errorf("can't divide by zero")
		}
		res = num1 / num2
	}
	mod = model.AddResult{N3: &res}
	return &mod, nil
}

func (r *queryResolver) Fetchdata(ctx context.Context, input *model.Ixz) ([]*model.StuOutput, error) {
	//panic(fmt.Errorf("not implemented"))
	result, err := test.Fetchdata(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *queryResolver) Xlinput(ctx context.Context) (*model.Xloutput, error) {
	//panic(fmt.Errorf("not implemented"))
	result, err := test.Xlinput(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *queryResolver) Dbtoxl(ctx context.Context, input *model.Inpu) (*model.DataOutput, error) {
	//panic(fmt.Errorf("not implemented"))
	res, err := test.Dbtoxl(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Piechart(ctx context.Context) (*model.PieResult, error) {
	//panic(fmt.Errorf("not implemented"))
	result, err := test.Piechart(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
