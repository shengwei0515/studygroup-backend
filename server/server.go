package server

import "studygroup"

func Init(addr string, session studygroup.WebSessionConfig) {
	r := NewRouter(session)
	r.Run(addr)
}
