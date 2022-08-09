package java

import (
	"code-analysis/java/parser"
	"code-analysis/utils"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type AnnotationListener struct {
	parser.BaseJavaParserListener
	indent int
	path   string
}

func NewAnnotationListener(path string) *AnnotationListener {
	return &AnnotationListener{
		path: path,
	}
}

func printIndent(s *AnnotationListener, sign bool) {
	if !sign {
		s.indent--
	}

	builder := strings.Builder{}
	for i := 0; i < s.indent; i++ {
		builder.WriteString("\t")
	}
	fmt.Printf(builder.String())

	if sign {
		s.indent++
	}
}

func (s *AnnotationListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	printIndent(s, true)
	fmt.Printf("进入注解:%s\n", ctx.GetText())
}

func (s *AnnotationListener) ExitAnnotation(ctx *parser.AnnotationContext) {
	printIndent(s, false)
	fmt.Printf("退出注解:%s\n", ctx.GetText())
}

func (s *AnnotationListener) EnterElementValue(ctx *parser.ElementValueContext) {
	printIndent(s, true)
	fmt.Printf("进入数据:%s\n", ctx.GetText())
}

func (s *AnnotationListener) ExitElementValue(ctx *parser.ElementValueContext) {
	printIndent(s, false)
	fmt.Printf("退出数据:%s\n", ctx.GetText())
}

func (s *AnnotationListener) EnterElementValuePairs(ctx *parser.ElementValuePairsContext) {
	printIndent(s, true)
	fmt.Printf("进入数据kv注解:%s\n", ctx.GetText())
}

func (s *AnnotationListener) ExitElementValuePairs(ctx *parser.ElementValuePairsContext) {
	printIndent(s, false)
	fmt.Printf("退出数据kv注解:%s\n", ctx.GetText())
}

func ParseJava(path string) {
	extractor := NewAnnotationListener(path)
	data := utils.ReadFile(extractor.path)
	stream := antlr.NewInputStream(data)
	lexer := parser.NewJavaLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	javaParser := parser.NewJavaParser(tokens)
	tree := javaParser.CompilationUnit()

	antlr.NewParseTreeWalker().Walk(extractor, tree)
}
