# Usage
`filter-log` can be used to select lines from `mc admin trace ...` output (non-verbose) using `--since` and `--until` flags as appropriate

For example,

``` shell
cat admin-trace-nover.txt | filter-log -s "2023-03-02T21:19:15" -u "2023-03-02T21:22:36"
```

Note: Time values passed to `--since` and --`until` need to be of the format in the example.
