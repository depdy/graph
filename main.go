package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/goccy/go-graphviz"
	"gopkg.in/yaml.v2"

	"github.com/depdy/graph/github"
)

func renderDOTGraph() ([]byte, error) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	n, err := graph.CreateNode("n")
	if err != nil {
		return nil, err
	}
	m, err := graph.CreateNode("m121211")
	if err != nil {
		return nil, err
	}
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		return nil, err
	}
	e.SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func _main() error {
	graphBytes, err := renderDOTGraph()
	if err != nil {
		return err
	}
	graph, err := graphviz.ParseBytes(graphBytes)
	if err != nil {
		return err
	}
	n, err := graph.Node("n")
	if err != nil {
		return err
	}
	l, err := graph.CreateNode("l3232323")
	if err != nil {
		return err
	}
	e2, err := graph.CreateEdge("e2", n, l)
	if err != nil {
		return err
	}
	e2.SetLabel("e2")
	g := graphviz.New()
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	g.RenderFilename(graph, "png", "static/rw.png")
	return nil
}

func main() {
	// if err := _main(); err != nil {
	// 	log.Fatalf("%+v", err)
	// }
	files, err := readFiles(".", ".*[.]y[a]*ml")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	repo, err := NewRepo(".")
	if err != nil {
		log.Fatal(err)
	}

	r, err := repo.Show("88d21d3", "actions/last/action.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer log.Fatal(r.Close())

	var a github.Xaction
	if err := yaml.NewDecoder(r).Decode(&a); err != nil {
		log.Fatal(err)
	}

	tm, err := github.NewRepository(files)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// It takes a path and a regex pattern, and returns a list of files that match the pattern
	_ = tm

}

func getModels() {

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
