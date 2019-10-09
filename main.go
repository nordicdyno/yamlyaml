package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	flag.Parse()

	var input io.Reader
	switch flag.NArg() {
	case 0:
		input = os.Stdin
	case 1:
		yamlFile, err := os.Open(flag.Arg(0))
		checkerr(err)
		defer yamlFile.Close()
		input = yamlFile
	default:
		fmt.Println("input must be from stdin or file")
		os.Exit(1)
	}

	dec := yaml.NewDecoder(input)

	for n := 1; ; n++ {
		var doc yaml.Node
		err := dec.Decode(&doc)
		if err != nil {
			if err != io.EOF {
				checkerr(err, fmt.Sprintf("DECODE ERROR: (type=%T)", err))
			}
			break
		}

		root := doc.Content[0]
		scan(root)

		out := os.Stdout
		enc := yaml.NewEncoder(out)
		enc.SetIndent(2)
		fmt.Fprintln(out, "---")
		err = enc.Encode(root)
		checkerr(err)
	}
}

func scan(n *yaml.Node) {
	switch n.Kind {
	case yaml.ScalarNode:
		tryToUnYAML(n)
	case yaml.MappingNode:
		scanMap(n)
	}
}

func scanMap(n *yaml.Node) {
	for i, item := range n.Content {
		if n.Kind == yaml.MappingNode && (i%2 == 0) {
			continue
		}
		scan(item)
	}
}

func tryToUnYAML(n *yaml.Node) {
	var item yaml.Node
	if err := yaml.Unmarshal([]byte(n.Value), &item); err != nil {
		// fmt.Printf("unmarshal failed with error: %v\n", err.Error())
		return
	}

	if item.Content[0].Kind == yaml.MappingNode {
		scanMap(item.Content[0])
	}
	*n = *item.Content[0]
}

func checkerr(err error, a ...interface{}) {
	if err == nil {
		return
	}
	if len(a) > 0 {
		fmt.Print(a...)
	} else {
		fmt.Print("ERROR: ")
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
