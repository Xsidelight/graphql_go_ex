package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Xsidelight/graphql_go_ex/graph/generated"
	"github.com/Xsidelight/graphql_go_ex/graph/model"
)

func (r *queryResolver) Places(ctx context.Context) ([]*model.Place, error) {
	var places []*model.Place
	dummyPlace1 := model.Place{
		Name: "Hotel Atlas Abashidze",
		LatLng: &model.LatLng{
			Latitude: 44.757001,
			Longitude: 41.708337,
		},
	}
	dummyPlace2 := model.Place{
		Name: "TBC Bank Saakadze Branch",
		LatLng: &model.LatLng{
			Latitude: 44.775456,
			Longitude: 41.724666,
		},
	}
	places = append(places, &dummyPlace1, &dummyPlace2)
	return places, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
