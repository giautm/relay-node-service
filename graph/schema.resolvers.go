package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strings"

	"giautm.dev/relay-node-service/graph/generated"
	"giautm.dev/relay-node-service/graph/model"
)

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	prefix := strings.Split(id, "_")[0]
	switch prefix {
	case "TD":
		return &model.Todo{ID: id}, nil
	case "DO":
		return &model.Document{ID: id}, nil
	}

	return nil, errors.New("unknown node type")
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]model.Node, error) {
	nodes := make([]model.Node, len(ids))
	for i, id := range ids {
		node, err := r.Node(ctx, id)
		if err != nil {
			return nil, err
		}

		nodes[i] = node
	}
	return nodes, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
