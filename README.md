go get -u github.com/golang/lint/golint 
go get -u github.com/alecthomas/gometalinter

gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow src/joke/   

/home/albina/go/src/gitlab.com/jokebot/src/joke/joke.go:26:23:warning: error return value not checked (defer resp.Body.Close()) (errcheck)
/home/albina/go/src/gitlab.com/jokebot/src/joke/joke_test.go:14:10:warning: error return value not checked (w.Write([]byte(resp))) (errcheck)

