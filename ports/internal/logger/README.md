## Logger

This package defines methods for logging information in this project. The implementation is quite straightforward, it provides a standard logger, a warning logger and an error logger.

There is one caveat, though:

- Global logger vs. explicit dependency vs. new logger for every log message

I would have normally gone by passing the logger as an explicit dependency to all functions that write logs, but that extra parameter is kindof ugly and I wanted to avoid for no other reason than I like my functions to have fewer parameters. I could have also gone with global variable, that is thread-safe and there's nothing particularly wrong about that except the generally accepted principle that global variables are discouraged, but I do usually prefer avoiding global variables.

Thus I went with creating a new logger for each log message. Since logging, even from multiple loggers, appears to be thread-safe*, I felt that creating a package and exporting a few functions made the code look the nicest, which was the main reason behind the approach. There are logging packages available such as [zap](https://github.com/uber-go/zap) and [hclog](https://github.com/hashicorp/go-hclog), but since my approach seemed to work fine and thread-safe, I couldn't have cared less about improving the logging for this demo project.

Creating a new logger obviously adds some overhead, but that never seemed to be an issue. Of course, I would've changed this approach if there were any performance or synchronization issues, but I never encountered one.

\* Even if you do go with either of the other approaches, you still generally have multiple loggers (info, warn etc.) which raises the same synchronization concern. That is probably handleled if using the logging packages, idk.