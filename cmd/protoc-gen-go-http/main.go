package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var (
		flags        flag.FlagSet
		importPrefix = flags.String("import_prefix", "", "prefix to prepend to import paths")
	)

	gen := protogen.Options{
		ParamFunc:         flags.Set,
		ImportRewriteFunc: func(path protogen.GoImportPath) protogen.GoImportPath {
			switch path {
			case "context", "fmt", "math":
				return path
			}

			if *importPrefix != "" {
				return protogen.GoImportPath(*importPrefix) + path
			}

			return path
		},
	}

	gen.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = 0

		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			GenerateFile(plugin, file)
		}

		return nil
	})
}