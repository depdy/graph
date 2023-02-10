package gt

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Repo struct {
	gr *git.Repository
}

func NewRepo(path string) (*Repo, error) {
	repo, err := git.PlainOpen(path)
	return &Repo{repo}, err
}

func (r Repo) Show(commit, path string) (io.ReadCloser, error) {
	h, err := r.gr.ResolveRevision(plumbing.Revision(commit))
	if err != nil {
		return nil, err
	}
	if h == nil {
		return nil, errors.New("no such commit")
	}

	obj, err := r.gr.Object(plumbing.AnyObject, *h)
	if err != nil {
		return nil, err
	}

	blob, err := resolve(obj, "actions/last/action.yaml")
	if err != nil {
		return nil, err
	}

	return blob.Reader()
}

func resolve(obj object.Object, path string) (*object.Blob, error) {
	switch o := obj.(type) {
	case *object.Commit:
		t, err := o.Tree()
		if err != nil {
			return nil, err
		}
		return resolve(t, path)
	case *object.Tag:
		target, err := o.Object()
		if err != nil {
			return nil, err
		}
		return resolve(target, path)
	case *object.Tree:
		file, err := o.File(path)
		if err != nil {
			return nil, err
		}
		return &file.Blob, nil
	case *object.Blob:
		return o, nil
	default:
		return nil, object.ErrUnsupportedObject
	}
}
