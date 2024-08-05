package utils

import "strings"

type ParamsValue = map[string]*string

type Param struct {
	Name         string  // Utilisez des chaînes directes au lieu des pointeurs pour simplifier
	DefaultValue *string // Gardez un pointeur ici si la valeur par défaut peut être absente
}

func NewParamsManager(args []string, paramsWaited []Param, lastArgInf bool) *ParamsValue {
	paramMap := make(ParamsValue)
	for i, param := range paramsWaited {
		if i < len(args) {
			paramMap[param.Name] = &args[i] // Assignez la valeur de args si disponible
		} else if param.DefaultValue != nil {
			paramMap[param.Name] = param.DefaultValue // Utilisez DefaultValue s'il y en a un
		} else {
			paramMap[param.Name] = nil // Assurez-vous qu'une clé est toujours présente, même si nil
		}
	}

	// Gestion du dernier paramètre pour prendre tous les arguments restants
	if lastArgInf && len(paramsWaited) > 0 && len(args) > len(paramsWaited) {
		lastParam := paramsWaited[len(paramsWaited)-1]
		remainingArgs := args[len(paramsWaited)-1:]
		concatenatedArgs := strings.Join(remainingArgs, " ") // Concatène tous les arguments restants avec un espace
		paramMap[lastParam.Name] = &concatenatedArgs
	}

	return &paramMap // Retourne le paramMap correctement rempli
}

func Get(values ParamsValue, key string) (value *string, null bool) {
	if val, exists := values[key]; !exists || val == nil {
		return nil, true
	}
	return values[key], false
}
