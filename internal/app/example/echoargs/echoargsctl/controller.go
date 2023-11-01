package echoargsctl

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/echoargs/echoargssrv"
	"example.com/internal/pkg/myconfig"
	"example.com/internal/pkg/mylog"
)

func New(config myconfig.MyConfigIface, logger mylog.MyLogIface, db *sql.DB) map[string]interface{} {
	srv := echoargssrv.New(config, logger)
	return map[string]interface{}{
		"GET /example/api/v1/echoargs/echo_query": func(ctx vo.SessionInfo, args echoargssrv.EhcoReqDto) (interface{}, error) {
			return srv.Echo(ctx, args)
		},

		"POST /example/api/v1/echoargs/echo_form": func(ctx vo.SessionInfo, args echoargssrv.EhcoReqDto) (interface{}, error) {
			return srv.Echo(ctx, args)
		},

		"POST /example/api/v1/echoargs/echo_multipart_form": func(ctx vo.SessionInfo, args echoargssrv.EhcoReqDto) (interface{}, error) {
			return srv.Echo(ctx, args)
		},

		"POST /example/api/v1/echoargs/echo_json": func(ctx vo.SessionInfo, args echoargssrv.EhcoReqDto) (interface{}, error) {
			return srv.Echo(ctx, args)
		},

		"POST /example/api/v2/echoargs/echo_json": func(ctx vo.SessionInfo, w http.ResponseWriter, r *http.Request) (interface{}, error) {
			arg := echoargssrv.EhcoReqDto{}
			if r.Body != nil {
				defer r.Body.Close()
				json.NewDecoder(r.Body).Decode(&arg)
			}
			b, _ := json.Marshal(arg)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", b)
			return w, nil
		},

		"GET /example/api/v2/echoargs/echo_query": func(ctx vo.SessionInfo, w http.ResponseWriter, r *http.Request) (interface{}, error) {
			r.ParseForm()
			ids := r.Form["id"]
			f32s := r.Form["f32"]
			f64s := r.Form["f64"]
			emails := r.Form["email"]
			sis := r.Form["si"]
			sf32s := r.Form["sf32"]
			sf64s := r.Form["sf64"]
			sss := r.Form["ss"]
			tms := r.Form["tm"]
			vs := r.Form["v"]
			
			w.Header().Set("Content-Type", "text/plain; charset=utf-8;")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "id: %s\n, f32: %s\n, f64: %s\n, email: %s\n, si: %s\n, sf32: %s\n, sf64: %s\n, ss: %s\n, tm: %s\n, v: %s\n",
				strings.Join(ids, ","), strings.Join(f32s, ", "), strings.Join(f64s, ", "), strings.Join(emails, ", "), strings.Join(sis, ", "),
				strings.Join(sf32s, ", "), strings.Join(sf64s, ", "), strings.Join(sss, ", "), strings.Join(tms, ", "), strings.Join(vs, ", "),
			)
			return w, nil
		},

		"POST /example/api/v2/echoargs/echo_form": func(ctx vo.SessionInfo, w http.ResponseWriter, r *http.Request) (interface{}, error) {
			r.ParseForm()
			ids := r.Form["id"]
			f32s := r.Form["f32"]
			f64s := r.Form["f64"]
			emails := r.Form["email"]
			sis := r.Form["si"]
			sf32s := r.Form["sf32"]
			sf64s := r.Form["sf64"]
			sss := r.Form["ss"]
			tms := r.Form["tm"]
			vs := r.Form["v"]
			
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "id: %s\n, f32: %s\n, f64: %s\n, email: %s\n, si: %s\n, sf32: %s\n, sf64: %s\n, ss: %s\n, tm: %s\n, v: %s\n",
				strings.Join(ids, ","), strings.Join(f32s, ", "), strings.Join(f64s, ", "), strings.Join(emails, ", "), strings.Join(sis, ", "),
				strings.Join(sf32s, ", "), strings.Join(sf64s, ", "), strings.Join(sss, ", "), strings.Join(tms, ", "), strings.Join(vs, ", "),
			)
			return w, nil
		},

		"POST /example/api/v2/echoargs/echo_multipart_form": func(ctx vo.SessionInfo, w http.ResponseWriter, r *http.Request) (interface{}, error) {
			r.ParseMultipartForm(1 << 20)
			ids := r.Form["id"]
			f32s := r.Form["f32"]
			f64s := r.Form["f64"]
			emails := r.Form["email"]
			sis := r.Form["si"]
			sf32s := r.Form["sf32"]
			sf64s := r.Form["sf64"]
			sss := r.Form["ss"]
			tms := r.Form["tm"]
			vs := r.Form["v"]
			
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "id: %s\n, f32: %s\n, f64: %s\n, email: %s\n, si: %s\n, sf32: %s\n, sf64: %s\n, ss: %s\n, tm: %s\n, v: %s\n",
				strings.Join(ids, ","), strings.Join(f32s, ", "), strings.Join(f64s, ", "), strings.Join(emails, ", "), strings.Join(sis, ", "),
				strings.Join(sf32s, ", "), strings.Join(sf64s, ", "), strings.Join(sss, ", "), strings.Join(tms, ", "), strings.Join(vs, ", "),
			)
			return w, nil
		},

		
	}
}
