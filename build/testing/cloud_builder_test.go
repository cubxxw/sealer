package testing

import (
	"testing"

	"github.com/alibaba/sealer/build"
)

func TestCloudBuilder_Build(t *testing.T) {
	conf := &build.Config{}
	builder := build.NewBuilder(conf, "")
	err := builder.Build("dashboard-test:latest", ".", "kubefile")
	if err != nil {
		t.Errorf("exec build error %v\n", err)
	}
}