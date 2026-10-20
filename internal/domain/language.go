package domain

import (
	"fmt"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

type Language string

const (
	LanguageGo         = Language("go")
	LanguagePython     = Language("py")
	LanguageJava       = Language("java")
	LanguageCpp        = Language("cpp")
	LanguageJavaScript = Language("js")
	LanguageTypeScript = Language("ts")
	LanguageCSharp     = Language("csharp")
	LanguageKotlin     = Language("ko")
	LanguageSwift      = Language("sw")
	LanguageRust       = Language("rust")
)

var (
	languageModelToGRPCModel = map[Language]submissions_v1.Language{
		LanguageGo:         submissions_v1.Language_LANGUAGE_GO,
		LanguagePython:     submissions_v1.Language_LANGUAGE_PYTHON,
		LanguageJava:       submissions_v1.Language_LANGUAGE_JAVA,
		LanguageCpp:        submissions_v1.Language_LANGUAGE_CPP,
		LanguageJavaScript: submissions_v1.Language_LANGUAGE_JAVA_SCRIPT,
		LanguageTypeScript: submissions_v1.Language_LANGUAGE_TYPE_SCRIPT,
		LanguageCSharp:     submissions_v1.Language_LANGUAGE_CSHARP,
		LanguageKotlin:     submissions_v1.Language_LANGUAGE_KOTLIN,
		LanguageSwift:      submissions_v1.Language_LANGUAGE_SWIFT,
		LanguageRust:       submissions_v1.Language_LANGUAGE_RUST,
	}

	//nolint:exhaustive // unspecified not needed in map
	languageGRPCModelToModel = map[submissions_v1.Language]Language{
		submissions_v1.Language_LANGUAGE_GO:          LanguageGo,
		submissions_v1.Language_LANGUAGE_PYTHON:      LanguagePython,
		submissions_v1.Language_LANGUAGE_JAVA:        LanguageJava,
		submissions_v1.Language_LANGUAGE_CPP:         LanguageCpp,
		submissions_v1.Language_LANGUAGE_JAVA_SCRIPT: LanguageJavaScript,
		submissions_v1.Language_LANGUAGE_TYPE_SCRIPT: LanguageTypeScript,
		submissions_v1.Language_LANGUAGE_CSHARP:      LanguageCSharp,
		submissions_v1.Language_LANGUAGE_KOTLIN:      LanguageKotlin,
		submissions_v1.Language_LANGUAGE_SWIFT:       LanguageSwift,
		submissions_v1.Language_LANGUAGE_RUST:        LanguageRust,
	}
)

func NewLanguage(lang string) (Language, error) {
	language := Language(lang)
	if !language.IsValid() {
		return "", fmt.Errorf("invalid language: %s: %w", lang, errs.ErrInvalidArgument)
	}

	return language, nil
}

func NewLanguageFromGRPC(grpcModel submissions_v1.Language) (Language, error) {
	language, ok := languageGRPCModelToModel[grpcModel]
	if !ok {
		return "", fmt.Errorf("invalid language: %s: %w", grpcModel.String(), errs.ErrInvalidArgument)
	}
	return language, nil
}

func (l Language) String() string {
	return string(l)
}

func (l Language) IsValid() bool {
	switch l {
	case LanguageGo, LanguagePython, LanguageJava, LanguageCpp, LanguageJavaScript,
		LanguageTypeScript, LanguageCSharp, LanguageKotlin, LanguageSwift, LanguageRust:
		return true
	}

	return false
}

func (l Language) GRPCModel() (submissions_v1.Language, error) {
	grpcModel, ok := languageModelToGRPCModel[l]
	if !ok {
		return submissions_v1.Language_LANGUAGE_UNSPECIFIED, fmt.Errorf("no such language: %s", l.String())
	}

	return grpcModel, nil
}
