<a name="unreleased"></a>
## [Unreleased]

### Chore
- setup git-chglog and generate CHANGELOG.md


<a name="0.7.0"></a>
## [0.7.0] - 2024-10-30
### Feat
- init git, git clone
- init fzf
- init tmux

### Lint
- **gomn:** check err before deferring file.Close()


<a name="0.6.3"></a>
## [0.6.3] - 2024-10-29
### Fix
- **aws/inboundIp:** nil != [] empty slice when comparing IpPerms


<a name="0.6.2"></a>
## [0.6.2] - 2024-10-28
### Fix
- **aws/inboundIp:** edit IP and keep port/protocol when editing SG


<a name="0.6.1"></a>
## [0.6.1] - 2024-10-25
### Fix
- **aws/inboundIp:** revoke and authorize the whole group


<a name="0.6.0"></a>
## [0.6.0] - 2024-10-25
### Feat
- **aws/inboundIp:** set inbound IPs for multiple profiles


<a name="0.5.0"></a>
## [0.5.0] - 2024-10-22
### Feat
- **macos:** setWallpaper


<a name="0.4.2"></a>
## [0.4.2] - 2024-10-17
### Fix
- **aws/inboundIp:** update security groups asynchronously


<a name="0.4.1"></a>
## [0.4.1] - 2024-10-17
### Fix
- **aws/inboundIp:** set ipv4/ipv6 independently of each other


<a name="0.4.0"></a>
## [0.4.0] - 2024-10-15
### Feat
- **inboundIp:** init


<a name="0.3.1"></a>
## [0.3.1] - 2024-10-13
### Test
- **tfgrep:** extract UnmarshalRegexArray and test it


<a name="0.3.0"></a>
## [0.3.0] - 2024-10-13
### Feat
- gomn (Generate One Month of Notes) ([#6](https://github.com/gforien/gf/issues/6))


<a name="0.2.0"></a>
## [0.2.0] - 2024-10-03
### Feat
- macos/hideAllWindows


<a name="0.1.0"></a>
## 0.1.0 - 2024-09-24
### Feat
- tfgrep ([#3](https://github.com/gforien/gf/issues/3))
- ecy ([#1](https://github.com/gforien/gf/issues/1))
- **kcp:** karabiner change profile
- **tfgrep:** discard empty lines ([#5](https://github.com/gforien/gf/issues/5))
- **tfgrep:** hide patterns ([#4](https://github.com/gforien/gf/issues/4))

### Fix
- **ecy:** configuration file path changed ([#2](https://github.com/gforien/gf/issues/2))


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
