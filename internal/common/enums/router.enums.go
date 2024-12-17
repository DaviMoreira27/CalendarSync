package enums

type HttpMethod int

const (
	Get HttpMethod = iota
	Post
	Put
	Delete
)