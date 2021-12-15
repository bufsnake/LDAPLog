package ldap_server

import (
	"fmt"
	"github.com/bufsnake/ldap-server/pkg/datas"
	ldap "github.com/vjeantet/ldapserver"
	"log"
	"os"
)

type ldapserver struct {
	data      *datas.Data
	bind_addr string
	server    *ldap.Server
}

func NewLDAPServer(data *datas.Data, bind string) ldapserver {
	return ldapserver{data: data, bind_addr: bind}
}

func (l *ldapserver) Listen() error {
	ldap.Logger = log.New(os.Stdout, "[ldap server] ", log.LstdFlags)
	l.server = ldap.NewServer()
	routes := ldap.NewRouteMux()
	_ = routes.Bind(l.bind)
	_ = routes.Search(l.search)
	l.server.Handle(routes)
	err := l.server.ListenAndServe(l.bind_addr)
	if err != nil {
		return err
	}
	return nil
}

func (l *ldapserver) bind(w ldap.ResponseWriter, m *ldap.Message) {
	res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
	// 设置bind成功
	w.Write(res)
}

func (l *ldapserver) search(w ldap.ResponseWriter, m *ldap.Message) {
	search_req := m.GetSearchRequest()
	l.data.AddData(fmt.Sprintf("%s", search_req.BaseObject()))
	w.Write(ldap.NewSearchResultDoneResponse(1))
}

func (l *ldapserver) Stop() {
	l.server.Stop()
}
