package main

import (
	"flag"
	"fmt"
)

func StringSlice(name string, value []string, usage string) *[]string {
	var s []string
	StringSliceVar(&s, name, value, usage)
	return &s
}

func StringSliceVar(p *[]string, name string, value []string, usage string) {
	flag.Var(&stringSliceFlag{
		value:        p,
		defaultValue: value,
	}, name, usage)
}

type stringSliceFlag struct {
	value        *[]string
	defaultValue []string
}

func (s *stringSliceFlag) String() string {
	if s == nil || s.value == nil {
		return "[]"
	}
	if len(*s.value) == 0 {
		return fmt.Sprintf("%v", s.defaultValue)
	}
	return fmt.Sprintf("%v", s.value)
}

func (s *stringSliceFlag) Set(value string) error {
	*s.value = append(*s.value, value)
	return nil
}
