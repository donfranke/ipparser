#IP Parser

This is a utility for taking a list of IP addresses in varying formats and extracts only the IP addresses. It is written in go. The result is a Splunk Processing Language (SPL) query.

# Usage

```
go ./ipparser i=[path to ip address file] t=[types of fields to include in result (src=src_ip, dest=dest_ip, both=src_ip and dest_ip)]
Example: ./ipparser -i=/tmp/iplist.txt -t=both
```

# Requirements
Go
