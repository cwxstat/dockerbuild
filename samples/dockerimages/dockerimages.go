package dockerimages

var golang = `FROM golang:latest AS build

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

`

// FIXME: make meaningful
func Images(lang string) string {
	if lang == "golang" {
		return golang
	}
	return golang
}
