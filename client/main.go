package main

import (
    "net/http"
    "fmt"
    "log"
)

const htmlcode = `
<html>
<head>
</head>
<body>
<script type='text/javascript' src='/static/jwplayer.js'></script>

<div id='mediaspace'>This text will be replaced</div>

<script type='text/javascript'>
  jwplayer('mediaspace').setup({
    'file': 'test.flv',
    'flashplayer': '/static/player.swf',
    'streamer': 'rtmp://127.0.0.1/flvplayback/',
    'autostart': 'true',
    'controlbar': 'bottom',
    'width': '470',
    'height': '290'
  });
</script>
</body>
</html>
`

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "%s", htmlcode)
    })
    http.Handle("/static/", http.FileServer(http.Dir("")))
    log.Fatal(http.ListenAndServe(":80", nil))
}
