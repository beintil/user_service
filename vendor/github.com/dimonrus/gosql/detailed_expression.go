package gosql

// expression with detail
type detailedExpression struct {
	// expression
	expression expression
	// detail
	detail string
	// is detail first order
	rightAlign bool
}

// SetRightAlign set detail first
func (d *detailedExpression) SetRightAlign(isRight bool) *detailedExpression {
	d.rightAlign = isRight
	return d
}

// Expression get expression
func (d *detailedExpression) Expression() *expression {
	return &d.expression
}

// SetDetail set detail
func (d *detailedExpression) SetDetail(detail string) *detailedExpression {
	d.detail = detail
	return d
}

// GetDetail get detail
func (d *detailedExpression) GetDetail() string {
	return d.detail
}

// ResetDetail reset detail
func (d *detailedExpression) ResetDetail() *detailedExpression {
	d.detail = ""
	return d
}

// IsEmpty check if empty
func (d *detailedExpression) IsEmpty() bool {
	return d == nil || (d.detail == "" && d.expression.Len() == 0)
}

// String render detailed expression
func (d *detailedExpression) String() string {
	if d.IsEmpty() || d.expression.Len() == 0 {
		return ""
	}
	detail := d.detail
	if d.rightAlign {
		if d.detail != "" {
			detail = detail + " "
		}
		return detail + "(" + d.expression.String(", ") + ")"
	}
	if d.detail != "" {
		detail = " " + detail
	}
	return "(" + d.expression.String(", ") + ")" + detail
}

// NewDetailedExpression init detailed expression
func NewDetailedExpression() *detailedExpression {
	return &detailedExpression{}
}
