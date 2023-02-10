package github

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/depdy/graph/gt"
	"gopkg.in/yaml.v2"
)

type Repository struct {
	fNames    []string
	git       *gt.Repo
	Workflows []Workflow
	Actions   []Action
}

// NewRepository parses the given files and returns a Repository.
func NewRepository(path string) (*Repository, error) {
	fNames, err := readFiles(".", ".*[.]y[a]*ml")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	ws, as, err := parseFiles(fNames)
	repo, err := gt.NewRepo(ws, as)
	if err != nil {
		return nil, err
	}
	return &Repository{
		fNames:    fNames,
		git:       repo,
		Workflows: ws,
		Actions:   as,
	}, err
}

func readFiles(path, pattern string) ([]string, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	var files []string
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if !r.MatchString(info.Name()) {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files, err
}

// It takes a list of files, reads them, and returns a list of workflows and actions
func parseFiles(files []string) (ws []Workflow, as []Action, err error) {
	for _, fName := range files {
		file, err := os.ReadFile(fName)

		ft, err := getFileType(file)
		if err != nil {
			return nil, nil, err
		}

		switch ft {
		case WorkflowType:
			var w workflow
			if err := yaml.Unmarshal(file, &w); err != nil {
				return nil, nil, err
			}
			ws = append(ws, *NewWorkflow(strings.TrimSuffix(filepath.Base(fName), filepath.Ext(fName)), ""))
		case ActionType:
			var a action
			if err := yaml.Unmarshal(file, &a); err != nil {
				return nil, nil, err
			}
			as = append(as, Action{
				Metadata: Metadata{
					Name: a.Name,
				},
			})
		}
	}

	return ws, as, err
}

type FileType int

const (
	Unkown FileType = iota
	ActionType
	WorkflowType
)

func getFileType(file []byte) (FileType, error) {
	var t map[string]interface{}
	if err := yaml.Unmarshal(file, &t); err != nil {
		return Unkown, err
	}

	if _, ok := t["jobs"]; ok {
		return WorkflowType, nil
	}

	if _, ok := t["runs"]; ok {
		return ActionType, nil
	}

	return Unkown, errors.New("can't deserialize, unknown filetype")

}

type Workflow struct {
	Metadata
	Actions   map[string]Action
	Workflows map[string]Workflow
}

func NewWorkflow(name, version string) *Workflow {
	return &Workflow{
		Metadata: Metadata{
			Name:    name,
			Version: version,
		},
		Actions:   map[string]Action{},
		Workflows: map[string]Workflow{},
	}
}

func (w *Workflow) AddWorkflow(add Workflow) {
	w.Workflows[w.key()] = add
}
func (w *Workflow) AddAction(add Action) {
	w.Actions[w.key()] = add
}

type Action struct {
	Metadata
	Actions map[string]Action
}

func (a Action) PopulateDependencies() {
	if len(a.Actions) != 0 {
		return
	}

}

type Metadata struct {
	Version string
	Name    string
}

func (m Metadata) key() string {
	return m.Name + "_" + m.Version
}
