package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1561128309, 0),

		Content: string("document.getElementById(\"text\").innerHTML =\n  \"Hello World<br /><p>Created by Xavier Olivares</p>\";\ndocument.getElementById(\"test\").innerHTML = \"Hi There<br/><p>Adding more stuff to test</p>\";\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "custom.css",
		FileModTime: time.Unix(1561128356, 0),

		Content: string(".container {\n    width: 400px;\n    height: 400px;\n    background-color: #009ACD;\n    margin: auto;\n    margin-top: auto;\n}\n\n.text {\n    position: relative;\n    top: 50%;\n    transform: translateY(-50%);\n    text-align: center;\n    color: #FFFFFF;\n}\n.test {\n    position: relative;\n    top: 50%;\n    transform: translateY(-50%);\n    text-align: center;\n    color:red;\n}\n\nhtml {\n    background-color: black\n}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1561128241, 0),

		Content: string("<html>\n    <head>\n        <title>Golang Practice</title>\n        <link rel=\"stylesheet\" href=\"custom.css\" />\n    </head>\n    <body>\n        <div class=\"container\">\n            <p id=\"text\" class=\"text\"></p>\n            <p id=\"test\" class=\"test\"></p>\n        </div>\n        <script src=\"app.js\"></script>\n    </body>\n</html>"),
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
