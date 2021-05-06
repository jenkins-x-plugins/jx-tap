module github.com/jenkins-x-plugins/jx-tap

require (
	github.com/cpuguy83/go-md2man v1.0.10
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gomarkdown/markdown v0.0.0-20201024011455-45c732cc8a6b
	github.com/jenkins-x/go-scm v1.8.2
	github.com/jenkins-x/jx-helpers/v3 v3.0.111
	github.com/jenkins-x/jx-logging/v3 v3.0.6
	github.com/mpontillo/tap13 v1.0.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	sigs.k8s.io/structured-merge-diff/v4 v4.0.3 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.2
	k8s.io/client-go => k8s.io/client-go v0.20.2
)

go 1.15
