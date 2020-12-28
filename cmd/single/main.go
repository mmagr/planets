package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mmagr/planets/internal/config"
	"github.com/mmagr/planets/internal/model"
	"github.com/mmagr/planets/internal/model/conditions"
	"github.com/mmagr/planets/internal/service"
	"github.com/spf13/cobra"
)

func main() {
	cmdForecast := &cobra.Command{
		Use:   "forecast [period to forecast in years]",
		Short: "Weather forecast summary for given period, starting at d0",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a period argument")
			}

			if _, err := strconv.Atoi(args[0]); err != nil {
				return errors.New("period must be a valid integer")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			days, err := strconv.Atoi(args[0])
			if err != nil {
				// just as precaution - should have been already checked by `args`
				fmt.Println("period must be an integer")
				return
			}

			// BUG: get year size from user - this is assuming the longest from problem definition
			summary(days * 360)
		},
	}

	rootCmd := &cobra.Command{Use: "planets"}
	rootCmd.AddCommand(cmdForecast)
	rootCmd.Execute()
}

func summary(days int) {
	parsedConfig := config.Init()
	p1 := config.Planet(parsedConfig.Sub("planet.vulcano"))
	p2 := config.Planet(parsedConfig.Sub("planet.ferengi"))
	p3 := config.Planet(parsedConfig.Sub("planet.betasoide"))
	calculator := service.NewClimatempo(p1, p2, p3, service.ILineFactory{}, service.TriangleFactory{})

	ctx := model.BatchContext{}
	for day := 0; day <= days; day++ {
		condition, metric := calculator.ConditionsOn(day)
		ctx.Observe(day, condition, metric)
	}

	fmt.Print(ctxToString(ctx))
}

// formats the context as string
func ctxToString(ctx model.BatchContext) string {
	var buffer strings.Builder
	for k, v := range ctx.Conditions {
		fmt.Fprintf(&buffer, "%s - %d periods\n", k, len(v))
		if k == conditions.Rain {
			tally := model.Tally{}
			for _, p := range v {
				tally.Observe(p.Tally.Day, p.Tally.Metric)
			}
			fmt.Fprintf(&buffer, "Maximum rain happened on %d (perimeter %.2f)\n", tally.Day, tally.Metric)
		}
	}
	return buffer.String()
}
