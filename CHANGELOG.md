<a name="unreleased"></a>
## [Unreleased]

### Chore
- setup git-chglog and generate CHANGELOG.md
- **git-chglog:** fix git-chglog default regex and regenerate CHANGELOG


<a name="0.7.0"></a>
## [0.7.0] - 2024-10-30
### Feat


- init git, git clone ([`93bd768`](https://github.com/gforien/gf/commit/93bd7680f5a5e45ac8e67b865ebf7974ac514baa))


- init fzf ([`4d60e58`](https://github.com/gforien/gf/commit/4d60e583db59b6fbccdf694a4c4146af8299088b))


- init tmux ([`4c7a837`](https://github.com/gforien/gf/commit/4c7a837d718eb402925f1065b2ef5fcf1cbc5c61))

### Lint


- **gomn:** check err before deferring file.Close() ([`724c8e6`](https://github.com/gforien/gf/commit/724c8e648e93654ce1af87e218d5e721fdcf1774))


<a name="0.6.3"></a>
## [0.6.3] - 2024-10-29
### Fix


- **aws/inboundIp:** nil != [] empty slice when comparing IpPerms ([`3de6cf6`](https://github.com/gforien/gf/commit/3de6cf61726b78c25eec1841b0da134ccd1e2fa8))


<a name="0.6.2"></a>
## [0.6.2] - 2024-10-28
### Fix


- **aws/inboundIp:** edit IP and keep port/protocol when editing SG ([`55c4b8c`](https://github.com/gforien/gf/commit/55c4b8caa1a837023aa8e6ac3adb48fb5a0c6dec))


<a name="0.6.1"></a>
## [0.6.1] - 2024-10-25
### Fix


- **aws/inboundIp:** revoke and authorize the whole group ([`1c124c0`](https://github.com/gforien/gf/commit/1c124c024186489748cc8e8abeb8929fee3521ae))


<a name="0.6.0"></a>
## [0.6.0] - 2024-10-25
### Feat


- **aws/inboundIp:** set inbound IPs for multiple profiles ([`99f4cc7`](https://github.com/gforien/gf/commit/99f4cc7cb636f65392bc502c83e2c56805e32bfb))


<a name="0.5.0"></a>
## [0.5.0] - 2024-10-22
### Feat


- **macos:** setWallpaper ([`e5b05ce`](https://github.com/gforien/gf/commit/e5b05ce93bb66bb86763bb45e367358ac893083e))


<a name="0.4.2"></a>
## [0.4.2] - 2024-10-17
### Fix


- **aws/inboundIp:** update security groups asynchronously ([`f08c47a`](https://github.com/gforien/gf/commit/f08c47a0b17c0d35083238210f9f8183f78716d6))


<a name="0.4.1"></a>
## [0.4.1] - 2024-10-17
### Fix


- **aws/inboundIp:** set ipv4/ipv6 independently of each other ([`cb70898`](https://github.com/gforien/gf/commit/cb7089873cc565e6fb7f137eca14185696cf35ba))


<a name="0.4.0"></a>
## [0.4.0] - 2024-10-15
### Feat


- **inboundIp:** init ([`1ea17fd`](https://github.com/gforien/gf/commit/1ea17fd11b4697bb3e69209e3d9eb0f4f762475a))


<a name="0.3.1"></a>
## [0.3.1] - 2024-10-13
### Test


- **tfgrep:** extract UnmarshalRegexArray and test it ([`0e80188`](https://github.com/gforien/gf/commit/0e8018801b8c57c2a35dd86a3e8e9ffbde5f2b92))


<a name="0.3.0"></a>
## [0.3.0] - 2024-10-13
### Feat


- gomn (Generate One Month of Notes) ([#6](https://github.com/gforien/gf/issues/6)) ([`b12360d`](https://github.com/gforien/gf/commit/b12360da980bcf486e3d4b0972930b75e619b338))


<a name="0.2.0"></a>
## [0.2.0] - 2024-10-03
### Feat


- macos/hideAllWindows ([`3126bf4`](https://github.com/gforien/gf/commit/3126bf4807845a375bab608947d0db9808afff8c))


<a name="0.1.0"></a>
## 0.1.0 - 2024-09-24
### Feat


- tfgrep ([#3](https://github.com/gforien/gf/issues/3)) ([`1261180`](https://github.com/gforien/gf/commit/126118083a465c54949f04f74822539bc32cf120))


- ecy ([#1](https://github.com/gforien/gf/issues/1)) ([`2f8bafe`](https://github.com/gforien/gf/commit/2f8bafe91aef35bccb457c307c963878a24a312e))


- **kcp:** karabiner change profile ([`eb94f08`](https://github.com/gforien/gf/commit/eb94f082b0dd3accf04da0d58046e441ee306fbc))


- **tfgrep:** discard empty lines ([#5](https://github.com/gforien/gf/issues/5)) ([`38fd47e`](https://github.com/gforien/gf/commit/38fd47e24b2ee42a100347331903d1fb08609f69))


- **tfgrep:** hide patterns ([#4](https://github.com/gforien/gf/issues/4)) ([`3585af5`](https://github.com/gforien/gf/commit/3585af513fc2b0940eedc3646d89c6645a5066c8))

### Fix


- **ecy:** configuration file path changed ([#2](https://github.com/gforien/gf/issues/2)) ([`14750d0`](https://github.com/gforien/gf/commit/14750d0224917b76d67cbf02b2a2c4230085c096))


[Unreleased]: https://github.com/gforien/gf/compare/0.7.0...HEAD
[0.7.0]: https://github.com/gforien/gf/compare/0.6.3...0.7.0
[0.6.3]: https://github.com/gforien/gf/compare/0.6.2...0.6.3
[0.6.2]: https://github.com/gforien/gf/compare/0.6.1...0.6.2
[0.6.1]: https://github.com/gforien/gf/compare/0.6.0...0.6.1
[0.6.0]: https://github.com/gforien/gf/compare/0.5.0...0.6.0
[0.5.0]: https://github.com/gforien/gf/compare/0.4.2...0.5.0
[0.4.2]: https://github.com/gforien/gf/compare/0.4.1...0.4.2
[0.4.1]: https://github.com/gforien/gf/compare/0.4.0...0.4.1
[0.4.0]: https://github.com/gforien/gf/compare/0.3.1...0.4.0
[0.3.1]: https://github.com/gforien/gf/compare/0.3.0...0.3.1
[0.3.0]: https://github.com/gforien/gf/compare/0.2.0...0.3.0
[0.2.0]: https://github.com/gforien/gf/compare/0.1.0...0.2.0
