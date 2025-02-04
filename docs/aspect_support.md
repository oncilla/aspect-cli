---
sidebar_label: "support"
---
## aspect support

Interactive, human-escalated support for Bazel problems

### Synopsis

support collects recent Bazel invocations and collects relevant log files.

It then posts a message to a Slack channel on behalf of the user, posting the problem report in
a form that makes it easier for responders to understand the context and reproduce the problem.

```
aspect support [flags]
```

### Options

```
  -h, --help   help for support
```

### Options inherited from parent commands

```
      --aspect:config string   User-specified Aspect CLI config file. /dev/null indicates that all further --aspect:config flags will be ignored.
      --aspect:interactive     Interactive mode (e.g. prompts for user input)
```

### SEE ALSO

* [aspect](aspect.md)	 - Aspect CLI

