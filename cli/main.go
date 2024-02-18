package main

import (
	"flag"
	"fmt"
	"github.com/greenscale-ai/genai-carbon-footprint/carbonfootprint"
	"os"
)

func main() {
	carbonIntensity := flag.Float64("carbon-intensity", carbonfootprint.WorldAvgCarbonIntensity, "carbon intensity (gCO2e/kWh) of electricity of your country or cloud provider region (default: 475 - world average)")
	totalInferenceLatency := flag.Float64("total-inference-latency", 0.0, "total time of one inference procedure in mini-seconds")
	firstLatency := flag.Float64("first-latency", 0.0, "first token latency in mini-seconds")
	nextLatency := flag.Float64("next-latency", 0.0, "next token latency in mini-seconds")
	tokenSize := flag.Int("token-size", 32, "output token number in one inference (default: 32)")
	tdp := flag.Int("tdp", 0, "device TDP (Thermal Design Power) in Watts, it could be CPU/GPU/Accelerators")
	mem := flag.Float64("mem", 0.0, "memory consumption in MB")
	flag.Parse()

	params := carbonfootprint.Params{
		CarbonIntensity:       *carbonIntensity,
		TotalInferenceLatency: *totalInferenceLatency,
		FirstLatency:          *firstLatency,
		NextLatency:           *nextLatency,
		TokenSize:             *tokenSize,
		TDP:                   *tdp,
		Mem:                   *mem,
	}

	energy, carbon, err := carbonfootprint.CalculateUsageAndEmission(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Total energy usage of one inference (kWh):", energy)
	fmt.Println("Carbon emission in one inference (gCO2e):", carbon)
}
