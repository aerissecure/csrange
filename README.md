# csrange

CSRange is a simple commandline utility for working with what I call comma separated ranges (CSR). These are typically found in programs like nmap or nessus and commonly used when referring to a group of ports. CSR provides formatting and counting operations.

Pass a CSR into csrange to get an ordered, unique, and compact CSR in return. Using the `-i` option will return a comma separated list of values instead that are ordered and unique. The `-c` option will also provide a count of the total number of integers the CSR represents (the same value with and without `-i`).