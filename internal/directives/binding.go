package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	en := en2.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}

	err = validate.Var(val, constraint)
	if err != nil {
		fieldName := graphql.GetPathContext(ctx).Field
		validationErrors := err.(validator.ValidationErrors)
		transErr := fmt.Errorf("%s%+v", *fieldName, validationErrors[0].Translate(trans))
		return val, transErr
	}

	return val, nil
}
