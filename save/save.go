package save

import (
	"path/filepath"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
)

func Init(gitPath string) (*git.Repository, error) {
	worktreeFS := osfs.New(gitPath)
	gitDirFS := osfs.New(filepath.Join(gitPath, ".git"))
	repo, err := git.Init(filesystem.NewStorage(gitDirFS, cache.NewObjectLRUDefault()), worktreeFS)
	return repo, err
}

func FileSaveCommit(repo *git.Repository, file string) error {

}
