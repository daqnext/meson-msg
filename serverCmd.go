package meson_msg

type SaveFileMsg struct {
	NameHash  string
	OriginUrl []string
	SizeLimit int64
}
