package meson_msg

type CertMsg struct {
	Chain []byte
	Key   []byte
	Hash  uint32
}
