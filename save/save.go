package save

import (
	"github.com/cwxstat/dopt/file"
	"path/filepath"
	"errors"
	"strings"
	

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

func FileSaveCommit(repo *git.Repository, repoDir, fileSource, fileDest string) error {
	
	file.Copy(fileSource, fileDest)
	r, err := git.PlainOpen(repoDir)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	fileCommit := strings.Replace(fileDest,repoDir +"/","",1)
	if fileCommit == "" {
		return ErrCommit
	}
	_, err = w.Add(fileCommit)
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
