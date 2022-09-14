package server

func Init(addr string) {
	r := NewRouter()
	r.Run(addr)
}
