package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ThermalSimulationParameters thermal simulation parameters
// swagger:model ThermalSimulationParameters
type ThermalSimulationParameters struct {

	// should be a number between 0.5 and 1.5
	BladeCrashThreshold float64 `json:"bladeCrashThreshold,omitempty"`

	// detect blade crash
	// Required: true
	DetectBladeCrash *bool `json:"detectBladeCrash"`

	// should be a number between 0.5 and 1.5 mm
	DynamicVirtualSensorRadius float64 `json:"dynamicVirtualSensorRadius,omitempty"`

	// elastic modulus
	// Required: true
	ElasticModulus *float64 `json:"elasticModulus"`

	// Must be between 0.00001 to 0.001 meters
	// Required: true
	// Maximum: 0.001
	// Minimum: 1e-05
	HatchSpacing *float64 `json:"hatchSpacing"`

	// false indicates that only the thermal solver will run, while true indicates that the mechanics solver will run after the thermal solver
	// Required: true
	IncludeStressAnalysis *bool `json:"includeStressAnalysis"`

	// Must be between 10 to 1000 watts
	// Required: true
	// Maximum: 1000
	// Minimum: 10
	LaserWattage *float64 `json:"laserWattage"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	LayerRotationAngle *float64 `json:"layerRotationAngle"`

	// Must be between 0.00001 to 0.0001 meters
	// Required: true
	// Maximum: 0.0001
	// Minimum: 1e-05
	LayerThickness *float64 `json:"layerThickness"`

	// Must be between 0 to 0.005 meters, Must be greater than minimumWallDistance
	// Required: true
	// Maximum: 0.005
	// Minimum: 0
	MaximumWallDistance *float64 `json:"maximumWallDistance"`

	// Must be between 0.00015 to 0.002 meters, Must be greater than minimumWallThickness
	// Required: true
	// Maximum: 0.002
	// Minimum: 0.00015
	MaximumWallThickness *float64 `json:"maximumWallThickness"`

	// Distance to move the part off the base plate for supports, Must be between 0 to 0.005 meters
	// Required: true
	// Maximum: 0.005
	// Minimum: 0
	MinimumSupportHeight *float64 `json:"minimumSupportHeight"`

	// Must be between 0 to 0.0003 meters, Must be less than maximumWallDistance
	// Required: true
	// Maximum: 0.0003
	// Minimum: 0
	MinimumWallDistance *float64 `json:"minimumWallDistance"`

	// Must be between 0.00005 to 0.0003 meters, Must be less than maximumWallThickness
	// Required: true
	// Maximum: 0.0003
	// Minimum: 5e-05
	MinimumWallThickness *float64 `json:"minimumWallThickness"`

	// output displacement after cutoff
	// Required: true
	OutputDisplacementAfterCutoff *bool `json:"outputDisplacementAfterCutoff"`

	// if true, mechanics solver output will include a zip file with the stress / distortion state at the end of each voxel layer
	// Required: true
	OutputLayerVTK *bool `json:"outputLayerVTK"`

	// for each slectedPoint, a series of vtk files will output thermal history around that point with a radius of staticVirtualSensorRadius.
	// Required: true
	OutputPointThermalHistory *bool `json:"outputPointThermalHistory"`

	// output shrinkage
	// Required: true
	OutputShrinkage *bool `json:"outputShrinkage"`

	// output state map
	// Required: true
	OutputStateMap *bool `json:"outputStateMap"`

	// output thermal vtk
	// Required: true
	OutputThermalVtk *bool `json:"outputThermalVtk"`

	// output thermal vtk layers
	OutputThermalVtkLayers string `json:"outputThermalVtkLayers,omitempty"`

	// Outputs meltpool dimensions and thermal history within a VTK file. Thermal history is output as the average temperature within the user-specified dynamicVirtualSensorRadius, centered at the middle of the laser beam.
	// Required: true
	OutputVirtualSensorData *bool `json:"outputVirtualSensorData"`

	// poisson ratio
	// Required: true
	PoissonRatio *float64 `json:"poissonRatio"`

	// Must be between 0.01 to 10 meters/second
	// Required: true
	// Maximum: 10
	// Minimum: 0.01
	ScanSpeed *float64 `json:"scanSpeed"`

	// List of points where the thermal solver will collect thermal history - limit 10
	SelectedPoints []*SelectedPoint `json:"selectedPoints"`

	// List of parts to simulate (current limit is one part, imposed by server)
	SimulationParts []*SimulationPart `json:"simulationParts"`

	// Must be between 0.001 to 0.1 meters
	// Required: true
	// Maximum: 0.1
	// Minimum: 0.001
	SlicingStripeWidth *float64 `json:"slicingStripeWidth"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	StartingLayerAngle *float64 `json:"startingLayerAngle"`

	// should be a number between 0.05 and 5.0 mm
	StaticVirtualSensorRadius float64 `json:"staticVirtualSensorRadius,omitempty"`

	// strain scaling factor
	// Required: true
	StrainScalingFactor *float64 `json:"strainScalingFactor"`

	// stress mode
	// Required: true
	StressMode *string `json:"stressMode"`

	// Must be between 1 to 89 degrees
	// Required: true
	// Maximum: 89
	// Minimum: 1
	SupportAngle *float64 `json:"supportAngle"`

	// Multiplier for support calculations, Must be between 0.1 to 10
	// Required: true
	// Maximum: 10
	// Minimum: 0.1
	SupportFactorOfSafety *float64 `json:"supportFactorOfSafety"`

	// support optimization
	// Required: true
	SupportOptimization *bool `json:"supportOptimization"`

	// support yield strength
	// Required: true
	SupportYieldStrength *float64 `json:"supportYieldStrength"`

	// support yield strength ratio
	// Required: true
	SupportYieldStrengthRatio *float64 `json:"supportYieldStrengthRatio"`

	// Must be between 0.00002 to 0.002 meters
	// Required: true
	// Maximum: 0.002
	// Minimum: 2e-05
	VoxelSize *float64 `json:"voxelSize"`
}

// Validate validates this thermal simulation parameters
func (m *ThermalSimulationParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetectBladeCrash(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateElasticModulus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHatchSpacing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIncludeStressAnalysis(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLaserWattage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerRotationAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaximumWallDistance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaximumWallThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumSupportHeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumWallDistance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumWallThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputDisplacementAfterCutoff(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputLayerVTK(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputPointThermalHistory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputShrinkage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputStateMap(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputThermalVtk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputVirtualSensorData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePoissonRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScanSpeed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSelectedPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSimulationParts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlicingStripeWidth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartingLayerAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStrainScalingFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStressMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportFactorOfSafety(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportOptimization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportYieldStrength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportYieldStrengthRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoxelSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThermalSimulationParameters) validateDetectBladeCrash(formats strfmt.Registry) error {

	if err := validate.Required("detectBladeCrash", "body", m.DetectBladeCrash); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateElasticModulus(formats strfmt.Registry) error {

	if err := validate.Required("elasticModulus", "body", m.ElasticModulus); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateHatchSpacing(formats strfmt.Registry) error {

	if err := validate.Required("hatchSpacing", "body", m.HatchSpacing); err != nil {
		return err
	}

	if err := validate.Minimum("hatchSpacing", "body", float64(*m.HatchSpacing), 1e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("hatchSpacing", "body", float64(*m.HatchSpacing), 0.001, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateIncludeStressAnalysis(formats strfmt.Registry) error {

	if err := validate.Required("includeStressAnalysis", "body", m.IncludeStressAnalysis); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateLaserWattage(formats strfmt.Registry) error {

	if err := validate.Required("laserWattage", "body", m.LaserWattage); err != nil {
		return err
	}

	if err := validate.Minimum("laserWattage", "body", float64(*m.LaserWattage), 10, false); err != nil {
		return err
	}

	if err := validate.Maximum("laserWattage", "body", float64(*m.LaserWattage), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateLayerRotationAngle(formats strfmt.Registry) error {

	if err := validate.Required("layerRotationAngle", "body", m.LayerRotationAngle); err != nil {
		return err
	}

	if err := validate.Minimum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 180, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateLayerThickness(formats strfmt.Registry) error {

	if err := validate.Required("layerThickness", "body", m.LayerThickness); err != nil {
		return err
	}

	if err := validate.Minimum("layerThickness", "body", float64(*m.LayerThickness), 1e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("layerThickness", "body", float64(*m.LayerThickness), 0.0001, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateMaximumWallDistance(formats strfmt.Registry) error {

	if err := validate.Required("maximumWallDistance", "body", m.MaximumWallDistance); err != nil {
		return err
	}

	if err := validate.Minimum("maximumWallDistance", "body", float64(*m.MaximumWallDistance), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("maximumWallDistance", "body", float64(*m.MaximumWallDistance), 0.005, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateMaximumWallThickness(formats strfmt.Registry) error {

	if err := validate.Required("maximumWallThickness", "body", m.MaximumWallThickness); err != nil {
		return err
	}

	if err := validate.Minimum("maximumWallThickness", "body", float64(*m.MaximumWallThickness), 0.00015, false); err != nil {
		return err
	}

	if err := validate.Maximum("maximumWallThickness", "body", float64(*m.MaximumWallThickness), 0.002, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateMinimumSupportHeight(formats strfmt.Registry) error {

	if err := validate.Required("minimumSupportHeight", "body", m.MinimumSupportHeight); err != nil {
		return err
	}

	if err := validate.Minimum("minimumSupportHeight", "body", float64(*m.MinimumSupportHeight), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumSupportHeight", "body", float64(*m.MinimumSupportHeight), 0.005, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateMinimumWallDistance(formats strfmt.Registry) error {

	if err := validate.Required("minimumWallDistance", "body", m.MinimumWallDistance); err != nil {
		return err
	}

	if err := validate.Minimum("minimumWallDistance", "body", float64(*m.MinimumWallDistance), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumWallDistance", "body", float64(*m.MinimumWallDistance), 0.0003, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateMinimumWallThickness(formats strfmt.Registry) error {

	if err := validate.Required("minimumWallThickness", "body", m.MinimumWallThickness); err != nil {
		return err
	}

	if err := validate.Minimum("minimumWallThickness", "body", float64(*m.MinimumWallThickness), 5e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumWallThickness", "body", float64(*m.MinimumWallThickness), 0.0003, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputDisplacementAfterCutoff(formats strfmt.Registry) error {

	if err := validate.Required("outputDisplacementAfterCutoff", "body", m.OutputDisplacementAfterCutoff); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputLayerVTK(formats strfmt.Registry) error {

	if err := validate.Required("outputLayerVTK", "body", m.OutputLayerVTK); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputPointThermalHistory(formats strfmt.Registry) error {

	if err := validate.Required("outputPointThermalHistory", "body", m.OutputPointThermalHistory); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputShrinkage(formats strfmt.Registry) error {

	if err := validate.Required("outputShrinkage", "body", m.OutputShrinkage); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputStateMap(formats strfmt.Registry) error {

	if err := validate.Required("outputStateMap", "body", m.OutputStateMap); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputThermalVtk(formats strfmt.Registry) error {

	if err := validate.Required("outputThermalVtk", "body", m.OutputThermalVtk); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateOutputVirtualSensorData(formats strfmt.Registry) error {

	if err := validate.Required("outputVirtualSensorData", "body", m.OutputVirtualSensorData); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validatePoissonRatio(formats strfmt.Registry) error {

	if err := validate.Required("poissonRatio", "body", m.PoissonRatio); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateScanSpeed(formats strfmt.Registry) error {

	if err := validate.Required("scanSpeed", "body", m.ScanSpeed); err != nil {
		return err
	}

	if err := validate.Minimum("scanSpeed", "body", float64(*m.ScanSpeed), 0.01, false); err != nil {
		return err
	}

	if err := validate.Maximum("scanSpeed", "body", float64(*m.ScanSpeed), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateSelectedPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.SelectedPoints); i++ {

		if swag.IsZero(m.SelectedPoints[i]) { // not required
			continue
		}

		if m.SelectedPoints[i] != nil {

			if err := m.SelectedPoints[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ThermalSimulationParameters) validateSimulationParts(formats strfmt.Registry) error {

	if swag.IsZero(m.SimulationParts) { // not required
		return nil
	}

	for i := 0; i < len(m.SimulationParts); i++ {

		if swag.IsZero(m.SimulationParts[i]) { // not required
			continue
		}

		if m.SimulationParts[i] != nil {

			if err := m.SimulationParts[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ThermalSimulationParameters) validateSlicingStripeWidth(formats strfmt.Registry) error {

	if err := validate.Required("slicingStripeWidth", "body", m.SlicingStripeWidth); err != nil {
		return err
	}

	if err := validate.Minimum("slicingStripeWidth", "body", float64(*m.SlicingStripeWidth), 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("slicingStripeWidth", "body", float64(*m.SlicingStripeWidth), 0.1, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateStartingLayerAngle(formats strfmt.Registry) error {

	if err := validate.Required("startingLayerAngle", "body", m.StartingLayerAngle); err != nil {
		return err
	}

	if err := validate.Minimum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 180, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateStrainScalingFactor(formats strfmt.Registry) error {

	if err := validate.Required("strainScalingFactor", "body", m.StrainScalingFactor); err != nil {
		return err
	}

	return nil
}

var thermalSimulationParametersTypeStressModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LinearElastic","J2Plasticity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		thermalSimulationParametersTypeStressModePropEnum = append(thermalSimulationParametersTypeStressModePropEnum, v)
	}
}

const (
	// ThermalSimulationParametersStressModeLinearElastic captures enum value "LinearElastic"
	ThermalSimulationParametersStressModeLinearElastic string = "LinearElastic"
	// ThermalSimulationParametersStressModeJ2Plasticity captures enum value "J2Plasticity"
	ThermalSimulationParametersStressModeJ2Plasticity string = "J2Plasticity"
)

// prop value enum
func (m *ThermalSimulationParameters) validateStressModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, thermalSimulationParametersTypeStressModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ThermalSimulationParameters) validateStressMode(formats strfmt.Registry) error {

	if err := validate.Required("stressMode", "body", m.StressMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateStressModeEnum("stressMode", "body", *m.StressMode); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateSupportAngle(formats strfmt.Registry) error {

	if err := validate.Required("supportAngle", "body", m.SupportAngle); err != nil {
		return err
	}

	if err := validate.Minimum("supportAngle", "body", float64(*m.SupportAngle), 1, false); err != nil {
		return err
	}

	if err := validate.Maximum("supportAngle", "body", float64(*m.SupportAngle), 89, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateSupportFactorOfSafety(formats strfmt.Registry) error {

	if err := validate.Required("supportFactorOfSafety", "body", m.SupportFactorOfSafety); err != nil {
		return err
	}

	if err := validate.Minimum("supportFactorOfSafety", "body", float64(*m.SupportFactorOfSafety), 0.1, false); err != nil {
		return err
	}

	if err := validate.Maximum("supportFactorOfSafety", "body", float64(*m.SupportFactorOfSafety), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateSupportOptimization(formats strfmt.Registry) error {

	if err := validate.Required("supportOptimization", "body", m.SupportOptimization); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateSupportYieldStrength(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrength", "body", m.SupportYieldStrength); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateSupportYieldStrengthRatio(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrengthRatio", "body", m.SupportYieldStrengthRatio); err != nil {
		return err
	}

	return nil
}

func (m *ThermalSimulationParameters) validateVoxelSize(formats strfmt.Registry) error {

	if err := validate.Required("voxelSize", "body", m.VoxelSize); err != nil {
		return err
	}

	if err := validate.Minimum("voxelSize", "body", float64(*m.VoxelSize), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("voxelSize", "body", float64(*m.VoxelSize), 0.002, false); err != nil {
		return err
	}

	return nil
}
