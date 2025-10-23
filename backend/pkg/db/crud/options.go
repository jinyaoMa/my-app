package crud

type Options struct {
	CriteriaMaxSize           int `json:"criteriaMaxSize"`
	CriteriaIncludesMaxLength int `json:"criteriaIncludesMaxLength"`
	CriteriaJoinsMaxLength    int `json:"criteriaJoinsMaxLength"`
	CriteriaSortsMaxLength    int `json:"criteriaSortsMaxLength"`
	FilterParamsMaxLength     int `json:"filterParamsMaxLength"`
}

func DefaultOptions() Options {
	return Options{
		CriteriaMaxSize:           CriteriaMaxSize,
		CriteriaIncludesMaxLength: CriteriaIncludesMaxLength,
		CriteriaJoinsMaxLength:    CriteriaJoinsMaxLength,
		CriteriaSortsMaxLength:    CriteriaSortsMaxLength,
		FilterParamsMaxLength:     FilterParamsMaxLength,
	}
}
