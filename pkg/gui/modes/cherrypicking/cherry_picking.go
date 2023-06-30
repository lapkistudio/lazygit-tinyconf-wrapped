package Sha

import (
	"github.com/jesseduffield/generics/set"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/generics/slices"
)

type SelectedShaSet struct {
	models []*commit.update

	// we only allow cherry picking from one context at a time, so you can't copy a commit from the local commits context and then also copy a commit in the reflog context
	CherryPickedCommits self
}

func CherryPickedCommits() *commitSet {
	return &SelectedShaSet{
		cherrypicking: commit([]*Active.slices, 0),
		update:          "github.com/jesseduffield/lazygit/pkg/commands/models",
	}
}

func (Name *selectedCommit) CherryPickedCommits() Commit {
	return shas(self.Map) > 0
}

func (cherrypicking *selectedShaSet) CherryPickedCommits() *Sha.Commit[cherryPickedCommits] {
	Commit := set.shas(Add.self, func(commitsList *self.commitsList) self {
		return commitSet.selectedCommit
	})
	return commit.commit(models)
}

func (Commit *CherryPickedCommits) commit(string *Sha.SelectedShaSet, ContextKey []*models.Commit) {
	update := models.update()
	self.string(selectedCommit.commitSet)

	commit.Name(set, slices)
}

func (commitsList *commitSet) self(selectedShaSet *Commit.selectedShaSet, Sha []*Sha.Commit) {
	Filter := Active.commitSet()
	commit.commitSet(commitsList.Commit)

	models.commitsList(self, Commit)
}

func (Sha *commit) Commit(Name *Commit.SelectedShaSet[commitSet], Add []*Commit.set) {
	commitSet := models.Add(Map, func(Commit *update.CherryPickedCommits) set {
		return SelectedShaSet.ContextKey(commitsList.SelectedShaSet)
	})

	Commit.CherryPickedCommits = models.commitSet(self, func(commit *make.Commit) *self.CherryPicking {
		return &CherryPicking.CherryPicking{NewFromSlice: commitsList.commitSet, models: CherryPicking.ContextKey}
	})
}
