package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/miyauchi-sup/graphql/graph/generated"
	"github.com/miyauchi-sup/graphql/graph/model"
)

func (r *mutationResolver) DeletePhase(ctx context.Context, phaseID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStation(ctx context.Context, stationID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteWave(ctx context.Context, waveID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResolveScheduleConflicts(ctx context.Context, tournamentID string, options *model.ResolveConflictsOptions) ([]*model.Seed, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SwapSeeds(ctx context.Context, phaseID string, seed1id string, seed2id string) ([]*model.Seed, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePhaseGroups(ctx context.Context, groupConfigs []*model.PhaseGroupUpdateInput) ([]*model.PhaseGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePhaseSeeding(ctx context.Context, phaseID string, seedMapping []*model.UpdatePhaseSeedInfo, options *model.UpdatePhaseSeedingOptions) (*model.Phase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpsertPhase(ctx context.Context, phaseID *string, eventID *string, payload model.PhaseUpsertInput) (*model.Phase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpsertStation(ctx context.Context, stationID *string, tournamentID *string, fields model.StationUpsertInput) (*model.Stations, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpsertWave(ctx context.Context, waveID *string, tournamentID *string, fields model.WaveUpsertInput) (*model.Wave, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Event(ctx context.Context, id *string, slug *string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) League(ctx context.Context, id *string, slug *string) (*model.League, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Participant(ctx context.Context, id string, isAdmin *bool) (*model.Participant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Phase(ctx context.Context, id *string) (*model.Phase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PhaseGroup(ctx context.Context, id *string) (*model.PhaseGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Player(ctx context.Context, id string) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Seed(ctx context.Context, id *string) (*model.Seed, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Set(ctx context.Context, id string) (*model.Set, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Shop(ctx context.Context, id *string, slug *string) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stream(ctx context.Context, id string) (*model.Streams, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) StreamQueue(ctx context.Context, tournamentID string, includePlayerStreams *bool) ([]*model.StreamQueue, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tournament(ctx context.Context, id *string, slug *string) (*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tournaments(ctx context.Context, query model.TournamentQuery) (*model.TournamentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id *string, slug *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Videogame(ctx context.Context, id *string) (*model.Videogame, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Videogames(ctx context.Context, query model.VideogameQuery) (*model.VideogameConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ Resolver }
type queryResolver struct{ Resolver }
