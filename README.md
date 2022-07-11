# Legacy XMPP

[![Issue Tracker][badge]](https://mellium.im/issue)
[![Docs](https://pkg.go.dev/badge/mellium.im/legacy)](https://pkg.go.dev/mellium.im/legacy)
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

## Compatibility

This module does not follow any sort of backwards compatibility promise.
We will attempt to follow semver and keep the API backwards compatible on a
best-effort basis, but changes may be made at any time to fix bugs or make the
API easier to use.


## License

The package may be used under the terms of the BSD 2-Clause License a copy of
which may be found in the file "[LICENSE]".

Unless you explicitly state otherwise, any contribution submitted for inclusion
in the work by you shall be licensed as above, without any additional terms or
conditions.


[badge]: https://img.shields.io/badge/style-mellium%2fxmpp-green.svg?longCache=true&style=popout-square&label=issues
[`mellium.im/xmpp`]: https://mellium.im/xmpp/
[CONTRIBUTING.md]: https://mellium.im/docs/CONTRIBUTING
[LICENSE]: https://codeberg.org/mellium/xmpp/src/branch/main/LICENSE
