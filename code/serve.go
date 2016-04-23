func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	if fn := testHookServerServe; fn != nil {
		fn(srv, l)
	}
	var tempDelay time.Duration // how long to sleep on accept failure
	if err := srv.setupHTTP2(); err != nil {
		return err
	}
	for {
		rw, e := l.Accept()
		if e != nil {
		....

		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve()
	}
}
