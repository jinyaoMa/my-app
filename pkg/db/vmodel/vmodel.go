package vmodel

type QueryCondition func(where func(query any, args ...any))
