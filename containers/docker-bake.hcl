group "default" {
  targets = ["http"]
}

target "http" {
  context = "./http"
  dockerfile = "Dockerfile"
  args = {
    GO_VERSION = "1.25.3"
  }
  tags = ["hpotter/http:latest"]
}
