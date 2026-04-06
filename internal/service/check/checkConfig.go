package check

func (cs *CheckService) CheckConfig(config map[string]interface{}) ([]string, []error) {
	return cs.checker.CheckConfig(config)
}
