# CloudInit Host Key Parser

This library parses SSH Host Keys from the output of CloudInit as provided by (for example) the AWS `ec2 get-console-output` command, in order that they can be used to verify host keys for SSH connections. Comments are removed from keys, though no validation of the actual keys is performed.

It is important to note that not all AMIs set the console correctly at boot time so this method is not 100% reliable, though if the output is present the parsing should succeed.

## Usage

```
// Use the AWS SDK for this
consoleOutput := GetConsoleOutput("i-12345")
if consoleOutput == "" {
    // No console output to read
}

hostKeys, err := hostkeys.Parse(consoleOutput)
if err != nil {
    if err == hostkeys.ErrNoStartHostKeysBlock {
        // No host keys in console output
    }
}

// hostKeys is a []string with the keys in the same order as read from the console output.
```
