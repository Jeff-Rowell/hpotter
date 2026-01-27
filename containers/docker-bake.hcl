group "default" {
  targets = ["http", "react2shell"]
}

target "http" {
  context = "./http"
  dockerfile = "Dockerfile"
  args = {
    GO_VERSION = "1.25.3"
  }
  tags = ["hpotter/http:latest"]
}

target "react2shell" {
  context = "react2shell"
  dockerfile = "Dockerfile"
  args = {
    NODE_VERSION = "25.5.0"
  }
  tags = ["hpotter/react2shell:latest"]
}
