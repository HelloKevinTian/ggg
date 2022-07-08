package generator

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

type PackageInfo struct {
	ProjectName string
}

var dir = "generator/"

var outPath = dir + "out/main.go"
var f1 = dir + "tpl/main.go.tpl"
var f2 = dir + "tpl/go.mod.tpl"

func Test() {
	f, err := os.Open(f1)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fd, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(fd))

	// p := PackageInfo{ProjectName: "sdk-ddns"}

	// outFile, err := os.Open(outPath)
	// if err != nil {
	// 	panic(err)
	// }

	// t1, _ := template.ParseFiles(f1, f2)
	// t1.ExecuteTemplate(os.Stdout, "go.mod", p)

	// t2, _ := template.New("test2").Parse(string(fd))
	// t2.ExecuteTemplate(outFile, "test2", p)
	// // t2.Execute(os.Stdout, p)

	data := PackageInfo{ProjectName: "sdk-ddns"}
	task := &Task{Name: "main.go", Path: dir + "out/main.go", Text: string(fd)}
	ff, err := task.Render(data)
	if err != nil {
		panic(err)
	}
	files := []*File{}
	files = append(files, ff)
	CreateDir(files, dir+"out")
}

type File struct {
	Name    string
	Content string
}

type Task struct {
	Name string
	Path string
	Text string
	*template.Template
}

var funcs = map[string]interface{}{
	"ToLower": strings.ToLower,
	// "LowerFirst": util.LowerFirst,
	// "UpperFirst": util.UpperFirst,
	// "NotPtr":     util.NotPtr,
	// "HasFeature": HasFeature,
}

// Build .
func (t *Task) Build() error {
	x, err := template.New(t.Name).Funcs(funcs).Parse(t.Text)
	if err != nil {
		return err
	}
	t.Template = x
	return nil
}

func (t *Task) Render(data interface{}) (*File, error) {
	if t.Template == nil {
		err := t.Build()
		if err != nil {
			return nil, err
		}
	}

	var buf bytes.Buffer
	err := t.ExecuteTemplate(&buf, t.Name, data)
	if err != nil {
		return nil, err
	}
	return &File{t.Path, buf.String()}, nil
}

func CreateDir(files []*File, dir string) string {
	os.RemoveAll(dir)

	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.Content)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}
