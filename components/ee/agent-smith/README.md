# Agent Smith 🕵️‍♂️

Agent smith is the component that takes care of policing workspaces
against threats and noisy neighbours to make sure that everyone gets a safe and
smooth Gitpod experience.

## How to add new signatures?
```
# find something to match, e.g. using elfdump or by inspecting scripts
agent-smith signature elfdump <binary>

# create a new signature
agent-smith signature new ...
```

Hint: do not use the `base64` to encode a signature

## How can I check if a signature matches?
```
agent-smith signature new <signature-args> | agent-smith signature match <test-binary>
```




# Scratchpad

```
leeway build components/agent-smith:falco-bpf-probe --serve 0.0.0.0:8081
ssh -p 2222 root@127.0.0.1  curl -L -o /root/probe.o http://10.0.2.2:8081/probe.o
```
