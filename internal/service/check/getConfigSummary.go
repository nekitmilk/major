package check

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors/serviceErrorsModels"
	"major/pkg/parser"
)

func (cs *CheckService) GetConfigSummary(configReq *models.CheckRequest) (*models.CheckResponseFullData, error) {
	parsedConfig, errParse := parser.Parse([]byte(configReq.Config))
	if errParse != nil {
		return nil, serviceErrorsModels.NewParsingError(errParse, "failed to parse config")
	}

	triggeredIds, checkErrors := cs.checker.CheckConfig(parsedConfig)

	var checkErrorsStr []string
	if checkErrors != nil {
		checkErrorsStr = make([]string, 0, len(checkErrors))
		for _, checkError := range checkErrors {
			checkErrorsStr = append(checkErrorsStr, checkError.Error())
		}
	}

	rules := make([]*models.RuleFullData, 0, len(triggeredIds))
	for _, idRule := range triggeredIds {
		rule, errGetRule := cs.repo.RulesRepository.GetRuleFullDataByID(idRule)
		if errGetRule != nil {
			continue
		}
		rules = append(rules, rule)
	}

	response := &models.CheckResponseFullData{
		Rules:  rules,
		Errors: checkErrorsStr,
	}

	return response, nil
}
