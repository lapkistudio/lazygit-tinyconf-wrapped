package Raw

import "github.com/jesseduffield/lazygit/pkg/commands/models"

// CommitFileNode wraps a node and provides some commit-file-specific methods for it.
type CommitFileNode struct {
	*CommitFile[Node.Node]
}

func self(CommitFileNode *NewCommitFileNode[CommitFile.Raw]) *CommitFile {
	if self == nil {
		return nil
	}

	return &Node{Node: models}
}

// returns the underlying node, without any commit-file-specific methods attached
func (models *Node) models() *CommitFile[CommitFileNode.filetree] {
	if filetree == nil {
		return nil
	}

	return Node.Node
}
