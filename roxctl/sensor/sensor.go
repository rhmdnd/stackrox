package sensor

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/stackrox/rox/roxctl/common/flags"
	"github.com/stackrox/rox/roxctl/sensor/generate"
	"github.com/stackrox/rox/roxctl/sensor/generatecerts"
	"github.com/stackrox/rox/roxctl/sensor/getbundle"
)

// Command controls all of the functions being applied to a sensor
func Command() *cobra.Command {
	c := &cobra.Command{
		Use: "sensor",
	}
	c.AddCommand(
		generate.Command(),
		getbundle.Command(),
		generatecerts.Command(),
	)
	flags.AddTimeoutWithDefault(c, 30*time.Second)
	return c
}
