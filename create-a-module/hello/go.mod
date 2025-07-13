module github.com/28604/go-doc-tutorial/create-a-module/hello

go 1.24.4

replace github.com/28604/go-doc-tutorial/create-a-module/greetings => ../greetings

require github.com/28604/go-doc-tutorial/create-a-module/greetings v0.0.0-00010101000000-000000000000
