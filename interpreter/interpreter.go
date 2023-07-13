package interpreter

type Expression interface {
	Interpret() int
}

// 数字表达式
type NumberExpression struct {
	Number int
}

func NewNumberExpression(number int) *NumberExpression {
	return &NumberExpression{Number: number}
}

func (ne *NumberExpression) Interpret() int {
	return ne.Number
}

// 加法表达式
type AddExpression struct {
	Left  Expression
	Right Expression
}

func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{Left: left, Right: right}
}

func (ae *AddExpression) Interpret() int {
	return ae.Left.Interpret() + ae.Right.Interpret()
}

// 减法表达式
type SubExpression struct {
	Left  Expression
	Right Expression
}

func NewSubExpression(left, right Expression) *SubExpression {
	return &SubExpression{Left: left, Right: right}
}

func (se *SubExpression) Interpret() int {
	return se.Left.Interpret() - se.Right.Interpret()
}

// 乘法表达式
type MulExpression struct {
	Left  Expression
	Right Expression
}

func NewMulExpression(left, right Expression) *MulExpression {
	return &MulExpression{Left: left, Right: right}
}

func (me *MulExpression) Interpret() int {
	return me.Left.Interpret() * me.Right.Interpret()
}

// 除法表达式
type DivExpression struct {
	Left  Expression
	Right Expression
}

func NewDivExpression(left, right Expression) *DivExpression {
	return &DivExpression{Left: left, Right: right}
}

func (de *DivExpression) Interpret() int {
	return de.Left.Interpret() / de.Right.Interpret()
}
