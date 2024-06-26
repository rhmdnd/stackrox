// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protoassert"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestReportConfigurationSerialization(t *testing.T) {
	obj := &storage.ReportConfiguration{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertReportConfigurationFromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertReportConfigurationToProto(m)
	assert.NoError(t, err)
	protoassert.Equal(t, obj, conv)
}
