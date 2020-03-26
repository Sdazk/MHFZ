package dnsserver

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Andoryuuta/Erupe/config"
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

// Config struct allows configuring the server.
type Config struct {
	Logger      *zap.Logger
	ErupeConfig *config.Config
}

// Server is the MHF launcher HTTP server.
type Server struct {
	sync.Mutex
	logger          *zap.Logger
	erupeConfig     *config.Config
	dnsServer       *dns.Server
	forwardResolver net.Resolver
	isShuttingDown  bool
}

func isMHFHost(host string) bool {
	switch host {
	case
		"members.mhf-z.jp.",
		"cog-members.mhf-z.jp.",
		"dmm-members.mhf-z.jp.",
		"psvita-members.mhf-z.jp.",
		"ps3-members.mhf-z.jp.",
		"ps4-members.mhf-z.jp.",
		"wiiu-members.mhf-z.jp.",
		"xbox360-members.mhf-z.jp.",

		"cog.mhf-z.jp.",
		"dmm.mhf-z.jp.",
		"psvita.mhf-z.jp.",
		"ps3.mhf-z.jp",
		"ps4.mhf-z.jp.",
		"wiiu.mhf-z.jp.",
		"xbox360.mhf-z.jp.",

		"members.mhf-g.jp.",

		"www.capcom-onlinegames.jp.",

		"srv-mhf.capcom-networks.jp.",
		"sign-mhf.capcom-networks.jp.",
		"srv-mhf-psvita.capcom-networks.jp.",
		"auth-mhf-psvita.capcom-networks.jp.",
		"sign-mhf-psvita.capcom-networks.jp.",
		"srv-mhf-ps3.capcom-networks.jp.",
		"auth-mhf-ps3.capcom-networks.jp.",
		"sign-mhf-ps3.capcom-networks.jp.",
		"srv-mhf-ps4.capcom-networks.jp.",
		"auth-mhf-ps4.capcom-networks.jp.",
		"sign-mhf-ps4.capcom-networks.jp.",

		"mhfg.capcom.com.tw.",
		"mhfz.capcom.com.tw.",
		"mhf-n.capcom.com.tw.",
		//"img.capcom.com.tw.",
		"event.capcom.com.tw.",
		"test.capcom.com.tw.",
		"www.capcom.com.tw.",
		"ctc.capcom.com.tw.":

		return true
	}
	return false
}

// NewServer creates a new Server type.
func NewServer(config *Config) *Server {
	s := &Server{
		logger:      config.Logger,
		erupeConfig: config.ErupeConfig,
		dnsServer:   &dns.Server{},
		forwardResolver: net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, "udp", config.ErupeConfig.DNS.ForwardTo)
			},
		},
	}
	return s
}

// Start starts the server in a new goroutine.
func (s *Server) Start() error {
	s.dnsServer.Addr = fmt.Sprintf(":%d", s.erupeConfig.DNS.Port)
	s.dnsServer.Net = "udp"
	s.dnsServer.Handler = s

	serveError := make(chan error, 1)
	go func() {
		if err := s.dnsServer.ListenAndServe(); err != nil {
			// Send error if any.
			serveError <- err
		}
	}()

	// Get the error from calling ListenAndServe, otherwise assume it's good after 250 milliseconds.
	select {
	case err := <-serveError:
		return err
	case <-time.After(250 * time.Millisecond):
		return nil
	}
}

// Shutdown exits the server gracefully.
func (s *Server) Shutdown() {
	s.logger.Debug("Shutting down")

	s.Lock()
	s.isShuttingDown = true
	s.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.dnsServer.ShutdownContext(ctx); err != nil {
		// Just warn because we are shutting down the server anyway.
		s.logger.Warn("Got error on dnsServer shutdown", zap.Error(err))
	}
}

// ServeDNS serves a DNS request.
func (s *Server) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		s.logger.Info("Got DNS A request", zap.String("domain", domain))
		if isMHFHost(domain) {
			s.logger.Info("That domain is ours, resolving to us.")
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(s.erupeConfig.HostIP),
			})
		} else {
			s.logger.Info("That domain isn't ours, proxy DNS request to google.")
			// We don't have a local entry for it, lookup with our forward/proxied DNS resolver and return the response.
			ips, err := s.forwardResolver.LookupHost(context.Background(), domain)
			if err == nil {
				for _, ip := range ips {
					msg.Answer = append(msg.Answer, &dns.A{
						Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
						A:   net.ParseIP(ip),
					})
				}
			}
		}
	}
	w.WriteMsg(&msg)
}
