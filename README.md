Learning how to render HTML, CSS, and JS with Golang

To bundle the assets we need to use a separate application
export PATH=$PATH:$GOPATH/bin

If there are any changes to your files, rebundle:
rice embed-go

The above command will create a file called rice-box.go in your project.

Tutorial link: https://www.thepolyglotdeveloper.com/2017/03/bundle-html-css-javascript-served-golang-application/

To deploy:
go build && ./testproject