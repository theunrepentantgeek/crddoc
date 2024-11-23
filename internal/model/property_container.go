package model

type PropertyContainer interface {
	Properties() []*Property
	Package() *Package
}
