package carbonfootprint

import "errors"

type Params struct {
	CarbonIntensity       float64
	TotalInferenceLatency float64
	FirstLatency          float64
	NextLatency           float64
	TokenSize             int
	TDP                   int
	Mem                   float64
}

const (
	WorldAvgCarbonIntensity = 475 // gCO2e/kWh
	PowerPerGBMem           = 0.1 // ratio in DDR5
	JouleToKWh              = 2.78e-7
)

func CalculateUsageAndEmission(params Params) (energy, carbon float64, err error) {
	if params.TDP == 0 {
		err = errors.New("you need to specify tdp (Thermal Design Power) for the model")
		return 0.0, 0.0, err
	}

	var duration float64
	if params.TotalInferenceLatency != 0.0 {
		duration = params.TotalInferenceLatency
	} else if params.FirstLatency != 0.0 && params.NextLatency != 0.0 {
		duration = params.FirstLatency + params.NextLatency*float64(params.TokenSize-1)
	} else {
		err = errors.New("you need to specify either total time of one inference or both first token latency and next token latency")
		return 0.0, 0.0, err
	}

	if params.CarbonIntensity == 0 {
		params.CarbonIntensity = WorldAvgCarbonIntensity
	}

	energy = (float64(params.TDP) + params.Mem*0.001*PowerPerGBMem) * duration * 0.001 * JouleToKWh
	carbon = energy * params.CarbonIntensity
	return energy, carbon, nil
}
