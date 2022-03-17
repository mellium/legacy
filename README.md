# Legacy XMPP

[![GoDoc](https://godoc.org/mellium.im/legacy?status.svg)][docs]
[![Chat](https://img.shields.io/badge/XMPP-users@mellium.chat-orange.svg)](https://mellium.chat)
[![License](https://img.shields.io/badge/license-FreeBSD-blue.svg)](https://opensource.org/licenses/BSD-2-Clause)

<a href="https://opencollective.com/mellium" alt="Donate on Open Collective"><img src="https://opencollective.com/mellium/donate/button@2x.png?color=blue" width="200"/></a>

The `legacy` module contains XMPP related functionality that may still be useful
for backwards compatibility, but has been deprecated or obsoleted by other
functionality.
This module should be used selectively and with caution as many of its packages
may have security weaknesses, compatibility issues, or other major problems.

To use it in your project, import it (or any of its packages) like so:

```go
import mellium.im/legacy
```

If you'd like to contribute to the project, see [CONTRIBUTING.md].
As a general rule, no new functionality will be added to this package, this is
just a place to put existing functionality that previously existed in
[`mellium.im/xmpp`] that wasn't quite ready to be completely retired.


## License

The package may be used under the terms of the BSD 2-Clause License a copy of
which may be found in the file "[LICENSE]".

Unless you explicitly state otherwise, any contribution submitted for inclusion
in the work by you shall be licensed as above, without any additional terms or
conditions.


[docs]: https://pkg.go.dev/mellium.im/legacy
[`mellium.im/xmpp`]: https://mellium.im/xmpp/
[CONTRIBUTING.md]: https://mellium.im/docs/CONTRIBUTING
[LICENSE]: https://github.com/mellium/legacy/blob/main/LICENSE
