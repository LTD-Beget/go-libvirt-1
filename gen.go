// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type Proc struct {
	Name  string
	Value string
}

type Const struct {
	Name  string
	Value string
}

type ReqRes struct {
	Name string
	Args []string
}

func main() {
	re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")

	for _, fu := range []string{
		"https://raw.githubusercontent.com/libvirt/libvirt/master/src/remote/remote_protocol.x",
		"https://raw.githubusercontent.com/libvirt/libvirt/master/src/remote/qemu_protocol.x",
		"https://raw.githubusercontent.com/libvirt/libvirt/master/src/remote/lxc_protocol.x",
		"https://raw.githubusercontent.com/libvirt/libvirt/master/src/rpc/virnetprotocol.x",
		"https://raw.githubusercontent.com/libvirt/libvirt/master/src/rpc/virkeepaliveprotocol.x",
	} {
		procs := []Proc{}
		reqres := []ReqRes{}
		consts := []Const{}

		buf := bytes.NewBuffer(nil)
		res, err := http.Get(fu)
		die(err)
		_, err = io.Copy(buf, res.Body)
		die(err)

		res.Body.Close()

		sc := bufio.NewScanner(bytes.NewBuffer(re.ReplaceAll(buf.Bytes(), nil)))
		for sc.Scan() {

			if strings.TrimSpace(sc.Text()) == "" {
				continue
			}

			if strings.Contains(strings.TrimSpace(sc.Text()), "const ") {
				parts := strings.Split(strings.Replace(strings.TrimSpace(sc.Text()), "const ", "", 1), "=")
				name := strings.Replace(strings.TrimSpace(parts[0]), "VIR_", "", 1)
				name = strings.TrimSpace(strings.Replace(strings.Title(strings.Replace(strings.ToLower(name), "_", " ", -1)), " ", "", -1))
				value := strings.Replace(strings.Replace(strings.TrimSpace(parts[1]), ";", "", 1), "VIR_", "", 1)
				value = strings.TrimSpace(strings.Replace(strings.Title(strings.Replace(strings.ToLower(value), "_", " ", -1)), " ", "", -1))
				consts = append(consts, Const{Name: name, Value: value})
				continue
			}

			if strings.Contains(strings.TrimSpace(sc.Text()), "enum remote_procedure") {
				for sc.Scan() {
					// Procedures
					if strings.Contains(strings.TrimSpace(sc.Text()), "REMOTE_PROC_") {
						parts := strings.Split(strings.TrimSpace(sc.Text()), "=")
						name := strings.TrimSpace(parts[0])
						name = strings.TrimSpace(strings.Replace(strings.Title(strings.Replace(strings.ToLower(name), "_", " ", -1)), " ", "", -1))
						value := strings.TrimSpace(parts[1])
						value = value[:len(value)-1]
						procs = append(procs, Proc{Name: name, Value: value})
						continue
					}
					if strings.TrimSpace(sc.Text()) == "" {
						continue
					}
					if strings.HasPrefix(sc.Text(), "};") {
						break
					}
				}
				continue
			}

			if strings.HasSuffix(strings.TrimSpace(sc.Text()), "_args {") ||
				strings.HasSuffix(strings.TrimSpace(sc.Text()), "_ret {") ||
				strings.HasSuffix(strings.TrimSpace(sc.Text()), "_msg {") {
				parts := strings.Split(sc.Text(), " ")
				name := strings.TrimSpace(parts[1])
				name = strings.Replace(name, "_ret", "_Res", 1)
				name = strings.Replace(name, "_args", "_Req", 1)
				name = strings.Replace(name, "_msg", "_Msg", 1)
				name = strings.Replace(strings.Title(strings.Replace(strings.ToLower(name), "_", " ", -1)), " ", "", -1)
				qs := ReqRes{Name: name}
				for sc.Scan() {
					if strings.TrimSpace(sc.Text()) == "};" {
						break
					}
					vals := strings.Split(strings.TrimSpace(sc.Text()), " ")
					varname := strings.Replace(strings.Title(vals[len(vals)-1]), ";", "", 1)
					varname = strings.TrimSpace(strings.Replace(strings.Title(strings.Replace(varname, "_", " ", -1)), " ", "", -1))
					vartype := strings.Join(vals[:len(vals)-1], " ")
					vararr := false
					if idx := strings.Index(varname, "<"); idx > 0 {
						vararr = true
						varname = varname[:idx]
					}
					switch varname {
					case "Snap":
						varname = "Snapshot"
					case "Uuid":
						varname = "UUID"
					case "XmlDesc", "Dxml", "DomXml":
						varname = "XML"
					case "Dom", "Ddom":
						varname = "Domain"
					case "Doms":
						varname = "Domains"
					case "Net":
						varname = "Network"
					case "Dev":
						varname = "Device"
					}

					if strings.HasSuffix(vals[0], "_domain") {
						vartype = "*RemoteDomain"
					} else if strings.HasSuffix(vals[0], "_network") {
						vartype = "*RemoteNetwork"
					} else if strings.HasSuffix(vals[0], "_storage_vol") {
						vartype = "*RemoteStorageVolume"
					} else if strings.HasSuffix(vals[0], "_storage_pool") {
						vartype = "*RemoteStoragePool"
					} else if strings.HasSuffix(vals[0], "_string") {
						vartype = "string"
					} else if strings.HasSuffix(vals[0], "_nwfilter") {
						vartype = "*RemoteNwFilter"
					} else if strings.HasSuffix(vals[0], "_interface") {
						vartype = "*RemoteInterface"
					} else if strings.HasSuffix(vals[0], "_secret") {
						vartype = "*RemoteSecret"
					} else if strings.HasSuffix(vals[0], "_node_device") {
						vartype = "*RemoteNodeDevice"
					} else if strings.HasSuffix(vals[0], "_domain_snapshot") {
						vartype = "*RemoteDomainSnapshot"
					}
					switch vartype {
					case "char", "unsigned char":
						vartype = "byte"
					case "remote_uuid":
						vartype = "UUID"
					case "unsigned int":
						vartype = "uint32"
					case "unsigned short":
						vartype = "uint8"
					case "int":
						vartype = "int"
					case "unsigned hyper":
						vartype = "uint64"
					case "hyper":
						vartype = "int64"
					case "opaque":
						vartype = "byte"
					case "remote_typed_param":
						continue
					}
					if strings.HasPrefix(vartype, "remote_") {
						vartype = strings.Replace(strings.Title(strings.Replace(vartype, "_", " ", -1)), " ", "", -1)
						vartype = strings.Replace(vartype, "Ret", "Res", 1)
					}

					if vararr {
						vartype = "[]" + vartype
					}
					qs.Args = append(qs.Args, fmt.Sprintf("%s %s", varname, vartype))
				}
				reqres = append(reqres, qs)
				continue
			}

			if strings.TrimSpace(sc.Text()) == "" {
				continue
			}

			if strings.HasPrefix(strings.TrimSpace(sc.Text()), "struct remote_") {
				struct_name := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(sc.Text()), "struct remote_nonnull_"), " {")
				struct_name = strings.TrimSuffix(strings.TrimPrefix(struct_name, "struct remote_"), " {")
				switch struct_name {
				case "typed_param":
					continue
				case "domain", "network", "storage_pool", "storage_vol", "node_device", "domain_snapshot", "secret", "interface", "nwfilter", "snapshot":
					if struct_name == "nwfilter" {
						struct_name = "NwFilter"
					}

					if struct_name == "storage_vol" {
						struct_name = "storage_volume"
					}
					qs := ReqRes{Name: "Remote" + strings.Replace(strings.Title(strings.Replace(struct_name, "_", " ", -1)), " ", "", -1)}
					for sc.Scan() {
						if strings.TrimSpace(sc.Text()) == "};" {
							break
						}
						vals := strings.Split(strings.TrimSpace(sc.Text()), " ")
						varname := strings.Replace(strings.Title(vals[len(vals)-1]), ";", "", 1)
						varname = strings.Replace(strings.Title(strings.Replace(varname, "_", " ", -1)), " ", "", -1)
						varname = strings.Replace(varname, "Id", "ID", -1)

						switch varname {
						case "Snap":
							varname = "Snapshot"
						case "Uuid":
							varname = "UUID"
						case "XmlDesc", "Dxml", "DomXml":
							varname = "XML"
						case "Dom", "Ddom":
							varname = "Domain"
						case "Doms":
							varname = "Domains"
						case "Net":
							varname = "Network"
						case "Dev":
							varname = "Device"
						}

						vartype := strings.Join(vals[:len(vals)-1], " ")

						if strings.HasSuffix(vals[0], "_domain") {
							vartype = "*RemoteDomain"
						} else if strings.HasSuffix(vals[0], "_network") {
							vartype = "*RemoteNetwork"
						} else if strings.HasSuffix(vals[0], "_storage_vol") {
							vartype = "*RemoteStorageVolume"
						} else if strings.HasSuffix(vals[0], "_storage_pool") {
							vartype = "*RemoteStoragePool"
						} else if strings.HasSuffix(vals[0], "_string") {
							vartype = "string"
						} else if strings.HasSuffix(vals[0], "_nwfilter") {
							vartype = "*RemoteNwFilter"
						} else if strings.HasSuffix(vals[0], "_interface") {
							vartype = "*RemoteInterface"
						} else if strings.HasSuffix(vals[0], "_secret") {
							vartype = "*RemoteSecret"
						} else if strings.HasSuffix(vals[0], "_node_device") {
							vartype = "*RemoteNodeDevice"
						} else if strings.HasSuffix(vals[0], "_domain_snapshot") {
							vartype = "*RemoteDomainSnapshot"
						}
						switch vartype {
						case "char", "unsigned char":
							vartype = "byte"
						case "remote_uuid":
							vartype = "UUID"
						case "unsigned int":
							vartype = "uint32"
						case "unsigned short":
							vartype = "uint8"
						case "int":
							vartype = "int"
						case "unsigned hyper":
							vartype = "uint64"
						case "hyper":
							vartype = "int64"
						case "opaque":
							vartype = "byte"
						case "remote_typed_param":
							continue
						}
						if strings.HasPrefix(vartype, "remote_") {
							vartype = strings.Replace(strings.Title(strings.Replace(vartype, "_", " ", -1)), " ", "", -1)
						}
						if strings.Contains(varname, "<") {
							vartype = "[]" + vartype
						}
						switch varname {
						case "Snap":
							varname = "Snapshot"
						case "XmlDesc":
							varname = "Xml"
						case "Dom":
							varname = "Domain"
						case "Doms":
							varname = "Domains"
						case "Net":
							varname = "Network"
						case "Dev":
							varname = "Device"
						}
						if idx := strings.Index(varname, "<"); idx > 0 {
							varname = varname[:idx]
						}

						qs.Args = append(qs.Args, fmt.Sprintf("%s %s", varname, vartype))
					}
					reqres = append(reqres, qs)
				default:
					qs := ReqRes{Name: "Remote" + strings.Replace(strings.Title(strings.Replace(struct_name, "_", " ", -1)), " ", "", -1)}
					for sc.Scan() {
						if strings.TrimSpace(sc.Text()) == "};" {
							break
						}
						vals := strings.Split(strings.TrimSpace(sc.Text()), " ")
						varname := strings.Replace(strings.Title(vals[len(vals)-1]), ";", "", 1)
						varname = strings.Replace(strings.Title(strings.Replace(varname, "_", " ", -1)), " ", "", -1)
						varname = strings.Replace(varname, "Id", "ID", -1)

						switch varname {
						case "Domain":
							if qs.Name == "RemoteError" {
								varname = "DomainID"
							}
						case "Snap":
							varname = "Snapshot"
						case "Uuid":
							varname = "UUID"
						case "XmlDesc", "Dxml", "DomXml":
							varname = "XML"
						case "Dom", "Ddom":
							varname = "Domain"
						case "Doms":
							varname = "Domains"
						case "Net":
							varname = "Network"
						case "Dev":
							varname = "Device"
						}

						vartype := strings.Join(vals[:len(vals)-1], " ")

						if strings.HasSuffix(vals[0], "_domain") {
							vartype = "*RemoteDomain"
						} else if strings.HasSuffix(vals[0], "_network") {
							vartype = "*RemoteNetwork"
						} else if strings.HasSuffix(vals[0], "_storage_vol") {
							vartype = "*RemoteStorageVolume"
						} else if strings.HasSuffix(vals[0], "_storage_pool") {
							vartype = "*RemoteStoragePool"
						} else if strings.HasSuffix(vals[0], "_string") {
							vartype = "string"
						} else if strings.HasSuffix(vals[0], "_nwfilter") {
							vartype = "*RemoteNwFilter"
						} else if strings.HasSuffix(vals[0], "_interface") {
							vartype = "*RemoteInterface"
						} else if strings.HasSuffix(vals[0], "_secret") {
							vartype = "*RemoteSecret"
						} else if strings.HasSuffix(vals[0], "_node_device") {
							vartype = "*RemoteNodeDevice"
						} else if strings.HasSuffix(vals[0], "_domain_snapshot") {
							vartype = "*RemoteDomainSnapshot"
						}
						switch vartype {
						case "char", "unsigned char":
							vartype = "byte"
						case "remote_uuid":
							vartype = "UUID"
						case "unsigned int":
							vartype = "uint32"
						case "unsigned short":
							vartype = "uint8"
						case "int":
							vartype = "int"
						case "unsigned hyper":
							vartype = "uint64"
						case "hyper":
							vartype = "int64"
						case "opaque":
							vartype = "byte"
						case "remote_typed_param":
							continue
						}
						if strings.HasPrefix(vartype, "remote_") {
							vartype = strings.Replace(strings.Title(strings.Replace(vartype, "_", " ", -1)), " ", "", -1)
						}
						if strings.Contains(varname, "<") {
							vartype = "[]" + vartype
						}
						switch varname {
						case "Snap":
							varname = "Snapshot"
						case "XmlDesc":
							varname = "Xml"
						case "Dom":
							varname = "Domain"
						case "Doms":
							varname = "Domains"
						case "Net":
							varname = "Network"
						case "Dev":
							varname = "Device"
						}
						if idx := strings.Index(varname, "<"); idx > 0 {
							varname = varname[:idx]
						}

						qs.Args = append(qs.Args, fmt.Sprintf("%s %s", varname, vartype))
					}
					reqres = append(reqres, qs)

				}
			}
			fmt.Printf("ttt %s\n", sc.Text())
		}
		die(sc.Err())

		u, _ := url.Parse(fu)
		fname := strings.Replace(filepath.Base(u.Path), ".x", ".go", 1)
		f, err := os.Create("gen_" + fname)
		die(err)

		packageTemplate.Execute(f, struct {
			URL    string
			Procs  []Proc
			ReqRes []ReqRes
			Consts []Const
		}{
			URL:    fu,
			Procs:  procs,
			ReqRes: reqres,
			Consts: consts,
		})
		f.Close()
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`
package libvirt

const (
{{- range .Consts }}
		{{ printf "%s = %s" .Name .Value }}
{{- end }}
)

const (
{{- range .Procs }}
		{{ printf "%s RemoteProcedure = %s" .Name .Value }}
{{- end }}
)

{{- range .ReqRes }}
{{ printf "\ntype %s struct {" .Name }}
{{- range .Args }}
		{{ printf "%s" . }}
{{- end }}
{{ printf "}" }}

{{- end }}
`))
