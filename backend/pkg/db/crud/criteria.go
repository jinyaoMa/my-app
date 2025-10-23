package crud

import (
	"regexp"
	"strings"

	"gorm.io/gorm"
)

const (
	CriteriaMaxSize           = 20
	CriteriaIncludesMaxLength = 5
	CriteriaJoinsMaxLength    = 5
	CriteriaSortsMaxLength    = 2
)

var regField = regexp.MustCompile(`^([a-zA-Z_]|[a-zA-Z_][a-zA-Z0-9_\.]*[a-zA-Z0-9_])$`)

func FixLikeKeyword(param string) string {
	keepEscape := strings.ReplaceAll(param, "\\", "\\\\")
	keepPercentage := strings.ReplaceAll(keepEscape, "%", "\\%")
	keepUnderscore := strings.ReplaceAll(keepPercentage, "_", "\\_")
	return keepUnderscore
}

type Criteria struct {
	Choice
	Size     int      `json:"size" required:"true" doc:"Page Size"`
	Page     int      `json:"page" required:"true" doc:"Page Number"`
	Joins    Joins    `json:"joins" required:"false" doc:"Join 121/belong2 Associations"`
	Filters  Filters  `json:"filters" required:"false" doc:"List Filter Conditions"`
	Sorts    Sorts    `json:"sorts" required:"false" doc:"List Sort Conditions"`
	Includes Includes `json:"includes" required:"false" doc:"Included Associations"`
}

func (c *Criteria) Apply(tx *gorm.DB, options Options) (newTx *gorm.DB, total int64, err error) {
	tx = c.Includes.Apply(tx,
		options.CriteriaMaxSize,
		options.CriteriaIncludesMaxLength,
		options.CriteriaSortsMaxLength)
	tx = c.Joins.Apply(tx, options.CriteriaJoinsMaxLength)
	tx = c.Filters.Apply(tx, options.FilterParamsMaxLength)
	tx = c.Sorts.Apply(tx, options.CriteriaSortsMaxLength)

	err = tx.Count(&total).Error
	if err != nil {
		return
	}

	tx = tx.Limit(c.Limit()).Offset(c.Offset())
	return c.Choice.Apply(tx), total, nil
}

func (c *Criteria) Limit() int {
	if c.Size < 1 || c.Size > CriteriaMaxSize {
		return 1
	}
	return c.Size
}

func (c *Criteria) Offset() int {
	page := max(c.Page, 1)
	return c.Limit() * (page - 1)
}
