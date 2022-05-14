package constants

var (
	TAG_BEGIN      = "<dopt: version v0.0.0>"
	TAG_END        = "</dopt: version v0.0.0>"
	DOPT_CONFIG    = ".dopt"
	DOCKER_DFT_CMD = "docker buildx build --no-cache --progress=plain --platform linux/amd64 --no-cache -t %s -f Dockerfile ."
)
