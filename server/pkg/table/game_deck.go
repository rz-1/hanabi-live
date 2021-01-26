package table

import (
	"math/rand"

	"github.com/Zamiell/hanabi-live/server/pkg/options"
	"github.com/Zamiell/hanabi-live/server/pkg/variants"
)

func (g *game) initDeck() {
	// Local variables
	t := g.table

	// If a custom deck was provided along with the game options,
	// then we can simply add every card to the deck as specified
	if t.ExtraOptions.CustomDeck != nil &&
		len(t.ExtraOptions.CustomDeck) != 0 &&
		t.ExtraOptions.CustomSeed == "" { // Custom seeds override custom decks

		for _, c := range t.ExtraOptions.CustomDeck {
			g.Deck = append(g.Deck, newCard(c.SuitIndex, c.Rank))
			g.CardIdentities = append(g.CardIdentities, &options.CardIdentity{
				SuitIndex: c.SuitIndex,
				Rank:      c.Rank,
			})
		}

		return
	}

	// Suits are represented as a slice of integers from 0 to the number of suits - 1
	// (e.g. [0, 1, 2, 3, 4] for a "No Variant" game)
	for suitIndex, suit := range t.Variant.Suits {
		// Ranks are represented as a slice of integers
		// (e.g. [1, 2, 3, 4, 5] for a "No Variant" game)
		for _, rank := range t.Variant.Ranks {
			// In a normal suit, there are:
			// - three 1's
			// - two 2's
			// - two 3's
			// - two 4's
			// - one five
			var amountToAdd int
			if rank == 1 {
				amountToAdd = 3
				if t.Variant.IsUpOrDown() || suit.Reversed {
					amountToAdd = 1
				}
			} else if rank == 5 { // nolint: gomnd
				amountToAdd = 1
				if suit.Reversed {
					amountToAdd = 3
				}
			} else if rank == variants.StartCardRank {
				amountToAdd = 1
			} else {
				amountToAdd = 2
			}
			if suit.OneOfEach {
				amountToAdd = 1
			}

			for i := 0; i < amountToAdd; i++ {
				// Add the card to the deck
				g.Deck = append(g.Deck, newCard(suitIndex, rank))
				g.CardIdentities = append(g.CardIdentities, &options.CardIdentity{
					SuitIndex: suitIndex,
					Rank:      rank,
				})
			}
		}
	}
}

func (g *game) shuffleDeck() {
	// It is assumed that "rand.Seed()" is already set before getting here
	// From: https://stackoverflow.com/questions/12264789/shuffle-array-in-go
	for i := range g.Deck {
		j := rand.Intn(i + 1) // nolint: gosec
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
		g.CardIdentities[i], g.CardIdentities[j] = g.CardIdentities[j], g.CardIdentities[i]
	}
}