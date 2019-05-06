# Netfilter queue DHCP packet handler

## Installation

-   Make sure all dependencies are installed and ready for handler to use
-   Run **make deps** for dependency installation and check. You need to have go-dep installed.
-   Build all binary files by running **make build**. It will build them in their respective folders.
-   In orderd for the packets to queue, IPTABLES rule must be invoked so that the packets can be routed to the designated queue..
Run **make rules** to invoke the IPTABLES rule. Command will ask you for sudo password. After the command is invoked, you will get a list of all IPTABLES rules. **-A INPUT -p udp -m udp --dport 67 -j NFQUEUE --queue-num 0** should be listed. Also, it will create a JSON configuration file in **/etc/netfilter-queue**. You need to populate empty values in order for the systems to work.

-   All systems can be built and ran separately
    - For DHCP packet handler run **make queue**
    - For UDP client run **make client**
    - FOR UDP server run **make server**
    - All commands will build the binary files and run them in the end

Every system can be ran separately by just going to their respective folders and run the created binary file. (example: cd client/ && sudo ./client)

