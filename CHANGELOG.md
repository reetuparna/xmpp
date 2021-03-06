# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

### Added

- internal/xmpptest: new `Tokens` type for testing invalid XML


## v0.17.1 — 2020-11-21

### Breaking

- roster: remove workaround for a bug in Go versions prior to 1.14 which is now
  the earliest supported version
- xmpp: the `Encode` and `EncodeElement` methods now take a context and respect
  its deadline

### Added

- internal/integration: new package for writing integration tests
- internal/integration/ejabberd: [Ejabberd] support for integration tests
- internal/integration/prosody: [Prosŏdy] support for integration tests
- internal/xmpptest: new `ClientServer` for testing two connected sessions
- xmpp: new `EncodeIQ` and `EncodeIQElement` methods

[Ejabberd]: https://www.ejabberd.im/
[Prosŏdy]: https://prosody.im/


### Fixed

- stanza: converting stanzas with empty to/from attributes no longer fails
- xmpp: fixed data race that could result in invalid session state and lead to
  writes on a closed session and other state related issues
- xmpp: the `Send` and `SendElement` methods now respect the context deadline


## v0.17.0 — 2020-11-11

### Breaking

- sasl2: removed experimental package
- xmpp: removed option to make STARTTLS feature optional


### Added

- styling: new package implementing [XEP-0393: Message Styling]
- xmpp: `ConnectionState` method


### Fixed

- roster: iters that have errored no longer panic when closed
- sasl2: using TeeIn/TeeOut no longer breaks SCRAM based SASL mechanisms
- xmpp: using TeeIn/TeeOut no longer breaks SCRAM based SASL mechanisms
- xmpp: stream negotiation no longer fails when the only required features
  cannot yet be negotiated because they depend on optional features


[XEP-0393: Message Styling]: https://xmpp.org/extensions/xep-0393.html


## v0.16.0 — 2020-03-08

### Breaking

- xmpp: the end element is now included in the token stream passed to handlers
- xmpp: SendElement now wraps the entire stream, not just the first element


### Added

- receipts: new package implementing [XEP-0333: Chat Markers]
- roster: add handler and mux option for roster pushes


[XEP-0333: Chat Markers]: https://xmpp.org/extensions/xep-0333.html


### Fixed

- mux: fix broken `Decode` and possible infinite loop due to cutting off the
  last token in a buffered XML token stream
- roster: work around a bug in Go 1.13 where `io.EOF` may be returned from the
  XML decoder


## v0.15.0 — 2020-02-28

### Breaking

- all: dropped support for versions of Go before 1.13
- mux: move `Wrap{IQ,Presence,Message}` functions to methods on the stanza types


### Added

- mux: ability to select handlers by stanza payload
- mux: new handler types and API
- ping: a function for easily encoding pings and handling errors
- ping: a handler and mux option for responding to pings
- stanza: ability to convert stanzas to/from `xml.StartElement`
- stanza: API to simplify replying to IQs
- uri: new package for parsing XMPP URI's and IRI's
- xtime: new package for handling [XEP-0202: Entity Time] and [XEP-0082: XMPP Date and Time Profiles]


[XEP-0202: Entity Time]: https://xmpp.org/extensions/xep-0202.html
[XEP-0082: XMPP Date and Time Profiles]: https://xmpp.org/extensions/xep-0082.html


### Fixed

- dial: if a port number is present in a JID it was previously ignored


## v0.14.0 — 2019-08-18

### Breaking

- ping: remove `IQ` function and replace with struct based API


### Added

- ping: add `IQ` struct based encoding API


### Changed

- stanza: a zero value `IQType` now marshals as "get"
- xmpp: read timeouts are now returned instead of ignored


### Fixed

- dial: fix broken fallback to domainpart
- xmpp: allow whitespace keepalives
- roster: the iterator now correctly closes the underlying TokenReadCloser
- xmpp: fix bug where stream processing could stop after an IQ was received


## v0.13.0 — 2019-07-27

### Breaking

- xmpp: change `Handler` to take an `xmlstream.TokenReadEncoder`
- xmpp: replace `EncodeToken` and `Flush` with `TokenWriter`
- xmpp: replace `Token` with `TokenReader`


### Added

- examples/echobot: add graceful shutdown on SIGINT
- xmpp: `Encode` and `EncodeElement` methods


### Changed

- xmpp: calls to `Serve` no longer return `io.EOF` on success


### Fixed

- examples/echobot: calling `Send` from within the handler resulted in deadlock
- xmpp: closing the input stream was racy, resulting in invalid XML


## v0.12.0

### Breaking

- dial: moved network dialing types and functions into new package.
- dial: use underlying net.Dialer's DNS Resolver in Dialer.
- stanza: change API of `WrapIQ` and `WrapPresence` to not abuse pointers
- xmpp: add new `SendIQ` API and remove response from `Send` and `SendElement`
- xmpp: new API for writing custom tokens to a session

### Fixed

- xmpp: let `Session.Close` operate concurrently with `SendElement` et al.
