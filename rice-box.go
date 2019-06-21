package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1561074732, 0),

		Content: string("document.getElementById(\"text\").innerHTML = \"Hello World<br /><p>Created by Xavier Olivares</p>\""),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "custom.css",
		FileModTime: time.Unix(1561075032, 0),

		Content: string(".container {\n    width: 400px;\n    height: 100px;\n    background-color: #009ACD;\n    margin: auto;\n}\n\n.text {\n    position: relative;\n    top: 50%;\n    transform: translateY(-50%);\n    text-align: center;\n    color: #FFFFFF;\n}\n\nhtml {\n    background-color: black\n}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1561074191, 0),

		Content: string("<html>\n    <head>\n        <title>The Polyglot Develop</title>\n        <link rel=\"stylesheet\" href=\"custom.css\" />\n    </head>\n    <body>\n        <div class=\"container\">\n            <p id=\"text\" class=\"text\"></p>\n        </div>\n        <script src=\"app.js\"></script>\n    </body>\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1561075032, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "app.js"
			file3, // "custom.css"
			file4, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`website`, &embedded.EmbeddedBox{
		Name: `website`,
		Time: time.Unix(1561075032, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"custom.css": file3,
			"index.html": file4,
		},
	})
}
