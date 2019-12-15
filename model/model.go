package model

import (
	"github.com/jinzhu/gorm"
)

// Election represents an election.
type Election struct {
	gorm.Model
	Name       string      // The name of this election.
	Candidates []Candidate // Candidates that run in this election.
	Voters     []Voter     // Voters that are allowed to vote in this election.
	Votes      []Vote      // Votes in this election.
}

// Candidate represents a candidate that run in an election.
type Candidate struct {
	gorm.Model
	ElectionId uint   // The id of the election this candidate run in.
	Name       string // Name of candidate.
}

// Voter represents a user whom are eligible to vote in an election.
type Voter struct {
	gorm.Model
	ElectionID uint   // The id of the election this voter may vote in.
	Name       string // Name of
}

// Vote represents a vote in an election.
type Vote struct {
	gorm.Model
	ElectionID  uint        // The id of the election this votes belongs to.
	VoteEntries []VoteEntry // Vote entries belonging to this vote.
}

// A vote entry is more or less a many to many mapping with the added rank
// for keeping track of internal order of candidates within one vote.
type VoteEntry struct {
	gorm.Model
	VoteID      uint      // The id of the vote this vote entry belongs to.
	CandidateID uint      // The id of the candidate this entry belongs to.
	Candidate   Candidate // The candidate this vote entry belongs to.
	Rank        uint      // The rank within the vote this entry has.
}
