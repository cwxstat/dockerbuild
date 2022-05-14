package save

import (
	"errors"
	"github.com/cwxstat/dopt/file"
	"path/filepath"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"time"
)

var ErrCommit = errors.New("no commit")

func Init(gitPath string) (*git.Repository, error) {
	worktreeFS := osfs.New(gitPath)
	gitDirFS := osfs.New(filepath.Join(gitPath, ".git"))
	repo, err := git.Init(filesystem.NewStorage(gitDirFS, cache.NewObjectLRUDefault()), worktreeFS)
	return repo, err
}

func FileSaveCommit(repo *git.Repository, repoDir, fileSource string) error {

	dirFileToCommit, err := file.CopyR(fileSource, repoDir)
	if err != nil {
		return err
	}

	r, err := git.PlainOpen(repoDir)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	_, err = w.Add(dirFileToCommit)
	if err != nil {
		return err
	}

	_, err = w.Commit("dopt go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "dopt",
			Email: "dopt@cwxstat.io",
			When:  time.Now(),
		},
	})
	return err

}
