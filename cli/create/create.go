package create

import (
	"errors"
	"path"
	"runtime"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	AppRoot = "internal"

	TemplateDir = "../../template"

	ApplicationDir = "application"
	EntityDir      = "domain/entity"
	RepositoryDir  = "domain/repository"
	PersistenceDir = "infrastructure/persistence"
	http           = "infrastructure/transport/http"
)

var AppName string

var Create = &cobra.Command{
	Use:  "create",
	Args: cobra.MinimumNArgs(2),
	RunE: Run,
}

func Run(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("not enough arguments")
	}
	AppName = args[0]
	name := args[1]
	fs := afero.NewBasePathFs(afero.NewOsFs(), AppRoot+"/")
	if err := createFolders(fs, name); err != nil {
		return err
	}
	if err := createFiles(fs, name); err != nil {
		return err
	}
	return nil
}

func createFolders(fs afero.Fs, name string) error {
	fs.Mkdir(name, 0755)
	dirs := []string{ApplicationDir, EntityDir, RepositoryDir, PersistenceDir, http}
	for _, dir := range dirs {
		if err := fs.MkdirAll(path.Join(name, dir), 0755); err != nil {
			return err
		}
	}
	return nil
}

func createFiles(fs afero.Fs, name string) error {
	createFile(fs, name, path.Join(TemplateDir, "application.stub"), path.Join(name, ApplicationDir, name+".go"))
	createFile(fs, name, path.Join(TemplateDir, "entity.stub"), path.Join(name, EntityDir, name+".go"))
	createFile(fs, name, path.Join(TemplateDir, "repository.stub"), path.Join(name, RepositoryDir, name+".go"))
	createFile(fs, name, path.Join(TemplateDir, "persistence.stub"), path.Join(name, PersistenceDir, name+".go"))
	createFile(fs, name, path.Join(TemplateDir, "handler.stub"), path.Join(name, http, "handler.go"))
	createFile(fs, name, path.Join(TemplateDir, "route.stub"), path.Join(name, http, "route.go"))

	return nil
}

func createFile(fs afero.Fs, name, stubPath, filePath string) error {
	fs.Create(filePath)

	_, filename, _, _ := runtime.Caller(1)
	stubPath = path.Join(path.Dir(filename), stubPath)

	contents, err := fileContents(stubPath)
	if err != nil {
		return err
	}
	contents = replaceStub(contents, name)

	if err := overwrite(AppRoot+"/"+filePath, contents); err != nil {
		return err
	}
	return nil
}

func fileContents(file string) (string, error) {
	a := afero.NewOsFs()
	contents, err := afero.ReadFile(a, file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func overwrite(file string, message string) error {
	a := afero.NewOsFs()
	return afero.WriteFile(a, file, []byte(message), 0666)
}

func replaceStub(content string, name string) string {

	content = strings.Replace(content, "{{TitleName}}", Title(name), -1)
	content = strings.Replace(content, "{{PluralLowerName}}", Lower(Plural(name)), -1)
	content = strings.Replace(content, "{{SingularLowerName}}", Lower(Singular(name)), -1)
	content = strings.Replace(content, "{{SingularCapitalName}}", UpperCamelCase(Singular(name)), -1)
	content = strings.Replace(content, "{{PluralCapitalName}}", UpperCamelCase(Plural(name)), -1)
	content = strings.Replace(content, "{{AppName}}", AppName, -1)
	content = strings.Replace(content, "{{AppRoot}}", AppRoot, -1)
	return content
}

func Plural(name string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Plural(name)
}

func Singular(name string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Singular(name)
}

func Lower(name string) string {
	return strings.ToLower(name)
}

func Title(name string) string {
	return strings.Title(Lower(name))
}
func UpperCamelCase(name string) string {
	return strings.Title(name)
}
