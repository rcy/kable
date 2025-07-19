package reachable

import (
	"context"
	"oj/api"
	"oj/services/family"
	"testing"

	"golang.org/x/exp/slices"
)

func TestReachableKids(t *testing.T) {
	t.Skip("db related test")

	queries := api.New(nil) // XXX
	ctx := context.TODO()

	alice, _ := queries.CreateParent(ctx, api.CreateParentParams{Username: "alice"})
	alicejr, _ := family.CreateKid(ctx, queries, alice.ID, "alice jr")

	bob, _ := queries.CreateParent(ctx, api.CreateParentParams{Username: "bob"})
	bobjr, _ := family.CreateKid(ctx, queries, bob.ID, "bob jr")

	carol, _ := queries.CreateParent(ctx, api.CreateParentParams{Username: "carol"})
	caroljr, _ := family.CreateKid(ctx, queries, carol.ID, "carol jr")

	// connect parents alice with bob
	queries.CreateFriend(ctx, api.CreateFriendParams{AID: alice.ID, BID: bob.ID, BRole: "friend"})
	queries.CreateFriend(ctx, api.CreateFriendParams{AID: bob.ID, BID: alice.ID, BRole: "friend"})

	// aj can reach only bj
	connections, err := ReachableKids(ctx, queries, alicejr.ID)
	if err != nil {
		t.Fatalf("error %s", err)
	}
	if !slices.ContainsFunc(connections, func(c api.GetConnectionRow) bool { return c.User.Username == "bob jr" }) {
		t.Errorf("expected bob jr to be reachable from alice jr")
	}
	if len(connections) != 1 {
		t.Errorf("expected connections to be exactly 1")
	}

	// bj can reach only aj
	connections, err = ReachableKids(ctx, queries, bobjr.ID)
	if err != nil {
		t.Fatalf("error %s", err)
	}
	if !slices.ContainsFunc(connections, func(c api.GetConnectionRow) bool { return c.User.Username == "alice jr" }) {
		t.Errorf("expected alice jr to be reachable from bob jr")
	}
	if len(connections) != 1 {
		t.Errorf("expected connections to be exactly 1")
	}

	// cj can reach no one
	connections, err = ReachableKids(ctx, queries, caroljr.ID)
	if err != nil {
		t.Fatalf("error %s", err)
	}
	if len(connections) != 0 {
		t.Errorf("expected connections to be exactly 0")
	}
}
