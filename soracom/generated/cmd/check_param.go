package cmd

import (
	"fmt"
)

func checkIfRequiredStringParameterIsSupplied(propName, optionName, in string, parsedBody interface{}, varValue string) error {
	if in == "body" {
		contains := doesBodyContainParameter(parsedBody, propName)
		if !contains && varValue == "" {
			return fmt.Errorf("required parameter '%s' in body (or command line option '%s') is not specified", propName, optionName)
		}
		return nil
	}

	if varValue == "" {
		return fmt.Errorf("required parameter '%s' is not specified", optionName)
	}
	return nil
}

//lint:ignore U1000 we want to keep this function for the future use.
func checkIfRequiredStringSliceParameterIsSupplied(propName, optionName, in string, parsedBody interface{}, varValue []string) error {
	if in == "body" {
		contains := doesBodyContainParameter(parsedBody, propName)
		if !contains && len(varValue) == 0 {
			return fmt.Errorf("required parameter '%s' in body (or command line option '%s') is not specified", propName, optionName)
		}
		return nil
	}

	if len(varValue) == 0 {
		return fmt.Errorf("required parameter '%s' is not specified", optionName)
	}
	return nil
}

func checkIfRequiredIntegerParameterIsSupplied(propName, optionName, in string, parsedBody interface{}, varValue int64) error {
	if in == "body" {
		contains := doesBodyContainParameter(parsedBody, propName)
		if !contains && varValue == 0 {
			return fmt.Errorf("required parameter '%s' in body (or command line option '%s') is not specified", propName, optionName)
		}
		return nil
	}

	if varValue == 0 {
		return fmt.Errorf("required parameter '%s' is not specified", optionName)
	}
	return nil
}

func checkIfRequiredFloatParameterIsSupplied(propName, optionName, in string, parsedBody interface{}, varValue float64) error {
	if in == "body" {
		contains := doesBodyContainParameter(parsedBody, propName)
		if !contains && varValue == 0.0 {
			return fmt.Errorf("required parameter '%s' in body (or command line option '%s') is not specified", propName, optionName)
		}
		return nil
	}

	if varValue == 0.0 {
		return fmt.Errorf("required parameter '%s' is not specified", optionName)
	}
	return nil
}

//lint:ignore U1000 we want to keep this function for the future use.
func checkIfRequiredBoolParameterIsSupplied(propName, optionName, in string, parsedBody interface{}, varValue bool) error {
	if in == "body" {
		contains := doesBodyContainParameter(parsedBody, propName)
		if !contains && !varValue {
			return fmt.Errorf("required parameter '%s' in body (or command line option '%s') is not specified", propName, optionName)
		}
		return nil
	}

	if !varValue {
		return fmt.Errorf("required parameter '%s' is not specified", optionName)
	}
	return nil
}

func doesBodyContainParameter(parsedBody interface{}, parameterName string) bool {
	m, ok := parsedBody.(map[string]interface{})
	if !ok {
		return false
	}

	_, found := m[parameterName]
	return found
}
