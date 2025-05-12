package model

type PropertyContainer interface {
	Embeds() PropertyList
	Properties() PropertyList
	Package() *Package
}
