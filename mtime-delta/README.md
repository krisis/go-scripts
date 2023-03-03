# Usage
`mtime-delta` can be used with xl-meta to compute the relative times at which versions of an object were created.

For example,
``` shell
xl-meta node/path/to/disk/obj/xl.meta | jq  -r '.Versions[]|.Header.ModTime' | mtime-delta 
```

