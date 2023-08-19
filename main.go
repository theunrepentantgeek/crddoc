package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/theunrepentantgeek/crddoc/internal/generator"
	"github.com/theunrepentantgeek/crddoc/internal/model"
)

func main() {
	log := CreateLogger()
	pkg := model.NewPackage(log)

	//err := pkg.LoadDirectory("C:\\GitHub\\azure-service-operator\\v2\\api\\network\\v1api20201101")
	//err := pkg.LoadDirectory("C:\\GitHub\\azure-service-operator\\v2\\api\\compute\\v1api20220301")
	err := pkg.LoadDirectory("C:\\GitHub\\azure-service-operator\\v2\\api\\containerservice\\v1api20210501")
	if err != nil {
		fmt.Println(err.Error())
	}

	gen := generator.New(log)
	err = gen.LoadTemplates("C:\\GitHub\\crddoc\\templates\\crd")
	if err != nil {
		fmt.Println(err.Error())
	}

	// Render the template and write to output.md
	f, err := os.Create("output.md")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	err = gen.Generate(pkg, w)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Flush()
}

func CreateLogger() logr.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05.999",
	}

	zl := zerolog.New(output).
		With().Timestamp().
		Logger()

	// Use standard interface for logging
	zerologr.VerbosityFieldName = "" // Don't include verbosity in output
	log := zerologr.New(&zl)

	return log
}
