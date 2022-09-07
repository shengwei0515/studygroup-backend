package server

func Init() {
	r := NewRouter()
	r.Run("0.0.0.0:8080")
}
