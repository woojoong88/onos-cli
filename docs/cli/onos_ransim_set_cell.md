## onos ransim set cell

Update a cell

```
onos ransim set cell <enbid> [field options] [flags]
```

### Options

```
      --arc int32         angle width of the coverage arc (default 120)
      --azimuth int32     azimuth of the coverage arc
      --color string      color label (default "blue")
  -h, --help              help for cell
      --lat float         geo location latitude (default 11)
      --lng float         geo location longitude (default 11)
      --max-ues uint32    maximum number of UEs connected (default 10000)
      --neighbors uints   neighbor cell ECGIs (default [])
      --tx-power float    transmit power (dB) (default 11)
```

### Options inherited from parent commands

```
      --auth-header string       Auth header in the form 'Bearer <base64>'
      --no-tls                   if present, do not use TLS
      --service-address string   the gRPC endpoint (default "ran-simulator:5150")
      --tls-cert-path string     the path to the TLS certificate
      --tls-key-path string      the path to the TLS key
```

### SEE ALSO

* [onos ransim set](onos_ransim_set.md)	 - Commands for setting RAN simulator model metrics and other information

###### Auto generated by spf13/cobra on 2-Apr-2021
