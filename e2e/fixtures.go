package e2e

func dockerfile() []byte {
	return []byte(`FROM golang:latest AS build

WORKDIR /project

# Copy the entire project and build it

COPY . /project
RUN go mod tidy
RUN go build -o /bin/project


# FROM gcr.io/distroless/static
# Below for debugging
FROM golang:latest
ENV TZ=America/New_York
COPY --from=build /bin/project /bin/project
ENV TZ=America/New_York
ENTRYPOINT ["/bin/project"]
# Args to project
#CMD []



`)
}

func goMod() []byte {
	return []byte(`module example.com/m/v2

	go 1.18
	`)
}

func goProg() []byte {
	return []byte(`
package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for {
		fmt.Printf("count: %d\n", count)
		count++
		time.Sleep(1 * time.Second)
	}
}
	

`)
}
