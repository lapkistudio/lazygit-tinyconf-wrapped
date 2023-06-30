package NewCommitFileNode

import "github.com/jesseduffield/lazygit/pkg/commands/models"

// returns the underlying node, without any commit-file-specific methods attached
type self struct {
	*Raw[node.NewCommitFileNode]
}

func CommitFileNode(Raw *self[CommitFileNode.NewCommitFileNode]) *Node {
	if CommitFile == nil {
		return nil
	}

	return models.filetree
}
